package syncer

import (
	"context"

	"github.com/rotisserie/eris"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ratelimit"
	"github.com/solo-io/gloo/projects/gloo/pkg/syncer/sanitizer"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"

	"github.com/hashicorp/go-multierror"
	gwsyncer "github.com/solo-io/gloo/projects/gateway/pkg/syncer"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	v1snap "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/gloosnapshot"
	"github.com/solo-io/gloo/projects/gloo/pkg/translator"
	envoycache "github.com/solo-io/solo-kit/pkg/api/v1/control-plane/cache"
	"github.com/solo-io/solo-kit/pkg/api/v2/reporter"
)

type translatorSyncer struct {
	translator translator.Translator
	sanitizer  sanitizer.XdsSanitizer
	xdsCache   envoycache.SnapshotCache
	xdsHasher  envoycache.NodeHash
	reporter   reporter.StatusReporter
	// used for debugging purposes only
	latestSnap *v1snap.ApiSnapshot
	extensions []TranslatorSyncerExtension
	// used to track which envoy node IDs exist without belonging to a proxy
	extensionKeys  map[string]struct{}
	settings       *v1.Settings
	gatewaySyncer  *gwsyncer.TranslatorSyncer
	proxyClient    v1.ProxyClient
	writeNamespace string
	// Chan used to track the reports that need to be written to etcd.
	reportChan chan reporter.ResourceReports
}

type TranslatorSyncerExtensionParams struct {
	RateLimitServiceSettings ratelimit.ServiceSettings
}

type TranslatorSyncerExtensionFactory func(context.Context, TranslatorSyncerExtensionParams) (TranslatorSyncerExtension, error)

type UpgradeableTranslatorSyncerExtension interface {
	ExtensionName() string
	IsUpgrade() bool
}

// TODO(kdorosh) in follow up PR, update this interface so it can never error
// It is logically invalid for us to return an error here (translation of resources always needs to
// result in a xds snapshot, so we are resilient to pod restarts)
type TranslatorSyncerExtension interface {
	Sync(
		ctx context.Context,
		snap *v1snap.ApiSnapshot,
		settings *v1.Settings,
		xdsCache envoycache.SnapshotCache,
		reports reporter.ResourceReports,
	) (string, error)
}

func NewTranslatorSyncer(
	translator translator.Translator,
	xdsCache envoycache.SnapshotCache,
	xdsHasher envoycache.NodeHash,
	sanitizer sanitizer.XdsSanitizer,
	statusReporter reporter.StatusReporter,
	devMode bool,
	extensions []TranslatorSyncerExtension,
	settings *v1.Settings,
	gatewaySyncer *gwsyncer.TranslatorSyncer,
	proxyClient v1.ProxyClient,
	writeNamespace string,
) v1snap.ApiSyncer {
	s := &translatorSyncer{
		translator:     translator,
		xdsCache:       xdsCache,
		xdsHasher:      xdsHasher,
		reporter:       statusReporter,
		extensions:     extensions,
		sanitizer:      sanitizer,
		settings:       settings,
		gatewaySyncer:  gatewaySyncer,
		proxyClient:    proxyClient,
		writeNamespace: writeNamespace,
	}
	if devMode {
		// TODO(ilackarms): move this somewhere else?
		go func() {
			_ = s.ServeXdsSnapshots()
		}()
	}
	return s
}

func (s *translatorSyncer) Sync(ctx context.Context, snap *v1snap.ApiSnapshot) error {
	logger := contextutils.LoggerFrom(ctx)
	reports := make(reporter.ResourceReports)
	var multiErr *multierror.Error

	// If gateway controller is enabled, run the gateway translation to generate proxies.
	// Use the ProxyClient interface to persist them either to an in-memory store or etcd as configured at startup.
	if s.gatewaySyncer != nil {
		logger.Debugf("getting proxies from gateway translation")
		if err := s.gatewaySyncer.Sync(ctx, snap); err != nil {
			multiErr = multierror.Append(multiErr, err)
		}
		proxyList, err := s.proxyClient.List(s.writeNamespace, clients.ListOpts{})
		if err != nil {
			multiErr = multierror.Append(multiErr, err)
		}
		snap.Proxies = proxyList
	}

	err := s.syncEnvoy(ctx, snap, reports)
	if err != nil {
		multiErr = multierror.Append(multiErr, err)
	}
	s.extensionKeys = map[string]struct{}{}
	for _, extension := range s.extensions {
		intermediateReports := make(reporter.ResourceReports)
		nodeID, err := extension.Sync(ctx, snap, s.settings, s.xdsCache, intermediateReports)
		if err != nil {
			multiErr = multierror.Append(multiErr, err)
		}
		reports.Merge(intermediateReports)
		s.extensionKeys[nodeID] = struct{}{}
	}
	if err := s.reporter.WriteReports(ctx, reports, nil); err != nil {
		logger.Debugf("Failed writing report for proxies: %v", err)
		multiErr = multierror.Append(multiErr, eris.Wrapf(err, "writing reports"))
	}
	//After reports are written for proxies, save in gateway syncer (previously gw watched for status changes to proxies)
	if s.gatewaySyncer != nil {
		s.gatewaySyncer.UpdateProxies(ctx)
	}
	return multiErr.ErrorOrNil()
}

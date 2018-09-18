package templates

import (
	"text/template"
)

var ResourceGroupEmitterTemplate = template.Must(template.New("resource_group_emitter").Funcs(funcs).Parse(`package {{ .Project.PackageName }}

{{- $client_declarations := new_str_slice }}
{{- $clients := new_str_slice }}
{{- range .Resources}}
{{- $client_declarations := (append_str_slice $client_declarations (printf "%vClient %vClient"  (lower_camel .Name) .Name)) }}
{{- $clients := (append_str_slice $clients (printf "%vClient"  (lower_camel .Name))) }}
{{- end}}
{{- $client_declarations := (join_str_slice $client_declarations ", ") }}
{{- $clients := (join_str_slice $clients ", ") }}

import (
	"sync"
	"time"

	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/errors"
	"github.com/solo-io/solo-kit/pkg/utils/errutils"
)

var (
	m{{ .GoName }}SnapshotIn  = stats.Int64("{{ lower_camel .GoName }}_snap_emitter/snap_in", "The number of snapshots in", "1")
	m{{ .GoName }}SnapshotOut = stats.Int64("{{ lower_camel .GoName }}_snap_emitter/snap_out", "The number of snapshots out", "1")

	{{ lower_camel .GoName }}snapshotInView = &view.View{
		Name:        "{{ lower_camel .GoName }}_snap_emitter/snap_in",
		Measure:     m{{ .GoName }}SnapshotIn,
		Description: "The number of snapshots updates coming in",
		Aggregation: view.Count(),
		TagKeys:     []tag.Key{
		},
	}
	{{ lower_camel .GoName }}snapshotOutView = &view.View{
		Name:        "{{ lower_camel .GoName }}_snap_emitter/snap_out",
		Measure:     m{{ .GoName }}SnapshotOut,
		Description: "The number of snapshots updates going out",
		Aggregation: view.Count(),
		TagKeys:     []tag.Key{
		},
	}
)

func init() {
	view.Register({{ lower_camel .GoName }}snapshotInView, {{ lower_camel .GoName }}snapshotOutView)
}

type {{ .GoName }}Emitter interface {
	Register() error
{{- range .Resources}}
	{{ .Name }}() {{ .Name }}Client
{{- end}}
	Snapshots(watchNamespaces []string, opts clients.WatchOpts) (<-chan *{{ .GoName }}Snapshot, <-chan error, error)
}

func New{{ .GoName }}Emitter({{ $client_declarations }}) {{ .GoName }}Emitter {
	return New{{ .GoName }}EmitterWithEmit({{ $clients }}, make(chan struct{}))
}

func New{{ .GoName }}EmitterWithEmit({{ $client_declarations }}, emit <-chan struct{}) {{ .GoName }}Emitter {
	return &{{ lower_camel .GoName }}Emitter{
{{- range .Resources}}
		{{ lower_camel .Name }}: {{ lower_camel .Name }}Client,
{{- end}}
		forceEmit: emit,
	}
}

type {{ lower_camel .GoName }}Emitter struct {
	forceEmit <- chan struct{}
{{- range .Resources}}
	{{ lower_camel .Name }} {{ .Name }}Client
{{- end}}
}

func (c *{{ lower_camel .GoName }}Emitter) Register() error {
{{- range .Resources}}
	if err := c.{{ lower_camel .Name }}.Register(); err != nil {
		return err
	}
{{- end}}
	return nil
}

{{- range .Resources}}

func (c *{{ lower_camel $.GoName }}Emitter) {{ .Name }}() {{ .Name }}Client {
	return c.{{ lower_camel .Name }}
}
{{- end}}

func (c *{{ lower_camel .GoName }}Emitter) Snapshots(watchNamespaces []string, opts clients.WatchOpts) (<-chan *{{ .GoName }}Snapshot, <-chan error, error) {
	errs := make(chan error)
	var done sync.WaitGroup
	ctx := opts.Ctx


{{- range .Resources}}
	/* Create channel for {{ .Name }} */
	type {{ lower_camel .Name }}ListWithNamespace struct {
		list {{ .Name }}List
		namespace string
	}
	{{ lower_camel .Name }}Chan := make(chan {{ lower_camel .Name }}ListWithNamespace)
{{- end}}

	for _, namespace := range watchNamespaces {
{{- range .Resources}}
		/* Setup watch for {{ .Name }} */
		{{ lower_camel .Name }}NamespacesChan, {{ lower_camel .Name }}Errs, err := c.{{ lower_camel .Name }}.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting {{ .Name }} watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, {{ lower_camel .Name }}Errs, namespace+"-{{ lower_camel .PluralName }}")
		}(namespace)

{{- end}}


		/* Watch for changes and update snapshot */
		go func(namespace string) {
			for {
				select {
				case <-ctx.Done():
					return
{{- range .Resources}}
				case {{ lower_camel .Name }}List := <- {{ lower_camel .Name }}NamespacesChan:
					select {
					case <-ctx.Done():
						return
					case {{ lower_camel .Name }}Chan <- {{ lower_camel .Name }}ListWithNamespace{list:{{ lower_camel .Name }}List, namespace:namespace}:
					}
{{- end}}
				}
			}
		}(namespace)
	}

	
	snapshots := make(chan *{{ .GoName }}Snapshot)
	go func() {
		originalSnapshot := ApiSnapshot{}
		currentSnapshot := originalSnapshot.Clone()
		timer := time.NewTicker(time.Second * 5)
		sync := func() {
			if originalSnapshot.Hash() == currentSnapshot.Hash() {
				return
			}
			originalSnapshot = currentSnapshot.Clone()
			sentSnapshot := currentSnapshot.Clone()
			snapshots <- &sentSnapshot
		}

/* TODO (yuval-k): figure out how to make this work to avoid a stale snapshot.
		// construct the first snapshot from all the configs that are currently there
		// that guarantees that the first snapshot contains all the data.
		for range watchNamespaces {
{{- range .Resources}}
   {{ lower_camel .Name }}NamespacedList := <- {{ lower_camel .Name }}Chan:
	namespace := {{ lower_camel .Name }}NamespacedList.namespace
	{{ lower_camel .Name }}List := {{ lower_camel .Name }}NamespacedList.list

	currentSnapshot.{{ .PluralName }}.Clear(namespace)
	currentSnapshot.{{ .PluralName }}.Add({{ lower_camel .Name }}List...)

{{- end}}
		}
*/

		for {
			select {
			case <-timer.C:
				sync()
			case <-ctx.Done():
				close(snapshots)
				done.Wait()
				close(errs)
				return
			case <-c.forceEmit:
				sentSnapshot := currentSnapshot.Clone()
				snapshots <- &sentSnapshot
{{- range .Resources}}
			case {{ lower_camel .Name }}NamespacedList := <- {{ lower_camel .Name }}Chan:
				namespace := {{ lower_camel .Name }}NamespacedList.namespace
				{{ lower_camel .Name }}List := {{ lower_camel .Name }}NamespacedList.list

				currentSnapshot.{{ .PluralName }}.Clear(namespace)
				currentSnapshot.{{ .PluralName }}.Add({{ lower_camel .Name }}List...)

{{- end}}
			}

			// if we got here its because a new entry in the channel
			stats.Record(ctx, m{{ .GoName }}SnapshotIn.M(1))
		}
	}()
	return snapshots, errs, nil
}
`))

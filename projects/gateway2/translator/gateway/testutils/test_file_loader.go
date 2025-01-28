package testutils

import (
	"bytes"
	"context"
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/kgateway-dev/kgateway/pkg/schemes"

	"github.com/golang/protobuf/proto"
	"github.com/kgateway-dev/kgateway/pkg/utils/protoutils"
	"github.com/kgateway-dev/kgateway/projects/gateway2/api/v1alpha1"
	"github.com/kgateway-dev/kgateway/projects/gateway2/translator/irtranslator"
	"github.com/rotisserie/eris"

	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/solo-io/go-utils/contextutils"
)

var (
	NoFilesFound = errors.New("no k8s files found")
)

func LoadFromFiles(ctx context.Context, filename string) ([]client.Object, error) {

	fileOrDir, err := os.Stat(filename)
	if err != nil {
		return nil, err
	}

	var yamlFiles []string
	if fileOrDir.IsDir() {
		contextutils.LoggerFrom(ctx).Infof("looking for YAML files in directory tree rooted at: %s", fileOrDir.Name())
		err := filepath.WalkDir(filename, func(path string, d fs.DirEntry, _ error) error {
			if strings.HasSuffix(path, ".yml") || strings.HasSuffix(path, ".yaml") {
				yamlFiles = append(yamlFiles, path)
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
	} else {
		yamlFiles = append(yamlFiles, filename)
	}

	if len(yamlFiles) == 0 {
		return nil, NoFilesFound
	}

	contextutils.LoggerFrom(ctx).Infow("user configuration YAML files found", zap.Strings("files", yamlFiles))

	var resources []client.Object
	for _, file := range yamlFiles {
		objs, err := parseFile(ctx, file)
		if err != nil {
			return nil, err
		}

		for _, obj := range objs {
			clientObj, ok := obj.(client.Object)
			if !ok {
				return nil, errors.Errorf("cannot convert runtime.Object to client.Object: %+v", obj)
			}
			if clientObj.GetNamespace() == "" {
				// fill in default namespace
				clientObj.SetNamespace("default")
			}
			resources = append(resources, clientObj)
		}

	}

	return resources, nil
}

func parseFile(ctx context.Context, filename string) ([]runtime.Object, error) {
	scheme := schemes.GatewayScheme()
	if err := v1alpha1.Install(scheme); err != nil {
		return nil, err
	}
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	type metaOnly struct {
		metav1.TypeMeta   `json:",inline"`
		metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	}

	// Split into individual YAML documents
	resourceYamlStrings := bytes.Split(file, []byte("\n---\n"))

	// Create resources from YAML documents
	var genericResources []runtime.Object
	for _, objYaml := range resourceYamlStrings {

		// Skip empty documents
		if len(bytes.TrimSpace(objYaml)) == 0 {
			continue
		}

		var meta metaOnly
		if err := yaml.Unmarshal(objYaml, &meta); err != nil {
			contextutils.LoggerFrom(ctx).Warnw("failed to parse resource metadata, skipping YAML document",
				zap.String("filename", filename),
				zap.String("truncatedYamlDoc", truncateString(string(objYaml), 100)),
			)
			continue
		}

		gvk := schema.FromAPIVersionAndKind(meta.APIVersion, meta.Kind)
		obj, err := scheme.New(gvk)
		if err != nil {
			contextutils.LoggerFrom(ctx).Warnw("unknown resource kind",
				zap.String("filename", filename),
				zap.String("resourceKind", gvk.String()),
				zap.String("truncatedYamlDoc", truncateString(string(objYaml), 100)),
			)
			continue
		}
		if err := yaml.Unmarshal(objYaml, obj); err != nil {
			contextutils.LoggerFrom(ctx).Warnw("failed to parse resource YAML",
				zap.Error(err),
				zap.String("filename", filename),
				zap.String("resourceKind", gvk.String()),
				zap.String("resourceId", obj.(client.Object).GetName()+"."+obj.(client.Object).GetNamespace()),
				zap.String("truncatedYamlDoc", truncateString(string(objYaml), 100)),
			)
			continue
		}

		genericResources = append(genericResources, obj)
	}

	return genericResources, err
}

func truncateString(str string, num int) string {
	result := str
	if len(str) > num {
		result = str[0:num] + "..."
	}
	return result
}

func ReadProxyFromFile(filename string) (*irtranslator.TranslationResult, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, eris.Wrapf(err, "reading proxy file")
	}
	var proxy irtranslator.TranslationResult

	if err := UnmarshalAnyYaml(data, &proxy); err != nil {
		return nil, eris.Wrapf(err, "parsing proxy from file")
	}
	return &proxy, nil
}

func MarshalYaml(m proto.Message) ([]byte, error) {
	jsn, err := protoutils.MarshalBytes(m)
	if err != nil {
		return nil, err
	}
	return yaml.JSONToYAML(jsn)
}
func MarshalAnyYaml(m any) ([]byte, error) {
	jsn, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return yaml.JSONToYAML(jsn)
}

func UnmarshalAnyYaml(data []byte, into any) error {
	jsn, err := yaml.YAMLToJSON(data)
	if err != nil {
		return err
	}

	return json.Unmarshal(jsn, into)
}

//go:build ignore

package debugprint

import (
	"fmt"
	"log"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/onsi/ginkgo/v2"
	"sigs.k8s.io/yaml"

	"github.com/kgateway-dev/kgateway/v2/pkg/utils/protoutils"
)

func PrintYaml(ress ...proto.Message) {
	log.Printf(SprintYaml(ress...))
}

func GinkgoPrintYaml(ress ...proto.Message) {
	fmt.Fprint(ginkgo.GinkgoWriter, SprintYaml(ress...))
}

func PrintAny(anyVals ...interface{}) {
	for _, res := range anyVals {
		yam, _ := yaml.Marshal(res)
		log.Printf("%s", yam)
	}
}

func SprintAny(anyVals ...interface{}) string {
	var yams []string
	for _, res := range anyVals {
		yam, _ := yaml.Marshal(res)
		yams = append(yams, string(yam))
	}
	return strings.Join(yams, "\n---\n")
}

func SprintYaml(ress ...proto.Message) string {
	var yams []string
	for _, res := range ress {
		js, _ := protoutils.MarshalBytes(res)
		yam, _ := yaml.JSONToYAML(js)

		yams = append(yams, string(yam))
	}
	return strings.Join(yams, "\n---\n")
}

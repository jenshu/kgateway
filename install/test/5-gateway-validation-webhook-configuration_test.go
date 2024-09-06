package test

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	gloostringutils "github.com/solo-io/gloo/pkg/utils/stringutils"
	"github.com/solo-io/go-utils/stringutils"
	. "github.com/solo-io/k8s-utils/manifesttestutils"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/yaml"
)

var _ = Describe("WebhookValidationConfiguration helm test", func() {
	var allTests = func(rendererTestCase renderTestCase) {

		var (
			testManifest TestManifest
			//expectedChart *unstructured.Unstructured
		)
		prepareMakefile := func(namespace string, values helmValues) {
			tm, err := rendererTestCase.renderer.RenderManifest(namespace, values)
			ExpectWithOffset(1, err).NotTo(HaveOccurred(), "Failed to render manifest")
			testManifest = tm
		}

		//
		DescribeTable("Can remove DELETEs from webhook rules", func(resources []string, expectedRemoved int) {
			timeoutSeconds := 5

			// Count the "DELETES" as a sanity check.
			expectedDeletes := 6 - expectedRemoved
			expectedChart := generateExpectedChart(timeoutSeconds, resources, expectedDeletes)

			prepareMakefile(namespace, helmValues{
				valuesArgs: []string{
					fmt.Sprintf(`gateway.validation.webhook.timeoutSeconds=%d`, timeoutSeconds),
					`gateway.validation.webhook.skipDeleteValidationResources={` + strings.Join(resources, ",") + `}`,
				},
			})
			testManifest.ExpectUnstructured(expectedChart.GetKind(), expectedChart.GetNamespace(), expectedChart.GetName()).To(BeEquivalentTo(expectedChart))
		},
			Entry("virtualservices", []string{"virtualservices"}, 1),
			Entry("routetables", []string{"routetables"}, 1),
			Entry("gateways", []string{"gateways"}, 0),
			Entry("upstreams", []string{"upstreams"}, 1),
			Entry("secrets", []string{"secrets"}, 1),
			Entry("namespaces", []string{"namespaces"}, 1),
			Entry("ratelimitconfigs", []string{"ratelimitconfigs"}, 1),
			Entry("mutiple", []string{"virtualservices", "routetables", "secrets"}, 3),
			Entry("mutiple (with removal)", []string{"ratelimitconfigs", "secrets"}, 2),
			Entry("all", []string{"*"}, 6),
			Entry("empty", []string{}, 0),
		)
	}
	runTests(allTests)
})

func generateExpectedChart(timeoutSeconds int, skipDeletes []string, expectedDeletes int) *unstructured.Unstructured {
	rules := generateRules(skipDeletes)

	// indent "rules"
	m1 := regexp.MustCompile("\n")
	rules = m1.ReplaceAllString(rules, "\n    ")

	// Check that we have the expected number of DELETEs
	m2 := regexp.MustCompile(`DELETE`)
	deletes := m2.FindAllStringIndex(rules, -1)
	Expect(len(deletes)).To(Equal(expectedDeletes))

	return makeUnstructured(`

apiVersion: admissionregistration.k8s.io/v1
kind: ValidatingWebhookConfiguration
metadata:
  name: gloo-gateway-validation-webhook-` + namespace + `
  labels:
    app: gloo
    gloo: gateway
  annotations:
    "helm.sh/hook": pre-install, pre-upgrade
    "helm.sh/hook-weight": "5" # should come before cert-gen job
webhooks:
- name: gloo.` + namespace + `.svc  # must be a domain with at least three segments separated by dots
  clientConfig:
    service:
      name: gloo
      namespace: ` + namespace + `
      path: "/validation"
    caBundle: "" # update manually or use certgen job
  rules:
    ` + rules + `
  sideEffects: None
  matchPolicy: Exact
  timeoutSeconds: ` + strconv.Itoa(timeoutSeconds) + `
  admissionReviewVersions:
    - v1beta1
  failurePolicy: Ignore
`)
}

func generateRules(skipDeleteReources []string) string {
	rules := []map[string][]string{
		{
			"operations":  {"CREATE", "UPDATE", "DELETE"},
			"apiGroups":   {"gateway.solo.io"},
			"apiVersions": {"v1"},
			"resources":   {"virtualservices"},
		},
		{
			"operations":  {"CREATE", "UPDATE", "DELETE"},
			"apiGroups":   {"gateway.solo.io"},
			"apiVersions": {"v1"},
			"resources":   {"routetables"},
		},
		{
			"operations":  {"CREATE", "UPDATE"},
			"apiGroups":   {"gateway.solo.io"},
			"apiVersions": {"v1"},
			"resources":   {"gateways"},
		},
		{
			"operations":  {"CREATE", "UPDATE", "DELETE"},
			"apiGroups":   {"gloo.solo.io"},
			"apiVersions": {"v1"},
			"resources":   {"upstreams"},
		},
		{
			"operations":  {"DELETE"},
			"apiGroups":   {""},
			"apiVersions": {"v1"},
			"resources":   {"secrets"},
		},
		{
			"operations":  {"UPDATE", "DELETE"},
			"apiGroups":   {""},
			"apiVersions": {"v1"},
			"resources":   {"namespaces"},
		},
		{
			"operations":  {"CREATE", "UPDATE", "DELETE"},
			"apiGroups":   {"ratelimit.solo.io"},
			"apiVersions": {"v1alpha1"},
			"resources":   {"ratelimitconfigs"},
		},
	}

	finalRules := []map[string][]string{}
	for i, rule := range rules {
		if stringutils.ContainsAny([]string{rule["resources"][0], "*"}, skipDeleteReources) {
			rule["operations"] = gloostringutils.DeleteOneByValue(rule["operations"], "DELETE")
			// A namespace with an updated to a label can cause it to no longer be watched,
			// equivalent to deleting it from the controller's perspective
			if rule["resources"][0] == "namespaces" {
				rule["operations"] = gloostringutils.DeleteOneByValue(rule["operations"], "UPDATE")
			}
		}

		if len(rule["operations"]) != 0 {
			finalRules = append(finalRules, rules[i])
		}
	}

	str, err := yaml.Marshal(finalRules)
	Expect(err).NotTo(HaveOccurred())
	return string(str)
}

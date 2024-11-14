/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/solo-io/gloo/projects/controller/pkg/api/v1/kube/apis/gloo.solo.io/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// EndpointLister helps list Endpoints.
// All objects returned here must be treated as read-only.
type EndpointLister interface {
	// List lists all Endpoints in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Endpoint, err error)
	// Endpoints returns an object that can list and get Endpoints.
	Endpoints(namespace string) EndpointNamespaceLister
	EndpointListerExpansion
}

// endpointLister implements the EndpointLister interface.
type endpointLister struct {
	listers.ResourceIndexer[*v1.Endpoint]
}

// NewEndpointLister returns a new EndpointLister.
func NewEndpointLister(indexer cache.Indexer) EndpointLister {
	return &endpointLister{listers.New[*v1.Endpoint](indexer, v1.Resource("endpoint"))}
}

// Endpoints returns an object that can list and get Endpoints.
func (s *endpointLister) Endpoints(namespace string) EndpointNamespaceLister {
	return endpointNamespaceLister{listers.NewNamespaced[*v1.Endpoint](s.ResourceIndexer, namespace)}
}

// EndpointNamespaceLister helps list and get Endpoints.
// All objects returned here must be treated as read-only.
type EndpointNamespaceLister interface {
	// List lists all Endpoints in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Endpoint, err error)
	// Get retrieves the Endpoint from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.Endpoint, error)
	EndpointNamespaceListerExpansion
}

// endpointNamespaceLister implements the EndpointNamespaceLister
// interface.
type endpointNamespaceLister struct {
	listers.ResourceIndexer[*v1.Endpoint]
}

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

// UpstreamGroupLister helps list UpstreamGroups.
// All objects returned here must be treated as read-only.
type UpstreamGroupLister interface {
	// List lists all UpstreamGroups in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.UpstreamGroup, err error)
	// UpstreamGroups returns an object that can list and get UpstreamGroups.
	UpstreamGroups(namespace string) UpstreamGroupNamespaceLister
	UpstreamGroupListerExpansion
}

// upstreamGroupLister implements the UpstreamGroupLister interface.
type upstreamGroupLister struct {
	listers.ResourceIndexer[*v1.UpstreamGroup]
}

// NewUpstreamGroupLister returns a new UpstreamGroupLister.
func NewUpstreamGroupLister(indexer cache.Indexer) UpstreamGroupLister {
	return &upstreamGroupLister{listers.New[*v1.UpstreamGroup](indexer, v1.Resource("upstreamgroup"))}
}

// UpstreamGroups returns an object that can list and get UpstreamGroups.
func (s *upstreamGroupLister) UpstreamGroups(namespace string) UpstreamGroupNamespaceLister {
	return upstreamGroupNamespaceLister{listers.NewNamespaced[*v1.UpstreamGroup](s.ResourceIndexer, namespace)}
}

// UpstreamGroupNamespaceLister helps list and get UpstreamGroups.
// All objects returned here must be treated as read-only.
type UpstreamGroupNamespaceLister interface {
	// List lists all UpstreamGroups in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.UpstreamGroup, err error)
	// Get retrieves the UpstreamGroup from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.UpstreamGroup, error)
	UpstreamGroupNamespaceListerExpansion
}

// upstreamGroupNamespaceLister implements the UpstreamGroupNamespaceLister
// interface.
type upstreamGroupNamespaceLister struct {
	listers.ResourceIndexer[*v1.UpstreamGroup]
}

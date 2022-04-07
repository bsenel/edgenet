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

// Code generated by informer-gen. DO NOT EDIT.

package externalversions

import (
	"fmt"

	v1alpha1 "github.com/EdgeNet-project/edgenet/pkg/apis/apps/v1alpha1"
	corev1alpha1 "github.com/EdgeNet-project/edgenet/pkg/apis/core/v1alpha1"
	networkingv1alpha1 "github.com/EdgeNet-project/edgenet/pkg/apis/networking/v1alpha1"
	registrationv1alpha1 "github.com/EdgeNet-project/edgenet/pkg/apis/registration/v1alpha1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=apps.edgenet.io, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithResource("selectivedeployments"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1alpha1().SelectiveDeployments().Informer()}, nil

		// Group=core.edgenet.io, Version=v1alpha1
	case corev1alpha1.SchemeGroupVersion.WithResource("nodecontributions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1alpha1().NodeContributions().Informer()}, nil
	case corev1alpha1.SchemeGroupVersion.WithResource("slices"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1alpha1().Slices().Informer()}, nil
	case corev1alpha1.SchemeGroupVersion.WithResource("sliceclaims"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1alpha1().SliceClaims().Informer()}, nil
	case corev1alpha1.SchemeGroupVersion.WithResource("subnamespaces"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1alpha1().SubNamespaces().Informer()}, nil
	case corev1alpha1.SchemeGroupVersion.WithResource("tenants"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1alpha1().Tenants().Informer()}, nil
	case corev1alpha1.SchemeGroupVersion.WithResource("tenantresourcequotas"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1alpha1().TenantResourceQuotas().Informer()}, nil

		// Group=networking.edgenet.io, Version=v1alpha1
	case networkingv1alpha1.SchemeGroupVersion.WithResource("vpnpeers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Networking().V1alpha1().VPNPeers().Informer()}, nil

		// Group=registration.edgenet.io, Version=v1alpha1
	case registrationv1alpha1.SchemeGroupVersion.WithResource("clusterrolerequests"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Registration().V1alpha1().ClusterRoleRequests().Informer()}, nil
	case registrationv1alpha1.SchemeGroupVersion.WithResource("rolerequests"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Registration().V1alpha1().RoleRequests().Informer()}, nil
	case registrationv1alpha1.SchemeGroupVersion.WithResource("tenantrequests"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Registration().V1alpha1().TenantRequests().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}

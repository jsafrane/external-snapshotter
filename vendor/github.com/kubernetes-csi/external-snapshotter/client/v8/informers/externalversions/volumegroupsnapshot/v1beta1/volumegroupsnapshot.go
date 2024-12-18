/*
Copyright 2024 The Kubernetes Authors.

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

package v1beta1

import (
	"context"
	time "time"

	volumegroupsnapshotv1beta1 "github.com/kubernetes-csi/external-snapshotter/client/v8/apis/volumegroupsnapshot/v1beta1"
	versioned "github.com/kubernetes-csi/external-snapshotter/client/v8/clientset/versioned"
	internalinterfaces "github.com/kubernetes-csi/external-snapshotter/client/v8/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/kubernetes-csi/external-snapshotter/client/v8/listers/volumegroupsnapshot/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// VolumeGroupSnapshotInformer provides access to a shared informer and lister for
// VolumeGroupSnapshots.
type VolumeGroupSnapshotInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.VolumeGroupSnapshotLister
}

type volumeGroupSnapshotInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewVolumeGroupSnapshotInformer constructs a new informer for VolumeGroupSnapshot type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewVolumeGroupSnapshotInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredVolumeGroupSnapshotInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredVolumeGroupSnapshotInformer constructs a new informer for VolumeGroupSnapshot type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredVolumeGroupSnapshotInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.GroupsnapshotV1beta1().VolumeGroupSnapshots(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.GroupsnapshotV1beta1().VolumeGroupSnapshots(namespace).Watch(context.TODO(), options)
			},
		},
		&volumegroupsnapshotv1beta1.VolumeGroupSnapshot{},
		resyncPeriod,
		indexers,
	)
}

func (f *volumeGroupSnapshotInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredVolumeGroupSnapshotInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *volumeGroupSnapshotInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&volumegroupsnapshotv1beta1.VolumeGroupSnapshot{}, f.defaultInformer)
}

func (f *volumeGroupSnapshotInformer) Lister() v1beta1.VolumeGroupSnapshotLister {
	return v1beta1.NewVolumeGroupSnapshotLister(f.Informer().GetIndexer())
}

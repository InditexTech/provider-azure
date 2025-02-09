// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SpringCloudStorageInitParameters struct {

	// The access key of the Azure Storage Account.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Account
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("primary_access_key",true)
	StorageAccountKey *string `json:"storageAccountKey,omitempty" tf:"storage_account_key,omitempty"`

	// Reference to a Account in storage to populate storageAccountKey.
	// +kubebuilder:validation:Optional
	StorageAccountKeyRef *v1.Reference `json:"storageAccountKeyRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate storageAccountKey.
	// +kubebuilder:validation:Optional
	StorageAccountKeySelector *v1.Selector `json:"storageAccountKeySelector,omitempty" tf:"-"`

	// The account name of the Azure Storage Account.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Account
	StorageAccountName *string `json:"storageAccountName,omitempty" tf:"storage_account_name,omitempty"`

	// Reference to a Account in storage to populate storageAccountName.
	// +kubebuilder:validation:Optional
	StorageAccountNameRef *v1.Reference `json:"storageAccountNameRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate storageAccountName.
	// +kubebuilder:validation:Optional
	StorageAccountNameSelector *v1.Selector `json:"storageAccountNameSelector,omitempty" tf:"-"`
}

type SpringCloudStorageObservation struct {

	// The ID of the Spring Cloud Storage.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the Spring Cloud Service where the Spring Cloud Storage should exist. Changing this forces a new Spring Cloud Storage to be created.
	SpringCloudServiceID *string `json:"springCloudServiceId,omitempty" tf:"spring_cloud_service_id,omitempty"`

	// The access key of the Azure Storage Account.
	StorageAccountKey *string `json:"storageAccountKey,omitempty" tf:"storage_account_key,omitempty"`

	// The account name of the Azure Storage Account.
	StorageAccountName *string `json:"storageAccountName,omitempty" tf:"storage_account_name,omitempty"`
}

type SpringCloudStorageParameters struct {

	// The ID of the Spring Cloud Service where the Spring Cloud Storage should exist. Changing this forces a new Spring Cloud Storage to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/appplatform/v1beta1.SpringCloudService
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SpringCloudServiceID *string `json:"springCloudServiceId,omitempty" tf:"spring_cloud_service_id,omitempty"`

	// Reference to a SpringCloudService in appplatform to populate springCloudServiceId.
	// +kubebuilder:validation:Optional
	SpringCloudServiceIDRef *v1.Reference `json:"springCloudServiceIdRef,omitempty" tf:"-"`

	// Selector for a SpringCloudService in appplatform to populate springCloudServiceId.
	// +kubebuilder:validation:Optional
	SpringCloudServiceIDSelector *v1.Selector `json:"springCloudServiceIdSelector,omitempty" tf:"-"`

	// The access key of the Azure Storage Account.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Account
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("primary_access_key",true)
	// +kubebuilder:validation:Optional
	StorageAccountKey *string `json:"storageAccountKey,omitempty" tf:"storage_account_key,omitempty"`

	// Reference to a Account in storage to populate storageAccountKey.
	// +kubebuilder:validation:Optional
	StorageAccountKeyRef *v1.Reference `json:"storageAccountKeyRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate storageAccountKey.
	// +kubebuilder:validation:Optional
	StorageAccountKeySelector *v1.Selector `json:"storageAccountKeySelector,omitempty" tf:"-"`

	// The account name of the Azure Storage Account.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Account
	// +kubebuilder:validation:Optional
	StorageAccountName *string `json:"storageAccountName,omitempty" tf:"storage_account_name,omitempty"`

	// Reference to a Account in storage to populate storageAccountName.
	// +kubebuilder:validation:Optional
	StorageAccountNameRef *v1.Reference `json:"storageAccountNameRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate storageAccountName.
	// +kubebuilder:validation:Optional
	StorageAccountNameSelector *v1.Selector `json:"storageAccountNameSelector,omitempty" tf:"-"`
}

// SpringCloudStorageSpec defines the desired state of SpringCloudStorage
type SpringCloudStorageSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SpringCloudStorageParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider SpringCloudStorageInitParameters `json:"initProvider,omitempty"`
}

// SpringCloudStorageStatus defines the observed state of SpringCloudStorage.
type SpringCloudStorageStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SpringCloudStorageObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SpringCloudStorage is the Schema for the SpringCloudStorages API. Manages a Spring Cloud Storage.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SpringCloudStorage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SpringCloudStorageSpec   `json:"spec"`
	Status            SpringCloudStorageStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SpringCloudStorageList contains a list of SpringCloudStorages
type SpringCloudStorageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpringCloudStorage `json:"items"`
}

// Repository type metadata.
var (
	SpringCloudStorage_Kind             = "SpringCloudStorage"
	SpringCloudStorage_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SpringCloudStorage_Kind}.String()
	SpringCloudStorage_KindAPIVersion   = SpringCloudStorage_Kind + "." + CRDGroupVersion.String()
	SpringCloudStorage_GroupVersionKind = CRDGroupVersion.WithKind(SpringCloudStorage_Kind)
)

func init() {
	SchemeBuilder.Register(&SpringCloudStorage{}, &SpringCloudStorageList{})
}

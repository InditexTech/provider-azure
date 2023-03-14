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

type AutomaticUpdateObservation struct {
}

type AutomaticUpdateParameters struct {

	// The automation account ID which holds the automatic update runbook and authenticates to Azure resources.
	// +kubebuilder:validation:Optional
	AutomationAccountID *string `json:"automationAccountId,omitempty" tf:"automation_account_id,omitempty"`

	// Should the Mobility service installed on Azure virtual machines be automatically updated. Defaults to false.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type SiteRecoveryProtectionContainerMappingObservation struct {

	// The ID of the Site Recovery Protection Container Mapping.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SiteRecoveryProtectionContainerMappingParameters struct {

	// a automatic_update block defined as below.
	// +kubebuilder:validation:Optional
	AutomaticUpdate []AutomaticUpdateParameters `json:"automaticUpdate,omitempty" tf:"automatic_update,omitempty"`

	// Name of fabric that should contains the protection container to map. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/recoveryservices/v1beta1.SiteRecoveryFabric
	// +kubebuilder:validation:Optional
	RecoveryFabricName *string `json:"recoveryFabricName,omitempty" tf:"recovery_fabric_name,omitempty"`

	// Reference to a SiteRecoveryFabric in recoveryservices to populate recoveryFabricName.
	// +kubebuilder:validation:Optional
	RecoveryFabricNameRef *v1.Reference `json:"recoveryFabricNameRef,omitempty" tf:"-"`

	// Selector for a SiteRecoveryFabric in recoveryservices to populate recoveryFabricName.
	// +kubebuilder:validation:Optional
	RecoveryFabricNameSelector *v1.Selector `json:"recoveryFabricNameSelector,omitempty" tf:"-"`

	// Id of the policy to use for this mapping. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/recoveryservices/v1beta1.SiteRecoveryReplicationPolicy
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	RecoveryReplicationPolicyID *string `json:"recoveryReplicationPolicyId,omitempty" tf:"recovery_replication_policy_id,omitempty"`

	// Reference to a SiteRecoveryReplicationPolicy in recoveryservices to populate recoveryReplicationPolicyId.
	// +kubebuilder:validation:Optional
	RecoveryReplicationPolicyIDRef *v1.Reference `json:"recoveryReplicationPolicyIdRef,omitempty" tf:"-"`

	// Selector for a SiteRecoveryReplicationPolicy in recoveryservices to populate recoveryReplicationPolicyId.
	// +kubebuilder:validation:Optional
	RecoveryReplicationPolicyIDSelector *v1.Selector `json:"recoveryReplicationPolicyIdSelector,omitempty" tf:"-"`

	// Name of the source protection container to map. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/recoveryservices/v1beta1.SiteRecoveryProtectionContainer
	// +kubebuilder:validation:Optional
	RecoverySourceProtectionContainerName *string `json:"recoverySourceProtectionContainerName,omitempty" tf:"recovery_source_protection_container_name,omitempty"`

	// Reference to a SiteRecoveryProtectionContainer in recoveryservices to populate recoverySourceProtectionContainerName.
	// +kubebuilder:validation:Optional
	RecoverySourceProtectionContainerNameRef *v1.Reference `json:"recoverySourceProtectionContainerNameRef,omitempty" tf:"-"`

	// Selector for a SiteRecoveryProtectionContainer in recoveryservices to populate recoverySourceProtectionContainerName.
	// +kubebuilder:validation:Optional
	RecoverySourceProtectionContainerNameSelector *v1.Selector `json:"recoverySourceProtectionContainerNameSelector,omitempty" tf:"-"`

	// Id of target protection container to map to. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/recoveryservices/v1beta1.SiteRecoveryProtectionContainer
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	RecoveryTargetProtectionContainerID *string `json:"recoveryTargetProtectionContainerId,omitempty" tf:"recovery_target_protection_container_id,omitempty"`

	// Reference to a SiteRecoveryProtectionContainer in recoveryservices to populate recoveryTargetProtectionContainerId.
	// +kubebuilder:validation:Optional
	RecoveryTargetProtectionContainerIDRef *v1.Reference `json:"recoveryTargetProtectionContainerIdRef,omitempty" tf:"-"`

	// Selector for a SiteRecoveryProtectionContainer in recoveryservices to populate recoveryTargetProtectionContainerId.
	// +kubebuilder:validation:Optional
	RecoveryTargetProtectionContainerIDSelector *v1.Selector `json:"recoveryTargetProtectionContainerIdSelector,omitempty" tf:"-"`

	// The name of the vault that should be updated. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/recoveryservices/v1beta1.Vault
	// +kubebuilder:validation:Optional
	RecoveryVaultName *string `json:"recoveryVaultName,omitempty" tf:"recovery_vault_name,omitempty"`

	// Reference to a Vault in recoveryservices to populate recoveryVaultName.
	// +kubebuilder:validation:Optional
	RecoveryVaultNameRef *v1.Reference `json:"recoveryVaultNameRef,omitempty" tf:"-"`

	// Selector for a Vault in recoveryservices to populate recoveryVaultName.
	// +kubebuilder:validation:Optional
	RecoveryVaultNameSelector *v1.Selector `json:"recoveryVaultNameSelector,omitempty" tf:"-"`

	// Name of the resource group where the vault that should be updated is located. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`
}

// SiteRecoveryProtectionContainerMappingSpec defines the desired state of SiteRecoveryProtectionContainerMapping
type SiteRecoveryProtectionContainerMappingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SiteRecoveryProtectionContainerMappingParameters `json:"forProvider"`
}

// SiteRecoveryProtectionContainerMappingStatus defines the observed state of SiteRecoveryProtectionContainerMapping.
type SiteRecoveryProtectionContainerMappingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SiteRecoveryProtectionContainerMappingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SiteRecoveryProtectionContainerMapping is the Schema for the SiteRecoveryProtectionContainerMappings API. Manages a Site Recovery protection container mapping on Azure.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SiteRecoveryProtectionContainerMapping struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SiteRecoveryProtectionContainerMappingSpec   `json:"spec"`
	Status            SiteRecoveryProtectionContainerMappingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SiteRecoveryProtectionContainerMappingList contains a list of SiteRecoveryProtectionContainerMappings
type SiteRecoveryProtectionContainerMappingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SiteRecoveryProtectionContainerMapping `json:"items"`
}

// Repository type metadata.
var (
	SiteRecoveryProtectionContainerMapping_Kind             = "SiteRecoveryProtectionContainerMapping"
	SiteRecoveryProtectionContainerMapping_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SiteRecoveryProtectionContainerMapping_Kind}.String()
	SiteRecoveryProtectionContainerMapping_KindAPIVersion   = SiteRecoveryProtectionContainerMapping_Kind + "." + CRDGroupVersion.String()
	SiteRecoveryProtectionContainerMapping_GroupVersionKind = CRDGroupVersion.WithKind(SiteRecoveryProtectionContainerMapping_Kind)
)

func init() {
	SchemeBuilder.Register(&SiteRecoveryProtectionContainerMapping{}, &SiteRecoveryProtectionContainerMappingList{})
}

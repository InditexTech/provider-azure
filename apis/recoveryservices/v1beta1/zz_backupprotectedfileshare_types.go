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

type BackupProtectedFileShareObservation struct {

	// Specifies the ID of the backup policy to use. The policy must be an Azure File Share backup policy. Other types are not supported.
	BackupPolicyID *string `json:"backupPolicyId,omitempty" tf:"backup_policy_id,omitempty"`

	// The ID of the Backup File Share.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
	RecoveryVaultName *string `json:"recoveryVaultName,omitempty" tf:"recovery_vault_name,omitempty"`

	// The name of the resource group in which to create the Azure Backup Protected File Share. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Specifies the name of the file share to backup. Changing this forces a new resource to be created.
	SourceFileShareName *string `json:"sourceFileShareName,omitempty" tf:"source_file_share_name,omitempty"`

	// Specifies the ID of the storage account of the file share to backup. Changing this forces a new resource to be created.
	SourceStorageAccountID *string `json:"sourceStorageAccountId,omitempty" tf:"source_storage_account_id,omitempty"`
}

type BackupProtectedFileShareParameters struct {

	// Specifies the ID of the backup policy to use. The policy must be an Azure File Share backup policy. Other types are not supported.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/recoveryservices/v1beta1.BackupPolicyFileShare
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	BackupPolicyID *string `json:"backupPolicyId,omitempty" tf:"backup_policy_id,omitempty"`

	// Reference to a BackupPolicyFileShare in recoveryservices to populate backupPolicyId.
	// +kubebuilder:validation:Optional
	BackupPolicyIDRef *v1.Reference `json:"backupPolicyIdRef,omitempty" tf:"-"`

	// Selector for a BackupPolicyFileShare in recoveryservices to populate backupPolicyId.
	// +kubebuilder:validation:Optional
	BackupPolicyIDSelector *v1.Selector `json:"backupPolicyIdSelector,omitempty" tf:"-"`

	// Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/recoveryservices/v1beta1.Vault
	// +kubebuilder:validation:Optional
	RecoveryVaultName *string `json:"recoveryVaultName,omitempty" tf:"recovery_vault_name,omitempty"`

	// Reference to a Vault in recoveryservices to populate recoveryVaultName.
	// +kubebuilder:validation:Optional
	RecoveryVaultNameRef *v1.Reference `json:"recoveryVaultNameRef,omitempty" tf:"-"`

	// Selector for a Vault in recoveryservices to populate recoveryVaultName.
	// +kubebuilder:validation:Optional
	RecoveryVaultNameSelector *v1.Selector `json:"recoveryVaultNameSelector,omitempty" tf:"-"`

	// The name of the resource group in which to create the Azure Backup Protected File Share. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Specifies the name of the file share to backup. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Share
	// +kubebuilder:validation:Optional
	SourceFileShareName *string `json:"sourceFileShareName,omitempty" tf:"source_file_share_name,omitempty"`

	// Reference to a Share in storage to populate sourceFileShareName.
	// +kubebuilder:validation:Optional
	SourceFileShareNameRef *v1.Reference `json:"sourceFileShareNameRef,omitempty" tf:"-"`

	// Selector for a Share in storage to populate sourceFileShareName.
	// +kubebuilder:validation:Optional
	SourceFileShareNameSelector *v1.Selector `json:"sourceFileShareNameSelector,omitempty" tf:"-"`

	// Specifies the ID of the storage account of the file share to backup. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/recoveryservices/v1beta1.BackupContainerStorageAccount
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("storage_account_id",false)
	// +kubebuilder:validation:Optional
	SourceStorageAccountID *string `json:"sourceStorageAccountId,omitempty" tf:"source_storage_account_id,omitempty"`

	// Reference to a BackupContainerStorageAccount in recoveryservices to populate sourceStorageAccountId.
	// +kubebuilder:validation:Optional
	SourceStorageAccountIDRef *v1.Reference `json:"sourceStorageAccountIdRef,omitempty" tf:"-"`

	// Selector for a BackupContainerStorageAccount in recoveryservices to populate sourceStorageAccountId.
	// +kubebuilder:validation:Optional
	SourceStorageAccountIDSelector *v1.Selector `json:"sourceStorageAccountIdSelector,omitempty" tf:"-"`
}

// BackupProtectedFileShareSpec defines the desired state of BackupProtectedFileShare
type BackupProtectedFileShareSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BackupProtectedFileShareParameters `json:"forProvider"`
}

// BackupProtectedFileShareStatus defines the observed state of BackupProtectedFileShare.
type BackupProtectedFileShareStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BackupProtectedFileShareObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BackupProtectedFileShare is the Schema for the BackupProtectedFileShares API. Manages an Azure Backup Protected File Share.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type BackupProtectedFileShare struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BackupProtectedFileShareSpec   `json:"spec"`
	Status            BackupProtectedFileShareStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackupProtectedFileShareList contains a list of BackupProtectedFileShares
type BackupProtectedFileShareList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupProtectedFileShare `json:"items"`
}

// Repository type metadata.
var (
	BackupProtectedFileShare_Kind             = "BackupProtectedFileShare"
	BackupProtectedFileShare_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BackupProtectedFileShare_Kind}.String()
	BackupProtectedFileShare_KindAPIVersion   = BackupProtectedFileShare_Kind + "." + CRDGroupVersion.String()
	BackupProtectedFileShare_GroupVersionKind = CRDGroupVersion.WithKind(BackupProtectedFileShare_Kind)
)

func init() {
	SchemeBuilder.Register(&BackupProtectedFileShare{}, &BackupProtectedFileShareList{})
}

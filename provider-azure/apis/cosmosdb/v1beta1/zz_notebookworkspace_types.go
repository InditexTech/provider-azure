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

type NotebookWorkspaceObservation struct {

	// The ID of the SQL Notebook Workspace.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the endpoint of Notebook server.
	ServerEndpoint *string `json:"serverEndpoint,omitempty" tf:"server_endpoint,omitempty"`
}

type NotebookWorkspaceParameters struct {

	// The name of the Cosmos DB Account to create the SQL Notebook Workspace within. Changing this forces a new SQL Notebook Workspace to be created.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/cosmosdb/v1beta1.Account
	// +kubebuilder:validation:Optional
	AccountName *string `json:"accountName,omitempty" tf:"account_name,omitempty"`

	// Reference to a Account in cosmosdb to populate accountName.
	// +kubebuilder:validation:Optional
	AccountNameRef *v1.Reference `json:"accountNameRef,omitempty" tf:"-"`

	// Selector for a Account in cosmosdb to populate accountName.
	// +kubebuilder:validation:Optional
	AccountNameSelector *v1.Selector `json:"accountNameSelector,omitempty" tf:"-"`

	// The name of the Resource Group where the SQL Notebook Workspace should exist. Changing this forces a new SQL Notebook Workspace to be created.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`
}

// NotebookWorkspaceSpec defines the desired state of NotebookWorkspace
type NotebookWorkspaceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NotebookWorkspaceParameters `json:"forProvider"`
}

// NotebookWorkspaceStatus defines the observed state of NotebookWorkspace.
type NotebookWorkspaceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NotebookWorkspaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NotebookWorkspace is the Schema for the NotebookWorkspaces API. Manages an SQL Notebook Workspace.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type NotebookWorkspace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NotebookWorkspaceSpec   `json:"spec"`
	Status            NotebookWorkspaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NotebookWorkspaceList contains a list of NotebookWorkspaces
type NotebookWorkspaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NotebookWorkspace `json:"items"`
}

// Repository type metadata.
var (
	NotebookWorkspace_Kind             = "NotebookWorkspace"
	NotebookWorkspace_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NotebookWorkspace_Kind}.String()
	NotebookWorkspace_KindAPIVersion   = NotebookWorkspace_Kind + "." + CRDGroupVersion.String()
	NotebookWorkspace_GroupVersionKind = CRDGroupVersion.WithKind(NotebookWorkspace_Kind)
)

func init() {
	SchemeBuilder.Register(&NotebookWorkspace{}, &NotebookWorkspaceList{})
}

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

type DNSCNAMERecordObservation struct {

	// The FQDN of the DNS CName Record.
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	// The DNS CName Record ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The target of the CNAME.
	Record *string `json:"record,omitempty" tf:"record,omitempty"`

	// Specifies the resource group where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The Time To Live (TTL) of the DNS record in seconds.
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The Azure resource id of the target object. Conflicts with record.
	TargetResourceID *string `json:"targetResourceId,omitempty" tf:"target_resource_id,omitempty"`

	// Specifies the DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName *string `json:"zoneName,omitempty" tf:"zone_name,omitempty"`
}

type DNSCNAMERecordParameters struct {

	// The target of the CNAME.
	// +kubebuilder:validation:Optional
	Record *string `json:"record,omitempty" tf:"record,omitempty"`

	// Specifies the resource group where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The Time To Live (TTL) of the DNS record in seconds.
	// +kubebuilder:validation:Optional
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The Azure resource id of the target object. Conflicts with record.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.DNSCNAMERecord
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	TargetResourceID *string `json:"targetResourceId,omitempty" tf:"target_resource_id,omitempty"`

	// Reference to a DNSCNAMERecord in network to populate targetResourceId.
	// +kubebuilder:validation:Optional
	TargetResourceIDRef *v1.Reference `json:"targetResourceIdRef,omitempty" tf:"-"`

	// Selector for a DNSCNAMERecord in network to populate targetResourceId.
	// +kubebuilder:validation:Optional
	TargetResourceIDSelector *v1.Selector `json:"targetResourceIdSelector,omitempty" tf:"-"`

	// Specifies the DNS Zone where the resource exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=DNSZone
	// +kubebuilder:validation:Optional
	ZoneName *string `json:"zoneName,omitempty" tf:"zone_name,omitempty"`

	// Reference to a DNSZone to populate zoneName.
	// +kubebuilder:validation:Optional
	ZoneNameRef *v1.Reference `json:"zoneNameRef,omitempty" tf:"-"`

	// Selector for a DNSZone to populate zoneName.
	// +kubebuilder:validation:Optional
	ZoneNameSelector *v1.Selector `json:"zoneNameSelector,omitempty" tf:"-"`
}

// DNSCNAMERecordSpec defines the desired state of DNSCNAMERecord
type DNSCNAMERecordSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DNSCNAMERecordParameters `json:"forProvider"`
}

// DNSCNAMERecordStatus defines the observed state of DNSCNAMERecord.
type DNSCNAMERecordStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DNSCNAMERecordObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DNSCNAMERecord is the Schema for the DNSCNAMERecords API. Manages a DNS CNAME Record.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type DNSCNAMERecord struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.ttl)",message="ttl is a required parameter"
	Spec   DNSCNAMERecordSpec   `json:"spec"`
	Status DNSCNAMERecordStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DNSCNAMERecordList contains a list of DNSCNAMERecords
type DNSCNAMERecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DNSCNAMERecord `json:"items"`
}

// Repository type metadata.
var (
	DNSCNAMERecord_Kind             = "DNSCNAMERecord"
	DNSCNAMERecord_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DNSCNAMERecord_Kind}.String()
	DNSCNAMERecord_KindAPIVersion   = DNSCNAMERecord_Kind + "." + CRDGroupVersion.String()
	DNSCNAMERecord_GroupVersionKind = CRDGroupVersion.WithKind(DNSCNAMERecord_Kind)
)

func init() {
	SchemeBuilder.Register(&DNSCNAMERecord{}, &DNSCNAMERecordList{})
}

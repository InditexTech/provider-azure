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

type FrontdoorCustomDomainAssociationObservation struct {

	// The ID of the Front Door Custom Domain that should be managed by the association resource. Changing this forces a new association resource to be created.
	CdnFrontdoorCustomDomainID *string `json:"cdnFrontdoorCustomDomainId,omitempty" tf:"cdn_frontdoor_custom_domain_id,omitempty"`

	// One or more IDs of the Front Door Route to which the Front Door Custom Domain is associated with.
	CdnFrontdoorRouteIds []*string `json:"cdnFrontdoorRouteIds,omitempty" tf:"cdn_frontdoor_route_ids,omitempty"`

	// The ID of the Front Door Custom Domain Association.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type FrontdoorCustomDomainAssociationParameters struct {

	// The ID of the Front Door Custom Domain that should be managed by the association resource. Changing this forces a new association resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/cdn/v1beta1.FrontdoorCustomDomain
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	CdnFrontdoorCustomDomainID *string `json:"cdnFrontdoorCustomDomainId,omitempty" tf:"cdn_frontdoor_custom_domain_id,omitempty"`

	// Reference to a FrontdoorCustomDomain in cdn to populate cdnFrontdoorCustomDomainId.
	// +kubebuilder:validation:Optional
	CdnFrontdoorCustomDomainIDRef *v1.Reference `json:"cdnFrontdoorCustomDomainIdRef,omitempty" tf:"-"`

	// Selector for a FrontdoorCustomDomain in cdn to populate cdnFrontdoorCustomDomainId.
	// +kubebuilder:validation:Optional
	CdnFrontdoorCustomDomainIDSelector *v1.Selector `json:"cdnFrontdoorCustomDomainIdSelector,omitempty" tf:"-"`

	// One or more IDs of the Front Door Route to which the Front Door Custom Domain is associated with.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/cdn/v1beta1.FrontdoorRoute
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	CdnFrontdoorRouteIds []*string `json:"cdnFrontdoorRouteIds,omitempty" tf:"cdn_frontdoor_route_ids,omitempty"`

	// References to FrontdoorRoute in cdn to populate cdnFrontdoorRouteIds.
	// +kubebuilder:validation:Optional
	CdnFrontdoorRouteIdsRefs []v1.Reference `json:"cdnFrontdoorRouteIdsRefs,omitempty" tf:"-"`

	// Selector for a list of FrontdoorRoute in cdn to populate cdnFrontdoorRouteIds.
	// +kubebuilder:validation:Optional
	CdnFrontdoorRouteIdsSelector *v1.Selector `json:"cdnFrontdoorRouteIdsSelector,omitempty" tf:"-"`
}

// FrontdoorCustomDomainAssociationSpec defines the desired state of FrontdoorCustomDomainAssociation
type FrontdoorCustomDomainAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FrontdoorCustomDomainAssociationParameters `json:"forProvider"`
}

// FrontdoorCustomDomainAssociationStatus defines the observed state of FrontdoorCustomDomainAssociation.
type FrontdoorCustomDomainAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FrontdoorCustomDomainAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FrontdoorCustomDomainAssociation is the Schema for the FrontdoorCustomDomainAssociations API. Manages the association between a Front Door (standard/premium) Custom Domain and one or more Front Door (standard/premium) Routes.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type FrontdoorCustomDomainAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FrontdoorCustomDomainAssociationSpec   `json:"spec"`
	Status            FrontdoorCustomDomainAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FrontdoorCustomDomainAssociationList contains a list of FrontdoorCustomDomainAssociations
type FrontdoorCustomDomainAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FrontdoorCustomDomainAssociation `json:"items"`
}

// Repository type metadata.
var (
	FrontdoorCustomDomainAssociation_Kind             = "FrontdoorCustomDomainAssociation"
	FrontdoorCustomDomainAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FrontdoorCustomDomainAssociation_Kind}.String()
	FrontdoorCustomDomainAssociation_KindAPIVersion   = FrontdoorCustomDomainAssociation_Kind + "." + CRDGroupVersion.String()
	FrontdoorCustomDomainAssociation_GroupVersionKind = CRDGroupVersion.WithKind(FrontdoorCustomDomainAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&FrontdoorCustomDomainAssociation{}, &FrontdoorCustomDomainAssociationList{})
}

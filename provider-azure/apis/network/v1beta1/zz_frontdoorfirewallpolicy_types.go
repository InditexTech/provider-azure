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

type CustomRuleObservation struct {
}

type CustomRuleParameters struct {

	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// One or more match_condition block defined below. Can support up to 10 match_condition blocks.
	// +kubebuilder:validation:Optional
	MatchCondition []MatchConditionParameters `json:"matchCondition,omitempty" tf:"match_condition,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The priority of the rule. Rules with a lower value will be evaluated before rules with a higher value. Defaults to 1.
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// The rate limit duration in minutes. Defaults to 1.
	// +kubebuilder:validation:Optional
	RateLimitDurationInMinutes *float64 `json:"rateLimitDurationInMinutes,omitempty" tf:"rate_limit_duration_in_minutes,omitempty"`

	// The rate limit threshold. Defaults to 10.
	// +kubebuilder:validation:Optional
	RateLimitThreshold *float64 `json:"rateLimitThreshold,omitempty" tf:"rate_limit_threshold,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type FrontdoorFirewallPolicyObservation struct {

	// The Frontend Endpoints associated with this Front Door Web Application Firewall policy.
	FrontendEndpointIds []*string `json:"frontendEndpointIds,omitempty" tf:"frontend_endpoint_ids,omitempty"`

	// The ID of the FrontDoor Firewall Policy.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Azure Region where this FrontDoor Firewall Policy exists.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`
}

type FrontdoorFirewallPolicyParameters struct {

	// If a custom_rule block's action type is block, this is the response body. The body must be specified in base64 encoding.
	// +kubebuilder:validation:Optional
	CustomBlockResponseBody *string `json:"customBlockResponseBody,omitempty" tf:"custom_block_response_body,omitempty"`

	// If a custom_rule block's action type is block, this is the response status code. Possible values are 200, 403, 405, 406, or 429.
	// +kubebuilder:validation:Optional
	CustomBlockResponseStatusCode *float64 `json:"customBlockResponseStatusCode,omitempty" tf:"custom_block_response_status_code,omitempty"`

	// One or more custom_rule blocks as defined below.
	// +kubebuilder:validation:Optional
	CustomRule []CustomRuleParameters `json:"customRule,omitempty" tf:"custom_rule,omitempty"`

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// One or more managed_rule blocks as defined below.
	// +kubebuilder:validation:Optional
	ManagedRule []ManagedRuleParameters `json:"managedRule,omitempty" tf:"managed_rule,omitempty"`

	// The firewall policy mode. Possible values are Detection, Prevention and defaults to Prevention.
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// If action type is redirect, this field represents redirect URL for the client.
	// +kubebuilder:validation:Optional
	RedirectURL *string `json:"redirectUrl,omitempty" tf:"redirect_url,omitempty"`

	// The name of the resource group. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A mapping of tags to assign to the Web Application Firewall Policy.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ManagedRuleExclusionObservation struct {
}

type ManagedRuleExclusionParameters struct {

	// +kubebuilder:validation:Required
	MatchVariable *string `json:"matchVariable" tf:"match_variable,omitempty"`

	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// +kubebuilder:validation:Required
	Selector *string `json:"selector" tf:"selector,omitempty"`
}

type ManagedRuleObservation struct {
}

type ManagedRuleParameters struct {

	// +kubebuilder:validation:Optional
	Exclusion []ManagedRuleExclusionParameters `json:"exclusion,omitempty" tf:"exclusion,omitempty"`

	// One or more override blocks as defined below.
	// +kubebuilder:validation:Optional
	Override []OverrideParameters `json:"override,omitempty" tf:"override,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// The version on the managed rule to use with this resource.
	// +kubebuilder:validation:Required
	Version *string `json:"version" tf:"version,omitempty"`
}

type MatchConditionObservation struct {
}

type MatchConditionParameters struct {

	// Up to 600 possible values to match. Limit is in total across all match_condition blocks and match_values arguments. String value itself can be up to 256 characters long.
	// +kubebuilder:validation:Required
	MatchValues []*string `json:"matchValues" tf:"match_values,omitempty"`

	// +kubebuilder:validation:Required
	MatchVariable *string `json:"matchVariable" tf:"match_variable,omitempty"`

	// Should the result of the condition be negated.
	// +kubebuilder:validation:Optional
	NegationCondition *bool `json:"negationCondition,omitempty" tf:"negation_condition,omitempty"`

	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// +kubebuilder:validation:Optional
	Selector *string `json:"selector,omitempty" tf:"selector,omitempty"`

	// Up to 5 transforms to apply. Possible values are Lowercase, RemoveNulls, Trim, Uppercase, URLDecode orURLEncode.
	// +kubebuilder:validation:Optional
	Transforms []*string `json:"transforms,omitempty" tf:"transforms,omitempty"`
}

type OverrideExclusionObservation struct {
}

type OverrideExclusionParameters struct {

	// +kubebuilder:validation:Required
	MatchVariable *string `json:"matchVariable" tf:"match_variable,omitempty"`

	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// +kubebuilder:validation:Required
	Selector *string `json:"selector" tf:"selector,omitempty"`
}

type OverrideObservation struct {
}

type OverrideParameters struct {

	// +kubebuilder:validation:Optional
	Exclusion []OverrideExclusionParameters `json:"exclusion,omitempty" tf:"exclusion,omitempty"`

	// One or more rule blocks as defined below. If none are specified, all of the rules in the group will be disabled.
	// +kubebuilder:validation:Optional
	Rule []OverrideRuleParameters `json:"rule,omitempty" tf:"rule,omitempty"`

	// The managed rule group to override.
	// +kubebuilder:validation:Required
	RuleGroupName *string `json:"ruleGroupName" tf:"rule_group_name,omitempty"`
}

type OverrideRuleObservation struct {
}

type OverrideRuleParameters struct {

	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	Exclusion []RuleExclusionParameters `json:"exclusion,omitempty" tf:"exclusion,omitempty"`

	// Identifier for the managed rule.
	// +kubebuilder:validation:Required
	RuleID *string `json:"ruleId" tf:"rule_id,omitempty"`
}

type RuleExclusionObservation struct {
}

type RuleExclusionParameters struct {

	// +kubebuilder:validation:Required
	MatchVariable *string `json:"matchVariable" tf:"match_variable,omitempty"`

	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// +kubebuilder:validation:Required
	Selector *string `json:"selector" tf:"selector,omitempty"`
}

// FrontdoorFirewallPolicySpec defines the desired state of FrontdoorFirewallPolicy
type FrontdoorFirewallPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FrontdoorFirewallPolicyParameters `json:"forProvider"`
}

// FrontdoorFirewallPolicyStatus defines the observed state of FrontdoorFirewallPolicy.
type FrontdoorFirewallPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FrontdoorFirewallPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FrontdoorFirewallPolicy is the Schema for the FrontdoorFirewallPolicys API. Manages an Azure Front Door Web Application Firewall Policy instance.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type FrontdoorFirewallPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FrontdoorFirewallPolicySpec   `json:"spec"`
	Status            FrontdoorFirewallPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FrontdoorFirewallPolicyList contains a list of FrontdoorFirewallPolicys
type FrontdoorFirewallPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FrontdoorFirewallPolicy `json:"items"`
}

// Repository type metadata.
var (
	FrontdoorFirewallPolicy_Kind             = "FrontdoorFirewallPolicy"
	FrontdoorFirewallPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FrontdoorFirewallPolicy_Kind}.String()
	FrontdoorFirewallPolicy_KindAPIVersion   = FrontdoorFirewallPolicy_Kind + "." + CRDGroupVersion.String()
	FrontdoorFirewallPolicy_GroupVersionKind = CRDGroupVersion.WithKind(FrontdoorFirewallPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&FrontdoorFirewallPolicy{}, &FrontdoorFirewallPolicyList{})
}

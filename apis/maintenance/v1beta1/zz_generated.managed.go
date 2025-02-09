/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this MaintenanceAssignmentDedicatedHost.
func (mg *MaintenanceAssignmentDedicatedHost) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MaintenanceAssignmentDedicatedHost.
func (mg *MaintenanceAssignmentDedicatedHost) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this MaintenanceAssignmentDedicatedHost.
func (mg *MaintenanceAssignmentDedicatedHost) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this MaintenanceAssignmentDedicatedHost.
func (mg *MaintenanceAssignmentDedicatedHost) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this MaintenanceAssignmentDedicatedHost.
func (mg *MaintenanceAssignmentDedicatedHost) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this MaintenanceAssignmentDedicatedHost.
func (mg *MaintenanceAssignmentDedicatedHost) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MaintenanceAssignmentDedicatedHost.
func (mg *MaintenanceAssignmentDedicatedHost) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MaintenanceAssignmentDedicatedHost.
func (mg *MaintenanceAssignmentDedicatedHost) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this MaintenanceAssignmentDedicatedHost.
func (mg *MaintenanceAssignmentDedicatedHost) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this MaintenanceAssignmentDedicatedHost.
func (mg *MaintenanceAssignmentDedicatedHost) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this MaintenanceAssignmentDedicatedHost.
func (mg *MaintenanceAssignmentDedicatedHost) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this MaintenanceAssignmentDedicatedHost.
func (mg *MaintenanceAssignmentDedicatedHost) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MaintenanceAssignmentVirtualMachine.
func (mg *MaintenanceAssignmentVirtualMachine) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MaintenanceAssignmentVirtualMachine.
func (mg *MaintenanceAssignmentVirtualMachine) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this MaintenanceAssignmentVirtualMachine.
func (mg *MaintenanceAssignmentVirtualMachine) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this MaintenanceAssignmentVirtualMachine.
func (mg *MaintenanceAssignmentVirtualMachine) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this MaintenanceAssignmentVirtualMachine.
func (mg *MaintenanceAssignmentVirtualMachine) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this MaintenanceAssignmentVirtualMachine.
func (mg *MaintenanceAssignmentVirtualMachine) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MaintenanceAssignmentVirtualMachine.
func (mg *MaintenanceAssignmentVirtualMachine) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MaintenanceAssignmentVirtualMachine.
func (mg *MaintenanceAssignmentVirtualMachine) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this MaintenanceAssignmentVirtualMachine.
func (mg *MaintenanceAssignmentVirtualMachine) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this MaintenanceAssignmentVirtualMachine.
func (mg *MaintenanceAssignmentVirtualMachine) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this MaintenanceAssignmentVirtualMachine.
func (mg *MaintenanceAssignmentVirtualMachine) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this MaintenanceAssignmentVirtualMachine.
func (mg *MaintenanceAssignmentVirtualMachine) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MaintenanceConfiguration.
func (mg *MaintenanceConfiguration) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MaintenanceConfiguration.
func (mg *MaintenanceConfiguration) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this MaintenanceConfiguration.
func (mg *MaintenanceConfiguration) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this MaintenanceConfiguration.
func (mg *MaintenanceConfiguration) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this MaintenanceConfiguration.
func (mg *MaintenanceConfiguration) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this MaintenanceConfiguration.
func (mg *MaintenanceConfiguration) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MaintenanceConfiguration.
func (mg *MaintenanceConfiguration) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MaintenanceConfiguration.
func (mg *MaintenanceConfiguration) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this MaintenanceConfiguration.
func (mg *MaintenanceConfiguration) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this MaintenanceConfiguration.
func (mg *MaintenanceConfiguration) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this MaintenanceConfiguration.
func (mg *MaintenanceConfiguration) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this MaintenanceConfiguration.
func (mg *MaintenanceConfiguration) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

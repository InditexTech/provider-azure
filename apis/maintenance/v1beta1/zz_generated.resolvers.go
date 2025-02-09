/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	v1beta11 "github.com/upbound/provider-azure/apis/azure/v1beta1"
	v1beta1 "github.com/upbound/provider-azure/apis/compute/v1beta1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this MaintenanceAssignmentDedicatedHost.
func (mg *MaintenanceAssignmentDedicatedHost) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DedicatedHostID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.DedicatedHostIDRef,
		Selector:     mg.Spec.ForProvider.DedicatedHostIDSelector,
		To: reference.To{
			List:    &v1beta1.DedicatedHostList{},
			Managed: &v1beta1.DedicatedHost{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DedicatedHostID")
	}
	mg.Spec.ForProvider.DedicatedHostID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DedicatedHostIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.MaintenanceConfigurationID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.MaintenanceConfigurationIDRef,
		Selector:     mg.Spec.ForProvider.MaintenanceConfigurationIDSelector,
		To: reference.To{
			List:    &MaintenanceConfigurationList{},
			Managed: &MaintenanceConfiguration{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.MaintenanceConfigurationID")
	}
	mg.Spec.ForProvider.MaintenanceConfigurationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.MaintenanceConfigurationIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DedicatedHostID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.DedicatedHostIDRef,
		Selector:     mg.Spec.InitProvider.DedicatedHostIDSelector,
		To: reference.To{
			List:    &v1beta1.DedicatedHostList{},
			Managed: &v1beta1.DedicatedHost{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.DedicatedHostID")
	}
	mg.Spec.InitProvider.DedicatedHostID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.DedicatedHostIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.MaintenanceConfigurationID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.MaintenanceConfigurationIDRef,
		Selector:     mg.Spec.InitProvider.MaintenanceConfigurationIDSelector,
		To: reference.To{
			List:    &MaintenanceConfigurationList{},
			Managed: &MaintenanceConfiguration{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.MaintenanceConfigurationID")
	}
	mg.Spec.InitProvider.MaintenanceConfigurationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.MaintenanceConfigurationIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this MaintenanceAssignmentVirtualMachine.
func (mg *MaintenanceAssignmentVirtualMachine) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.MaintenanceConfigurationID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.MaintenanceConfigurationIDRef,
		Selector:     mg.Spec.ForProvider.MaintenanceConfigurationIDSelector,
		To: reference.To{
			List:    &MaintenanceConfigurationList{},
			Managed: &MaintenanceConfiguration{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.MaintenanceConfigurationID")
	}
	mg.Spec.ForProvider.MaintenanceConfigurationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.MaintenanceConfigurationIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VirtualMachineID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.VirtualMachineIDRef,
		Selector:     mg.Spec.ForProvider.VirtualMachineIDSelector,
		To: reference.To{
			List:    &v1beta1.LinuxVirtualMachineList{},
			Managed: &v1beta1.LinuxVirtualMachine{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VirtualMachineID")
	}
	mg.Spec.ForProvider.VirtualMachineID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VirtualMachineIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.MaintenanceConfigurationID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.MaintenanceConfigurationIDRef,
		Selector:     mg.Spec.InitProvider.MaintenanceConfigurationIDSelector,
		To: reference.To{
			List:    &MaintenanceConfigurationList{},
			Managed: &MaintenanceConfiguration{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.MaintenanceConfigurationID")
	}
	mg.Spec.InitProvider.MaintenanceConfigurationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.MaintenanceConfigurationIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this MaintenanceConfiguration.
func (mg *MaintenanceConfiguration) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta11.ResourceGroupList{},
			Managed: &v1beta11.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

//go:build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleAssignment) DeepCopyInto(out *RoleAssignment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleAssignment.
func (in *RoleAssignment) DeepCopy() *RoleAssignment {
	if in == nil {
		return nil
	}
	out := new(RoleAssignment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RoleAssignment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleAssignmentInitParameters) DeepCopyInto(out *RoleAssignmentInitParameters) {
	*out = *in
	if in.AppRoleID != nil {
		in, out := &in.AppRoleID, &out.AppRoleID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleAssignmentInitParameters.
func (in *RoleAssignmentInitParameters) DeepCopy() *RoleAssignmentInitParameters {
	if in == nil {
		return nil
	}
	out := new(RoleAssignmentInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleAssignmentList) DeepCopyInto(out *RoleAssignmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RoleAssignment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleAssignmentList.
func (in *RoleAssignmentList) DeepCopy() *RoleAssignmentList {
	if in == nil {
		return nil
	}
	out := new(RoleAssignmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RoleAssignmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleAssignmentObservation) DeepCopyInto(out *RoleAssignmentObservation) {
	*out = *in
	if in.AppRoleID != nil {
		in, out := &in.AppRoleID, &out.AppRoleID
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.PrincipalDisplayName != nil {
		in, out := &in.PrincipalDisplayName, &out.PrincipalDisplayName
		*out = new(string)
		**out = **in
	}
	if in.PrincipalObjectID != nil {
		in, out := &in.PrincipalObjectID, &out.PrincipalObjectID
		*out = new(string)
		**out = **in
	}
	if in.PrincipalType != nil {
		in, out := &in.PrincipalType, &out.PrincipalType
		*out = new(string)
		**out = **in
	}
	if in.ResourceDisplayName != nil {
		in, out := &in.ResourceDisplayName, &out.ResourceDisplayName
		*out = new(string)
		**out = **in
	}
	if in.ResourceObjectID != nil {
		in, out := &in.ResourceObjectID, &out.ResourceObjectID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleAssignmentObservation.
func (in *RoleAssignmentObservation) DeepCopy() *RoleAssignmentObservation {
	if in == nil {
		return nil
	}
	out := new(RoleAssignmentObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleAssignmentParameters) DeepCopyInto(out *RoleAssignmentParameters) {
	*out = *in
	if in.AppRoleID != nil {
		in, out := &in.AppRoleID, &out.AppRoleID
		*out = new(string)
		**out = **in
	}
	if in.PrincipalObjectID != nil {
		in, out := &in.PrincipalObjectID, &out.PrincipalObjectID
		*out = new(string)
		**out = **in
	}
	if in.PrincipalObjectIDRef != nil {
		in, out := &in.PrincipalObjectIDRef, &out.PrincipalObjectIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.PrincipalObjectIDSelector != nil {
		in, out := &in.PrincipalObjectIDSelector, &out.PrincipalObjectIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceObjectID != nil {
		in, out := &in.ResourceObjectID, &out.ResourceObjectID
		*out = new(string)
		**out = **in
	}
	if in.ResourceObjectIDRef != nil {
		in, out := &in.ResourceObjectIDRef, &out.ResourceObjectIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceObjectIDSelector != nil {
		in, out := &in.ResourceObjectIDSelector, &out.ResourceObjectIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleAssignmentParameters.
func (in *RoleAssignmentParameters) DeepCopy() *RoleAssignmentParameters {
	if in == nil {
		return nil
	}
	out := new(RoleAssignmentParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleAssignmentSpec) DeepCopyInto(out *RoleAssignmentSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleAssignmentSpec.
func (in *RoleAssignmentSpec) DeepCopy() *RoleAssignmentSpec {
	if in == nil {
		return nil
	}
	out := new(RoleAssignmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleAssignmentStatus) DeepCopyInto(out *RoleAssignmentStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleAssignmentStatus.
func (in *RoleAssignmentStatus) DeepCopy() *RoleAssignmentStatus {
	if in == nil {
		return nil
	}
	out := new(RoleAssignmentStatus)
	in.DeepCopyInto(out)
	return out
}

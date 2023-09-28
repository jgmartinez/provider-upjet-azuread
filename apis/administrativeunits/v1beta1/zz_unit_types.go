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

type UnitInitParameters struct {

	// The description of the administrative unit.
	// The description for the administrative unit
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The display name of the administrative unit.
	// The display name for the administrative unit
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Whether the administrative unit and its members are hidden or publicly viewable in the directory.
	// Whether the administrative unit and its members are hidden or publicly viewable in the directory
	HiddenMembershipEnabled *bool `json:"hiddenMembershipEnabled,omitempty" tf:"hidden_membership_enabled,omitempty"`

	// A set of object IDs of members who should be present in this administrative unit. Supported object types are Users or Groups.
	// A set of object IDs of members who should be present in this administrative unit. Supported object types are Users or Groups
	Members []*string `json:"members,omitempty" tf:"members,omitempty"`

	// If `true`, will return an error if an existing administrative unit is found with the same name
	PreventDuplicateNames *bool `json:"preventDuplicateNames,omitempty" tf:"prevent_duplicate_names,omitempty"`
}

type UnitObservation struct {

	// The description of the administrative unit.
	// The description for the administrative unit
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The display name of the administrative unit.
	// The display name for the administrative unit
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Whether the administrative unit and its members are hidden or publicly viewable in the directory.
	// Whether the administrative unit and its members are hidden or publicly viewable in the directory
	HiddenMembershipEnabled *bool `json:"hiddenMembershipEnabled,omitempty" tf:"hidden_membership_enabled,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A set of object IDs of members who should be present in this administrative unit. Supported object types are Users or Groups.
	// A set of object IDs of members who should be present in this administrative unit. Supported object types are Users or Groups
	Members []*string `json:"members,omitempty" tf:"members,omitempty"`

	// The object ID of the administrative unit.
	// The object ID of the administrative unit
	ObjectID *string `json:"objectId,omitempty" tf:"object_id,omitempty"`

	// If `true`, will return an error if an existing administrative unit is found with the same name
	PreventDuplicateNames *bool `json:"preventDuplicateNames,omitempty" tf:"prevent_duplicate_names,omitempty"`
}

type UnitParameters struct {

	// The description of the administrative unit.
	// The description for the administrative unit
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The display name of the administrative unit.
	// The display name for the administrative unit
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Whether the administrative unit and its members are hidden or publicly viewable in the directory.
	// Whether the administrative unit and its members are hidden or publicly viewable in the directory
	// +kubebuilder:validation:Optional
	HiddenMembershipEnabled *bool `json:"hiddenMembershipEnabled,omitempty" tf:"hidden_membership_enabled,omitempty"`

	// A set of object IDs of members who should be present in this administrative unit. Supported object types are Users or Groups.
	// A set of object IDs of members who should be present in this administrative unit. Supported object types are Users or Groups
	// +kubebuilder:validation:Optional
	Members []*string `json:"members,omitempty" tf:"members,omitempty"`

	// If `true`, will return an error if an existing administrative unit is found with the same name
	// +kubebuilder:validation:Optional
	PreventDuplicateNames *bool `json:"preventDuplicateNames,omitempty" tf:"prevent_duplicate_names,omitempty"`
}

// UnitSpec defines the desired state of Unit
type UnitSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UnitParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider UnitInitParameters `json:"initProvider,omitempty"`
}

// UnitStatus defines the observed state of Unit.
type UnitStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UnitObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Unit is the Schema for the Units API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azuread}
type Unit struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.displayName) || (has(self.initProvider) && has(self.initProvider.displayName))",message="spec.forProvider.displayName is a required parameter"
	Spec   UnitSpec   `json:"spec"`
	Status UnitStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UnitList contains a list of Units
type UnitList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Unit `json:"items"`
}

// Repository type metadata.
var (
	Unit_Kind             = "Unit"
	Unit_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Unit_Kind}.String()
	Unit_KindAPIVersion   = Unit_Kind + "." + CRDGroupVersion.String()
	Unit_GroupVersionKind = CRDGroupVersion.WithKind(Unit_Kind)
)

func init() {
	SchemeBuilder.Register(&Unit{}, &UnitList{})
}

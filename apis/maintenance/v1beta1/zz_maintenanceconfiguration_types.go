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

type InstallPatchesInitParameters struct {

	// A linux block as defined above. This property only applies when scope is set to InGuestPatch
	Linux []LinuxInitParameters `json:"linux,omitempty" tf:"linux,omitempty"`

	// Possible reboot preference as defined by the user based on which it would be decided to reboot the machine or not after the patch operation is completed. Possible values are Always, IfRequired and Never. This property only applies when scope is set to InGuestPatch.
	Reboot *string `json:"reboot,omitempty" tf:"reboot,omitempty"`

	// A windows block as defined above. This property only applies when scope is set to InGuestPatch
	Windows []WindowsInitParameters `json:"windows,omitempty" tf:"windows,omitempty"`
}

type InstallPatchesObservation struct {

	// A linux block as defined above. This property only applies when scope is set to InGuestPatch
	Linux []LinuxObservation `json:"linux,omitempty" tf:"linux,omitempty"`

	// Possible reboot preference as defined by the user based on which it would be decided to reboot the machine or not after the patch operation is completed. Possible values are Always, IfRequired and Never. This property only applies when scope is set to InGuestPatch.
	Reboot *string `json:"reboot,omitempty" tf:"reboot,omitempty"`

	// A windows block as defined above. This property only applies when scope is set to InGuestPatch
	Windows []WindowsObservation `json:"windows,omitempty" tf:"windows,omitempty"`
}

type InstallPatchesParameters struct {

	// A linux block as defined above. This property only applies when scope is set to InGuestPatch
	// +kubebuilder:validation:Optional
	Linux []LinuxParameters `json:"linux,omitempty" tf:"linux,omitempty"`

	// Possible reboot preference as defined by the user based on which it would be decided to reboot the machine or not after the patch operation is completed. Possible values are Always, IfRequired and Never. This property only applies when scope is set to InGuestPatch.
	// +kubebuilder:validation:Optional
	Reboot *string `json:"reboot,omitempty" tf:"reboot,omitempty"`

	// A windows block as defined above. This property only applies when scope is set to InGuestPatch
	// +kubebuilder:validation:Optional
	Windows []WindowsParameters `json:"windows,omitempty" tf:"windows,omitempty"`
}

type LinuxInitParameters struct {

	// List of Classification category of patches to be patched. Possible values are Critical, Security, UpdateRollup, FeaturePack, ServicePack, Definition, Tools and Updates.
	ClassificationsToInclude []*string `json:"classificationsToInclude,omitempty" tf:"classifications_to_include,omitempty"`

	// List of package names to be excluded from patching.
	PackageNamesMaskToExclude []*string `json:"packageNamesMaskToExclude,omitempty" tf:"package_names_mask_to_exclude,omitempty"`

	// List of package names to be included for patching.
	PackageNamesMaskToInclude []*string `json:"packageNamesMaskToInclude,omitempty" tf:"package_names_mask_to_include,omitempty"`
}

type LinuxObservation struct {

	// List of Classification category of patches to be patched. Possible values are Critical, Security, UpdateRollup, FeaturePack, ServicePack, Definition, Tools and Updates.
	ClassificationsToInclude []*string `json:"classificationsToInclude,omitempty" tf:"classifications_to_include,omitempty"`

	// List of package names to be excluded from patching.
	PackageNamesMaskToExclude []*string `json:"packageNamesMaskToExclude,omitempty" tf:"package_names_mask_to_exclude,omitempty"`

	// List of package names to be included for patching.
	PackageNamesMaskToInclude []*string `json:"packageNamesMaskToInclude,omitempty" tf:"package_names_mask_to_include,omitempty"`
}

type LinuxParameters struct {

	// List of Classification category of patches to be patched. Possible values are Critical, Security, UpdateRollup, FeaturePack, ServicePack, Definition, Tools and Updates.
	// +kubebuilder:validation:Optional
	ClassificationsToInclude []*string `json:"classificationsToInclude,omitempty" tf:"classifications_to_include,omitempty"`

	// List of package names to be excluded from patching.
	// +kubebuilder:validation:Optional
	PackageNamesMaskToExclude []*string `json:"packageNamesMaskToExclude,omitempty" tf:"package_names_mask_to_exclude,omitempty"`

	// List of package names to be included for patching.
	// +kubebuilder:validation:Optional
	PackageNamesMaskToInclude []*string `json:"packageNamesMaskToInclude,omitempty" tf:"package_names_mask_to_include,omitempty"`
}

type MaintenanceConfigurationInitParameters struct {

	// The in guest user patch mode. Possible values are Platform or User. Must be specified when scope is InGuestPatch.
	InGuestUserPatchMode *string `json:"inGuestUserPatchMode,omitempty" tf:"in_guest_user_patch_mode,omitempty"`

	// An install_patches block as defined below.
	InstallPatches []InstallPatchesInitParameters `json:"installPatches,omitempty" tf:"install_patches,omitempty"`

	// Specified the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A mapping of properties to assign to the resource.
	Properties map[string]*string `json:"properties,omitempty" tf:"properties,omitempty"`

	// The scope of the Maintenance Configuration. Possible values are Extension, Host, InGuestPatch, OSImage, SQLDB or SQLManagedInstance.
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// A mapping of tags to assign to the resource. The key could not contain upper case letter.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The visibility of the Maintenance Configuration. The only allowable value is Custom.
	Visibility *string `json:"visibility,omitempty" tf:"visibility,omitempty"`

	// A window block as defined below.
	Window []WindowInitParameters `json:"window,omitempty" tf:"window,omitempty"`
}

type MaintenanceConfigurationObservation struct {

	// The ID of the Maintenance Configuration.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The in guest user patch mode. Possible values are Platform or User. Must be specified when scope is InGuestPatch.
	InGuestUserPatchMode *string `json:"inGuestUserPatchMode,omitempty" tf:"in_guest_user_patch_mode,omitempty"`

	// An install_patches block as defined below.
	InstallPatches []InstallPatchesObservation `json:"installPatches,omitempty" tf:"install_patches,omitempty"`

	// Specified the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A mapping of properties to assign to the resource.
	Properties map[string]*string `json:"properties,omitempty" tf:"properties,omitempty"`

	// The name of the Resource Group where the Maintenance Configuration should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The scope of the Maintenance Configuration. Possible values are Extension, Host, InGuestPatch, OSImage, SQLDB or SQLManagedInstance.
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// A mapping of tags to assign to the resource. The key could not contain upper case letter.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The visibility of the Maintenance Configuration. The only allowable value is Custom.
	Visibility *string `json:"visibility,omitempty" tf:"visibility,omitempty"`

	// A window block as defined below.
	Window []WindowObservation `json:"window,omitempty" tf:"window,omitempty"`
}

type MaintenanceConfigurationParameters struct {

	// The in guest user patch mode. Possible values are Platform or User. Must be specified when scope is InGuestPatch.
	// +kubebuilder:validation:Optional
	InGuestUserPatchMode *string `json:"inGuestUserPatchMode,omitempty" tf:"in_guest_user_patch_mode,omitempty"`

	// An install_patches block as defined below.
	// +kubebuilder:validation:Optional
	InstallPatches []InstallPatchesParameters `json:"installPatches,omitempty" tf:"install_patches,omitempty"`

	// Specified the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A mapping of properties to assign to the resource.
	// +kubebuilder:validation:Optional
	Properties map[string]*string `json:"properties,omitempty" tf:"properties,omitempty"`

	// The name of the Resource Group where the Maintenance Configuration should exist. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The scope of the Maintenance Configuration. Possible values are Extension, Host, InGuestPatch, OSImage, SQLDB or SQLManagedInstance.
	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// A mapping of tags to assign to the resource. The key could not contain upper case letter.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The visibility of the Maintenance Configuration. The only allowable value is Custom.
	// +kubebuilder:validation:Optional
	Visibility *string `json:"visibility,omitempty" tf:"visibility,omitempty"`

	// A window block as defined below.
	// +kubebuilder:validation:Optional
	Window []WindowParameters `json:"window,omitempty" tf:"window,omitempty"`
}

type WindowInitParameters struct {

	// The duration of the maintenance window in HH:mm format.
	Duration *string `json:"duration,omitempty" tf:"duration,omitempty"`

	// Effective expiration date of the maintenance window in YYYY-MM-DD hh:mm format.
	ExpirationDateTime *string `json:"expirationDateTime,omitempty" tf:"expiration_date_time,omitempty"`

	// The rate at which a maintenance window is expected to recur. The rate can be expressed as daily, weekly, or monthly schedules.
	RecurEvery *string `json:"recurEvery,omitempty" tf:"recur_every,omitempty"`

	// Effective start date of the maintenance window in YYYY-MM-DD hh:mm format.
	StartDateTime *string `json:"startDateTime,omitempty" tf:"start_date_time,omitempty"`

	// The time zone for the maintenance window. A list of timezones can be obtained by executing [System.TimeZoneInfo]::GetSystemTimeZones() in PowerShell.
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`
}

type WindowObservation struct {

	// The duration of the maintenance window in HH:mm format.
	Duration *string `json:"duration,omitempty" tf:"duration,omitempty"`

	// Effective expiration date of the maintenance window in YYYY-MM-DD hh:mm format.
	ExpirationDateTime *string `json:"expirationDateTime,omitempty" tf:"expiration_date_time,omitempty"`

	// The rate at which a maintenance window is expected to recur. The rate can be expressed as daily, weekly, or monthly schedules.
	RecurEvery *string `json:"recurEvery,omitempty" tf:"recur_every,omitempty"`

	// Effective start date of the maintenance window in YYYY-MM-DD hh:mm format.
	StartDateTime *string `json:"startDateTime,omitempty" tf:"start_date_time,omitempty"`

	// The time zone for the maintenance window. A list of timezones can be obtained by executing [System.TimeZoneInfo]::GetSystemTimeZones() in PowerShell.
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`
}

type WindowParameters struct {

	// The duration of the maintenance window in HH:mm format.
	// +kubebuilder:validation:Optional
	Duration *string `json:"duration,omitempty" tf:"duration,omitempty"`

	// Effective expiration date of the maintenance window in YYYY-MM-DD hh:mm format.
	// +kubebuilder:validation:Optional
	ExpirationDateTime *string `json:"expirationDateTime,omitempty" tf:"expiration_date_time,omitempty"`

	// The rate at which a maintenance window is expected to recur. The rate can be expressed as daily, weekly, or monthly schedules.
	// +kubebuilder:validation:Optional
	RecurEvery *string `json:"recurEvery,omitempty" tf:"recur_every,omitempty"`

	// Effective start date of the maintenance window in YYYY-MM-DD hh:mm format.
	// +kubebuilder:validation:Optional
	StartDateTime *string `json:"startDateTime,omitempty" tf:"start_date_time,omitempty"`

	// The time zone for the maintenance window. A list of timezones can be obtained by executing [System.TimeZoneInfo]::GetSystemTimeZones() in PowerShell.
	// +kubebuilder:validation:Optional
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`
}

type WindowsInitParameters struct {

	// List of Classification category of patches to be patched. Possible values are Critical, Security, UpdateRollup, FeaturePack, ServicePack, Definition, Tools and Updates.
	ClassificationsToInclude []*string `json:"classificationsToInclude,omitempty" tf:"classifications_to_include,omitempty"`

	// List of KB numbers to be excluded from patching.
	KbNumbersToExclude []*string `json:"kbNumbersToExclude,omitempty" tf:"kb_numbers_to_exclude,omitempty"`

	// List of KB numbers to be included for patching.
	KbNumbersToInclude []*string `json:"kbNumbersToInclude,omitempty" tf:"kb_numbers_to_include,omitempty"`
}

type WindowsObservation struct {

	// List of Classification category of patches to be patched. Possible values are Critical, Security, UpdateRollup, FeaturePack, ServicePack, Definition, Tools and Updates.
	ClassificationsToInclude []*string `json:"classificationsToInclude,omitempty" tf:"classifications_to_include,omitempty"`

	// List of KB numbers to be excluded from patching.
	KbNumbersToExclude []*string `json:"kbNumbersToExclude,omitempty" tf:"kb_numbers_to_exclude,omitempty"`

	// List of KB numbers to be included for patching.
	KbNumbersToInclude []*string `json:"kbNumbersToInclude,omitempty" tf:"kb_numbers_to_include,omitempty"`
}

type WindowsParameters struct {

	// List of Classification category of patches to be patched. Possible values are Critical, Security, UpdateRollup, FeaturePack, ServicePack, Definition, Tools and Updates.
	// +kubebuilder:validation:Optional
	ClassificationsToInclude []*string `json:"classificationsToInclude,omitempty" tf:"classifications_to_include,omitempty"`

	// List of KB numbers to be excluded from patching.
	// +kubebuilder:validation:Optional
	KbNumbersToExclude []*string `json:"kbNumbersToExclude,omitempty" tf:"kb_numbers_to_exclude,omitempty"`

	// List of KB numbers to be included for patching.
	// +kubebuilder:validation:Optional
	KbNumbersToInclude []*string `json:"kbNumbersToInclude,omitempty" tf:"kb_numbers_to_include,omitempty"`
}

// MaintenanceConfigurationSpec defines the desired state of MaintenanceConfiguration
type MaintenanceConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MaintenanceConfigurationParameters `json:"forProvider"`
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
	InitProvider MaintenanceConfigurationInitParameters `json:"initProvider,omitempty"`
}

// MaintenanceConfigurationStatus defines the observed state of MaintenanceConfiguration.
type MaintenanceConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MaintenanceConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MaintenanceConfiguration is the Schema for the MaintenanceConfigurations API. Manages a Maintenance Configuration.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type MaintenanceConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || has(self.initProvider.location)",message="location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.scope) || has(self.initProvider.scope)",message="scope is a required parameter"
	Spec   MaintenanceConfigurationSpec   `json:"spec"`
	Status MaintenanceConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MaintenanceConfigurationList contains a list of MaintenanceConfigurations
type MaintenanceConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MaintenanceConfiguration `json:"items"`
}

// Repository type metadata.
var (
	MaintenanceConfiguration_Kind             = "MaintenanceConfiguration"
	MaintenanceConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MaintenanceConfiguration_Kind}.String()
	MaintenanceConfiguration_KindAPIVersion   = MaintenanceConfiguration_Kind + "." + CRDGroupVersion.String()
	MaintenanceConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(MaintenanceConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&MaintenanceConfiguration{}, &MaintenanceConfigurationList{})
}

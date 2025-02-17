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

type CustomerManagedKeyInitParameters struct {

	// The ID of the User Assigned Identity that has access to the key.
	IdentityID *string `json:"identityId,omitempty" tf:"identity_id,omitempty"`

	// Used to specify whether enable Infrastructure Encryption (Double Encryption). Changing this forces a new resource to be created.
	InfrastructureEncryptionEnabled *bool `json:"infrastructureEncryptionEnabled,omitempty" tf:"infrastructure_encryption_enabled,omitempty"`

	// The ID of the Key Vault Key which should be used to Encrypt the data in this ServiceBus Namespace.
	KeyVaultKeyID *string `json:"keyVaultKeyId,omitempty" tf:"key_vault_key_id,omitempty"`
}

type CustomerManagedKeyObservation struct {

	// The ID of the User Assigned Identity that has access to the key.
	IdentityID *string `json:"identityId,omitempty" tf:"identity_id,omitempty"`

	// Used to specify whether enable Infrastructure Encryption (Double Encryption). Changing this forces a new resource to be created.
	InfrastructureEncryptionEnabled *bool `json:"infrastructureEncryptionEnabled,omitempty" tf:"infrastructure_encryption_enabled,omitempty"`

	// The ID of the Key Vault Key which should be used to Encrypt the data in this ServiceBus Namespace.
	KeyVaultKeyID *string `json:"keyVaultKeyId,omitempty" tf:"key_vault_key_id,omitempty"`
}

type CustomerManagedKeyParameters struct {

	// The ID of the User Assigned Identity that has access to the key.
	// +kubebuilder:validation:Optional
	IdentityID *string `json:"identityId,omitempty" tf:"identity_id,omitempty"`

	// Used to specify whether enable Infrastructure Encryption (Double Encryption). Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	InfrastructureEncryptionEnabled *bool `json:"infrastructureEncryptionEnabled,omitempty" tf:"infrastructure_encryption_enabled,omitempty"`

	// The ID of the Key Vault Key which should be used to Encrypt the data in this ServiceBus Namespace.
	// +kubebuilder:validation:Optional
	KeyVaultKeyID *string `json:"keyVaultKeyId,omitempty" tf:"key_vault_key_id,omitempty"`
}

type IdentityInitParameters struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this ServiceBus namespace.
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this ServiceBus Namespace. Possible values are SystemAssigned, UserAssigned, SystemAssigned, UserAssigned (to enable both).
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IdentityObservation struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this ServiceBus namespace.
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// The Principal ID for the Service Principal associated with the Managed Service Identity of this ServiceBus Namespace.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The Tenant ID for the Service Principal associated with the Managed Service Identity of this ServiceBus Namespace.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this ServiceBus Namespace. Possible values are SystemAssigned, UserAssigned, SystemAssigned, UserAssigned (to enable both).
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IdentityParameters struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this ServiceBus namespace.
	// +kubebuilder:validation:Optional
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this ServiceBus Namespace. Possible values are SystemAssigned, UserAssigned, SystemAssigned, UserAssigned (to enable both).
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ServiceBusNamespaceInitParameters struct {

	// Specifies the capacity. When sku is Premium, capacity can be 1, 2, 4, 8 or 16. When sku is Basic or Standard, capacity can be 0 only.
	Capacity *float64 `json:"capacity,omitempty" tf:"capacity,omitempty"`

	// An customer_managed_key block as defined below.
	CustomerManagedKey []CustomerManagedKeyInitParameters `json:"customerManagedKey,omitempty" tf:"customer_managed_key,omitempty"`

	// An identity block as defined below.
	Identity []IdentityInitParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// Whether or not SAS authentication is enabled for the Service Bus namespace. Defaults to true.
	LocalAuthEnabled *bool `json:"localAuthEnabled,omitempty" tf:"local_auth_enabled,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The minimum supported TLS version for this Service Bus Namespace. Valid values are: 1.0, 1.1 and 1.2. The current default minimum TLS version is 1.2.
	MinimumTLSVersion *string `json:"minimumTlsVersion,omitempty" tf:"minimum_tls_version,omitempty"`

	// Is public network access enabled for the Service Bus Namespace? Defaults to true.
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// Defines which tier to use. Options are Basic, Standard or Premium. Please note that setting this field to Premium will force the creation of a new resource.
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Whether or not this resource is zone redundant. sku needs to be Premium. Changing this forces a new resource to be created.
	ZoneRedundant *bool `json:"zoneRedundant,omitempty" tf:"zone_redundant,omitempty"`
}

type ServiceBusNamespaceObservation struct {

	// Specifies the capacity. When sku is Premium, capacity can be 1, 2, 4, 8 or 16. When sku is Basic or Standard, capacity can be 0 only.
	Capacity *float64 `json:"capacity,omitempty" tf:"capacity,omitempty"`

	// An customer_managed_key block as defined below.
	CustomerManagedKey []CustomerManagedKeyObservation `json:"customerManagedKey,omitempty" tf:"customer_managed_key,omitempty"`

	// The URL to access the ServiceBus Namespace.
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// The ServiceBus Namespace ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An identity block as defined below.
	Identity []IdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`

	// Whether or not SAS authentication is enabled for the Service Bus namespace. Defaults to true.
	LocalAuthEnabled *bool `json:"localAuthEnabled,omitempty" tf:"local_auth_enabled,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The minimum supported TLS version for this Service Bus Namespace. Valid values are: 1.0, 1.1 and 1.2. The current default minimum TLS version is 1.2.
	MinimumTLSVersion *string `json:"minimumTlsVersion,omitempty" tf:"minimum_tls_version,omitempty"`

	// Is public network access enabled for the Service Bus Namespace? Defaults to true.
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// The name of the resource group in which to Changing this forces a new resource to be created.
	// create the namespace.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Defines which tier to use. Options are Basic, Standard or Premium. Please note that setting this field to Premium will force the creation of a new resource.
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Whether or not this resource is zone redundant. sku needs to be Premium. Changing this forces a new resource to be created.
	ZoneRedundant *bool `json:"zoneRedundant,omitempty" tf:"zone_redundant,omitempty"`
}

type ServiceBusNamespaceParameters struct {

	// Specifies the capacity. When sku is Premium, capacity can be 1, 2, 4, 8 or 16. When sku is Basic or Standard, capacity can be 0 only.
	// +kubebuilder:validation:Optional
	Capacity *float64 `json:"capacity,omitempty" tf:"capacity,omitempty"`

	// An customer_managed_key block as defined below.
	// +kubebuilder:validation:Optional
	CustomerManagedKey []CustomerManagedKeyParameters `json:"customerManagedKey,omitempty" tf:"customer_managed_key,omitempty"`

	// An identity block as defined below.
	// +kubebuilder:validation:Optional
	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// Whether or not SAS authentication is enabled for the Service Bus namespace. Defaults to true.
	// +kubebuilder:validation:Optional
	LocalAuthEnabled *bool `json:"localAuthEnabled,omitempty" tf:"local_auth_enabled,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The minimum supported TLS version for this Service Bus Namespace. Valid values are: 1.0, 1.1 and 1.2. The current default minimum TLS version is 1.2.
	// +kubebuilder:validation:Optional
	MinimumTLSVersion *string `json:"minimumTlsVersion,omitempty" tf:"minimum_tls_version,omitempty"`

	// Is public network access enabled for the Service Bus Namespace? Defaults to true.
	// +kubebuilder:validation:Optional
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// The name of the resource group in which to Changing this forces a new resource to be created.
	// create the namespace.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Defines which tier to use. Options are Basic, Standard or Premium. Please note that setting this field to Premium will force the creation of a new resource.
	// +kubebuilder:validation:Optional
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Whether or not this resource is zone redundant. sku needs to be Premium. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	ZoneRedundant *bool `json:"zoneRedundant,omitempty" tf:"zone_redundant,omitempty"`
}

// ServiceBusNamespaceSpec defines the desired state of ServiceBusNamespace
type ServiceBusNamespaceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServiceBusNamespaceParameters `json:"forProvider"`
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
	InitProvider ServiceBusNamespaceInitParameters `json:"initProvider,omitempty"`
}

// ServiceBusNamespaceStatus defines the observed state of ServiceBusNamespace.
type ServiceBusNamespaceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServiceBusNamespaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceBusNamespace is the Schema for the ServiceBusNamespaces API. Manages a ServiceBus Namespace.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ServiceBusNamespace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || has(self.initProvider.location)",message="location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.sku) || has(self.initProvider.sku)",message="sku is a required parameter"
	Spec   ServiceBusNamespaceSpec   `json:"spec"`
	Status ServiceBusNamespaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceBusNamespaceList contains a list of ServiceBusNamespaces
type ServiceBusNamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceBusNamespace `json:"items"`
}

// Repository type metadata.
var (
	ServiceBusNamespace_Kind             = "ServiceBusNamespace"
	ServiceBusNamespace_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServiceBusNamespace_Kind}.String()
	ServiceBusNamespace_KindAPIVersion   = ServiceBusNamespace_Kind + "." + CRDGroupVersion.String()
	ServiceBusNamespace_GroupVersionKind = CRDGroupVersion.WithKind(ServiceBusNamespace_Kind)
)

func init() {
	SchemeBuilder.Register(&ServiceBusNamespace{}, &ServiceBusNamespaceList{})
}

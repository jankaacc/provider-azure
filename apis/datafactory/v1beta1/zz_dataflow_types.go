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

type DataFlowObservation struct {

	// The ID of the Data Factory Data Flow.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DataFlowParameters struct {

	// List of tags that can be used for describing the Data Factory Data Flow.
	// +kubebuilder:validation:Optional
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The ID of Data Factory in which to associate the Data Flow with. Changing this forces a new resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/datafactory/v1beta1.Factory
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DataFactoryID *string `json:"dataFactoryId,omitempty" tf:"data_factory_id,omitempty"`

	// Reference to a Factory in datafactory to populate dataFactoryId.
	// +kubebuilder:validation:Optional
	DataFactoryIDRef *v1.Reference `json:"dataFactoryIdRef,omitempty" tf:"-"`

	// Selector for a Factory in datafactory to populate dataFactoryId.
	// +kubebuilder:validation:Optional
	DataFactoryIDSelector *v1.Selector `json:"dataFactoryIdSelector,omitempty" tf:"-"`

	// The description for the Data Factory Data Flow.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The folder that this Data Flow is in. If not specified, the Data Flow will appear at the root level.
	// +kubebuilder:validation:Optional
	Folder *string `json:"folder,omitempty" tf:"folder,omitempty"`

	// The script for the Data Factory Data Flow.
	// +kubebuilder:validation:Optional
	Script *string `json:"script,omitempty" tf:"script,omitempty"`

	// The script lines for the Data Factory Data Flow.
	// +kubebuilder:validation:Optional
	ScriptLines []*string `json:"scriptLines,omitempty" tf:"script_lines,omitempty"`

	// One or more sink blocks as defined below.
	// +kubebuilder:validation:Required
	Sink []SinkParameters `json:"sink" tf:"sink,omitempty"`

	// One or more source blocks as defined below.
	// +kubebuilder:validation:Required
	Source []SourceParameters `json:"source" tf:"source,omitempty"`

	// One or more transformation blocks as defined below.
	// +kubebuilder:validation:Optional
	Transformation []TransformationParameters `json:"transformation,omitempty" tf:"transformation,omitempty"`
}

type DataSetObservation struct {
}

type DataSetParameters struct {

	// The name for the Data Flow transformation.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/datafactory/v1beta1.DataSetJSON
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Reference to a DataSetJSON in datafactory to populate name.
	// +kubebuilder:validation:Optional
	NameRef *v1.Reference `json:"nameRef,omitempty" tf:"-"`

	// Selector for a DataSetJSON in datafactory to populate name.
	// +kubebuilder:validation:Optional
	NameSelector *v1.Selector `json:"nameSelector,omitempty" tf:"-"`

	// A map of parameters to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type FlowletObservation struct {
}

type FlowletParameters struct {

	// Specifies the reference data flow parameters from dataset.
	// +kubebuilder:validation:Optional
	DataSetParameters *string `json:"datasetParameters,omitempty" tf:"dataset_parameters,omitempty"`

	// The name for the Data Flow transformation.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type RejectedLinkedServiceObservation struct {
}

type RejectedLinkedServiceParameters struct {

	// The name for the Data Flow transformation.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type SchemaLinkedServiceObservation struct {
}

type SchemaLinkedServiceParameters struct {

	// The name for the Data Flow transformation.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type SinkLinkedServiceObservation struct {
}

type SinkLinkedServiceParameters struct {

	// The name for the Data Flow transformation.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type SinkObservation struct {
}

type SinkParameters struct {

	// A dataset block as defined below.
	// +kubebuilder:validation:Optional
	DataSet []DataSetParameters `json:"dataset,omitempty" tf:"dataset,omitempty"`

	// The description for the Data Flow Source.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A flowlet block as defined below.
	// +kubebuilder:validation:Optional
	Flowlet []FlowletParameters `json:"flowlet,omitempty" tf:"flowlet,omitempty"`

	// A linked_service block as defined below.
	// +kubebuilder:validation:Optional
	LinkedService []SinkLinkedServiceParameters `json:"linkedService,omitempty" tf:"linked_service,omitempty"`

	// The name for the Data Flow Source.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// A rejected_linked_service block as defined below.
	// +kubebuilder:validation:Optional
	RejectedLinkedService []RejectedLinkedServiceParameters `json:"rejectedLinkedService,omitempty" tf:"rejected_linked_service,omitempty"`

	// A schema_linked_service block as defined below.
	// +kubebuilder:validation:Optional
	SchemaLinkedService []SchemaLinkedServiceParameters `json:"schemaLinkedService,omitempty" tf:"schema_linked_service,omitempty"`
}

type SourceDataSetObservation struct {
}

type SourceDataSetParameters struct {

	// The name for the Data Flow transformation.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/datafactory/v1beta1.DataSetJSON
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Reference to a DataSetJSON in datafactory to populate name.
	// +kubebuilder:validation:Optional
	NameRef *v1.Reference `json:"nameRef,omitempty" tf:"-"`

	// Selector for a DataSetJSON in datafactory to populate name.
	// +kubebuilder:validation:Optional
	NameSelector *v1.Selector `json:"nameSelector,omitempty" tf:"-"`

	// A map of parameters to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type SourceFlowletObservation struct {
}

type SourceFlowletParameters struct {

	// Specifies the reference data flow parameters from dataset.
	// +kubebuilder:validation:Optional
	DataSetParameters *string `json:"datasetParameters,omitempty" tf:"dataset_parameters,omitempty"`

	// The name for the Data Flow transformation.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type SourceLinkedServiceObservation struct {
}

type SourceLinkedServiceParameters struct {

	// The name for the Data Flow transformation.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type SourceObservation struct {
}

type SourceParameters struct {

	// A dataset block as defined below.
	// +kubebuilder:validation:Optional
	DataSet []SourceDataSetParameters `json:"dataset,omitempty" tf:"dataset,omitempty"`

	// The description for the Data Flow Source.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A flowlet block as defined below.
	// +kubebuilder:validation:Optional
	Flowlet []SourceFlowletParameters `json:"flowlet,omitempty" tf:"flowlet,omitempty"`

	// A linked_service block as defined below.
	// +kubebuilder:validation:Optional
	LinkedService []SourceLinkedServiceParameters `json:"linkedService,omitempty" tf:"linked_service,omitempty"`

	// The name for the Data Flow Source.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// A rejected_linked_service block as defined below.
	// +kubebuilder:validation:Optional
	RejectedLinkedService []SourceRejectedLinkedServiceParameters `json:"rejectedLinkedService,omitempty" tf:"rejected_linked_service,omitempty"`

	// A schema_linked_service block as defined below.
	// +kubebuilder:validation:Optional
	SchemaLinkedService []SourceSchemaLinkedServiceParameters `json:"schemaLinkedService,omitempty" tf:"schema_linked_service,omitempty"`
}

type SourceRejectedLinkedServiceObservation struct {
}

type SourceRejectedLinkedServiceParameters struct {

	// The name for the Data Flow transformation.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type SourceSchemaLinkedServiceObservation struct {
}

type SourceSchemaLinkedServiceParameters struct {

	// The name for the Data Flow transformation.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type TransformationDataSetObservation struct {
}

type TransformationDataSetParameters struct {

	// The name for the Data Flow transformation.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type TransformationFlowletObservation struct {
}

type TransformationFlowletParameters struct {

	// Specifies the reference data flow parameters from dataset.
	// +kubebuilder:validation:Optional
	DataSetParameters *string `json:"datasetParameters,omitempty" tf:"dataset_parameters,omitempty"`

	// The name for the Data Flow transformation.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type TransformationLinkedServiceObservation struct {
}

type TransformationLinkedServiceParameters struct {

	// The name for the Data Flow transformation.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type TransformationObservation struct {
}

type TransformationParameters struct {

	// A dataset block as defined below.
	// +kubebuilder:validation:Optional
	DataSet []TransformationDataSetParameters `json:"dataset,omitempty" tf:"dataset,omitempty"`

	// The description for the Data Flow transformation.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A flowlet block as defined below.
	// +kubebuilder:validation:Optional
	Flowlet []TransformationFlowletParameters `json:"flowlet,omitempty" tf:"flowlet,omitempty"`

	// A linked_service block as defined below.
	// +kubebuilder:validation:Optional
	LinkedService []TransformationLinkedServiceParameters `json:"linkedService,omitempty" tf:"linked_service,omitempty"`

	// The name for the Data Flow transformation.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

// DataFlowSpec defines the desired state of DataFlow
type DataFlowSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DataFlowParameters `json:"forProvider"`
}

// DataFlowStatus defines the observed state of DataFlow.
type DataFlowStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DataFlowObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DataFlow is the Schema for the DataFlows API. Manages a Data Flow inside an Azure Data Factory.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type DataFlow struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataFlowSpec   `json:"spec"`
	Status            DataFlowStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DataFlowList contains a list of DataFlows
type DataFlowList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataFlow `json:"items"`
}

// Repository type metadata.
var (
	DataFlow_Kind             = "DataFlow"
	DataFlow_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DataFlow_Kind}.String()
	DataFlow_KindAPIVersion   = DataFlow_Kind + "." + CRDGroupVersion.String()
	DataFlow_GroupVersionKind = CRDGroupVersion.WithKind(DataFlow_Kind)
)

func init() {
	SchemeBuilder.Register(&DataFlow{}, &DataFlowList{})
}
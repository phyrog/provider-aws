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

type RestAPIPolicyObservation struct {

	// The ID of the REST API
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RestAPIPolicyParameters struct {

	// JSON formatted policy document that controls access to the API Gateway
	// +kubebuilder:validation:Required
	Policy *string `json:"policy" tf:"policy,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The ID of the REST API.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/apigateway/v1beta1.RestAPI
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	RestAPIID *string `json:"restApiId,omitempty" tf:"rest_api_id,omitempty"`

	// Reference to a RestAPI in apigateway to populate restApiId.
	// +kubebuilder:validation:Optional
	RestAPIIDRef *v1.Reference `json:"restApiIdRef,omitempty" tf:"-"`

	// Selector for a RestAPI in apigateway to populate restApiId.
	// +kubebuilder:validation:Optional
	RestAPIIDSelector *v1.Selector `json:"restApiIdSelector,omitempty" tf:"-"`
}

// RestAPIPolicySpec defines the desired state of RestAPIPolicy
type RestAPIPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RestAPIPolicyParameters `json:"forProvider"`
}

// RestAPIPolicyStatus defines the observed state of RestAPIPolicy.
type RestAPIPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RestAPIPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RestAPIPolicy is the Schema for the RestAPIPolicys API. Provides an API Gateway REST API Policy.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type RestAPIPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RestAPIPolicySpec   `json:"spec"`
	Status            RestAPIPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RestAPIPolicyList contains a list of RestAPIPolicys
type RestAPIPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RestAPIPolicy `json:"items"`
}

// Repository type metadata.
var (
	RestAPIPolicy_Kind             = "RestAPIPolicy"
	RestAPIPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RestAPIPolicy_Kind}.String()
	RestAPIPolicy_KindAPIVersion   = RestAPIPolicy_Kind + "." + CRDGroupVersion.String()
	RestAPIPolicy_GroupVersionKind = CRDGroupVersion.WithKind(RestAPIPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&RestAPIPolicy{}, &RestAPIPolicyList{})
}

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

type CookiesConfigCookiesObservation struct {
}

type CookiesConfigCookiesParameters struct {

	// +kubebuilder:validation:Optional
	Items []*string `json:"items,omitempty" tf:"items,omitempty"`
}

type HeadersConfigHeadersObservation struct {
}

type HeadersConfigHeadersParameters struct {

	// +kubebuilder:validation:Optional
	Items []*string `json:"items,omitempty" tf:"items,omitempty"`
}

type OriginRequestPolicyCookiesConfigObservation struct {
}

type OriginRequestPolicyCookiesConfigParameters struct {

	// +kubebuilder:validation:Required
	CookieBehavior *string `json:"cookieBehavior" tf:"cookie_behavior,omitempty"`

	// +kubebuilder:validation:Optional
	Cookies []CookiesConfigCookiesParameters `json:"cookies,omitempty" tf:"cookies,omitempty"`
}

type OriginRequestPolicyHeadersConfigObservation struct {
}

type OriginRequestPolicyHeadersConfigParameters struct {

	// +kubebuilder:validation:Optional
	HeaderBehavior *string `json:"headerBehavior,omitempty" tf:"header_behavior,omitempty"`

	// +kubebuilder:validation:Optional
	Headers []HeadersConfigHeadersParameters `json:"headers,omitempty" tf:"headers,omitempty"`
}

type OriginRequestPolicyObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type OriginRequestPolicyParameters struct {

	// +kubebuilder:validation:Optional
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// +kubebuilder:validation:Required
	CookiesConfig []OriginRequestPolicyCookiesConfigParameters `json:"cookiesConfig" tf:"cookies_config,omitempty"`

	// +kubebuilder:validation:Required
	HeadersConfig []OriginRequestPolicyHeadersConfigParameters `json:"headersConfig" tf:"headers_config,omitempty"`

	// +kubebuilder:validation:Required
	QueryStringsConfig []OriginRequestPolicyQueryStringsConfigParameters `json:"queryStringsConfig" tf:"query_strings_config,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

type OriginRequestPolicyQueryStringsConfigObservation struct {
}

type OriginRequestPolicyQueryStringsConfigParameters struct {

	// +kubebuilder:validation:Required
	QueryStringBehavior *string `json:"queryStringBehavior" tf:"query_string_behavior,omitempty"`

	// +kubebuilder:validation:Optional
	QueryStrings []QueryStringsConfigQueryStringsParameters `json:"queryStrings,omitempty" tf:"query_strings,omitempty"`
}

type QueryStringsConfigQueryStringsObservation struct {
}

type QueryStringsConfigQueryStringsParameters struct {

	// +kubebuilder:validation:Optional
	Items []*string `json:"items,omitempty" tf:"items,omitempty"`
}

// OriginRequestPolicySpec defines the desired state of OriginRequestPolicy
type OriginRequestPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OriginRequestPolicyParameters `json:"forProvider"`
}

// OriginRequestPolicyStatus defines the observed state of OriginRequestPolicy.
type OriginRequestPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OriginRequestPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// OriginRequestPolicy is the Schema for the OriginRequestPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type OriginRequestPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OriginRequestPolicySpec   `json:"spec"`
	Status            OriginRequestPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OriginRequestPolicyList contains a list of OriginRequestPolicys
type OriginRequestPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OriginRequestPolicy `json:"items"`
}

// Repository type metadata.
var (
	OriginRequestPolicy_Kind             = "OriginRequestPolicy"
	OriginRequestPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OriginRequestPolicy_Kind}.String()
	OriginRequestPolicy_KindAPIVersion   = OriginRequestPolicy_Kind + "." + CRDGroupVersion.String()
	OriginRequestPolicy_GroupVersionKind = CRDGroupVersion.WithKind(OriginRequestPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&OriginRequestPolicy{}, &OriginRequestPolicyList{})
}
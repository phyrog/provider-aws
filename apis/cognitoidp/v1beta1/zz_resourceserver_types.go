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

type ResourceServerObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A list of all scopes configured for this resource server in the format identifier/scope_name.
	ScopeIdentifiers []*string `json:"scopeIdentifiers,omitempty" tf:"scope_identifiers,omitempty"`
}

type ResourceServerParameters struct {

	// An identifier for the resource server.
	// +kubebuilder:validation:Required
	Identifier *string `json:"identifier" tf:"identifier,omitempty"`

	// A name for the resource server.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// A list of Authorization Scope.
	// +kubebuilder:validation:Optional
	Scope []ScopeParameters `json:"scope,omitempty" tf:"scope,omitempty"`

	// +crossplane:generate:reference:type=UserPool
	// +kubebuilder:validation:Optional
	UserPoolID *string `json:"userPoolId,omitempty" tf:"user_pool_id,omitempty"`

	// Reference to a UserPool to populate userPoolId.
	// +kubebuilder:validation:Optional
	UserPoolIDRef *v1.Reference `json:"userPoolIdRef,omitempty" tf:"-"`

	// Selector for a UserPool to populate userPoolId.
	// +kubebuilder:validation:Optional
	UserPoolIDSelector *v1.Selector `json:"userPoolIdSelector,omitempty" tf:"-"`
}

type ScopeObservation struct {
}

type ScopeParameters struct {

	// The scope description.
	// +kubebuilder:validation:Required
	ScopeDescription *string `json:"scopeDescription" tf:"scope_description,omitempty"`

	// The scope name.
	// +kubebuilder:validation:Required
	ScopeName *string `json:"scopeName" tf:"scope_name,omitempty"`
}

// ResourceServerSpec defines the desired state of ResourceServer
type ResourceServerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ResourceServerParameters `json:"forProvider"`
}

// ResourceServerStatus defines the observed state of ResourceServer.
type ResourceServerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ResourceServerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceServer is the Schema for the ResourceServers API. Provides a Cognito Resource Server.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ResourceServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ResourceServerSpec   `json:"spec"`
	Status            ResourceServerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceServerList contains a list of ResourceServers
type ResourceServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourceServer `json:"items"`
}

// Repository type metadata.
var (
	ResourceServer_Kind             = "ResourceServer"
	ResourceServer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ResourceServer_Kind}.String()
	ResourceServer_KindAPIVersion   = ResourceServer_Kind + "." + CRDGroupVersion.String()
	ResourceServer_GroupVersionKind = CRDGroupVersion.WithKind(ResourceServer_Kind)
)

func init() {
	SchemeBuilder.Register(&ResourceServer{}, &ResourceServerList{})
}

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

type CorsObservation struct {
}

type CorsParameters struct {

	// Whether to allow cookies or other credentials in requests to the function URL. The default is false.
	// +kubebuilder:validation:Optional
	AllowCredentials *bool `json:"allowCredentials,omitempty" tf:"allow_credentials,omitempty"`

	// The HTTP headers that origins can include in requests to the function URL. For example: ["date", "keep-alive", "x-custom-header"].
	// +kubebuilder:validation:Optional
	AllowHeaders []*string `json:"allowHeaders,omitempty" tf:"allow_headers,omitempty"`

	// The HTTP methods that are allowed when calling the function URL. For example: ["GET", "POST", "DELETE"], or the wildcard character (["*"]).
	// +kubebuilder:validation:Optional
	AllowMethods []*string `json:"allowMethods,omitempty" tf:"allow_methods,omitempty"`

	// The origins that can access the function URL. You can list any number of specific origins (or the wildcard character ("*")), separated by a comma. For example: ["https://www.example.com", "http://localhost:60905"].
	// +kubebuilder:validation:Optional
	AllowOrigins []*string `json:"allowOrigins,omitempty" tf:"allow_origins,omitempty"`

	// The HTTP headers in your function response that you want to expose to origins that call the function URL.
	// +kubebuilder:validation:Optional
	ExposeHeaders []*string `json:"exposeHeaders,omitempty" tf:"expose_headers,omitempty"`

	// The maximum amount of time, in seconds, that web browsers can cache results of a preflight request. By default, this is set to 0, which means that the browser doesn't cache results. The maximum value is 86400.
	// +kubebuilder:validation:Optional
	MaxAge *float64 `json:"maxAge,omitempty" tf:"max_age,omitempty"`
}

type FunctionURLObservation struct {

	// The Amazon Resource Name (ARN) of the function.
	FunctionArn *string `json:"functionArn,omitempty" tf:"function_arn,omitempty"`

	// The HTTP URL endpoint for the function in the format https://<url_id>.lambda-url.<region>.on.aws.
	FunctionURL *string `json:"functionUrl,omitempty" tf:"function_url,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A generated ID for the endpoint.
	URLID *string `json:"urlId,omitempty" tf:"url_id,omitempty"`
}

type FunctionURLParameters struct {

	// The type of authentication that the function URL uses. Set to "AWS_IAM" to restrict access to authenticated IAM users only. Set to "NONE" to bypass IAM authentication and create a public endpoint. See the AWS documentation for more details.
	// +kubebuilder:validation:Required
	AuthorizationType *string `json:"authorizationType" tf:"authorization_type,omitempty"`

	// The cross-origin resource sharing (CORS) settings for the function URL. Documented below.
	// +kubebuilder:validation:Optional
	Cors []CorsParameters `json:"cors,omitempty" tf:"cors,omitempty"`

	// The name (or ARN) of the Lambda function.
	// +crossplane:generate:reference:type=Function
	// +kubebuilder:validation:Optional
	FunctionName *string `json:"functionName,omitempty" tf:"function_name,omitempty"`

	// Reference to a Function to populate functionName.
	// +kubebuilder:validation:Optional
	FunctionNameRef *v1.Reference `json:"functionNameRef,omitempty" tf:"-"`

	// Selector for a Function to populate functionName.
	// +kubebuilder:validation:Optional
	FunctionNameSelector *v1.Selector `json:"functionNameSelector,omitempty" tf:"-"`

	// The alias name or "$LATEST".
	// +kubebuilder:validation:Optional
	Qualifier *string `json:"qualifier,omitempty" tf:"qualifier,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// FunctionURLSpec defines the desired state of FunctionURL
type FunctionURLSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FunctionURLParameters `json:"forProvider"`
}

// FunctionURLStatus defines the observed state of FunctionURL.
type FunctionURLStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FunctionURLObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FunctionURL is the Schema for the FunctionURLs API. Provides a Lambda function URL resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type FunctionURL struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FunctionURLSpec   `json:"spec"`
	Status            FunctionURLStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FunctionURLList contains a list of FunctionURLs
type FunctionURLList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FunctionURL `json:"items"`
}

// Repository type metadata.
var (
	FunctionURL_Kind             = "FunctionURL"
	FunctionURL_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FunctionURL_Kind}.String()
	FunctionURL_KindAPIVersion   = FunctionURL_Kind + "." + CRDGroupVersion.String()
	FunctionURL_GroupVersionKind = CRDGroupVersion.WithKind(FunctionURL_Kind)
)

func init() {
	SchemeBuilder.Register(&FunctionURL{}, &FunctionURLList{})
}

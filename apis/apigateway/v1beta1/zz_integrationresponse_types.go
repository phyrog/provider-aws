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

type IntegrationResponseObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type IntegrationResponseParameters struct {

	// Specifies how to handle request payload content type conversions. Supported values are CONVERT_TO_BINARY and CONVERT_TO_TEXT. If this property is not defined, the response payload will be passed through from the integration response to the method response without modification.
	// +kubebuilder:validation:Optional
	ContentHandling *string `json:"contentHandling,omitempty" tf:"content_handling,omitempty"`

	// The HTTP method (GET, POST, PUT, DELETE, HEAD, OPTIONS, ANY)
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/apigateway/v1beta1.Method
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("http_method",false)
	// +kubebuilder:validation:Optional
	HTTPMethod *string `json:"httpMethod,omitempty" tf:"http_method,omitempty"`

	// Reference to a Method in apigateway to populate httpMethod.
	// +kubebuilder:validation:Optional
	HTTPMethodRef *v1.Reference `json:"httpMethodRef,omitempty" tf:"-"`

	// Selector for a Method in apigateway to populate httpMethod.
	// +kubebuilder:validation:Optional
	HTTPMethodSelector *v1.Selector `json:"httpMethodSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The API resource ID
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/apigateway/v1beta1.Resource
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// Reference to a Resource in apigateway to populate resourceId.
	// +kubebuilder:validation:Optional
	ResourceIDRef *v1.Reference `json:"resourceIdRef,omitempty" tf:"-"`

	// Selector for a Resource in apigateway to populate resourceId.
	// +kubebuilder:validation:Optional
	ResourceIDSelector *v1.Selector `json:"resourceIdSelector,omitempty" tf:"-"`

	// A map of response parameters that can be read from the backend response.
	// For example: response_parameters = { "method.response.header.X-Some-Header" = "integration.response.header.X-Some-Other-Header" }
	// +kubebuilder:validation:Optional
	ResponseParameters map[string]*string `json:"responseParameters,omitempty" tf:"response_parameters,omitempty"`

	// A map specifying the templates used to transform the integration response body
	// +kubebuilder:validation:Optional
	ResponseTemplates map[string]*string `json:"responseTemplates,omitempty" tf:"response_templates,omitempty"`

	// The ID of the associated REST API
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

	// Specifies the regular expression pattern used to choose
	// an integration response based on the response from the backend. Omit configuring this to make the integration the default one.
	// If the backend is an AWS Lambda function, the AWS Lambda function error header is matched.
	// For all other HTTP and AWS backends, the HTTP status code is matched.
	// +kubebuilder:validation:Optional
	SelectionPattern *string `json:"selectionPattern,omitempty" tf:"selection_pattern,omitempty"`

	// The HTTP status code
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/apigateway/v1beta1.MethodResponse
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("status_code",false)
	// +kubebuilder:validation:Optional
	StatusCode *string `json:"statusCode,omitempty" tf:"status_code,omitempty"`

	// Reference to a MethodResponse in apigateway to populate statusCode.
	// +kubebuilder:validation:Optional
	StatusCodeRef *v1.Reference `json:"statusCodeRef,omitempty" tf:"-"`

	// Selector for a MethodResponse in apigateway to populate statusCode.
	// +kubebuilder:validation:Optional
	StatusCodeSelector *v1.Selector `json:"statusCodeSelector,omitempty" tf:"-"`
}

// IntegrationResponseSpec defines the desired state of IntegrationResponse
type IntegrationResponseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IntegrationResponseParameters `json:"forProvider"`
}

// IntegrationResponseStatus defines the observed state of IntegrationResponse.
type IntegrationResponseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IntegrationResponseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IntegrationResponse is the Schema for the IntegrationResponses API. Provides an HTTP Method Integration Response for an API Gateway Resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type IntegrationResponse struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IntegrationResponseSpec   `json:"spec"`
	Status            IntegrationResponseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IntegrationResponseList contains a list of IntegrationResponses
type IntegrationResponseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IntegrationResponse `json:"items"`
}

// Repository type metadata.
var (
	IntegrationResponse_Kind             = "IntegrationResponse"
	IntegrationResponse_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IntegrationResponse_Kind}.String()
	IntegrationResponse_KindAPIVersion   = IntegrationResponse_Kind + "." + CRDGroupVersion.String()
	IntegrationResponse_GroupVersionKind = CRDGroupVersion.WithKind(IntegrationResponse_Kind)
)

func init() {
	SchemeBuilder.Register(&IntegrationResponse{}, &IntegrationResponseList{})
}

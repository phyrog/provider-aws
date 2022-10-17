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

type AllowedPublishersObservation struct {
}

type AllowedPublishersParameters struct {

	// The Amazon Resource Name (ARN) for each of the signing profiles. A signing profile defines a trusted user who can sign a code package.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/signer/v1beta1.SigningProfile
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	SigningProfileVersionArns []*string `json:"signingProfileVersionArns,omitempty" tf:"signing_profile_version_arns,omitempty"`

	// References to SigningProfile in signer to populate signingProfileVersionArns.
	// +kubebuilder:validation:Optional
	SigningProfileVersionArnsRefs []v1.Reference `json:"signingProfileVersionArnsRefs,omitempty" tf:"-"`

	// Selector for a list of SigningProfile in signer to populate signingProfileVersionArns.
	// +kubebuilder:validation:Optional
	SigningProfileVersionArnsSelector *v1.Selector `json:"signingProfileVersionArnsSelector,omitempty" tf:"-"`
}

type CodeSigningConfigObservation struct {

	// The Amazon Resource Name (ARN) of the code signing configuration.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Unique identifier for the code signing configuration.
	ConfigID *string `json:"configId,omitempty" tf:"config_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The date and time that the code signing configuration was last modified.
	LastModified *string `json:"lastModified,omitempty" tf:"last_modified,omitempty"`
}

type CodeSigningConfigParameters struct {

	// A configuration block of allowed publishers as signing profiles for this code signing configuration. Detailed below.
	// +kubebuilder:validation:Required
	AllowedPublishers []AllowedPublishersParameters `json:"allowedPublishers" tf:"allowed_publishers,omitempty"`

	// Descriptive name for this code signing configuration.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A configuration block of code signing policies that define the actions to take if the validation checks fail. Detailed below.
	// +kubebuilder:validation:Optional
	Policies []PoliciesParameters `json:"policies,omitempty" tf:"policies,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

type PoliciesObservation struct {
}

type PoliciesParameters struct {

	// Code signing configuration policy for deployment validation failure. If you set the policy to Enforce, Lambda blocks the deployment request if code-signing validation checks fail. If you set the policy to Warn, Lambda allows the deployment and creates a CloudWatch log. Valid values: Warn, Enforce. Default value: Warn.
	// +kubebuilder:validation:Required
	UntrustedArtifactOnDeployment *string `json:"untrustedArtifactOnDeployment" tf:"untrusted_artifact_on_deployment,omitempty"`
}

// CodeSigningConfigSpec defines the desired state of CodeSigningConfig
type CodeSigningConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CodeSigningConfigParameters `json:"forProvider"`
}

// CodeSigningConfigStatus defines the observed state of CodeSigningConfig.
type CodeSigningConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CodeSigningConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CodeSigningConfig is the Schema for the CodeSigningConfigs API. Provides a Lambda Code Signing Config resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type CodeSigningConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CodeSigningConfigSpec   `json:"spec"`
	Status            CodeSigningConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CodeSigningConfigList contains a list of CodeSigningConfigs
type CodeSigningConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CodeSigningConfig `json:"items"`
}

// Repository type metadata.
var (
	CodeSigningConfig_Kind             = "CodeSigningConfig"
	CodeSigningConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CodeSigningConfig_Kind}.String()
	CodeSigningConfig_KindAPIVersion   = CodeSigningConfig_Kind + "." + CRDGroupVersion.String()
	CodeSigningConfig_GroupVersionKind = CRDGroupVersion.WithKind(CodeSigningConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&CodeSigningConfig{}, &CodeSigningConfigList{})
}

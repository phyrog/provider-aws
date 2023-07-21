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

type UsageLimitObservation struct {

	// The limit amount. If time-based, this amount is in Redshift Processing Units (RPU) consumed per hour. If data-based, this amount is in terabytes (TB) of data transferred between Regions in cross-account sharing. The value must be a positive number.
	Amount *float64 `json:"amount,omitempty" tf:"amount,omitempty"`

	// Amazon Resource Name (ARN) of the Redshift Serverless Usage Limit.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The action that Amazon Redshift Serverless takes when the limit is reached. Valid values are log, emit-metric, and deactivate. The default is log.
	BreachAction *string `json:"breachAction,omitempty" tf:"breach_action,omitempty"`

	// The Redshift Usage Limit id.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The time period that the amount applies to. A weekly period begins on Sunday. Valid values are daily, weekly, and monthly. The default is monthly.
	Period *string `json:"period,omitempty" tf:"period,omitempty"`

	// The Amazon Resource Name (ARN) of the Amazon Redshift Serverless resource to create the usage limit for.
	ResourceArn *string `json:"resourceArn,omitempty" tf:"resource_arn,omitempty"`

	// The type of Amazon Redshift Serverless usage to create a usage limit for. Valid values are serverless-compute or cross-region-datasharing.
	UsageType *string `json:"usageType,omitempty" tf:"usage_type,omitempty"`
}

type UsageLimitParameters struct {

	// The limit amount. If time-based, this amount is in Redshift Processing Units (RPU) consumed per hour. If data-based, this amount is in terabytes (TB) of data transferred between Regions in cross-account sharing. The value must be a positive number.
	// +kubebuilder:validation:Optional
	Amount *float64 `json:"amount,omitempty" tf:"amount,omitempty"`

	// The action that Amazon Redshift Serverless takes when the limit is reached. Valid values are log, emit-metric, and deactivate. The default is log.
	// +kubebuilder:validation:Optional
	BreachAction *string `json:"breachAction,omitempty" tf:"breach_action,omitempty"`

	// The time period that the amount applies to. A weekly period begins on Sunday. Valid values are daily, weekly, and monthly. The default is monthly.
	// +kubebuilder:validation:Optional
	Period *string `json:"period,omitempty" tf:"period,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The Amazon Resource Name (ARN) of the Amazon Redshift Serverless resource to create the usage limit for.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/redshiftserverless/v1beta1.Workgroup
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	ResourceArn *string `json:"resourceArn,omitempty" tf:"resource_arn,omitempty"`

	// Reference to a Workgroup in redshiftserverless to populate resourceArn.
	// +kubebuilder:validation:Optional
	ResourceArnRef *v1.Reference `json:"resourceArnRef,omitempty" tf:"-"`

	// Selector for a Workgroup in redshiftserverless to populate resourceArn.
	// +kubebuilder:validation:Optional
	ResourceArnSelector *v1.Selector `json:"resourceArnSelector,omitempty" tf:"-"`

	// The type of Amazon Redshift Serverless usage to create a usage limit for. Valid values are serverless-compute or cross-region-datasharing.
	// +kubebuilder:validation:Optional
	UsageType *string `json:"usageType,omitempty" tf:"usage_type,omitempty"`
}

// UsageLimitSpec defines the desired state of UsageLimit
type UsageLimitSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UsageLimitParameters `json:"forProvider"`
}

// UsageLimitStatus defines the observed state of UsageLimit.
type UsageLimitStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UsageLimitObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// UsageLimit is the Schema for the UsageLimits API. Provides a Redshift Serverless Usage Limit resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type UsageLimit struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.amount)",message="amount is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.usageType)",message="usageType is a required parameter"
	Spec   UsageLimitSpec   `json:"spec"`
	Status UsageLimitStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UsageLimitList contains a list of UsageLimits
type UsageLimitList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UsageLimit `json:"items"`
}

// Repository type metadata.
var (
	UsageLimit_Kind             = "UsageLimit"
	UsageLimit_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: UsageLimit_Kind}.String()
	UsageLimit_KindAPIVersion   = UsageLimit_Kind + "." + CRDGroupVersion.String()
	UsageLimit_GroupVersionKind = CRDGroupVersion.WithKind(UsageLimit_Kind)
)

func init() {
	SchemeBuilder.Register(&UsageLimit{}, &UsageLimitList{})
}

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

type EBSVolumeObservation struct {

	// The volume ARN (e.g., arn:aws:ec2:us-east-1:0123456789012:volume/vol-59fcb34e).
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The volume ID (e.g., vol-59fcb34e).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type EBSVolumeParameters struct {

	// The AZ where the EBS volume will exist.
	// +kubebuilder:validation:Required
	AvailabilityZone *string `json:"availabilityZone" tf:"availability_zone,omitempty"`

	// If true, the disk will be encrypted.
	// +kubebuilder:validation:Optional
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// The amount of IOPS to provision for the disk. Only valid for type of io1, io2 or gp3.
	// +kubebuilder:validation:Optional
	Iops *float64 `json:"iops,omitempty" tf:"iops,omitempty"`

	// The ARN for the KMS encryption key. When specifying kms_key_id, encrypted needs to be set to true.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Reference to a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDRef *v1.Reference `json:"kmsKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDSelector *v1.Selector `json:"kmsKeyIdSelector,omitempty" tf:"-"`

	// Specifies whether to enable Amazon EBS Multi-Attach. Multi-Attach is supported on io1 and io2 volumes.
	// +kubebuilder:validation:Optional
	MultiAttachEnabled *bool `json:"multiAttachEnabled,omitempty" tf:"multi_attach_enabled,omitempty"`

	// The Amazon Resource Name (ARN) of the Outpost.
	// +kubebuilder:validation:Optional
	OutpostArn *string `json:"outpostArn,omitempty" tf:"outpost_arn,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The size of the drive in GiBs.
	// +kubebuilder:validation:Optional
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// A snapshot to base the EBS volume off of.
	// +kubebuilder:validation:Optional
	SnapshotID *string `json:"snapshotId,omitempty" tf:"snapshot_id,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The throughput that the volume supports, in MiB/s. Only valid for type of gp3.
	// +kubebuilder:validation:Optional
	Throughput *float64 `json:"throughput,omitempty" tf:"throughput,omitempty"`

	// The type of EBS volume. Can be standard, gp2, gp3, io1, io2, sc1 or st1 (Default: gp2).
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// EBSVolumeSpec defines the desired state of EBSVolume
type EBSVolumeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EBSVolumeParameters `json:"forProvider"`
}

// EBSVolumeStatus defines the observed state of EBSVolume.
type EBSVolumeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EBSVolumeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EBSVolume is the Schema for the EBSVolumes API. Provides an elastic block storage resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type EBSVolume struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EBSVolumeSpec   `json:"spec"`
	Status            EBSVolumeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EBSVolumeList contains a list of EBSVolumes
type EBSVolumeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EBSVolume `json:"items"`
}

// Repository type metadata.
var (
	EBSVolume_Kind             = "EBSVolume"
	EBSVolume_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EBSVolume_Kind}.String()
	EBSVolume_KindAPIVersion   = EBSVolume_Kind + "." + CRDGroupVersion.String()
	EBSVolume_GroupVersionKind = CRDGroupVersion.WithKind(EBSVolume_Kind)
)

func init() {
	SchemeBuilder.Register(&EBSVolume{}, &EBSVolumeList{})
}

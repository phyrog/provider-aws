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

type VPCObservation struct {

	// Amazon Resource Name (ARN) of VPC
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The ID of the VPC
	DHCPOptionsID *string `json:"dhcpOptionsId,omitempty" tf:"dhcp_options_id,omitempty"`

	// The ID of the network ACL created by default on VPC creation
	DefaultNetworkACLID *string `json:"defaultNetworkAclId,omitempty" tf:"default_network_acl_id,omitempty"`

	// The ID of the route table created by default on VPC creation
	DefaultRouteTableID *string `json:"defaultRouteTableId,omitempty" tf:"default_route_table_id,omitempty"`

	// The ID of the security group created by default on VPC creation
	DefaultSecurityGroupID *string `json:"defaultSecurityGroupId,omitempty" tf:"default_security_group_id,omitempty"`

	// The ID of the VPC
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The association ID for the IPv6 CIDR block.
	IPv6AssociationID *string `json:"ipv6AssociationId,omitempty" tf:"ipv6_association_id,omitempty"`

	// The ID of the main route table associated with
	// this VPC. Note that you can change a VPC's main route table by using an
	// aws_main_route_table_association.
	MainRouteTableID *string `json:"mainRouteTableId,omitempty" tf:"main_route_table_id,omitempty"`

	// The ID of the AWS account that owns the VPC.
	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type VPCParameters struct {

	// Requests an Amazon-provided IPv6 CIDR block with a /56 prefix length for the VPC. You cannot specify the range of IP addresses, or the size of the CIDR block. Default is false. Conflicts with ipv6_ipam_pool_id
	// +kubebuilder:validation:Optional
	AssignGeneratedIPv6CidrBlock *bool `json:"assignGeneratedIpv6CidrBlock,omitempty" tf:"assign_generated_ipv6_cidr_block,omitempty"`

	// The IPv4 CIDR block for the VPC. CIDR can be explicitly set or it can be derived from IPAM using ipv4_netmask_length.
	// +kubebuilder:validation:Optional
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`

	// A boolean flag to enable/disable ClassicLink
	// for the VPC. Only valid in regions and accounts that support EC2 Classic.
	// See the ClassicLink documentation for more information. Defaults false.
	// +kubebuilder:validation:Optional
	EnableClassiclink *bool `json:"enableClassiclink,omitempty" tf:"enable_classiclink,omitempty"`

	// A boolean flag to enable/disable ClassicLink DNS Support for the VPC.
	// Only valid in regions and accounts that support EC2 Classic.
	// +kubebuilder:validation:Optional
	EnableClassiclinkDNSSupport *bool `json:"enableClassiclinkDnsSupport,omitempty" tf:"enable_classiclink_dns_support,omitempty"`

	// A boolean flag to enable/disable DNS hostnames in the VPC. Defaults false.
	// +kubebuilder:validation:Optional
	EnableDNSHostnames *bool `json:"enableDnsHostnames,omitempty" tf:"enable_dns_hostnames,omitempty"`

	// A boolean flag to enable/disable DNS support in the VPC. Defaults true.
	// +kubebuilder:validation:Optional
	EnableDNSSupport *bool `json:"enableDnsSupport,omitempty" tf:"enable_dns_support,omitempty"`

	// The ID of an IPv4 IPAM pool you want to use for allocating this VPC's CIDR. IPAM is a VPC feature that you can use to automate your IP address management workflows including assigning, tracking, troubleshooting, and auditing IP addresses across AWS Regions and accounts. Using IPAM you can monitor IP address usage throughout your AWS Organization.
	// +kubebuilder:validation:Optional
	IPv4IpamPoolID *string `json:"ipv4IpamPoolId,omitempty" tf:"ipv4_ipam_pool_id,omitempty"`

	// The netmask length of the IPv4 CIDR you want to allocate to this VPC. Requires specifying a ipv4_ipam_pool_id.
	// +kubebuilder:validation:Optional
	IPv4NetmaskLength *float64 `json:"ipv4NetmaskLength,omitempty" tf:"ipv4_netmask_length,omitempty"`

	// IPv6 CIDR block to request from an IPAM Pool. Can be set explicitly or derived from IPAM using ipv6_netmask_length.
	// +kubebuilder:validation:Optional
	IPv6CidrBlock *string `json:"ipv6CidrBlock,omitempty" tf:"ipv6_cidr_block,omitempty"`

	// By default when an IPv6 CIDR is assigned to a VPC a default ipv6_cidr_block_network_border_group will be set to the region of the VPC. This can be changed to restrict advertisement of public addresses to specific Network Border Groups such as LocalZones.
	// +kubebuilder:validation:Optional
	IPv6CidrBlockNetworkBorderGroup *string `json:"ipv6CidrBlockNetworkBorderGroup,omitempty" tf:"ipv6_cidr_block_network_border_group,omitempty"`

	// IPAM Pool ID for a IPv6 pool. Conflicts with assign_generated_ipv6_cidr_block.
	// +kubebuilder:validation:Optional
	IPv6IpamPoolID *string `json:"ipv6IpamPoolId,omitempty" tf:"ipv6_ipam_pool_id,omitempty"`

	// Netmask length to request from IPAM Pool. Conflicts with ipv6_cidr_block. This can be omitted if IPAM pool as a allocation_default_netmask_length set. Valid values: 56.
	// +kubebuilder:validation:Optional
	IPv6NetmaskLength *float64 `json:"ipv6NetmaskLength,omitempty" tf:"ipv6_netmask_length,omitempty"`

	// A tenancy option for instances launched into the VPC. Default is default, which ensures that EC2 instances launched in this VPC use the EC2 instance tenancy attribute specified when the EC2 instance is launched. The only other option is dedicated, which ensures that EC2 instances launched in this VPC are run on dedicated tenancy instances regardless of the tenancy attribute specified at launch. This has a dedicated per region fee of $2 per hour, plus an hourly per instance usage fee.
	// +kubebuilder:validation:Optional
	InstanceTenancy *string `json:"instanceTenancy,omitempty" tf:"instance_tenancy,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// A map of tags to assign to the resource. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// VPCSpec defines the desired state of VPC
type VPCSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VPCParameters `json:"forProvider"`
}

// VPCStatus defines the observed state of VPC.
type VPCStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VPCObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VPC is the Schema for the VPCs API. Provides a VPC resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type VPC struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VPCSpec   `json:"spec"`
	Status            VPCStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPCList contains a list of VPCs
type VPCList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPC `json:"items"`
}

// Repository type metadata.
var (
	VPC_Kind             = "VPC"
	VPC_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VPC_Kind}.String()
	VPC_KindAPIVersion   = VPC_Kind + "." + CRDGroupVersion.String()
	VPC_GroupVersionKind = CRDGroupVersion.WithKind(VPC_Kind)
)

func init() {
	SchemeBuilder.Register(&VPC{}, &VPCList{})
}

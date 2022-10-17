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

type AttachmentObservation struct {

	// ID of the network interface.
	AttachmentID *string `json:"attachmentId,omitempty" tf:"attachment_id,omitempty"`

	// Integer to define the devices index.
	DeviceIndex *float64 `json:"deviceIndex,omitempty" tf:"device_index,omitempty"`

	// ID of the instance to attach to.
	Instance *string `json:"instance,omitempty" tf:"instance,omitempty"`
}

type AttachmentParameters struct {
}

type NetworkInterfaceObservation_2 struct {

	// ARN of the network interface.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Configuration block to define the attachment of the ENI. See Attachment below for more details!
	Attachment []AttachmentObservation `json:"attachment,omitempty" tf:"attachment,omitempty"`

	// ID of the network interface.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// MAC address of the network interface.
	MacAddress *string `json:"macAddress,omitempty" tf:"mac_address,omitempty"`

	// ARN of the network interface.
	OutpostArn *string `json:"outpostArn,omitempty" tf:"outpost_arn,omitempty"`

	// AWS account ID of the owner of the network interface.
	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	// Private DNS name of the network interface (IPv4).
	PrivateDNSName *string `json:"privateDnsName,omitempty" tf:"private_dns_name,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type NetworkInterfaceParameters_2 struct {

	// Description for the network interface.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Number of IPv4 prefixes that AWS automatically assigns to the network interface.
	// +kubebuilder:validation:Optional
	IPv4PrefixCount *float64 `json:"ipv4PrefixCount,omitempty" tf:"ipv4_prefix_count,omitempty"`

	// One or more IPv4 prefixes assigned to the network interface.
	// +kubebuilder:validation:Optional
	IPv4Prefixes []*string `json:"ipv4Prefixes,omitempty" tf:"ipv4_prefixes,omitempty"`

	// Number of IPv6 addresses to assign to a network interface. You can't use this option if specifying specific ipv6_addresses. If your subnet has the AssignIpv6AddressOnCreation attribute set to true, you can specify 0 to override this setting.
	// +kubebuilder:validation:Optional
	IPv6AddressCount *float64 `json:"ipv6AddressCount,omitempty" tf:"ipv6_address_count,omitempty"`

	// List of private IPs to assign to the ENI in sequential order.
	// +kubebuilder:validation:Optional
	IPv6AddressList []*string `json:"ipv6AddressList,omitempty" tf:"ipv6_address_list,omitempty"`

	// +kubebuilder:validation:Optional
	IPv6AddressListEnabled *bool `json:"ipv6AddressListEnabled,omitempty" tf:"ipv6_address_list_enabled,omitempty"`

	// One or more specific IPv6 addresses from the IPv6 CIDR block range of your subnet. Addresses are assigned without regard to order. You can't use this option if you're specifying ipv6_address_count.
	// +kubebuilder:validation:Optional
	IPv6Addresses []*string `json:"ipv6Addresses,omitempty" tf:"ipv6_addresses,omitempty"`

	// Number of IPv6 prefixes that AWS automatically assigns to the network interface.
	// +kubebuilder:validation:Optional
	IPv6PrefixCount *float64 `json:"ipv6PrefixCount,omitempty" tf:"ipv6_prefix_count,omitempty"`

	// One or more IPv6 prefixes assigned to the network interface.
	// +kubebuilder:validation:Optional
	IPv6Prefixes []*string `json:"ipv6Prefixes,omitempty" tf:"ipv6_prefixes,omitempty"`

	// Type of network interface to create. Set to efa for Elastic Fabric Adapter. Changing interface_type will cause the resource to be destroyed and re-created.
	// +kubebuilder:validation:Optional
	InterfaceType *string `json:"interfaceType,omitempty" tf:"interface_type,omitempty"`

	// +kubebuilder:validation:Optional
	PrivateIP *string `json:"privateIp,omitempty" tf:"private_ip,omitempty"`

	// List of private IPs to assign to the ENI in sequential order. Requires setting private_ip_list_enable to true.
	// +kubebuilder:validation:Optional
	PrivateIPList []*string `json:"privateIpList,omitempty" tf:"private_ip_list,omitempty"`

	// +kubebuilder:validation:Optional
	PrivateIPListEnabled *bool `json:"privateIpListEnabled,omitempty" tf:"private_ip_list_enabled,omitempty"`

	// List of private IPs to assign to the ENI without regard to order.
	// +kubebuilder:validation:Optional
	PrivateIps []*string `json:"privateIps,omitempty" tf:"private_ips,omitempty"`

	// Number of secondary private IPs to assign to the ENI. The total number of private IPs will be 1 + private_ips_count, as a primary private IP will be assiged to an ENI by default.
	// +kubebuilder:validation:Optional
	PrivateIpsCount *float64 `json:"privateIpsCount,omitempty" tf:"private_ips_count,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// References to SecurityGroup to populate securityGroups.
	// +kubebuilder:validation:Optional
	SecurityGroupRefs []v1.Reference `json:"securityGroupRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup to populate securityGroups.
	// +kubebuilder:validation:Optional
	SecurityGroupSelector *v1.Selector `json:"securityGroupSelector,omitempty" tf:"-"`

	// List of security group IDs to assign to the ENI.
	// +crossplane:generate:reference:type=SecurityGroup
	// +crossplane:generate:reference:refFieldName=SecurityGroupRefs
	// +crossplane:generate:reference:selectorFieldName=SecurityGroupSelector
	// +kubebuilder:validation:Optional
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// Whether to enable source destination checking for the ENI. Default true.
	// +kubebuilder:validation:Optional
	SourceDestCheck *bool `json:"sourceDestCheck,omitempty" tf:"source_dest_check,omitempty"`

	// Subnet ID to create the ENI in.
	// +crossplane:generate:reference:type=Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// Map of tags to assign to the resource. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// NetworkInterfaceSpec defines the desired state of NetworkInterface
type NetworkInterfaceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NetworkInterfaceParameters_2 `json:"forProvider"`
}

// NetworkInterfaceStatus defines the observed state of NetworkInterface.
type NetworkInterfaceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NetworkInterfaceObservation_2 `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkInterface is the Schema for the NetworkInterfaces API. Provides an Elastic network interface (ENI) resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type NetworkInterface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkInterfaceSpec   `json:"spec"`
	Status            NetworkInterfaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkInterfaceList contains a list of NetworkInterfaces
type NetworkInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkInterface `json:"items"`
}

// Repository type metadata.
var (
	NetworkInterface_Kind             = "NetworkInterface"
	NetworkInterface_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NetworkInterface_Kind}.String()
	NetworkInterface_KindAPIVersion   = NetworkInterface_Kind + "." + CRDGroupVersion.String()
	NetworkInterface_GroupVersionKind = CRDGroupVersion.WithKind(NetworkInterface_Kind)
)

func init() {
	SchemeBuilder.Register(&NetworkInterface{}, &NetworkInterfaceList{})
}

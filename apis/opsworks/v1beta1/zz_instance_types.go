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

type EBSBlockDeviceObservation struct {
}

type EBSBlockDeviceParameters struct {

	// Whether the volume should be destroyed on instance termination. Default is true.
	// +kubebuilder:validation:Optional
	DeleteOnTermination *bool `json:"deleteOnTermination,omitempty" tf:"delete_on_termination,omitempty"`

	// Name of the device to mount.
	// +kubebuilder:validation:Required
	DeviceName *string `json:"deviceName" tf:"device_name,omitempty"`

	// Amount of provisioned IOPS. This must be set with a volume_type of io1.
	// +kubebuilder:validation:Optional
	Iops *float64 `json:"iops,omitempty" tf:"iops,omitempty"`

	// Snapshot ID to mount.
	// +kubebuilder:validation:Optional
	SnapshotID *string `json:"snapshotId,omitempty" tf:"snapshot_id,omitempty"`

	// Size of the volume in gigabytes.
	// +kubebuilder:validation:Optional
	VolumeSize *float64 `json:"volumeSize,omitempty" tf:"volume_size,omitempty"`

	// Type of volume. Valid values are standard, gp2, or io1. Default is standard.
	// +kubebuilder:validation:Optional
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type EphemeralBlockDeviceObservation struct {
}

type EphemeralBlockDeviceParameters struct {

	// Name of the block device to mount on the instance.
	// +kubebuilder:validation:Required
	DeviceName *string `json:"deviceName" tf:"device_name,omitempty"`

	// The Instance Store Device Name (e.g., ephemeral0).
	// +kubebuilder:validation:Required
	VirtualName *string `json:"virtualName" tf:"virtual_name,omitempty"`
}

type InstanceObservation struct {

	// EC2 instance ID.
	EC2InstanceID *string `json:"ec2InstanceId,omitempty" tf:"ec2_instance_id,omitempty"`

	// ID of the OpsWorks instance.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// ID of the last service error.
	LastServiceErrorID *string `json:"lastServiceErrorId,omitempty" tf:"last_service_error_id,omitempty"`

	// Instance's platform.
	Platform *string `json:"platform,omitempty" tf:"platform,omitempty"`

	// Private DNS name assigned to the instance. Can only be used inside the Amazon EC2, and only available if you've enabled DNS hostnames for your VPC.
	PrivateDNS *string `json:"privateDns,omitempty" tf:"private_dns,omitempty"`

	// Private IP address assigned to the instance.
	PrivateIP *string `json:"privateIp,omitempty" tf:"private_ip,omitempty"`

	// Public DNS name assigned to the instance. For EC2-VPC, this is only available if you've enabled DNS hostnames for your VPC.
	PublicDNS *string `json:"publicDns,omitempty" tf:"public_dns,omitempty"`

	// Public IP address assigned to the instance, if applicable.
	PublicIP *string `json:"publicIp,omitempty" tf:"public_ip,omitempty"`

	// For registered instances, who performed the registration.
	RegisteredBy *string `json:"registeredBy,omitempty" tf:"registered_by,omitempty"`

	// Instance's reported AWS OpsWorks Stacks agent version.
	ReportedAgentVersion *string `json:"reportedAgentVersion,omitempty" tf:"reported_agent_version,omitempty"`

	// For registered instances, the reported operating system family.
	ReportedOsFamily *string `json:"reportedOsFamily,omitempty" tf:"reported_os_family,omitempty"`

	// For registered instances, the reported operating system name.
	ReportedOsName *string `json:"reportedOsName,omitempty" tf:"reported_os_name,omitempty"`

	// For registered instances, the reported operating system version.
	ReportedOsVersion *string `json:"reportedOsVersion,omitempty" tf:"reported_os_version,omitempty"`

	// Root device volume ID.
	RootDeviceVolumeID *string `json:"rootDeviceVolumeId,omitempty" tf:"root_device_volume_id,omitempty"`

	// SSH key's Deep Security Agent (DSA) fingerprint.
	SSHHostDsaKeyFingerprint *string `json:"sshHostDsaKeyFingerprint,omitempty" tf:"ssh_host_dsa_key_fingerprint,omitempty"`

	// SSH key's RSA fingerprint.
	SSHHostRsaKeyFingerprint *string `json:"sshHostRsaKeyFingerprint,omitempty" tf:"ssh_host_rsa_key_fingerprint,omitempty"`
}

type InstanceParameters struct {

	// AMI to use for the instance.  If an AMI is specified, os must be Custom.
	// +kubebuilder:validation:Optional
	AMIID *string `json:"amiId,omitempty" tf:"ami_id,omitempty"`

	// OpsWorks agent to install. Default is INHERIT.
	// +kubebuilder:validation:Optional
	AgentVersion *string `json:"agentVersion,omitempty" tf:"agent_version,omitempty"`

	// Machine architecture for created instances.  Valid values are x86_64 or i386. The default is x86_64.
	// +kubebuilder:validation:Optional
	Architecture *string `json:"architecture,omitempty" tf:"architecture,omitempty"`

	// Creates load-based or time-based instances.  Valid values are load, timer.
	// +kubebuilder:validation:Optional
	AutoScalingType *string `json:"autoScalingType,omitempty" tf:"auto_scaling_type,omitempty"`

	// Name of the availability zone where instances will be created by default.
	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// Time that the instance was created.
	// +kubebuilder:validation:Optional
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Whether to delete EBS volume on deletion. Default is true.
	// +kubebuilder:validation:Optional
	DeleteEBS *bool `json:"deleteEbs,omitempty" tf:"delete_ebs,omitempty"`

	// Whether to delete the Elastic IP on deletion.
	// +kubebuilder:validation:Optional
	DeleteEIP *bool `json:"deleteEip,omitempty" tf:"delete_eip,omitempty"`

	// Configuration block for additional EBS block devices to attach to the instance. See Block Devices below.
	// +kubebuilder:validation:Optional
	EBSBlockDevice []EBSBlockDeviceParameters `json:"ebsBlockDevice,omitempty" tf:"ebs_block_device,omitempty"`

	// Whether the launched EC2 instance will be EBS-optimized.
	// +kubebuilder:validation:Optional
	EBSOptimized *bool `json:"ebsOptimized,omitempty" tf:"ebs_optimized,omitempty"`

	// ECS cluster's ARN for container instances.
	// +kubebuilder:validation:Optional
	EcsClusterArn *string `json:"ecsClusterArn,omitempty" tf:"ecs_cluster_arn,omitempty"`

	// Instance Elastic IP address.
	// +kubebuilder:validation:Optional
	ElasticIP *string `json:"elasticIp,omitempty" tf:"elastic_ip,omitempty"`

	// Configuration block for ephemeral (also known as "Instance Store") volumes on the instance. See Block Devices below.
	// +kubebuilder:validation:Optional
	EphemeralBlockDevice []EphemeralBlockDeviceParameters `json:"ephemeralBlockDevice,omitempty" tf:"ephemeral_block_device,omitempty"`

	// Instance's host name.
	// +kubebuilder:validation:Optional
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// For registered instances, infrastructure class: ec2 or on-premises.
	// +kubebuilder:validation:Optional
	InfrastructureClass *string `json:"infrastructureClass,omitempty" tf:"infrastructure_class,omitempty"`

	// Controls where to install OS and package updates when the instance boots.  Default is true.
	// +kubebuilder:validation:Optional
	InstallUpdatesOnBoot *bool `json:"installUpdatesOnBoot,omitempty" tf:"install_updates_on_boot,omitempty"`

	// ARN of the instance's IAM profile.
	// +kubebuilder:validation:Optional
	InstanceProfileArn *string `json:"instanceProfileArn,omitempty" tf:"instance_profile_arn,omitempty"`

	// Type of instance to start.
	// +kubebuilder:validation:Optional
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// List of the layers the instance will belong to.
	// +crossplane:generate:reference:type=CustomLayer
	// +kubebuilder:validation:Optional
	LayerIds []*string `json:"layerIds,omitempty" tf:"layer_ids,omitempty"`

	// References to CustomLayer to populate layerIds.
	// +kubebuilder:validation:Optional
	LayerIdsRefs []v1.Reference `json:"layerIdsRefs,omitempty" tf:"-"`

	// Selector for a list of CustomLayer to populate layerIds.
	// +kubebuilder:validation:Optional
	LayerIdsSelector *v1.Selector `json:"layerIdsSelector,omitempty" tf:"-"`

	// Name of operating system that will be installed.
	// +kubebuilder:validation:Optional
	Os *string `json:"os,omitempty" tf:"os,omitempty"`

	// Configuration block for the root block device of the instance. See Block Devices below.
	// +kubebuilder:validation:Optional
	RootBlockDevice []RootBlockDeviceParameters `json:"rootBlockDevice,omitempty" tf:"root_block_device,omitempty"`

	// Name of the type of root device instances will have by default. Valid values are ebs or instance-store.
	// +kubebuilder:validation:Optional
	RootDeviceType *string `json:"rootDeviceType,omitempty" tf:"root_device_type,omitempty"`

	// Name of the SSH keypair that instances will have by default.
	// +kubebuilder:validation:Optional
	SSHKeyName *string `json:"sshKeyName,omitempty" tf:"ssh_key_name,omitempty"`

	// References to SecurityGroup in ec2 to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIDRefs []v1.Reference `json:"securityGroupIdRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in ec2 to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIDSelector *v1.Selector `json:"securityGroupIdSelector,omitempty" tf:"-"`

	// Associated security groups.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup
	// +crossplane:generate:reference:refFieldName=SecurityGroupIDRefs
	// +crossplane:generate:reference:selectorFieldName=SecurityGroupIDSelector
	// +kubebuilder:validation:Optional
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// Identifier of the stack the instance will belong to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/opsworks/v1beta1.Stack
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	StackID *string `json:"stackId,omitempty" tf:"stack_id,omitempty"`

	// Reference to a Stack in opsworks to populate stackId.
	// +kubebuilder:validation:Optional
	StackIDRef *v1.Reference `json:"stackIdRef,omitempty" tf:"-"`

	// Selector for a Stack in opsworks to populate stackId.
	// +kubebuilder:validation:Optional
	StackIDSelector *v1.Selector `json:"stackIdSelector,omitempty" tf:"-"`

	// Desired state of the instance. Valid values are running or stopped.
	// +kubebuilder:validation:Optional
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// Instance status. Will be one of booting, connection_lost, online, pending, rebooting, requested, running_setup, setup_failed, shutting_down, start_failed, stop_failed, stopped, stopping, terminated, or terminating.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Subnet ID to attach to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in ec2 to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in ec2 to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// Instance tenancy to use. Valid values are default, dedicated or host.
	// +kubebuilder:validation:Optional
	Tenancy *string `json:"tenancy,omitempty" tf:"tenancy,omitempty"`

	// Keyword to choose what virtualization mode created instances will use. Valid values are paravirtual or hvm.
	// +kubebuilder:validation:Optional
	VirtualizationType *string `json:"virtualizationType,omitempty" tf:"virtualization_type,omitempty"`
}

type RootBlockDeviceObservation struct {
}

type RootBlockDeviceParameters struct {

	// Whether the volume should be destroyed on instance termination. Default is true.
	// +kubebuilder:validation:Optional
	DeleteOnTermination *bool `json:"deleteOnTermination,omitempty" tf:"delete_on_termination,omitempty"`

	// Amount of provisioned IOPS. This must be set with a volume_type of io1.
	// +kubebuilder:validation:Optional
	Iops *float64 `json:"iops,omitempty" tf:"iops,omitempty"`

	// Size of the volume in gigabytes.
	// +kubebuilder:validation:Optional
	VolumeSize *float64 `json:"volumeSize,omitempty" tf:"volume_size,omitempty"`

	// Type of volume. Valid values are standard, gp2, or io1. Default is standard.
	// +kubebuilder:validation:Optional
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

// InstanceSpec defines the desired state of Instance
type InstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceParameters `json:"forProvider"`
}

// InstanceStatus defines the observed state of Instance.
type InstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Instance is the Schema for the Instances API. Provides an OpsWorks instance resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Instance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceSpec   `json:"spec"`
	Status            InstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceList contains a list of Instances
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Instance `json:"items"`
}

// Repository type metadata.
var (
	Instance_Kind             = "Instance"
	Instance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Instance_Kind}.String()
	Instance_KindAPIVersion   = Instance_Kind + "." + CRDGroupVersion.String()
	Instance_GroupVersionKind = CRDGroupVersion.WithKind(Instance_Kind)
)

func init() {
	SchemeBuilder.Register(&Instance{}, &InstanceList{})
}

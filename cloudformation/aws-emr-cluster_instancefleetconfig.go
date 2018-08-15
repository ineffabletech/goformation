package cloudformation

// AWSEMRCluster_InstanceFleetConfig AWS CloudFormation Resource (AWS::EMR::Cluster.InstanceFleetConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-instancefleetconfig.html
type AWSEMRCluster_InstanceFleetConfig struct {
	dependsOn []string

	// InstanceTypeConfigs AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-instancefleetconfig.html#cfn-elasticmapreduce-cluster-instancefleetconfig-instancetypeconfigs
	InstanceTypeConfigs []AWSEMRCluster_InstanceTypeConfig `json:"InstanceTypeConfigs,omitempty"`

	// LaunchSpecifications AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-instancefleetconfig.html#cfn-elasticmapreduce-cluster-instancefleetconfig-launchspecifications
	LaunchSpecifications *AWSEMRCluster_InstanceFleetProvisioningSpecifications `json:"LaunchSpecifications,omitempty"`

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-instancefleetconfig.html#cfn-elasticmapreduce-cluster-instancefleetconfig-name
	Name *String `json:"Name,omitempty"`

	// TargetOnDemandCapacity AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-instancefleetconfig.html#cfn-elasticmapreduce-cluster-instancefleetconfig-targetondemandcapacity
	TargetOnDemandCapacity *Integer `json:"TargetOnDemandCapacity,omitempty"`

	// TargetSpotCapacity AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-instancefleetconfig.html#cfn-elasticmapreduce-cluster-instancefleetconfig-targetspotcapacity
	TargetSpotCapacity *Integer `json:"TargetSpotCapacity,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSEMRCluster_InstanceFleetConfig) AddDependencies(v ...string) *AWSEMRCluster_InstanceFleetConfig {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSEMRCluster_InstanceFleetConfig) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRCluster_InstanceFleetConfig) AWSCloudFormationType() string {
	return "AWS::EMR::Cluster.InstanceFleetConfig"
}

package cloudformation

// AWSEKSCluster_ResourcesVpcConfig AWS CloudFormation Resource (AWS::EKS::Cluster.ResourcesVpcConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-resourcesvpcconfig.html
type AWSEKSCluster_ResourcesVpcConfig struct {
	dependsOn []string

	// SecurityGroupIds AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-resourcesvpcconfig.html#cfn-eks-cluster-resourcesvpcconfig-securitygroupids
	SecurityGroupIds []*String `json:"SecurityGroupIds,omitempty"`

	// SubnetIds AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-resourcesvpcconfig.html#cfn-eks-cluster-resourcesvpcconfig-subnetids
	SubnetIds []*String `json:"SubnetIds,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSEKSCluster_ResourcesVpcConfig) AddDependencies(v ...string) *AWSEKSCluster_ResourcesVpcConfig {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSEKSCluster_ResourcesVpcConfig) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEKSCluster_ResourcesVpcConfig) AWSCloudFormationType() string {
	return "AWS::EKS::Cluster.ResourcesVpcConfig"
}

package cloudformation

// AWSEMRCluster_BootstrapActionConfig AWS CloudFormation Resource (AWS::EMR::Cluster.BootstrapActionConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-bootstrapactionconfig.html
type AWSEMRCluster_BootstrapActionConfig struct {
	dependsOn []string

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-bootstrapactionconfig.html#cfn-elasticmapreduce-cluster-bootstrapactionconfig-name
	Name *String `json:"Name,omitempty"`

	// ScriptBootstrapAction AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-bootstrapactionconfig.html#cfn-elasticmapreduce-cluster-bootstrapactionconfig-scriptbootstrapaction
	ScriptBootstrapAction *AWSEMRCluster_ScriptBootstrapActionConfig `json:"ScriptBootstrapAction,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSEMRCluster_BootstrapActionConfig) AddDependencies(v ...string) *AWSEMRCluster_BootstrapActionConfig {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSEMRCluster_BootstrapActionConfig) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRCluster_BootstrapActionConfig) AWSCloudFormationType() string {
	return "AWS::EMR::Cluster.BootstrapActionConfig"
}

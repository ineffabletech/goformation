package cloudformation

// AWSEMRCluster_ScriptBootstrapActionConfig AWS CloudFormation Resource (AWS::EMR::Cluster.ScriptBootstrapActionConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-scriptbootstrapactionconfig.html
type AWSEMRCluster_ScriptBootstrapActionConfig struct {
	dependsOn []string

	// Args AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-scriptbootstrapactionconfig.html#cfn-elasticmapreduce-cluster-scriptbootstrapactionconfig-args
	Args []*String `json:"Args,omitempty"`

	// Path AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-scriptbootstrapactionconfig.html#cfn-elasticmapreduce-cluster-scriptbootstrapactionconfig-path
	Path *String `json:"Path,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSEMRCluster_ScriptBootstrapActionConfig) AddDependencies(v ...string) *AWSEMRCluster_ScriptBootstrapActionConfig {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSEMRCluster_ScriptBootstrapActionConfig) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRCluster_ScriptBootstrapActionConfig) AWSCloudFormationType() string {
	return "AWS::EMR::Cluster.ScriptBootstrapActionConfig"
}

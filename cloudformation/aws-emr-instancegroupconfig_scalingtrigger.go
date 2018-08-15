package cloudformation

// AWSEMRInstanceGroupConfig_ScalingTrigger AWS CloudFormation Resource (AWS::EMR::InstanceGroupConfig.ScalingTrigger)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-scalingtrigger.html
type AWSEMRInstanceGroupConfig_ScalingTrigger struct {
	dependsOn []string

	// CloudWatchAlarmDefinition AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-scalingtrigger.html#cfn-elasticmapreduce-instancegroupconfig-scalingtrigger-cloudwatchalarmdefinition
	CloudWatchAlarmDefinition *AWSEMRInstanceGroupConfig_CloudWatchAlarmDefinition `json:"CloudWatchAlarmDefinition,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSEMRInstanceGroupConfig_ScalingTrigger) AddDependencies(v ...string) *AWSEMRInstanceGroupConfig_ScalingTrigger {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSEMRInstanceGroupConfig_ScalingTrigger) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRInstanceGroupConfig_ScalingTrigger) AWSCloudFormationType() string {
	return "AWS::EMR::InstanceGroupConfig.ScalingTrigger"
}

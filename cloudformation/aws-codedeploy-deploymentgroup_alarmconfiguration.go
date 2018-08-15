package cloudformation

// AWSCodeDeployDeploymentGroup_AlarmConfiguration AWS CloudFormation Resource (AWS::CodeDeploy::DeploymentGroup.AlarmConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-alarmconfiguration.html
type AWSCodeDeployDeploymentGroup_AlarmConfiguration struct {
	dependsOn []string

	// Alarms AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-alarmconfiguration.html#cfn-codedeploy-deploymentgroup-alarmconfiguration-alarms
	Alarms []AWSCodeDeployDeploymentGroup_Alarm `json:"Alarms,omitempty"`

	// Enabled AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-alarmconfiguration.html#cfn-codedeploy-deploymentgroup-alarmconfiguration-enabled
	Enabled *Boolean `json:"Enabled,omitempty"`

	// IgnorePollAlarmFailure AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-alarmconfiguration.html#cfn-codedeploy-deploymentgroup-alarmconfiguration-ignorepollalarmfailure
	IgnorePollAlarmFailure *Boolean `json:"IgnorePollAlarmFailure,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSCodeDeployDeploymentGroup_AlarmConfiguration) AddDependencies(v ...string) *AWSCodeDeployDeploymentGroup_AlarmConfiguration {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSCodeDeployDeploymentGroup_AlarmConfiguration) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodeDeployDeploymentGroup_AlarmConfiguration) AWSCloudFormationType() string {
	return "AWS::CodeDeploy::DeploymentGroup.AlarmConfiguration"
}

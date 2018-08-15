package cloudformation

// AWSEC2LaunchTemplate_Monitoring AWS CloudFormation Resource (AWS::EC2::LaunchTemplate.Monitoring)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata-monitoring.html
type AWSEC2LaunchTemplate_Monitoring struct {
	dependsOn []string

	// Enabled AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata-monitoring.html#cfn-ec2-launchtemplate-launchtemplatedata-monitoring-enabled
	Enabled *Boolean `json:"Enabled,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSEC2LaunchTemplate_Monitoring) AddDependencies(v ...string) *AWSEC2LaunchTemplate_Monitoring {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSEC2LaunchTemplate_Monitoring) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2LaunchTemplate_Monitoring) AWSCloudFormationType() string {
	return "AWS::EC2::LaunchTemplate.Monitoring"
}

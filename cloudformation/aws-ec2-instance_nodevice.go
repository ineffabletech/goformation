package cloudformation

// AWSEC2Instance_NoDevice AWS CloudFormation Resource (AWS::EC2::Instance.NoDevice)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-nodevice.html
type AWSEC2Instance_NoDevice struct {
	dependsOn []string
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSEC2Instance_NoDevice) AddDependencies(v ...string) *AWSEC2Instance_NoDevice {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSEC2Instance_NoDevice) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2Instance_NoDevice) AWSCloudFormationType() string {
	return "AWS::EC2::Instance.NoDevice"
}

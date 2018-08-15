package cloudformation

// AWSEC2LaunchTemplate_ElasticGpuSpecification AWS CloudFormation Resource (AWS::EC2::LaunchTemplate.ElasticGpuSpecification)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-elasticgpuspecification.html
type AWSEC2LaunchTemplate_ElasticGpuSpecification struct {
	dependsOn []string

	// Type AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-elasticgpuspecification.html#cfn-ec2-launchtemplate-elasticgpuspecification-type
	Type *String `json:"Type,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSEC2LaunchTemplate_ElasticGpuSpecification) AddDependencies(v ...string) *AWSEC2LaunchTemplate_ElasticGpuSpecification {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSEC2LaunchTemplate_ElasticGpuSpecification) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2LaunchTemplate_ElasticGpuSpecification) AWSCloudFormationType() string {
	return "AWS::EC2::LaunchTemplate.ElasticGpuSpecification"
}

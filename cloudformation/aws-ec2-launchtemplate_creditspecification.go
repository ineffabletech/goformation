package cloudformation

// AWSEC2LaunchTemplate_CreditSpecification AWS CloudFormation Resource (AWS::EC2::LaunchTemplate.CreditSpecification)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata-creditspecification.html
type AWSEC2LaunchTemplate_CreditSpecification struct {
	dependsOn []string

	// CpuCredits AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata-creditspecification.html#cfn-ec2-launchtemplate-launchtemplatedata-creditspecification-cpucredits
	CpuCredits *String `json:"CpuCredits,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSEC2LaunchTemplate_CreditSpecification) AddDependencies(v ...string) *AWSEC2LaunchTemplate_CreditSpecification {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSEC2LaunchTemplate_CreditSpecification) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2LaunchTemplate_CreditSpecification) AWSCloudFormationType() string {
	return "AWS::EC2::LaunchTemplate.CreditSpecification"
}

package cloudformation

// AWSEC2Instance_CreditSpecification AWS CloudFormation Resource (AWS::EC2::Instance.CreditSpecification)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-creditspecification.html
type AWSEC2Instance_CreditSpecification struct {
	dependsOn []string

	// CPUCredits AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-creditspecification.html#cfn-ec2-instance-creditspecification-cpucredits
	CPUCredits *String `json:"CPUCredits,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSEC2Instance_CreditSpecification) AddDependencies(v ...string) *AWSEC2Instance_CreditSpecification {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSEC2Instance_CreditSpecification) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2Instance_CreditSpecification) AWSCloudFormationType() string {
	return "AWS::EC2::Instance.CreditSpecification"
}

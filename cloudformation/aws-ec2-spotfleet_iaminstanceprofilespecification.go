package cloudformation

// AWSEC2SpotFleet_IamInstanceProfileSpecification AWS CloudFormation Resource (AWS::EC2::SpotFleet.IamInstanceProfileSpecification)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-iaminstanceprofile.html
type AWSEC2SpotFleet_IamInstanceProfileSpecification struct {
	dependsOn []string

	// Arn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-iaminstanceprofile.html#cfn-ec2-spotfleet-iaminstanceprofilespecification-arn
	Arn *String `json:"Arn,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSEC2SpotFleet_IamInstanceProfileSpecification) AddDependencies(v ...string) *AWSEC2SpotFleet_IamInstanceProfileSpecification {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSEC2SpotFleet_IamInstanceProfileSpecification) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2SpotFleet_IamInstanceProfileSpecification) AWSCloudFormationType() string {
	return "AWS::EC2::SpotFleet.IamInstanceProfileSpecification"
}

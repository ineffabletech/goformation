package cloudformation

// AWSEC2SpotFleet_SpotFleetTagSpecification AWS CloudFormation Resource (AWS::EC2::SpotFleet.SpotFleetTagSpecification)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-tagspecifications.html
type AWSEC2SpotFleet_SpotFleetTagSpecification struct {
	dependsOn []string

	// ResourceType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-tagspecifications.html#cfn-ec2-spotfleet-spotfleettagspecification-resourcetype
	ResourceType *String `json:"ResourceType,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSEC2SpotFleet_SpotFleetTagSpecification) AddDependencies(v ...string) *AWSEC2SpotFleet_SpotFleetTagSpecification {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSEC2SpotFleet_SpotFleetTagSpecification) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2SpotFleet_SpotFleetTagSpecification) AWSCloudFormationType() string {
	return "AWS::EC2::SpotFleet.SpotFleetTagSpecification"
}

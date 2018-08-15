package cloudformation

// AWSEC2SpotFleet_SpotPlacement AWS CloudFormation Resource (AWS::EC2::SpotFleet.SpotPlacement)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-placement.html
type AWSEC2SpotFleet_SpotPlacement struct {
	dependsOn []string

	// AvailabilityZone AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-placement.html#cfn-ec2-spotfleet-spotplacement-availabilityzone
	AvailabilityZone *String `json:"AvailabilityZone,omitempty"`

	// GroupName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-placement.html#cfn-ec2-spotfleet-spotplacement-groupname
	GroupName *String `json:"GroupName,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSEC2SpotFleet_SpotPlacement) AddDependencies(v ...string) *AWSEC2SpotFleet_SpotPlacement {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSEC2SpotFleet_SpotPlacement) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2SpotFleet_SpotPlacement) AWSCloudFormationType() string {
	return "AWS::EC2::SpotFleet.SpotPlacement"
}

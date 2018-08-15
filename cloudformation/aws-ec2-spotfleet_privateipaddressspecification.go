package cloudformation

// AWSEC2SpotFleet_PrivateIpAddressSpecification AWS CloudFormation Resource (AWS::EC2::SpotFleet.PrivateIpAddressSpecification)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-networkinterfaces-privateipaddresses.html
type AWSEC2SpotFleet_PrivateIpAddressSpecification struct {
	dependsOn []string

	// Primary AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-networkinterfaces-privateipaddresses.html#cfn-ec2-spotfleet-privateipaddressspecification-primary
	Primary *Boolean `json:"Primary,omitempty"`

	// PrivateIpAddress AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-networkinterfaces-privateipaddresses.html#cfn-ec2-spotfleet-privateipaddressspecification-privateipaddress
	PrivateIpAddress *String `json:"PrivateIpAddress,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSEC2SpotFleet_PrivateIpAddressSpecification) AddDependencies(v ...string) *AWSEC2SpotFleet_PrivateIpAddressSpecification {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSEC2SpotFleet_PrivateIpAddressSpecification) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2SpotFleet_PrivateIpAddressSpecification) AWSCloudFormationType() string {
	return "AWS::EC2::SpotFleet.PrivateIpAddressSpecification"
}

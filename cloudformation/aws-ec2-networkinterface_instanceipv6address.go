package cloudformation

// AWSEC2NetworkInterface_InstanceIpv6Address AWS CloudFormation Resource (AWS::EC2::NetworkInterface.InstanceIpv6Address)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinterface-instanceipv6address.html
type AWSEC2NetworkInterface_InstanceIpv6Address struct {
	dependsOn []string

	// Ipv6Address AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinterface-instanceipv6address.html#cfn-ec2-networkinterface-instanceipv6address-ipv6address
	Ipv6Address *String `json:"Ipv6Address,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSEC2NetworkInterface_InstanceIpv6Address) AddDependencies(v ...string) *AWSEC2NetworkInterface_InstanceIpv6Address {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSEC2NetworkInterface_InstanceIpv6Address) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2NetworkInterface_InstanceIpv6Address) AWSCloudFormationType() string {
	return "AWS::EC2::NetworkInterface.InstanceIpv6Address"
}

package cloudformation

// AWSEC2NetworkAclEntry_Icmp AWS CloudFormation Resource (AWS::EC2::NetworkAclEntry.Icmp)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkaclentry-icmp.html
type AWSEC2NetworkAclEntry_Icmp struct {
	dependsOn []string

	// Code AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkaclentry-icmp.html#cfn-ec2-networkaclentry-icmp-code
	Code *Integer `json:"Code,omitempty"`

	// Type AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkaclentry-icmp.html#cfn-ec2-networkaclentry-icmp-type
	Type *Integer `json:"Type,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSEC2NetworkAclEntry_Icmp) AddDependencies(v ...string) *AWSEC2NetworkAclEntry_Icmp {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSEC2NetworkAclEntry_Icmp) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2NetworkAclEntry_Icmp) AWSCloudFormationType() string {
	return "AWS::EC2::NetworkAclEntry.Icmp"
}

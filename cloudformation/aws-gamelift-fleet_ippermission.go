package cloudformation

// AWSGameLiftFleet_IpPermission AWS CloudFormation Resource (AWS::GameLift::Fleet.IpPermission)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-ec2inboundpermission.html
type AWSGameLiftFleet_IpPermission struct {
	dependsOn []string

	// FromPort AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-ec2inboundpermission.html#cfn-gamelift-fleet-ec2inboundpermissions-fromport
	FromPort *Integer `json:"FromPort,omitempty"`

	// IpRange AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-ec2inboundpermission.html#cfn-gamelift-fleet-ec2inboundpermissions-iprange
	IpRange *String `json:"IpRange,omitempty"`

	// Protocol AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-ec2inboundpermission.html#cfn-gamelift-fleet-ec2inboundpermissions-protocol
	Protocol *String `json:"Protocol,omitempty"`

	// ToPort AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-ec2inboundpermission.html#cfn-gamelift-fleet-ec2inboundpermissions-toport
	ToPort *Integer `json:"ToPort,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSGameLiftFleet_IpPermission) AddDependencies(v ...string) *AWSGameLiftFleet_IpPermission {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSGameLiftFleet_IpPermission) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSGameLiftFleet_IpPermission) AWSCloudFormationType() string {
	return "AWS::GameLift::Fleet.IpPermission"
}

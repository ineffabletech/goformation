package cloudformation

// AWSEC2VPNConnection_VpnTunnelOptionsSpecification AWS CloudFormation Resource (AWS::EC2::VPNConnection.VpnTunnelOptionsSpecification)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-vpntunneloptionsspecification.html
type AWSEC2VPNConnection_VpnTunnelOptionsSpecification struct {
	dependsOn []string

	// PreSharedKey AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-vpntunneloptionsspecification.html#cfn-ec2-vpnconnection-vpntunneloptionsspecification-presharedkey
	PreSharedKey *String `json:"PreSharedKey,omitempty"`

	// TunnelInsideCidr AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-vpntunneloptionsspecification.html#cfn-ec2-vpnconnection-vpntunneloptionsspecification-tunnelinsidecidr
	TunnelInsideCidr *String `json:"TunnelInsideCidr,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSEC2VPNConnection_VpnTunnelOptionsSpecification) AddDependencies(v ...string) *AWSEC2VPNConnection_VpnTunnelOptionsSpecification {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSEC2VPNConnection_VpnTunnelOptionsSpecification) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2VPNConnection_VpnTunnelOptionsSpecification) AWSCloudFormationType() string {
	return "AWS::EC2::VPNConnection.VpnTunnelOptionsSpecification"
}

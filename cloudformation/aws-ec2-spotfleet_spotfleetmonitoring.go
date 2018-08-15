package cloudformation

// AWSEC2SpotFleet_SpotFleetMonitoring AWS CloudFormation Resource (AWS::EC2::SpotFleet.SpotFleetMonitoring)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-monitoring.html
type AWSEC2SpotFleet_SpotFleetMonitoring struct {
	dependsOn []string

	// Enabled AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-monitoring.html#cfn-ec2-spotfleet-spotfleetmonitoring-enabled
	Enabled *Boolean `json:"Enabled,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSEC2SpotFleet_SpotFleetMonitoring) AddDependencies(v ...string) *AWSEC2SpotFleet_SpotFleetMonitoring {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSEC2SpotFleet_SpotFleetMonitoring) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2SpotFleet_SpotFleetMonitoring) AWSCloudFormationType() string {
	return "AWS::EC2::SpotFleet.SpotFleetMonitoring"
}

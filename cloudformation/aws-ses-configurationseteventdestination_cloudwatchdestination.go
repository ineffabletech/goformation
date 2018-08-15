package cloudformation

// AWSSESConfigurationSetEventDestination_CloudWatchDestination AWS CloudFormation Resource (AWS::SES::ConfigurationSetEventDestination.CloudWatchDestination)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationseteventdestination-cloudwatchdestination.html
type AWSSESConfigurationSetEventDestination_CloudWatchDestination struct {
	dependsOn []string

	// DimensionConfigurations AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationseteventdestination-cloudwatchdestination.html#cfn-ses-configurationseteventdestination-cloudwatchdestination-dimensionconfigurations
	DimensionConfigurations []AWSSESConfigurationSetEventDestination_DimensionConfiguration `json:"DimensionConfigurations,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSSESConfigurationSetEventDestination_CloudWatchDestination) AddDependencies(v ...string) *AWSSESConfigurationSetEventDestination_CloudWatchDestination {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSSESConfigurationSetEventDestination_CloudWatchDestination) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSESConfigurationSetEventDestination_CloudWatchDestination) AWSCloudFormationType() string {
	return "AWS::SES::ConfigurationSetEventDestination.CloudWatchDestination"
}

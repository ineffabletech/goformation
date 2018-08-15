package cloudformation

// AWSSESConfigurationSetEventDestination_DimensionConfiguration AWS CloudFormation Resource (AWS::SES::ConfigurationSetEventDestination.DimensionConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationseteventdestination-dimensionconfiguration.html
type AWSSESConfigurationSetEventDestination_DimensionConfiguration struct {
	dependsOn []string

	// DefaultDimensionValue AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationseteventdestination-dimensionconfiguration.html#cfn-ses-configurationseteventdestination-dimensionconfiguration-defaultdimensionvalue
	DefaultDimensionValue *String `json:"DefaultDimensionValue,omitempty"`

	// DimensionName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationseteventdestination-dimensionconfiguration.html#cfn-ses-configurationseteventdestination-dimensionconfiguration-dimensionname
	DimensionName *String `json:"DimensionName,omitempty"`

	// DimensionValueSource AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationseteventdestination-dimensionconfiguration.html#cfn-ses-configurationseteventdestination-dimensionconfiguration-dimensionvaluesource
	DimensionValueSource *String `json:"DimensionValueSource,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSSESConfigurationSetEventDestination_DimensionConfiguration) AddDependencies(v ...string) *AWSSESConfigurationSetEventDestination_DimensionConfiguration {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSSESConfigurationSetEventDestination_DimensionConfiguration) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSESConfigurationSetEventDestination_DimensionConfiguration) AWSCloudFormationType() string {
	return "AWS::SES::ConfigurationSetEventDestination.DimensionConfiguration"
}

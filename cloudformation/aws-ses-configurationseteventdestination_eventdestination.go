package cloudformation

// AWSSESConfigurationSetEventDestination_EventDestination AWS CloudFormation Resource (AWS::SES::ConfigurationSetEventDestination.EventDestination)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationseteventdestination-eventdestination.html
type AWSSESConfigurationSetEventDestination_EventDestination struct {
	dependsOn []string

	// CloudWatchDestination AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationseteventdestination-eventdestination.html#cfn-ses-configurationseteventdestination-eventdestination-cloudwatchdestination
	CloudWatchDestination *AWSSESConfigurationSetEventDestination_CloudWatchDestination `json:"CloudWatchDestination,omitempty"`

	// Enabled AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationseteventdestination-eventdestination.html#cfn-ses-configurationseteventdestination-eventdestination-enabled
	Enabled *Boolean `json:"Enabled,omitempty"`

	// KinesisFirehoseDestination AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationseteventdestination-eventdestination.html#cfn-ses-configurationseteventdestination-eventdestination-kinesisfirehosedestination
	KinesisFirehoseDestination *AWSSESConfigurationSetEventDestination_KinesisFirehoseDestination `json:"KinesisFirehoseDestination,omitempty"`

	// MatchingEventTypes AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationseteventdestination-eventdestination.html#cfn-ses-configurationseteventdestination-eventdestination-matchingeventtypes
	MatchingEventTypes []*String `json:"MatchingEventTypes,omitempty"`

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationseteventdestination-eventdestination.html#cfn-ses-configurationseteventdestination-eventdestination-name
	Name *String `json:"Name,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSSESConfigurationSetEventDestination_EventDestination) AddDependencies(v ...string) *AWSSESConfigurationSetEventDestination_EventDestination {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSSESConfigurationSetEventDestination_EventDestination) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSESConfigurationSetEventDestination_EventDestination) AWSCloudFormationType() string {
	return "AWS::SES::ConfigurationSetEventDestination.EventDestination"
}

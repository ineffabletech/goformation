package cloudformation

// AWSSESConfigurationSetEventDestination_KinesisFirehoseDestination AWS CloudFormation Resource (AWS::SES::ConfigurationSetEventDestination.KinesisFirehoseDestination)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationseteventdestination-kinesisfirehosedestination.html
type AWSSESConfigurationSetEventDestination_KinesisFirehoseDestination struct {
	dependsOn []string

	// DeliveryStreamARN AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationseteventdestination-kinesisfirehosedestination.html#cfn-ses-configurationseteventdestination-kinesisfirehosedestination-deliverystreamarn
	DeliveryStreamARN *String `json:"DeliveryStreamARN,omitempty"`

	// IAMRoleARN AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationseteventdestination-kinesisfirehosedestination.html#cfn-ses-configurationseteventdestination-kinesisfirehosedestination-iamrolearn
	IAMRoleARN *String `json:"IAMRoleARN,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSSESConfigurationSetEventDestination_KinesisFirehoseDestination) AddDependencies(v ...string) *AWSSESConfigurationSetEventDestination_KinesisFirehoseDestination {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSSESConfigurationSetEventDestination_KinesisFirehoseDestination) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSESConfigurationSetEventDestination_KinesisFirehoseDestination) AWSCloudFormationType() string {
	return "AWS::SES::ConfigurationSetEventDestination.KinesisFirehoseDestination"
}

package cloudformation

// AWSKinesisAnalyticsApplication_Input AWS CloudFormation Resource (AWS::KinesisAnalytics::Application.Input)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-input.html
type AWSKinesisAnalyticsApplication_Input struct {
	dependsOn []string

	// InputParallelism AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-input.html#cfn-kinesisanalytics-application-input-inputparallelism
	InputParallelism *AWSKinesisAnalyticsApplication_InputParallelism `json:"InputParallelism,omitempty"`

	// InputProcessingConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-input.html#cfn-kinesisanalytics-application-input-inputprocessingconfiguration
	InputProcessingConfiguration *AWSKinesisAnalyticsApplication_InputProcessingConfiguration `json:"InputProcessingConfiguration,omitempty"`

	// InputSchema AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-input.html#cfn-kinesisanalytics-application-input-inputschema
	InputSchema *AWSKinesisAnalyticsApplication_InputSchema `json:"InputSchema,omitempty"`

	// KinesisFirehoseInput AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-input.html#cfn-kinesisanalytics-application-input-kinesisfirehoseinput
	KinesisFirehoseInput *AWSKinesisAnalyticsApplication_KinesisFirehoseInput `json:"KinesisFirehoseInput,omitempty"`

	// KinesisStreamsInput AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-input.html#cfn-kinesisanalytics-application-input-kinesisstreamsinput
	KinesisStreamsInput *AWSKinesisAnalyticsApplication_KinesisStreamsInput `json:"KinesisStreamsInput,omitempty"`

	// NamePrefix AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-input.html#cfn-kinesisanalytics-application-input-nameprefix
	NamePrefix *String `json:"NamePrefix,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSKinesisAnalyticsApplication_Input) AddDependencies(v ...string) *AWSKinesisAnalyticsApplication_Input {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSKinesisAnalyticsApplication_Input) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsApplication_Input) AWSCloudFormationType() string {
	return "AWS::KinesisAnalytics::Application.Input"
}

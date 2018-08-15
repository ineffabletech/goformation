package cloudformation

// AWSKinesisAnalyticsApplicationOutput_Output AWS CloudFormation Resource (AWS::KinesisAnalytics::ApplicationOutput.Output)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationoutput-output.html
type AWSKinesisAnalyticsApplicationOutput_Output struct {
	dependsOn []string

	// DestinationSchema AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationoutput-output.html#cfn-kinesisanalytics-applicationoutput-output-destinationschema
	DestinationSchema *AWSKinesisAnalyticsApplicationOutput_DestinationSchema `json:"DestinationSchema,omitempty"`

	// KinesisFirehoseOutput AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationoutput-output.html#cfn-kinesisanalytics-applicationoutput-output-kinesisfirehoseoutput
	KinesisFirehoseOutput *AWSKinesisAnalyticsApplicationOutput_KinesisFirehoseOutput `json:"KinesisFirehoseOutput,omitempty"`

	// KinesisStreamsOutput AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationoutput-output.html#cfn-kinesisanalytics-applicationoutput-output-kinesisstreamsoutput
	KinesisStreamsOutput *AWSKinesisAnalyticsApplicationOutput_KinesisStreamsOutput `json:"KinesisStreamsOutput,omitempty"`

	// LambdaOutput AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationoutput-output.html#cfn-kinesisanalytics-applicationoutput-output-lambdaoutput
	LambdaOutput *AWSKinesisAnalyticsApplicationOutput_LambdaOutput `json:"LambdaOutput,omitempty"`

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationoutput-output.html#cfn-kinesisanalytics-applicationoutput-output-name
	Name *String `json:"Name,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSKinesisAnalyticsApplicationOutput_Output) AddDependencies(v ...string) *AWSKinesisAnalyticsApplicationOutput_Output {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSKinesisAnalyticsApplicationOutput_Output) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsApplicationOutput_Output) AWSCloudFormationType() string {
	return "AWS::KinesisAnalytics::ApplicationOutput.Output"
}

package cloudformation

// AWSKinesisAnalyticsApplicationOutput_KinesisFirehoseOutput AWS CloudFormation Resource (AWS::KinesisAnalytics::ApplicationOutput.KinesisFirehoseOutput)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationoutput-kinesisfirehoseoutput.html
type AWSKinesisAnalyticsApplicationOutput_KinesisFirehoseOutput struct {
	dependsOn []string

	// ResourceARN AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationoutput-kinesisfirehoseoutput.html#cfn-kinesisanalytics-applicationoutput-kinesisfirehoseoutput-resourcearn
	ResourceARN *String `json:"ResourceARN,omitempty"`

	// RoleARN AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationoutput-kinesisfirehoseoutput.html#cfn-kinesisanalytics-applicationoutput-kinesisfirehoseoutput-rolearn
	RoleARN *String `json:"RoleARN,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSKinesisAnalyticsApplicationOutput_KinesisFirehoseOutput) AddDependencies(v ...string) *AWSKinesisAnalyticsApplicationOutput_KinesisFirehoseOutput {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSKinesisAnalyticsApplicationOutput_KinesisFirehoseOutput) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsApplicationOutput_KinesisFirehoseOutput) AWSCloudFormationType() string {
	return "AWS::KinesisAnalytics::ApplicationOutput.KinesisFirehoseOutput"
}

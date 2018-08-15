package cloudformation

// AWSKinesisAnalyticsApplication_KinesisStreamsInput AWS CloudFormation Resource (AWS::KinesisAnalytics::Application.KinesisStreamsInput)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-kinesisstreamsinput.html
type AWSKinesisAnalyticsApplication_KinesisStreamsInput struct {
	dependsOn []string

	// ResourceARN AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-kinesisstreamsinput.html#cfn-kinesisanalytics-application-kinesisstreamsinput-resourcearn
	ResourceARN *String `json:"ResourceARN,omitempty"`

	// RoleARN AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-kinesisstreamsinput.html#cfn-kinesisanalytics-application-kinesisstreamsinput-rolearn
	RoleARN *String `json:"RoleARN,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSKinesisAnalyticsApplication_KinesisStreamsInput) AddDependencies(v ...string) *AWSKinesisAnalyticsApplication_KinesisStreamsInput {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSKinesisAnalyticsApplication_KinesisStreamsInput) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsApplication_KinesisStreamsInput) AWSCloudFormationType() string {
	return "AWS::KinesisAnalytics::Application.KinesisStreamsInput"
}

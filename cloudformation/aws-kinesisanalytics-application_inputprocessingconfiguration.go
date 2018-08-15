package cloudformation

// AWSKinesisAnalyticsApplication_InputProcessingConfiguration AWS CloudFormation Resource (AWS::KinesisAnalytics::Application.InputProcessingConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-inputprocessingconfiguration.html
type AWSKinesisAnalyticsApplication_InputProcessingConfiguration struct {
	dependsOn []string

	// InputLambdaProcessor AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-inputprocessingconfiguration.html#cfn-kinesisanalytics-application-inputprocessingconfiguration-inputlambdaprocessor
	InputLambdaProcessor *AWSKinesisAnalyticsApplication_InputLambdaProcessor `json:"InputLambdaProcessor,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSKinesisAnalyticsApplication_InputProcessingConfiguration) AddDependencies(v ...string) *AWSKinesisAnalyticsApplication_InputProcessingConfiguration {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSKinesisAnalyticsApplication_InputProcessingConfiguration) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsApplication_InputProcessingConfiguration) AWSCloudFormationType() string {
	return "AWS::KinesisAnalytics::Application.InputProcessingConfiguration"
}

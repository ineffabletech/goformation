package cloudformation

// AWSKinesisAnalyticsApplication_JSONMappingParameters AWS CloudFormation Resource (AWS::KinesisAnalytics::Application.JSONMappingParameters)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-jsonmappingparameters.html
type AWSKinesisAnalyticsApplication_JSONMappingParameters struct {
	dependsOn []string

	// RecordRowPath AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-jsonmappingparameters.html#cfn-kinesisanalytics-application-jsonmappingparameters-recordrowpath
	RecordRowPath *String `json:"RecordRowPath,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSKinesisAnalyticsApplication_JSONMappingParameters) AddDependencies(v ...string) *AWSKinesisAnalyticsApplication_JSONMappingParameters {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSKinesisAnalyticsApplication_JSONMappingParameters) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsApplication_JSONMappingParameters) AWSCloudFormationType() string {
	return "AWS::KinesisAnalytics::Application.JSONMappingParameters"
}

package cloudformation

// AWSAppSyncGraphQLApi_LogConfig AWS CloudFormation Resource (AWS::AppSync::GraphQLApi.LogConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-graphqlapi-logconfig.html
type AWSAppSyncGraphQLApi_LogConfig struct {
	dependsOn []string

	// CloudWatchLogsRoleArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-graphqlapi-logconfig.html#cfn-appsync-graphqlapi-logconfig-cloudwatchlogsrolearn
	CloudWatchLogsRoleArn *String `json:"CloudWatchLogsRoleArn,omitempty"`

	// FieldLogLevel AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-graphqlapi-logconfig.html#cfn-appsync-graphqlapi-logconfig-fieldloglevel
	FieldLogLevel *String `json:"FieldLogLevel,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSAppSyncGraphQLApi_LogConfig) AddDependencies(v ...string) *AWSAppSyncGraphQLApi_LogConfig {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSAppSyncGraphQLApi_LogConfig) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSAppSyncGraphQLApi_LogConfig) AWSCloudFormationType() string {
	return "AWS::AppSync::GraphQLApi.LogConfig"
}

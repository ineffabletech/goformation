package cloudformation

// AWSAppSyncDataSource_DynamoDBConfig AWS CloudFormation Resource (AWS::AppSync::DataSource.DynamoDBConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-dynamodbconfig.html
type AWSAppSyncDataSource_DynamoDBConfig struct {
	dependsOn []string

	// AwsRegion AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-dynamodbconfig.html#cfn-appsync-datasource-dynamodbconfig-awsregion
	AwsRegion *String `json:"AwsRegion,omitempty"`

	// TableName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-dynamodbconfig.html#cfn-appsync-datasource-dynamodbconfig-tablename
	TableName *String `json:"TableName,omitempty"`

	// UseCallerCredentials AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-dynamodbconfig.html#cfn-appsync-datasource-dynamodbconfig-usecallercredentials
	UseCallerCredentials *Boolean `json:"UseCallerCredentials,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSAppSyncDataSource_DynamoDBConfig) AddDependencies(v ...string) *AWSAppSyncDataSource_DynamoDBConfig {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSAppSyncDataSource_DynamoDBConfig) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSAppSyncDataSource_DynamoDBConfig) AWSCloudFormationType() string {
	return "AWS::AppSync::DataSource.DynamoDBConfig"
}

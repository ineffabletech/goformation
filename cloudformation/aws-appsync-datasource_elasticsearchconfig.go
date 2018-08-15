package cloudformation

// AWSAppSyncDataSource_ElasticsearchConfig AWS CloudFormation Resource (AWS::AppSync::DataSource.ElasticsearchConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-elasticsearchconfig.html
type AWSAppSyncDataSource_ElasticsearchConfig struct {
	dependsOn []string

	// AwsRegion AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-elasticsearchconfig.html#cfn-appsync-datasource-elasticsearchconfig-awsregion
	AwsRegion *String `json:"AwsRegion,omitempty"`

	// Endpoint AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-elasticsearchconfig.html#cfn-appsync-datasource-elasticsearchconfig-endpoint
	Endpoint *String `json:"Endpoint,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSAppSyncDataSource_ElasticsearchConfig) AddDependencies(v ...string) *AWSAppSyncDataSource_ElasticsearchConfig {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSAppSyncDataSource_ElasticsearchConfig) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSAppSyncDataSource_ElasticsearchConfig) AWSCloudFormationType() string {
	return "AWS::AppSync::DataSource.ElasticsearchConfig"
}

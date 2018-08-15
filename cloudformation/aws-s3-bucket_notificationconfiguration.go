package cloudformation

// AWSS3Bucket_NotificationConfiguration AWS CloudFormation Resource (AWS::S3::Bucket.NotificationConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig.html
type AWSS3Bucket_NotificationConfiguration struct {
	dependsOn []string

	// LambdaConfigurations AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig.html#cfn-s3-bucket-notificationconfig-lambdaconfig
	LambdaConfigurations []AWSS3Bucket_LambdaConfiguration `json:"LambdaConfigurations,omitempty"`

	// QueueConfigurations AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig.html#cfn-s3-bucket-notificationconfig-queueconfig
	QueueConfigurations []AWSS3Bucket_QueueConfiguration `json:"QueueConfigurations,omitempty"`

	// TopicConfigurations AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig.html#cfn-s3-bucket-notificationconfig-topicconfig
	TopicConfigurations []AWSS3Bucket_TopicConfiguration `json:"TopicConfigurations,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSS3Bucket_NotificationConfiguration) AddDependencies(v ...string) *AWSS3Bucket_NotificationConfiguration {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSS3Bucket_NotificationConfiguration) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3Bucket_NotificationConfiguration) AWSCloudFormationType() string {
	return "AWS::S3::Bucket.NotificationConfiguration"
}

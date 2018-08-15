package cloudformation

// AWSServerlessFunction_S3NotificationFilter AWS CloudFormation Resource (AWS::Serverless::Function.S3NotificationFilter)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfiguration-config-filter.html
type AWSServerlessFunction_S3NotificationFilter struct {
	dependsOn []string

	// S3Key AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfiguration-config-filter.html
	S3Key *String `json:"S3Key,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSServerlessFunction_S3NotificationFilter) AddDependencies(v ...string) *AWSServerlessFunction_S3NotificationFilter {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSServerlessFunction_S3NotificationFilter) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSServerlessFunction_S3NotificationFilter) AWSCloudFormationType() string {
	return "AWS::Serverless::Function.S3NotificationFilter"
}

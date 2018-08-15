package cloudformation

// AWSS3Bucket_S3KeyFilter AWS CloudFormation Resource (AWS::S3::Bucket.S3KeyFilter)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfiguration-config-filter-s3key.html
type AWSS3Bucket_S3KeyFilter struct {
	dependsOn []string

	// Rules AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfiguration-config-filter-s3key.html#cfn-s3-bucket-notificationconfiguraiton-config-filter-s3key-rules
	Rules []AWSS3Bucket_FilterRule `json:"Rules,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSS3Bucket_S3KeyFilter) AddDependencies(v ...string) *AWSS3Bucket_S3KeyFilter {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSS3Bucket_S3KeyFilter) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3Bucket_S3KeyFilter) AWSCloudFormationType() string {
	return "AWS::S3::Bucket.S3KeyFilter"
}

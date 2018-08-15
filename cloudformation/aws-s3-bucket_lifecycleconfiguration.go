package cloudformation

// AWSS3Bucket_LifecycleConfiguration AWS CloudFormation Resource (AWS::S3::Bucket.LifecycleConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig.html
type AWSS3Bucket_LifecycleConfiguration struct {
	dependsOn []string

	// Rules AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig.html#cfn-s3-bucket-lifecycleconfig-rules
	Rules []AWSS3Bucket_Rule `json:"Rules,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSS3Bucket_LifecycleConfiguration) AddDependencies(v ...string) *AWSS3Bucket_LifecycleConfiguration {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSS3Bucket_LifecycleConfiguration) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3Bucket_LifecycleConfiguration) AWSCloudFormationType() string {
	return "AWS::S3::Bucket.LifecycleConfiguration"
}

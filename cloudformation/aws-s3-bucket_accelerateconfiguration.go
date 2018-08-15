package cloudformation

// AWSS3Bucket_AccelerateConfiguration AWS CloudFormation Resource (AWS::S3::Bucket.AccelerateConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-accelerateconfiguration.html
type AWSS3Bucket_AccelerateConfiguration struct {
	dependsOn []string

	// AccelerationStatus AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-accelerateconfiguration.html#cfn-s3-bucket-accelerateconfiguration-accelerationstatus
	AccelerationStatus *String `json:"AccelerationStatus,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSS3Bucket_AccelerateConfiguration) AddDependencies(v ...string) *AWSS3Bucket_AccelerateConfiguration {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSS3Bucket_AccelerateConfiguration) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3Bucket_AccelerateConfiguration) AWSCloudFormationType() string {
	return "AWS::S3::Bucket.AccelerateConfiguration"
}

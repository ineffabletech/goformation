package cloudformation

// AWSS3Bucket_SseKmsEncryptedObjects AWS CloudFormation Resource (AWS::S3::Bucket.SseKmsEncryptedObjects)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-ssekmsencryptedobjects.html
type AWSS3Bucket_SseKmsEncryptedObjects struct {
	dependsOn []string

	// Status AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-ssekmsencryptedobjects.html#cfn-s3-bucket-ssekmsencryptedobjects-status
	Status *String `json:"Status,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSS3Bucket_SseKmsEncryptedObjects) AddDependencies(v ...string) *AWSS3Bucket_SseKmsEncryptedObjects {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSS3Bucket_SseKmsEncryptedObjects) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3Bucket_SseKmsEncryptedObjects) AWSCloudFormationType() string {
	return "AWS::S3::Bucket.SseKmsEncryptedObjects"
}

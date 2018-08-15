package cloudformation

// AWSS3Bucket_SourceSelectionCriteria AWS CloudFormation Resource (AWS::S3::Bucket.SourceSelectionCriteria)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-sourceselectioncriteria.html
type AWSS3Bucket_SourceSelectionCriteria struct {
	dependsOn []string

	// SseKmsEncryptedObjects AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-sourceselectioncriteria.html#cfn-s3-bucket-sourceselectioncriteria-ssekmsencryptedobjects
	SseKmsEncryptedObjects *AWSS3Bucket_SseKmsEncryptedObjects `json:"SseKmsEncryptedObjects,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSS3Bucket_SourceSelectionCriteria) AddDependencies(v ...string) *AWSS3Bucket_SourceSelectionCriteria {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSS3Bucket_SourceSelectionCriteria) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3Bucket_SourceSelectionCriteria) AWSCloudFormationType() string {
	return "AWS::S3::Bucket.SourceSelectionCriteria"
}

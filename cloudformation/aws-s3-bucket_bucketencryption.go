package cloudformation

// AWSS3Bucket_BucketEncryption AWS CloudFormation Resource (AWS::S3::Bucket.BucketEncryption)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-bucketencryption.html
type AWSS3Bucket_BucketEncryption struct {
	dependsOn []string

	// ServerSideEncryptionConfiguration AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-bucketencryption.html#cfn-s3-bucket-bucketencryption-serversideencryptionconfiguration
	ServerSideEncryptionConfiguration []AWSS3Bucket_ServerSideEncryptionRule `json:"ServerSideEncryptionConfiguration,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSS3Bucket_BucketEncryption) AddDependencies(v ...string) *AWSS3Bucket_BucketEncryption {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSS3Bucket_BucketEncryption) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3Bucket_BucketEncryption) AWSCloudFormationType() string {
	return "AWS::S3::Bucket.BucketEncryption"
}

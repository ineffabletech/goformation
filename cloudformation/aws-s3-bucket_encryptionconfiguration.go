package cloudformation

// AWSS3Bucket_EncryptionConfiguration AWS CloudFormation Resource (AWS::S3::Bucket.EncryptionConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-encryptionconfiguration.html
type AWSS3Bucket_EncryptionConfiguration struct {
	dependsOn []string

	// ReplicaKmsKeyID AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-encryptionconfiguration.html#cfn-s3-bucket-encryptionconfiguration-replicakmskeyid
	ReplicaKmsKeyID *String `json:"ReplicaKmsKeyID,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSS3Bucket_EncryptionConfiguration) AddDependencies(v ...string) *AWSS3Bucket_EncryptionConfiguration {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSS3Bucket_EncryptionConfiguration) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3Bucket_EncryptionConfiguration) AWSCloudFormationType() string {
	return "AWS::S3::Bucket.EncryptionConfiguration"
}

package cloudformation

// AWSS3Bucket_VersioningConfiguration AWS CloudFormation Resource (AWS::S3::Bucket.VersioningConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-versioningconfig.html
type AWSS3Bucket_VersioningConfiguration struct {
	dependsOn []string

	// Status AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-versioningconfig.html#cfn-s3-bucket-versioningconfig-status
	Status *String `json:"Status,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSS3Bucket_VersioningConfiguration) AddDependencies(v ...string) *AWSS3Bucket_VersioningConfiguration {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSS3Bucket_VersioningConfiguration) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3Bucket_VersioningConfiguration) AWSCloudFormationType() string {
	return "AWS::S3::Bucket.VersioningConfiguration"
}

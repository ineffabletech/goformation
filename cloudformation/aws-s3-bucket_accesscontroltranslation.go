package cloudformation

// AWSS3Bucket_AccessControlTranslation AWS CloudFormation Resource (AWS::S3::Bucket.AccessControlTranslation)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-accesscontroltranslation.html
type AWSS3Bucket_AccessControlTranslation struct {
	dependsOn []string

	// Owner AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-accesscontroltranslation.html#cfn-s3-bucket-accesscontroltranslation-owner
	Owner *String `json:"Owner,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSS3Bucket_AccessControlTranslation) AddDependencies(v ...string) *AWSS3Bucket_AccessControlTranslation {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSS3Bucket_AccessControlTranslation) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3Bucket_AccessControlTranslation) AWSCloudFormationType() string {
	return "AWS::S3::Bucket.AccessControlTranslation"
}

package cloudformation

// AWSS3Bucket_ServerSideEncryptionRule AWS CloudFormation Resource (AWS::S3::Bucket.ServerSideEncryptionRule)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-serversideencryptionrule.html
type AWSS3Bucket_ServerSideEncryptionRule struct {
	dependsOn []string

	// ServerSideEncryptionByDefault AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-serversideencryptionrule.html#cfn-s3-bucket-serversideencryptionrule-serversideencryptionbydefault
	ServerSideEncryptionByDefault *AWSS3Bucket_ServerSideEncryptionByDefault `json:"ServerSideEncryptionByDefault,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSS3Bucket_ServerSideEncryptionRule) AddDependencies(v ...string) *AWSS3Bucket_ServerSideEncryptionRule {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSS3Bucket_ServerSideEncryptionRule) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3Bucket_ServerSideEncryptionRule) AWSCloudFormationType() string {
	return "AWS::S3::Bucket.ServerSideEncryptionRule"
}

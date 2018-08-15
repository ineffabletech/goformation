package cloudformation

// AWSS3Bucket_StorageClassAnalysis AWS CloudFormation Resource (AWS::S3::Bucket.StorageClassAnalysis)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-storageclassanalysis.html
type AWSS3Bucket_StorageClassAnalysis struct {
	dependsOn []string

	// DataExport AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-storageclassanalysis.html#cfn-s3-bucket-storageclassanalysis-dataexport
	DataExport *AWSS3Bucket_DataExport `json:"DataExport,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSS3Bucket_StorageClassAnalysis) AddDependencies(v ...string) *AWSS3Bucket_StorageClassAnalysis {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSS3Bucket_StorageClassAnalysis) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3Bucket_StorageClassAnalysis) AWSCloudFormationType() string {
	return "AWS::S3::Bucket.StorageClassAnalysis"
}

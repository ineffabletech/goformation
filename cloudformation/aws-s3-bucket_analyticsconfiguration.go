package cloudformation

// AWSS3Bucket_AnalyticsConfiguration AWS CloudFormation Resource (AWS::S3::Bucket.AnalyticsConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-analyticsconfiguration.html
type AWSS3Bucket_AnalyticsConfiguration struct {
	dependsOn []string

	// Id AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-analyticsconfiguration.html#cfn-s3-bucket-analyticsconfiguration-id
	Id *String `json:"Id,omitempty"`

	// Prefix AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-analyticsconfiguration.html#cfn-s3-bucket-analyticsconfiguration-prefix
	Prefix *String `json:"Prefix,omitempty"`

	// StorageClassAnalysis AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-analyticsconfiguration.html#cfn-s3-bucket-analyticsconfiguration-storageclassanalysis
	StorageClassAnalysis *AWSS3Bucket_StorageClassAnalysis `json:"StorageClassAnalysis,omitempty"`

	// TagFilters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-analyticsconfiguration.html#cfn-s3-bucket-analyticsconfiguration-tagfilters
	TagFilters []AWSS3Bucket_TagFilter `json:"TagFilters,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSS3Bucket_AnalyticsConfiguration) AddDependencies(v ...string) *AWSS3Bucket_AnalyticsConfiguration {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSS3Bucket_AnalyticsConfiguration) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3Bucket_AnalyticsConfiguration) AWSCloudFormationType() string {
	return "AWS::S3::Bucket.AnalyticsConfiguration"
}

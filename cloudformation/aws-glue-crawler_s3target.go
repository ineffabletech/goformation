package cloudformation

// AWSGlueCrawler_S3Target AWS CloudFormation Resource (AWS::Glue::Crawler.S3Target)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-s3target.html
type AWSGlueCrawler_S3Target struct {
	dependsOn []string

	// Exclusions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-s3target.html#cfn-glue-crawler-s3target-exclusions
	Exclusions []*String `json:"Exclusions,omitempty"`

	// Path AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-s3target.html#cfn-glue-crawler-s3target-path
	Path *String `json:"Path,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSGlueCrawler_S3Target) AddDependencies(v ...string) *AWSGlueCrawler_S3Target {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSGlueCrawler_S3Target) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSGlueCrawler_S3Target) AWSCloudFormationType() string {
	return "AWS::Glue::Crawler.S3Target"
}

package cloudformation

// AWSGlueCrawler_Targets AWS CloudFormation Resource (AWS::Glue::Crawler.Targets)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-targets.html
type AWSGlueCrawler_Targets struct {
	dependsOn []string

	// JdbcTargets AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-targets.html#cfn-glue-crawler-targets-jdbctargets
	JdbcTargets []AWSGlueCrawler_JdbcTarget `json:"JdbcTargets,omitempty"`

	// S3Targets AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-targets.html#cfn-glue-crawler-targets-s3targets
	S3Targets []AWSGlueCrawler_S3Target `json:"S3Targets,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSGlueCrawler_Targets) AddDependencies(v ...string) *AWSGlueCrawler_Targets {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSGlueCrawler_Targets) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSGlueCrawler_Targets) AWSCloudFormationType() string {
	return "AWS::Glue::Crawler.Targets"
}

package cloudformation

// AWSGlueCrawler_JdbcTarget AWS CloudFormation Resource (AWS::Glue::Crawler.JdbcTarget)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-jdbctarget.html
type AWSGlueCrawler_JdbcTarget struct {
	dependsOn []string

	// ConnectionName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-jdbctarget.html#cfn-glue-crawler-jdbctarget-connectionname
	ConnectionName *String `json:"ConnectionName,omitempty"`

	// Exclusions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-jdbctarget.html#cfn-glue-crawler-jdbctarget-exclusions
	Exclusions []*String `json:"Exclusions,omitempty"`

	// Path AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-jdbctarget.html#cfn-glue-crawler-jdbctarget-path
	Path *String `json:"Path,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSGlueCrawler_JdbcTarget) AddDependencies(v ...string) *AWSGlueCrawler_JdbcTarget {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSGlueCrawler_JdbcTarget) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSGlueCrawler_JdbcTarget) AWSCloudFormationType() string {
	return "AWS::Glue::Crawler.JdbcTarget"
}

package cloudformation

// AWSGlueCrawler_SchemaChangePolicy AWS CloudFormation Resource (AWS::Glue::Crawler.SchemaChangePolicy)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-schemachangepolicy.html
type AWSGlueCrawler_SchemaChangePolicy struct {
	dependsOn []string

	// DeleteBehavior AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-schemachangepolicy.html#cfn-glue-crawler-schemachangepolicy-deletebehavior
	DeleteBehavior *String `json:"DeleteBehavior,omitempty"`

	// UpdateBehavior AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-schemachangepolicy.html#cfn-glue-crawler-schemachangepolicy-updatebehavior
	UpdateBehavior *String `json:"UpdateBehavior,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSGlueCrawler_SchemaChangePolicy) AddDependencies(v ...string) *AWSGlueCrawler_SchemaChangePolicy {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSGlueCrawler_SchemaChangePolicy) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSGlueCrawler_SchemaChangePolicy) AWSCloudFormationType() string {
	return "AWS::Glue::Crawler.SchemaChangePolicy"
}

package cloudformation

// AWSGlueCrawler_Schedule AWS CloudFormation Resource (AWS::Glue::Crawler.Schedule)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-schedule.html
type AWSGlueCrawler_Schedule struct {
	dependsOn []string

	// ScheduleExpression AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-schedule.html#cfn-glue-crawler-schedule-scheduleexpression
	ScheduleExpression *String `json:"ScheduleExpression,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSGlueCrawler_Schedule) AddDependencies(v ...string) *AWSGlueCrawler_Schedule {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSGlueCrawler_Schedule) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSGlueCrawler_Schedule) AWSCloudFormationType() string {
	return "AWS::Glue::Crawler.Schedule"
}

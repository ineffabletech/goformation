package cloudformation

// AWSLogsMetricFilter_MetricTransformation AWS CloudFormation Resource (AWS::Logs::MetricFilter.MetricTransformation)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-metricfilter-metrictransformation.html
type AWSLogsMetricFilter_MetricTransformation struct {
	dependsOn []string

	// DefaultValue AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-metricfilter-metrictransformation.html#cfn-cwl-metricfilter-metrictransformation-defaultvalue
	DefaultValue *Double `json:"DefaultValue,omitempty"`

	// MetricName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-metricfilter-metrictransformation.html#cfn-cwl-metricfilter-metrictransformation-metricname
	MetricName *String `json:"MetricName,omitempty"`

	// MetricNamespace AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-metricfilter-metrictransformation.html#cfn-cwl-metricfilter-metrictransformation-metricnamespace
	MetricNamespace *String `json:"MetricNamespace,omitempty"`

	// MetricValue AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-metricfilter-metrictransformation.html#cfn-cwl-metricfilter-metrictransformation-metricvalue
	MetricValue *String `json:"MetricValue,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSLogsMetricFilter_MetricTransformation) AddDependencies(v ...string) *AWSLogsMetricFilter_MetricTransformation {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSLogsMetricFilter_MetricTransformation) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSLogsMetricFilter_MetricTransformation) AWSCloudFormationType() string {
	return "AWS::Logs::MetricFilter.MetricTransformation"
}

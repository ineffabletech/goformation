package cloudformation

// AWSAutoScalingPlansScalingPlan_ApplicationSource AWS CloudFormation Resource (AWS::AutoScalingPlans::ScalingPlan.ApplicationSource)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-applicationsource.html
type AWSAutoScalingPlansScalingPlan_ApplicationSource struct {
	dependsOn []string

	// CloudFormationStackARN AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-applicationsource.html#cfn-autoscalingplans-scalingplan-applicationsource-cloudformationstackarn
	CloudFormationStackARN *String `json:"CloudFormationStackARN,omitempty"`

	// TagFilters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-applicationsource.html#cfn-autoscalingplans-scalingplan-applicationsource-tagfilters
	TagFilters []AWSAutoScalingPlansScalingPlan_TagFilter `json:"TagFilters,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSAutoScalingPlansScalingPlan_ApplicationSource) AddDependencies(v ...string) *AWSAutoScalingPlansScalingPlan_ApplicationSource {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSAutoScalingPlansScalingPlan_ApplicationSource) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSAutoScalingPlansScalingPlan_ApplicationSource) AWSCloudFormationType() string {
	return "AWS::AutoScalingPlans::ScalingPlan.ApplicationSource"
}

package cloudformation

// AWSAutoScalingPlansScalingPlan_PredefinedScalingMetricSpecification AWS CloudFormation Resource (AWS::AutoScalingPlans::ScalingPlan.PredefinedScalingMetricSpecification)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-predefinedscalingmetricspecification.html
type AWSAutoScalingPlansScalingPlan_PredefinedScalingMetricSpecification struct {
	dependsOn []string

	// PredefinedScalingMetricType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-predefinedscalingmetricspecification.html#cfn-autoscalingplans-scalingplan-predefinedscalingmetricspecification-predefinedscalingmetrictype
	PredefinedScalingMetricType *String `json:"PredefinedScalingMetricType,omitempty"`

	// ResourceLabel AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-predefinedscalingmetricspecification.html#cfn-autoscalingplans-scalingplan-predefinedscalingmetricspecification-resourcelabel
	ResourceLabel *String `json:"ResourceLabel,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSAutoScalingPlansScalingPlan_PredefinedScalingMetricSpecification) AddDependencies(v ...string) *AWSAutoScalingPlansScalingPlan_PredefinedScalingMetricSpecification {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSAutoScalingPlansScalingPlan_PredefinedScalingMetricSpecification) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSAutoScalingPlansScalingPlan_PredefinedScalingMetricSpecification) AWSCloudFormationType() string {
	return "AWS::AutoScalingPlans::ScalingPlan.PredefinedScalingMetricSpecification"
}

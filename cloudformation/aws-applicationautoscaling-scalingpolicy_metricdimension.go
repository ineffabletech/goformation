package cloudformation

// AWSApplicationAutoScalingScalingPolicy_MetricDimension AWS CloudFormation Resource (AWS::ApplicationAutoScaling::ScalingPolicy.MetricDimension)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-metricdimension.html
type AWSApplicationAutoScalingScalingPolicy_MetricDimension struct {
	dependsOn []string

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-metricdimension.html#cfn-applicationautoscaling-scalingpolicy-metricdimension-name
	Name *String `json:"Name,omitempty"`

	// Value AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-metricdimension.html#cfn-applicationautoscaling-scalingpolicy-metricdimension-value
	Value *String `json:"Value,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSApplicationAutoScalingScalingPolicy_MetricDimension) AddDependencies(v ...string) *AWSApplicationAutoScalingScalingPolicy_MetricDimension {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSApplicationAutoScalingScalingPolicy_MetricDimension) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApplicationAutoScalingScalingPolicy_MetricDimension) AWSCloudFormationType() string {
	return "AWS::ApplicationAutoScaling::ScalingPolicy.MetricDimension"
}

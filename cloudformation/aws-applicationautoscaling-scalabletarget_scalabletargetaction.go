package cloudformation

// AWSApplicationAutoScalingScalableTarget_ScalableTargetAction AWS CloudFormation Resource (AWS::ApplicationAutoScaling::ScalableTarget.ScalableTargetAction)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalabletarget-scalabletargetaction.html
type AWSApplicationAutoScalingScalableTarget_ScalableTargetAction struct {
	dependsOn []string

	// MaxCapacity AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalabletarget-scalabletargetaction.html#cfn-applicationautoscaling-scalabletarget-scalabletargetaction-maxcapacity
	MaxCapacity *Integer `json:"MaxCapacity,omitempty"`

	// MinCapacity AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalabletarget-scalabletargetaction.html#cfn-applicationautoscaling-scalabletarget-scalabletargetaction-mincapacity
	MinCapacity *Integer `json:"MinCapacity,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSApplicationAutoScalingScalableTarget_ScalableTargetAction) AddDependencies(v ...string) *AWSApplicationAutoScalingScalableTarget_ScalableTargetAction {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSApplicationAutoScalingScalableTarget_ScalableTargetAction) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApplicationAutoScalingScalableTarget_ScalableTargetAction) AWSCloudFormationType() string {
	return "AWS::ApplicationAutoScaling::ScalableTarget.ScalableTargetAction"
}

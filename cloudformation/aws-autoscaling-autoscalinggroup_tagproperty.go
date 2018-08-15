package cloudformation

// AWSAutoScalingAutoScalingGroup_TagProperty AWS CloudFormation Resource (AWS::AutoScaling::AutoScalingGroup.TagProperty)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-tags.html
type AWSAutoScalingAutoScalingGroup_TagProperty struct {
	dependsOn []string

	// Key AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-tags.html#cfn-as-tags-Key
	Key *String `json:"Key,omitempty"`

	// PropagateAtLaunch AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-tags.html#cfn-as-tags-PropagateAtLaunch
	PropagateAtLaunch *Boolean `json:"PropagateAtLaunch,omitempty"`

	// Value AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-tags.html#cfn-as-tags-Value
	Value *String `json:"Value,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSAutoScalingAutoScalingGroup_TagProperty) AddDependencies(v ...string) *AWSAutoScalingAutoScalingGroup_TagProperty {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSAutoScalingAutoScalingGroup_TagProperty) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSAutoScalingAutoScalingGroup_TagProperty) AWSCloudFormationType() string {
	return "AWS::AutoScaling::AutoScalingGroup.TagProperty"
}

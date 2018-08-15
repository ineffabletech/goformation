package cloudformation

// AWSEventsRule_RunCommandTarget AWS CloudFormation Resource (AWS::Events::Rule.RunCommandTarget)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-runcommandtarget.html
type AWSEventsRule_RunCommandTarget struct {
	dependsOn []string

	// Key AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-runcommandtarget.html#cfn-events-rule-runcommandtarget-key
	Key *String `json:"Key,omitempty"`

	// Values AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-runcommandtarget.html#cfn-events-rule-runcommandtarget-values
	Values []*String `json:"Values,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSEventsRule_RunCommandTarget) AddDependencies(v ...string) *AWSEventsRule_RunCommandTarget {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSEventsRule_RunCommandTarget) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEventsRule_RunCommandTarget) AWSCloudFormationType() string {
	return "AWS::Events::Rule.RunCommandTarget"
}

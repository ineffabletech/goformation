package cloudformation

// AWSEventsRule_RunCommandParameters AWS CloudFormation Resource (AWS::Events::Rule.RunCommandParameters)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-runcommandparameters.html
type AWSEventsRule_RunCommandParameters struct {
	dependsOn []string

	// RunCommandTargets AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-runcommandparameters.html#cfn-events-rule-runcommandparameters-runcommandtargets
	RunCommandTargets []AWSEventsRule_RunCommandTarget `json:"RunCommandTargets,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSEventsRule_RunCommandParameters) AddDependencies(v ...string) *AWSEventsRule_RunCommandParameters {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSEventsRule_RunCommandParameters) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEventsRule_RunCommandParameters) AWSCloudFormationType() string {
	return "AWS::Events::Rule.RunCommandParameters"
}

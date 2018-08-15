package cloudformation

// AWSEventsRule_EcsParameters AWS CloudFormation Resource (AWS::Events::Rule.EcsParameters)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-ecsparameters.html
type AWSEventsRule_EcsParameters struct {
	dependsOn []string

	// TaskCount AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-ecsparameters.html#cfn-events-rule-ecsparameters-taskcount
	TaskCount *Integer `json:"TaskCount,omitempty"`

	// TaskDefinitionArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-ecsparameters.html#cfn-events-rule-ecsparameters-taskdefinitionarn
	TaskDefinitionArn *String `json:"TaskDefinitionArn,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSEventsRule_EcsParameters) AddDependencies(v ...string) *AWSEventsRule_EcsParameters {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSEventsRule_EcsParameters) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEventsRule_EcsParameters) AWSCloudFormationType() string {
	return "AWS::Events::Rule.EcsParameters"
}

package cloudformation

// AWSGlueTrigger_Action AWS CloudFormation Resource (AWS::Glue::Trigger.Action)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-action.html
type AWSGlueTrigger_Action struct {
	dependsOn []string

	// Arguments AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-action.html#cfn-glue-trigger-action-arguments
	Arguments interface{} `json:"Arguments,omitempty"`

	// JobName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-action.html#cfn-glue-trigger-action-jobname
	JobName *String `json:"JobName,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSGlueTrigger_Action) AddDependencies(v ...string) *AWSGlueTrigger_Action {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSGlueTrigger_Action) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSGlueTrigger_Action) AWSCloudFormationType() string {
	return "AWS::Glue::Trigger.Action"
}

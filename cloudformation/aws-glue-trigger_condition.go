package cloudformation

// AWSGlueTrigger_Condition AWS CloudFormation Resource (AWS::Glue::Trigger.Condition)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-condition.html
type AWSGlueTrigger_Condition struct {
	dependsOn []string

	// JobName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-condition.html#cfn-glue-trigger-condition-jobname
	JobName *String `json:"JobName,omitempty"`

	// LogicalOperator AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-condition.html#cfn-glue-trigger-condition-logicaloperator
	LogicalOperator *String `json:"LogicalOperator,omitempty"`

	// State AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-condition.html#cfn-glue-trigger-condition-state
	State *String `json:"State,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSGlueTrigger_Condition) AddDependencies(v ...string) *AWSGlueTrigger_Condition {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSGlueTrigger_Condition) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSGlueTrigger_Condition) AWSCloudFormationType() string {
	return "AWS::Glue::Trigger.Condition"
}

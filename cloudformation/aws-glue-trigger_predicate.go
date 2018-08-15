package cloudformation

// AWSGlueTrigger_Predicate AWS CloudFormation Resource (AWS::Glue::Trigger.Predicate)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-predicate.html
type AWSGlueTrigger_Predicate struct {
	dependsOn []string

	// Conditions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-predicate.html#cfn-glue-trigger-predicate-conditions
	Conditions []AWSGlueTrigger_Condition `json:"Conditions,omitempty"`

	// Logical AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-predicate.html#cfn-glue-trigger-predicate-logical
	Logical *String `json:"Logical,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSGlueTrigger_Predicate) AddDependencies(v ...string) *AWSGlueTrigger_Predicate {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSGlueTrigger_Predicate) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSGlueTrigger_Predicate) AWSCloudFormationType() string {
	return "AWS::Glue::Trigger.Predicate"
}

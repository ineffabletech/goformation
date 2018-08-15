package cloudformation

// AWSBudgetsBudget_TimePeriod AWS CloudFormation Resource (AWS::Budgets::Budget.TimePeriod)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-timeperiod.html
type AWSBudgetsBudget_TimePeriod struct {
	dependsOn []string

	// End AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-timeperiod.html#cfn-budgets-budget-timeperiod-end
	End *String `json:"End,omitempty"`

	// Start AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-timeperiod.html#cfn-budgets-budget-timeperiod-start
	Start *String `json:"Start,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSBudgetsBudget_TimePeriod) AddDependencies(v ...string) *AWSBudgetsBudget_TimePeriod {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSBudgetsBudget_TimePeriod) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSBudgetsBudget_TimePeriod) AWSCloudFormationType() string {
	return "AWS::Budgets::Budget.TimePeriod"
}

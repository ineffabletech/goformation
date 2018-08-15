package cloudformation

// AWSBudgetsBudget_CostTypes AWS CloudFormation Resource (AWS::Budgets::Budget.CostTypes)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-costtypes.html
type AWSBudgetsBudget_CostTypes struct {
	dependsOn []string

	// IncludeCredit AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-costtypes.html#cfn-budgets-budget-costtypes-includecredit
	IncludeCredit *Boolean `json:"IncludeCredit,omitempty"`

	// IncludeDiscount AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-costtypes.html#cfn-budgets-budget-costtypes-includediscount
	IncludeDiscount *Boolean `json:"IncludeDiscount,omitempty"`

	// IncludeOtherSubscription AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-costtypes.html#cfn-budgets-budget-costtypes-includeothersubscription
	IncludeOtherSubscription *Boolean `json:"IncludeOtherSubscription,omitempty"`

	// IncludeRecurring AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-costtypes.html#cfn-budgets-budget-costtypes-includerecurring
	IncludeRecurring *Boolean `json:"IncludeRecurring,omitempty"`

	// IncludeRefund AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-costtypes.html#cfn-budgets-budget-costtypes-includerefund
	IncludeRefund *Boolean `json:"IncludeRefund,omitempty"`

	// IncludeSubscription AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-costtypes.html#cfn-budgets-budget-costtypes-includesubscription
	IncludeSubscription *Boolean `json:"IncludeSubscription,omitempty"`

	// IncludeSupport AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-costtypes.html#cfn-budgets-budget-costtypes-includesupport
	IncludeSupport *Boolean `json:"IncludeSupport,omitempty"`

	// IncludeTax AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-costtypes.html#cfn-budgets-budget-costtypes-includetax
	IncludeTax *Boolean `json:"IncludeTax,omitempty"`

	// IncludeUpfront AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-costtypes.html#cfn-budgets-budget-costtypes-includeupfront
	IncludeUpfront *Boolean `json:"IncludeUpfront,omitempty"`

	// UseAmortized AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-costtypes.html#cfn-budgets-budget-costtypes-useamortized
	UseAmortized *Boolean `json:"UseAmortized,omitempty"`

	// UseBlended AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-costtypes.html#cfn-budgets-budget-costtypes-useblended
	UseBlended *Boolean `json:"UseBlended,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSBudgetsBudget_CostTypes) AddDependencies(v ...string) *AWSBudgetsBudget_CostTypes {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSBudgetsBudget_CostTypes) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSBudgetsBudget_CostTypes) AWSCloudFormationType() string {
	return "AWS::Budgets::Budget.CostTypes"
}

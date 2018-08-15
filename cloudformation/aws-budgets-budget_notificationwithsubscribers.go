package cloudformation

// AWSBudgetsBudget_NotificationWithSubscribers AWS CloudFormation Resource (AWS::Budgets::Budget.NotificationWithSubscribers)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-notificationwithsubscribers.html
type AWSBudgetsBudget_NotificationWithSubscribers struct {
	dependsOn []string

	// Notification AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-notificationwithsubscribers.html#cfn-budgets-budget-notificationwithsubscribers-notification
	Notification *AWSBudgetsBudget_Notification `json:"Notification,omitempty"`

	// Subscribers AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-notificationwithsubscribers.html#cfn-budgets-budget-notificationwithsubscribers-subscribers
	Subscribers []AWSBudgetsBudget_Subscriber `json:"Subscribers,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSBudgetsBudget_NotificationWithSubscribers) AddDependencies(v ...string) *AWSBudgetsBudget_NotificationWithSubscribers {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSBudgetsBudget_NotificationWithSubscribers) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSBudgetsBudget_NotificationWithSubscribers) AWSCloudFormationType() string {
	return "AWS::Budgets::Budget.NotificationWithSubscribers"
}

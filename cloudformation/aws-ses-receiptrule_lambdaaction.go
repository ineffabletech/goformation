package cloudformation

// AWSSESReceiptRule_LambdaAction AWS CloudFormation Resource (AWS::SES::ReceiptRule.LambdaAction)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-lambdaaction.html
type AWSSESReceiptRule_LambdaAction struct {
	dependsOn []string

	// FunctionArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-lambdaaction.html#cfn-ses-receiptrule-lambdaaction-functionarn
	FunctionArn *String `json:"FunctionArn,omitempty"`

	// InvocationType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-lambdaaction.html#cfn-ses-receiptrule-lambdaaction-invocationtype
	InvocationType *String `json:"InvocationType,omitempty"`

	// TopicArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-lambdaaction.html#cfn-ses-receiptrule-lambdaaction-topicarn
	TopicArn *String `json:"TopicArn,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSSESReceiptRule_LambdaAction) AddDependencies(v ...string) *AWSSESReceiptRule_LambdaAction {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSSESReceiptRule_LambdaAction) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSESReceiptRule_LambdaAction) AWSCloudFormationType() string {
	return "AWS::SES::ReceiptRule.LambdaAction"
}

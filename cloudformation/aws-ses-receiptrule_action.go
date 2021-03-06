package cloudformation

// AWSSESReceiptRule_Action AWS CloudFormation Resource (AWS::SES::ReceiptRule.Action)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-action.html
type AWSSESReceiptRule_Action struct {
	dependsOn []string

	// AddHeaderAction AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-action.html#cfn-ses-receiptrule-action-addheaderaction
	AddHeaderAction *AWSSESReceiptRule_AddHeaderAction `json:"AddHeaderAction,omitempty"`

	// BounceAction AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-action.html#cfn-ses-receiptrule-action-bounceaction
	BounceAction *AWSSESReceiptRule_BounceAction `json:"BounceAction,omitempty"`

	// LambdaAction AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-action.html#cfn-ses-receiptrule-action-lambdaaction
	LambdaAction *AWSSESReceiptRule_LambdaAction `json:"LambdaAction,omitempty"`

	// S3Action AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-action.html#cfn-ses-receiptrule-action-s3action
	S3Action *AWSSESReceiptRule_S3Action `json:"S3Action,omitempty"`

	// SNSAction AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-action.html#cfn-ses-receiptrule-action-snsaction
	SNSAction *AWSSESReceiptRule_SNSAction `json:"SNSAction,omitempty"`

	// StopAction AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-action.html#cfn-ses-receiptrule-action-stopaction
	StopAction *AWSSESReceiptRule_StopAction `json:"StopAction,omitempty"`

	// WorkmailAction AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-action.html#cfn-ses-receiptrule-action-workmailaction
	WorkmailAction *AWSSESReceiptRule_WorkmailAction `json:"WorkmailAction,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSSESReceiptRule_Action) AddDependencies(v ...string) *AWSSESReceiptRule_Action {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSSESReceiptRule_Action) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSESReceiptRule_Action) AWSCloudFormationType() string {
	return "AWS::SES::ReceiptRule.Action"
}

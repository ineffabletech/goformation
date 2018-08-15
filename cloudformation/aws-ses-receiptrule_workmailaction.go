package cloudformation

// AWSSESReceiptRule_WorkmailAction AWS CloudFormation Resource (AWS::SES::ReceiptRule.WorkmailAction)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-workmailaction.html
type AWSSESReceiptRule_WorkmailAction struct {
	dependsOn []string

	// OrganizationArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-workmailaction.html#cfn-ses-receiptrule-workmailaction-organizationarn
	OrganizationArn *String `json:"OrganizationArn,omitempty"`

	// TopicArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-workmailaction.html#cfn-ses-receiptrule-workmailaction-topicarn
	TopicArn *String `json:"TopicArn,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSSESReceiptRule_WorkmailAction) AddDependencies(v ...string) *AWSSESReceiptRule_WorkmailAction {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSSESReceiptRule_WorkmailAction) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSESReceiptRule_WorkmailAction) AWSCloudFormationType() string {
	return "AWS::SES::ReceiptRule.WorkmailAction"
}

package cloudformation

// AWSSESReceiptRule_AddHeaderAction AWS CloudFormation Resource (AWS::SES::ReceiptRule.AddHeaderAction)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-addheaderaction.html
type AWSSESReceiptRule_AddHeaderAction struct {
	dependsOn []string

	// HeaderName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-addheaderaction.html#cfn-ses-receiptrule-addheaderaction-headername
	HeaderName *String `json:"HeaderName,omitempty"`

	// HeaderValue AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-addheaderaction.html#cfn-ses-receiptrule-addheaderaction-headervalue
	HeaderValue *String `json:"HeaderValue,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSSESReceiptRule_AddHeaderAction) AddDependencies(v ...string) *AWSSESReceiptRule_AddHeaderAction {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSSESReceiptRule_AddHeaderAction) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSESReceiptRule_AddHeaderAction) AWSCloudFormationType() string {
	return "AWS::SES::ReceiptRule.AddHeaderAction"
}

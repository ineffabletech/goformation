package cloudformation

// AWSWAFWebACL_WafAction AWS CloudFormation Resource (AWS::WAF::WebACL.WafAction)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-webacl-action.html
type AWSWAFWebACL_WafAction struct {
	dependsOn []string

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-webacl-action.html#cfn-waf-webacl-action-type
	Type *String `json:"Type,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSWAFWebACL_WafAction) AddDependencies(v ...string) *AWSWAFWebACL_WafAction {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSWAFWebACL_WafAction) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFWebACL_WafAction) AWSCloudFormationType() string {
	return "AWS::WAF::WebACL.WafAction"
}

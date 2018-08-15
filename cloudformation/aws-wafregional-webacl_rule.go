package cloudformation

// AWSWAFRegionalWebACL_Rule AWS CloudFormation Resource (AWS::WAFRegional::WebACL.Rule)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-webacl-rule.html
type AWSWAFRegionalWebACL_Rule struct {
	dependsOn []string

	// Action AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-webacl-rule.html#cfn-wafregional-webacl-rule-action
	Action *AWSWAFRegionalWebACL_Action `json:"Action,omitempty"`

	// Priority AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-webacl-rule.html#cfn-wafregional-webacl-rule-priority
	Priority *Integer `json:"Priority,omitempty"`

	// RuleId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-webacl-rule.html#cfn-wafregional-webacl-rule-ruleid
	RuleId *String `json:"RuleId,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSWAFRegionalWebACL_Rule) AddDependencies(v ...string) *AWSWAFRegionalWebACL_Rule {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSWAFRegionalWebACL_Rule) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFRegionalWebACL_Rule) AWSCloudFormationType() string {
	return "AWS::WAFRegional::WebACL.Rule"
}

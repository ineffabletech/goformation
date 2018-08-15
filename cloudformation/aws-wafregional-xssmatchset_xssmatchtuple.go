package cloudformation

// AWSWAFRegionalXssMatchSet_XssMatchTuple AWS CloudFormation Resource (AWS::WAFRegional::XssMatchSet.XssMatchTuple)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-xssmatchset-xssmatchtuple.html
type AWSWAFRegionalXssMatchSet_XssMatchTuple struct {
	dependsOn []string

	// FieldToMatch AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-xssmatchset-xssmatchtuple.html#cfn-wafregional-xssmatchset-xssmatchtuple-fieldtomatch
	FieldToMatch *AWSWAFRegionalXssMatchSet_FieldToMatch `json:"FieldToMatch,omitempty"`

	// TextTransformation AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-xssmatchset-xssmatchtuple.html#cfn-wafregional-xssmatchset-xssmatchtuple-texttransformation
	TextTransformation *String `json:"TextTransformation,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSWAFRegionalXssMatchSet_XssMatchTuple) AddDependencies(v ...string) *AWSWAFRegionalXssMatchSet_XssMatchTuple {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSWAFRegionalXssMatchSet_XssMatchTuple) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFRegionalXssMatchSet_XssMatchTuple) AWSCloudFormationType() string {
	return "AWS::WAFRegional::XssMatchSet.XssMatchTuple"
}

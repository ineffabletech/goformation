package cloudformation

// AWSWAFXssMatchSet_FieldToMatch AWS CloudFormation Resource (AWS::WAF::XssMatchSet.FieldToMatch)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-xssmatchset-xssmatchtuple-fieldtomatch.html
type AWSWAFXssMatchSet_FieldToMatch struct {
	dependsOn []string

	// Data AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-xssmatchset-xssmatchtuple-fieldtomatch.html#cfn-waf-xssmatchset-xssmatchtuple-fieldtomatch-data
	Data *String `json:"Data,omitempty"`

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-xssmatchset-xssmatchtuple-fieldtomatch.html#cfn-waf-xssmatchset-xssmatchtuple-fieldtomatch-type
	Type *String `json:"Type,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSWAFXssMatchSet_FieldToMatch) AddDependencies(v ...string) *AWSWAFXssMatchSet_FieldToMatch {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSWAFXssMatchSet_FieldToMatch) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFXssMatchSet_FieldToMatch) AWSCloudFormationType() string {
	return "AWS::WAF::XssMatchSet.FieldToMatch"
}

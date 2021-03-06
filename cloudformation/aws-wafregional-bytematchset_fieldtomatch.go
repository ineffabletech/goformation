package cloudformation

// AWSWAFRegionalByteMatchSet_FieldToMatch AWS CloudFormation Resource (AWS::WAFRegional::ByteMatchSet.FieldToMatch)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-bytematchset-fieldtomatch.html
type AWSWAFRegionalByteMatchSet_FieldToMatch struct {
	dependsOn []string

	// Data AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-bytematchset-fieldtomatch.html#cfn-wafregional-bytematchset-fieldtomatch-data
	Data *String `json:"Data,omitempty"`

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-bytematchset-fieldtomatch.html#cfn-wafregional-bytematchset-fieldtomatch-type
	Type *String `json:"Type,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSWAFRegionalByteMatchSet_FieldToMatch) AddDependencies(v ...string) *AWSWAFRegionalByteMatchSet_FieldToMatch {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSWAFRegionalByteMatchSet_FieldToMatch) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFRegionalByteMatchSet_FieldToMatch) AWSCloudFormationType() string {
	return "AWS::WAFRegional::ByteMatchSet.FieldToMatch"
}

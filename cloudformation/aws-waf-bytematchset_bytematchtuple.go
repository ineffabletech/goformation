package cloudformation

// AWSWAFByteMatchSet_ByteMatchTuple AWS CloudFormation Resource (AWS::WAF::ByteMatchSet.ByteMatchTuple)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-bytematchset-bytematchtuples.html
type AWSWAFByteMatchSet_ByteMatchTuple struct {
	dependsOn []string

	// FieldToMatch AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-bytematchset-bytematchtuples.html#cfn-waf-bytematchset-bytematchtuples-fieldtomatch
	FieldToMatch *AWSWAFByteMatchSet_FieldToMatch `json:"FieldToMatch,omitempty"`

	// PositionalConstraint AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-bytematchset-bytematchtuples.html#cfn-waf-bytematchset-bytematchtuples-positionalconstraint
	PositionalConstraint *String `json:"PositionalConstraint,omitempty"`

	// TargetString AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-bytematchset-bytematchtuples.html#cfn-waf-bytematchset-bytematchtuples-targetstring
	TargetString *String `json:"TargetString,omitempty"`

	// TargetStringBase64 AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-bytematchset-bytematchtuples.html#cfn-waf-bytematchset-bytematchtuples-targetstringbase64
	TargetStringBase64 *String `json:"TargetStringBase64,omitempty"`

	// TextTransformation AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-bytematchset-bytematchtuples.html#cfn-waf-bytematchset-bytematchtuples-texttransformation
	TextTransformation *String `json:"TextTransformation,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSWAFByteMatchSet_ByteMatchTuple) AddDependencies(v ...string) *AWSWAFByteMatchSet_ByteMatchTuple {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSWAFByteMatchSet_ByteMatchTuple) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFByteMatchSet_ByteMatchTuple) AWSCloudFormationType() string {
	return "AWS::WAF::ByteMatchSet.ByteMatchTuple"
}

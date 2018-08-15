package cloudformation

// AWSSSMAssociation_ParameterValues AWS CloudFormation Resource (AWS::SSM::Association.ParameterValues)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-parametervalues.html
type AWSSSMAssociation_ParameterValues struct {
	dependsOn []string

	// ParameterValues AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-parametervalues.html#cfn-ssm-association-parametervalues-parametervalues
	ParameterValues []*String `json:"ParameterValues,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSSSMAssociation_ParameterValues) AddDependencies(v ...string) *AWSSSMAssociation_ParameterValues {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSSSMAssociation_ParameterValues) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSSMAssociation_ParameterValues) AWSCloudFormationType() string {
	return "AWS::SSM::Association.ParameterValues"
}

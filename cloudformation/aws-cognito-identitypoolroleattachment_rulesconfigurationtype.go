package cloudformation

// AWSCognitoIdentityPoolRoleAttachment_RulesConfigurationType AWS CloudFormation Resource (AWS::Cognito::IdentityPoolRoleAttachment.RulesConfigurationType)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypoolroleattachment-rulesconfigurationtype.html
type AWSCognitoIdentityPoolRoleAttachment_RulesConfigurationType struct {
	dependsOn []string

	// Rules AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypoolroleattachment-rulesconfigurationtype.html#cfn-cognito-identitypoolroleattachment-rulesconfigurationtype-rules
	Rules []AWSCognitoIdentityPoolRoleAttachment_MappingRule `json:"Rules,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSCognitoIdentityPoolRoleAttachment_RulesConfigurationType) AddDependencies(v ...string) *AWSCognitoIdentityPoolRoleAttachment_RulesConfigurationType {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSCognitoIdentityPoolRoleAttachment_RulesConfigurationType) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCognitoIdentityPoolRoleAttachment_RulesConfigurationType) AWSCloudFormationType() string {
	return "AWS::Cognito::IdentityPoolRoleAttachment.RulesConfigurationType"
}

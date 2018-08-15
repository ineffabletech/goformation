package cloudformation

// AWSCognitoUserPool_Policies AWS CloudFormation Resource (AWS::Cognito::UserPool.Policies)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-policies.html
type AWSCognitoUserPool_Policies struct {
	dependsOn []string

	// PasswordPolicy AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-policies.html#cfn-cognito-userpool-policies-passwordpolicy
	PasswordPolicy *AWSCognitoUserPool_PasswordPolicy `json:"PasswordPolicy,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSCognitoUserPool_Policies) AddDependencies(v ...string) *AWSCognitoUserPool_Policies {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSCognitoUserPool_Policies) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCognitoUserPool_Policies) AWSCloudFormationType() string {
	return "AWS::Cognito::UserPool.Policies"
}

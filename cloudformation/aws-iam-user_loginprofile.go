package cloudformation

// AWSIAMUser_LoginProfile AWS CloudFormation Resource (AWS::IAM::User.LoginProfile)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-user-loginprofile.html
type AWSIAMUser_LoginProfile struct {
	dependsOn []string

	// Password AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-user-loginprofile.html#cfn-iam-user-loginprofile-password
	Password *String `json:"Password,omitempty"`

	// PasswordResetRequired AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-user-loginprofile.html#cfn-iam-user-loginprofile-passwordresetrequired
	PasswordResetRequired *Boolean `json:"PasswordResetRequired,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSIAMUser_LoginProfile) AddDependencies(v ...string) *AWSIAMUser_LoginProfile {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSIAMUser_LoginProfile) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIAMUser_LoginProfile) AWSCloudFormationType() string {
	return "AWS::IAM::User.LoginProfile"
}

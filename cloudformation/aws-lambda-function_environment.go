package cloudformation

// AWSLambdaFunction_Environment AWS CloudFormation Resource (AWS::Lambda::Function.Environment)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-environment.html
type AWSLambdaFunction_Environment struct {
	dependsOn []string

	// Variables AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-environment.html#cfn-lambda-function-environment-variables
	Variables map[string]*String `json:"Variables,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSLambdaFunction_Environment) AddDependencies(v ...string) *AWSLambdaFunction_Environment {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSLambdaFunction_Environment) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSLambdaFunction_Environment) AWSCloudFormationType() string {
	return "AWS::Lambda::Function.Environment"
}

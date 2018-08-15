package cloudformation

// AWSLambdaFunction_DeadLetterConfig AWS CloudFormation Resource (AWS::Lambda::Function.DeadLetterConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-deadletterconfig.html
type AWSLambdaFunction_DeadLetterConfig struct {
	dependsOn []string

	// TargetArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-deadletterconfig.html#cfn-lambda-function-deadletterconfig-targetarn
	TargetArn *String `json:"TargetArn,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSLambdaFunction_DeadLetterConfig) AddDependencies(v ...string) *AWSLambdaFunction_DeadLetterConfig {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSLambdaFunction_DeadLetterConfig) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSLambdaFunction_DeadLetterConfig) AWSCloudFormationType() string {
	return "AWS::Lambda::Function.DeadLetterConfig"
}

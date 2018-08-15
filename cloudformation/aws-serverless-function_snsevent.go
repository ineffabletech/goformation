package cloudformation

// AWSServerlessFunction_SNSEvent AWS CloudFormation Resource (AWS::Serverless::Function.SNSEvent)
// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#sns
type AWSServerlessFunction_SNSEvent struct {
	dependsOn []string

	// Topic AWS CloudFormation Property
	// Required: true
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#sns
	Topic *String `json:"Topic,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSServerlessFunction_SNSEvent) AddDependencies(v ...string) *AWSServerlessFunction_SNSEvent {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSServerlessFunction_SNSEvent) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSServerlessFunction_SNSEvent) AWSCloudFormationType() string {
	return "AWS::Serverless::Function.SNSEvent"
}

package cloudformation

// AWSServerlessFunction_EventSource AWS CloudFormation Resource (AWS::Serverless::Function.EventSource)
// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#event-source-object
type AWSServerlessFunction_EventSource struct {
	dependsOn []string

	// Properties AWS CloudFormation Property
	// Required: true
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#event-source-types
	Properties *AWSServerlessFunction_Properties `json:"Properties,omitempty"`

	// Type AWS CloudFormation Property
	// Required: true
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#event-source-object
	Type *String `json:"Type,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSServerlessFunction_EventSource) AddDependencies(v ...string) *AWSServerlessFunction_EventSource {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSServerlessFunction_EventSource) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSServerlessFunction_EventSource) AWSCloudFormationType() string {
	return "AWS::Serverless::Function.EventSource"
}

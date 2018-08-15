package cloudformation

// AWSServerlessFunction_CloudWatchEventEvent AWS CloudFormation Resource (AWS::Serverless::Function.CloudWatchEventEvent)
// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#cloudwatchevent
type AWSServerlessFunction_CloudWatchEventEvent struct {
	dependsOn []string

	// Input AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#cloudwatchevent
	Input *String `json:"Input,omitempty"`

	// InputPath AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#cloudwatchevent
	InputPath *String `json:"InputPath,omitempty"`

	// Pattern AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AmazonCloudWatch/latest/events/CloudWatchEventsandEventPatterns.html
	Pattern interface{} `json:"Pattern,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSServerlessFunction_CloudWatchEventEvent) AddDependencies(v ...string) *AWSServerlessFunction_CloudWatchEventEvent {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSServerlessFunction_CloudWatchEventEvent) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSServerlessFunction_CloudWatchEventEvent) AWSCloudFormationType() string {
	return "AWS::Serverless::Function.CloudWatchEventEvent"
}

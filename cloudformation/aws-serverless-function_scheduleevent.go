package cloudformation

// AWSServerlessFunction_ScheduleEvent AWS CloudFormation Resource (AWS::Serverless::Function.ScheduleEvent)
// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#schedule
type AWSServerlessFunction_ScheduleEvent struct {
	dependsOn []string

	// Input AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#schedule
	Input *String `json:"Input,omitempty"`

	// Schedule AWS CloudFormation Property
	// Required: true
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#schedule
	Schedule *String `json:"Schedule,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSServerlessFunction_ScheduleEvent) AddDependencies(v ...string) *AWSServerlessFunction_ScheduleEvent {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSServerlessFunction_ScheduleEvent) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSServerlessFunction_ScheduleEvent) AWSCloudFormationType() string {
	return "AWS::Serverless::Function.ScheduleEvent"
}

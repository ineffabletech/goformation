package cloudformation

// AWSIoTTopicRule_SqsAction AWS CloudFormation Resource (AWS::IoT::TopicRule.SqsAction)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-sqsaction.html
type AWSIoTTopicRule_SqsAction struct {
	dependsOn []string

	// QueueUrl AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-sqsaction.html#cfn-iot-topicrule-sqsaction-queueurl
	QueueUrl *String `json:"QueueUrl,omitempty"`

	// RoleArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-sqsaction.html#cfn-iot-topicrule-sqsaction-rolearn
	RoleArn *String `json:"RoleArn,omitempty"`

	// UseBase64 AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-sqsaction.html#cfn-iot-topicrule-sqsaction-usebase64
	UseBase64 *Boolean `json:"UseBase64,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSIoTTopicRule_SqsAction) AddDependencies(v ...string) *AWSIoTTopicRule_SqsAction {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSIoTTopicRule_SqsAction) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIoTTopicRule_SqsAction) AWSCloudFormationType() string {
	return "AWS::IoT::TopicRule.SqsAction"
}

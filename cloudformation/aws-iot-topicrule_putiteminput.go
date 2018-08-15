package cloudformation

// AWSIoTTopicRule_PutItemInput AWS CloudFormation Resource (AWS::IoT::TopicRule.PutItemInput)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-putiteminput.html
type AWSIoTTopicRule_PutItemInput struct {
	dependsOn []string

	// TableName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-putiteminput.html#cfn-iot-topicrule-putiteminput-tablename
	TableName *String `json:"TableName,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSIoTTopicRule_PutItemInput) AddDependencies(v ...string) *AWSIoTTopicRule_PutItemInput {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSIoTTopicRule_PutItemInput) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIoTTopicRule_PutItemInput) AWSCloudFormationType() string {
	return "AWS::IoT::TopicRule.PutItemInput"
}

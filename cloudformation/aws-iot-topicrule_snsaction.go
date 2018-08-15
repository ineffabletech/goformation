package cloudformation

// AWSIoTTopicRule_SnsAction AWS CloudFormation Resource (AWS::IoT::TopicRule.SnsAction)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-snsaction.html
type AWSIoTTopicRule_SnsAction struct {
	dependsOn []string

	// MessageFormat AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-snsaction.html#cfn-iot-topicrule-snsaction-messageformat
	MessageFormat *String `json:"MessageFormat,omitempty"`

	// RoleArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-snsaction.html#cfn-iot-topicrule-snsaction-rolearn
	RoleArn *String `json:"RoleArn,omitempty"`

	// TargetArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-snsaction.html#cfn-iot-topicrule-snsaction-targetarn
	TargetArn *String `json:"TargetArn,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSIoTTopicRule_SnsAction) AddDependencies(v ...string) *AWSIoTTopicRule_SnsAction {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSIoTTopicRule_SnsAction) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIoTTopicRule_SnsAction) AWSCloudFormationType() string {
	return "AWS::IoT::TopicRule.SnsAction"
}

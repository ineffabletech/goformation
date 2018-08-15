package cloudformation

// AWSIoTTopicRule_TopicRulePayload AWS CloudFormation Resource (AWS::IoT::TopicRule.TopicRulePayload)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-topicrulepayload.html
type AWSIoTTopicRule_TopicRulePayload struct {
	dependsOn []string

	// Actions AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-topicrulepayload.html#cfn-iot-topicrule-topicrulepayload-actions
	Actions []AWSIoTTopicRule_Action `json:"Actions,omitempty"`

	// AwsIotSqlVersion AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-topicrulepayload.html#cfn-iot-topicrule-topicrulepayload-awsiotsqlversion
	AwsIotSqlVersion *String `json:"AwsIotSqlVersion,omitempty"`

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-topicrulepayload.html#cfn-iot-topicrule-topicrulepayload-description
	Description *String `json:"Description,omitempty"`

	// RuleDisabled AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-topicrulepayload.html#cfn-iot-topicrule-topicrulepayload-ruledisabled
	RuleDisabled *Boolean `json:"RuleDisabled,omitempty"`

	// Sql AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-topicrulepayload.html#cfn-iot-topicrule-topicrulepayload-sql
	Sql *String `json:"Sql,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSIoTTopicRule_TopicRulePayload) AddDependencies(v ...string) *AWSIoTTopicRule_TopicRulePayload {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSIoTTopicRule_TopicRulePayload) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIoTTopicRule_TopicRulePayload) AWSCloudFormationType() string {
	return "AWS::IoT::TopicRule.TopicRulePayload"
}

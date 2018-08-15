package cloudformation

// AWSEventsRule_KinesisParameters AWS CloudFormation Resource (AWS::Events::Rule.KinesisParameters)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-kinesisparameters.html
type AWSEventsRule_KinesisParameters struct {
	dependsOn []string

	// PartitionKeyPath AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-kinesisparameters.html#cfn-events-rule-kinesisparameters-partitionkeypath
	PartitionKeyPath *String `json:"PartitionKeyPath,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSEventsRule_KinesisParameters) AddDependencies(v ...string) *AWSEventsRule_KinesisParameters {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSEventsRule_KinesisParameters) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEventsRule_KinesisParameters) AWSCloudFormationType() string {
	return "AWS::Events::Rule.KinesisParameters"
}

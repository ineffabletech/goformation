package cloudformation

// AWSKinesisFirehoseDeliveryStream_KinesisStreamSourceConfiguration AWS CloudFormation Resource (AWS::KinesisFirehose::DeliveryStream.KinesisStreamSourceConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-kinesisstreamsourceconfiguration.html
type AWSKinesisFirehoseDeliveryStream_KinesisStreamSourceConfiguration struct {
	dependsOn []string

	// KinesisStreamARN AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-kinesisstreamsourceconfiguration.html#cfn-kinesisfirehose-deliverystream-kinesisstreamsourceconfiguration-kinesisstreamarn
	KinesisStreamARN *String `json:"KinesisStreamARN,omitempty"`

	// RoleARN AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-kinesisstreamsourceconfiguration.html#cfn-kinesisfirehose-deliverystream-kinesisstreamsourceconfiguration-rolearn
	RoleARN *String `json:"RoleARN,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSKinesisFirehoseDeliveryStream_KinesisStreamSourceConfiguration) AddDependencies(v ...string) *AWSKinesisFirehoseDeliveryStream_KinesisStreamSourceConfiguration {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSKinesisFirehoseDeliveryStream_KinesisStreamSourceConfiguration) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisFirehoseDeliveryStream_KinesisStreamSourceConfiguration) AWSCloudFormationType() string {
	return "AWS::KinesisFirehose::DeliveryStream.KinesisStreamSourceConfiguration"
}

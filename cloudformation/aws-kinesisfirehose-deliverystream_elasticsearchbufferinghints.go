package cloudformation

// AWSKinesisFirehoseDeliveryStream_ElasticsearchBufferingHints AWS CloudFormation Resource (AWS::KinesisFirehose::DeliveryStream.ElasticsearchBufferingHints)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-elasticsearchbufferinghints.html
type AWSKinesisFirehoseDeliveryStream_ElasticsearchBufferingHints struct {
	dependsOn []string

	// IntervalInSeconds AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-elasticsearchbufferinghints.html#cfn-kinesisfirehose-deliverystream-elasticsearchbufferinghints-intervalinseconds
	IntervalInSeconds *Integer `json:"IntervalInSeconds,omitempty"`

	// SizeInMBs AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-elasticsearchbufferinghints.html#cfn-kinesisfirehose-deliverystream-elasticsearchbufferinghints-sizeinmbs
	SizeInMBs *Integer `json:"SizeInMBs,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSKinesisFirehoseDeliveryStream_ElasticsearchBufferingHints) AddDependencies(v ...string) *AWSKinesisFirehoseDeliveryStream_ElasticsearchBufferingHints {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSKinesisFirehoseDeliveryStream_ElasticsearchBufferingHints) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisFirehoseDeliveryStream_ElasticsearchBufferingHints) AWSCloudFormationType() string {
	return "AWS::KinesisFirehose::DeliveryStream.ElasticsearchBufferingHints"
}

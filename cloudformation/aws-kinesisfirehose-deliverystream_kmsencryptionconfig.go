package cloudformation

// AWSKinesisFirehoseDeliveryStream_KMSEncryptionConfig AWS CloudFormation Resource (AWS::KinesisFirehose::DeliveryStream.KMSEncryptionConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-kmsencryptionconfig.html
type AWSKinesisFirehoseDeliveryStream_KMSEncryptionConfig struct {
	dependsOn []string

	// AWSKMSKeyARN AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-kmsencryptionconfig.html#cfn-kinesisfirehose-deliverystream-kmsencryptionconfig-awskmskeyarn
	AWSKMSKeyARN *String `json:"AWSKMSKeyARN,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSKinesisFirehoseDeliveryStream_KMSEncryptionConfig) AddDependencies(v ...string) *AWSKinesisFirehoseDeliveryStream_KMSEncryptionConfig {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSKinesisFirehoseDeliveryStream_KMSEncryptionConfig) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisFirehoseDeliveryStream_KMSEncryptionConfig) AWSCloudFormationType() string {
	return "AWS::KinesisFirehose::DeliveryStream.KMSEncryptionConfig"
}

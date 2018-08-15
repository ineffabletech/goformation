package cloudformation

// AWSKinesisFirehoseDeliveryStream_EncryptionConfiguration AWS CloudFormation Resource (AWS::KinesisFirehose::DeliveryStream.EncryptionConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-encryptionconfiguration.html
type AWSKinesisFirehoseDeliveryStream_EncryptionConfiguration struct {
	dependsOn []string

	// KMSEncryptionConfig AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-encryptionconfiguration.html#cfn-kinesisfirehose-deliverystream-encryptionconfiguration-kmsencryptionconfig
	KMSEncryptionConfig *AWSKinesisFirehoseDeliveryStream_KMSEncryptionConfig `json:"KMSEncryptionConfig,omitempty"`

	// NoEncryptionConfig AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-encryptionconfiguration.html#cfn-kinesisfirehose-deliverystream-encryptionconfiguration-noencryptionconfig
	NoEncryptionConfig *String `json:"NoEncryptionConfig,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSKinesisFirehoseDeliveryStream_EncryptionConfiguration) AddDependencies(v ...string) *AWSKinesisFirehoseDeliveryStream_EncryptionConfiguration {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSKinesisFirehoseDeliveryStream_EncryptionConfiguration) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisFirehoseDeliveryStream_EncryptionConfiguration) AWSCloudFormationType() string {
	return "AWS::KinesisFirehose::DeliveryStream.EncryptionConfiguration"
}

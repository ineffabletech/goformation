package cloudformation

// AWSKinesisStream_StreamEncryption AWS CloudFormation Resource (AWS::Kinesis::Stream.StreamEncryption)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-stream-streamencryption.html
type AWSKinesisStream_StreamEncryption struct {
	dependsOn []string

	// EncryptionType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-stream-streamencryption.html#cfn-kinesis-stream-streamencryption-encryptiontype
	EncryptionType *String `json:"EncryptionType,omitempty"`

	// KeyId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-stream-streamencryption.html#cfn-kinesis-stream-streamencryption-keyid
	KeyId *String `json:"KeyId,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSKinesisStream_StreamEncryption) AddDependencies(v ...string) *AWSKinesisStream_StreamEncryption {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSKinesisStream_StreamEncryption) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisStream_StreamEncryption) AWSCloudFormationType() string {
	return "AWS::Kinesis::Stream.StreamEncryption"
}

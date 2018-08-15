package cloudformation

// AWSElasticsearchDomain_EncryptionAtRestOptions AWS CloudFormation Resource (AWS::Elasticsearch::Domain.EncryptionAtRestOptions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticsearch-domain-encryptionatrestoptions.html
type AWSElasticsearchDomain_EncryptionAtRestOptions struct {
	dependsOn []string

	// Enabled AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticsearch-domain-encryptionatrestoptions.html#cfn-elasticsearch-domain-encryptionatrestoptions-enabled
	Enabled *Boolean `json:"Enabled,omitempty"`

	// KmsKeyId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticsearch-domain-encryptionatrestoptions.html#cfn-elasticsearch-domain-encryptionatrestoptions-kmskeyid
	KmsKeyId *String `json:"KmsKeyId,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSElasticsearchDomain_EncryptionAtRestOptions) AddDependencies(v ...string) *AWSElasticsearchDomain_EncryptionAtRestOptions {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSElasticsearchDomain_EncryptionAtRestOptions) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticsearchDomain_EncryptionAtRestOptions) AWSCloudFormationType() string {
	return "AWS::Elasticsearch::Domain.EncryptionAtRestOptions"
}

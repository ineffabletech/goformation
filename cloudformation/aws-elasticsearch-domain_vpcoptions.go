package cloudformation

// AWSElasticsearchDomain_VPCOptions AWS CloudFormation Resource (AWS::Elasticsearch::Domain.VPCOptions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticsearch-domain-vpcoptions.html
type AWSElasticsearchDomain_VPCOptions struct {
	dependsOn []string

	// SecurityGroupIds AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticsearch-domain-vpcoptions.html#cfn-elasticsearch-domain-vpcoptions-securitygroupids
	SecurityGroupIds []*String `json:"SecurityGroupIds,omitempty"`

	// SubnetIds AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticsearch-domain-vpcoptions.html#cfn-elasticsearch-domain-vpcoptions-subnetids
	SubnetIds []*String `json:"SubnetIds,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSElasticsearchDomain_VPCOptions) AddDependencies(v ...string) *AWSElasticsearchDomain_VPCOptions {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSElasticsearchDomain_VPCOptions) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticsearchDomain_VPCOptions) AWSCloudFormationType() string {
	return "AWS::Elasticsearch::Domain.VPCOptions"
}

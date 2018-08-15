package cloudformation

// AWSServiceDiscoveryService_DnsConfig AWS CloudFormation Resource (AWS::ServiceDiscovery::Service.DnsConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicediscovery-service-dnsconfig.html
type AWSServiceDiscoveryService_DnsConfig struct {
	dependsOn []string

	// DnsRecords AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicediscovery-service-dnsconfig.html#cfn-servicediscovery-service-dnsconfig-dnsrecords
	DnsRecords []AWSServiceDiscoveryService_DnsRecord `json:"DnsRecords,omitempty"`

	// NamespaceId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicediscovery-service-dnsconfig.html#cfn-servicediscovery-service-dnsconfig-namespaceid
	NamespaceId *String `json:"NamespaceId,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSServiceDiscoveryService_DnsConfig) AddDependencies(v ...string) *AWSServiceDiscoveryService_DnsConfig {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSServiceDiscoveryService_DnsConfig) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSServiceDiscoveryService_DnsConfig) AWSCloudFormationType() string {
	return "AWS::ServiceDiscovery::Service.DnsConfig"
}

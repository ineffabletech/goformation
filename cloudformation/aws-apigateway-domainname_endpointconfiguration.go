package cloudformation

// AWSApiGatewayDomainName_EndpointConfiguration AWS CloudFormation Resource (AWS::ApiGateway::DomainName.EndpointConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-domainname-endpointconfiguration.html
type AWSApiGatewayDomainName_EndpointConfiguration struct {
	dependsOn []string

	// Types AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-domainname-endpointconfiguration.html#cfn-apigateway-domainname-endpointconfiguration-types
	Types []*String `json:"Types,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSApiGatewayDomainName_EndpointConfiguration) AddDependencies(v ...string) *AWSApiGatewayDomainName_EndpointConfiguration {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSApiGatewayDomainName_EndpointConfiguration) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApiGatewayDomainName_EndpointConfiguration) AWSCloudFormationType() string {
	return "AWS::ApiGateway::DomainName.EndpointConfiguration"
}

package cloudformation

// AWSApiGatewayRestApi_EndpointConfiguration AWS CloudFormation Resource (AWS::ApiGateway::RestApi.EndpointConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-restapi-endpointconfiguration.html
type AWSApiGatewayRestApi_EndpointConfiguration struct {
	dependsOn []string

	// Types AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-restapi-endpointconfiguration.html#cfn-apigateway-restapi-endpointconfiguration-types
	Types []*String `json:"Types,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSApiGatewayRestApi_EndpointConfiguration) AddDependencies(v ...string) *AWSApiGatewayRestApi_EndpointConfiguration {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSApiGatewayRestApi_EndpointConfiguration) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApiGatewayRestApi_EndpointConfiguration) AWSCloudFormationType() string {
	return "AWS::ApiGateway::RestApi.EndpointConfiguration"
}

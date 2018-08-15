package cloudformation

// AWSApiGatewayUsagePlan_ThrottleSettings AWS CloudFormation Resource (AWS::ApiGateway::UsagePlan.ThrottleSettings)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-throttlesettings.html
type AWSApiGatewayUsagePlan_ThrottleSettings struct {
	dependsOn []string

	// BurstLimit AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-throttlesettings.html#cfn-apigateway-usageplan-throttlesettings-burstlimit
	BurstLimit *Integer `json:"BurstLimit,omitempty"`

	// RateLimit AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-throttlesettings.html#cfn-apigateway-usageplan-throttlesettings-ratelimit
	RateLimit *Double `json:"RateLimit,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSApiGatewayUsagePlan_ThrottleSettings) AddDependencies(v ...string) *AWSApiGatewayUsagePlan_ThrottleSettings {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSApiGatewayUsagePlan_ThrottleSettings) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApiGatewayUsagePlan_ThrottleSettings) AWSCloudFormationType() string {
	return "AWS::ApiGateway::UsagePlan.ThrottleSettings"
}

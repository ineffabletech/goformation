package cloudformation

// AWSApiGatewayUsagePlan_QuotaSettings AWS CloudFormation Resource (AWS::ApiGateway::UsagePlan.QuotaSettings)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-quotasettings.html
type AWSApiGatewayUsagePlan_QuotaSettings struct {
	dependsOn []string

	// Limit AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-quotasettings.html#cfn-apigateway-usageplan-quotasettings-limit
	Limit *Integer `json:"Limit,omitempty"`

	// Offset AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-quotasettings.html#cfn-apigateway-usageplan-quotasettings-offset
	Offset *Integer `json:"Offset,omitempty"`

	// Period AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-quotasettings.html#cfn-apigateway-usageplan-quotasettings-period
	Period *String `json:"Period,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSApiGatewayUsagePlan_QuotaSettings) AddDependencies(v ...string) *AWSApiGatewayUsagePlan_QuotaSettings {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSApiGatewayUsagePlan_QuotaSettings) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApiGatewayUsagePlan_QuotaSettings) AWSCloudFormationType() string {
	return "AWS::ApiGateway::UsagePlan.QuotaSettings"
}

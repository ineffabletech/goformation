package cloudformation

// AWSApiGatewayUsagePlan_ApiStage AWS CloudFormation Resource (AWS::ApiGateway::UsagePlan.ApiStage)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-apistage.html
type AWSApiGatewayUsagePlan_ApiStage struct {
	dependsOn []string

	// ApiId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-apistage.html#cfn-apigateway-usageplan-apistage-apiid
	ApiId *String `json:"ApiId,omitempty"`

	// Stage AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-apistage.html#cfn-apigateway-usageplan-apistage-stage
	Stage *String `json:"Stage,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSApiGatewayUsagePlan_ApiStage) AddDependencies(v ...string) *AWSApiGatewayUsagePlan_ApiStage {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSApiGatewayUsagePlan_ApiStage) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApiGatewayUsagePlan_ApiStage) AWSCloudFormationType() string {
	return "AWS::ApiGateway::UsagePlan.ApiStage"
}

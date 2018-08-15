package cloudformation

// AWSECSService_NetworkConfiguration AWS CloudFormation Resource (AWS::ECS::Service.NetworkConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-networkconfiguration.html
type AWSECSService_NetworkConfiguration struct {
	dependsOn []string

	// AwsvpcConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-networkconfiguration.html#cfn-ecs-service-networkconfiguration-awsvpcconfiguration
	AwsvpcConfiguration *AWSECSService_AwsVpcConfiguration `json:"AwsvpcConfiguration,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSECSService_NetworkConfiguration) AddDependencies(v ...string) *AWSECSService_NetworkConfiguration {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSECSService_NetworkConfiguration) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSECSService_NetworkConfiguration) AWSCloudFormationType() string {
	return "AWS::ECS::Service.NetworkConfiguration"
}

package cloudformation

// AWSECSService_DeploymentConfiguration AWS CloudFormation Resource (AWS::ECS::Service.DeploymentConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentconfiguration.html
type AWSECSService_DeploymentConfiguration struct {
	dependsOn []string

	// MaximumPercent AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentconfiguration.html#cfn-ecs-service-deploymentconfiguration-maximumpercent
	MaximumPercent *Integer `json:"MaximumPercent,omitempty"`

	// MinimumHealthyPercent AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentconfiguration.html#cfn-ecs-service-deploymentconfiguration-minimumhealthypercent
	MinimumHealthyPercent *Integer `json:"MinimumHealthyPercent,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSECSService_DeploymentConfiguration) AddDependencies(v ...string) *AWSECSService_DeploymentConfiguration {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSECSService_DeploymentConfiguration) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSECSService_DeploymentConfiguration) AWSCloudFormationType() string {
	return "AWS::ECS::Service.DeploymentConfiguration"
}

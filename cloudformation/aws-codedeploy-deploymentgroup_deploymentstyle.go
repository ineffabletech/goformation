package cloudformation

// AWSCodeDeployDeploymentGroup_DeploymentStyle AWS CloudFormation Resource (AWS::CodeDeploy::DeploymentGroup.DeploymentStyle)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deploymentstyle.html
type AWSCodeDeployDeploymentGroup_DeploymentStyle struct {
	dependsOn []string

	// DeploymentOption AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deploymentstyle.html#cfn-codedeploy-deploymentgroup-deploymentstyle-deploymentoption
	DeploymentOption *String `json:"DeploymentOption,omitempty"`

	// DeploymentType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deploymentstyle.html#cfn-codedeploy-deploymentgroup-deploymentstyle-deploymenttype
	DeploymentType *String `json:"DeploymentType,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSCodeDeployDeploymentGroup_DeploymentStyle) AddDependencies(v ...string) *AWSCodeDeployDeploymentGroup_DeploymentStyle {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSCodeDeployDeploymentGroup_DeploymentStyle) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodeDeployDeploymentGroup_DeploymentStyle) AWSCloudFormationType() string {
	return "AWS::CodeDeploy::DeploymentGroup.DeploymentStyle"
}

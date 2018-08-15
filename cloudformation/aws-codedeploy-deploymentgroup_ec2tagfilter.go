package cloudformation

// AWSCodeDeployDeploymentGroup_EC2TagFilter AWS CloudFormation Resource (AWS::CodeDeploy::DeploymentGroup.EC2TagFilter)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-ec2tagfilters.html
type AWSCodeDeployDeploymentGroup_EC2TagFilter struct {
	dependsOn []string

	// Key AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-ec2tagfilters.html#cfn-properties-codedeploy-deploymentgroup-ec2tagfilters-key
	Key *String `json:"Key,omitempty"`

	// Type AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-ec2tagfilters.html#cfn-properties-codedeploy-deploymentgroup-ec2tagfilters-type
	Type *String `json:"Type,omitempty"`

	// Value AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-ec2tagfilters.html#cfn-properties-codedeploy-deploymentgroup-ec2tagfilters-value
	Value *String `json:"Value,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSCodeDeployDeploymentGroup_EC2TagFilter) AddDependencies(v ...string) *AWSCodeDeployDeploymentGroup_EC2TagFilter {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSCodeDeployDeploymentGroup_EC2TagFilter) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodeDeployDeploymentGroup_EC2TagFilter) AWSCloudFormationType() string {
	return "AWS::CodeDeploy::DeploymentGroup.EC2TagFilter"
}

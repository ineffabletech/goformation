package cloudformation

// AWSCodeDeployDeploymentGroup_RevisionLocation AWS CloudFormation Resource (AWS::CodeDeploy::DeploymentGroup.RevisionLocation)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deployment-revision.html
type AWSCodeDeployDeploymentGroup_RevisionLocation struct {
	dependsOn []string

	// GitHubLocation AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deployment-revision.html#cfn-properties-codedeploy-deploymentgroup-deployment-revision-githublocation
	GitHubLocation *AWSCodeDeployDeploymentGroup_GitHubLocation `json:"GitHubLocation,omitempty"`

	// RevisionType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deployment-revision.html#cfn-properties-codedeploy-deploymentgroup-deployment-revision-revisiontype
	RevisionType *String `json:"RevisionType,omitempty"`

	// S3Location AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deployment-revision.html#cfn-properties-codedeploy-deploymentgroup-deployment-revision-s3location
	S3Location *AWSCodeDeployDeploymentGroup_S3Location `json:"S3Location,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSCodeDeployDeploymentGroup_RevisionLocation) AddDependencies(v ...string) *AWSCodeDeployDeploymentGroup_RevisionLocation {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSCodeDeployDeploymentGroup_RevisionLocation) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodeDeployDeploymentGroup_RevisionLocation) AWSCloudFormationType() string {
	return "AWS::CodeDeploy::DeploymentGroup.RevisionLocation"
}

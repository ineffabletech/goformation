package cloudformation

// AWSCodePipelinePipeline_ActionDeclaration AWS CloudFormation Resource (AWS::CodePipeline::Pipeline.ActionDeclaration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages-actions.html
type AWSCodePipelinePipeline_ActionDeclaration struct {
	dependsOn []string

	// ActionTypeId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages-actions.html#cfn-codepipeline-pipeline-stages-actions-actiontypeid
	ActionTypeId *AWSCodePipelinePipeline_ActionTypeId `json:"ActionTypeId,omitempty"`

	// Configuration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages-actions.html#cfn-codepipeline-pipeline-stages-actions-configuration
	Configuration interface{} `json:"Configuration,omitempty"`

	// InputArtifacts AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages-actions.html#cfn-codepipeline-pipeline-stages-actions-inputartifacts
	InputArtifacts []AWSCodePipelinePipeline_InputArtifact `json:"InputArtifacts,omitempty"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages-actions.html#cfn-codepipeline-pipeline-stages-actions-name
	Name *String `json:"Name,omitempty"`

	// OutputArtifacts AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages-actions.html#cfn-codepipeline-pipeline-stages-actions-outputartifacts
	OutputArtifacts []AWSCodePipelinePipeline_OutputArtifact `json:"OutputArtifacts,omitempty"`

	// RoleArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages-actions.html#cfn-codepipeline-pipeline-stages-actions-rolearn
	RoleArn *String `json:"RoleArn,omitempty"`

	// RunOrder AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages-actions.html#cfn-codepipeline-pipeline-stages-actions-runorder
	RunOrder *Integer `json:"RunOrder,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSCodePipelinePipeline_ActionDeclaration) AddDependencies(v ...string) *AWSCodePipelinePipeline_ActionDeclaration {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSCodePipelinePipeline_ActionDeclaration) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodePipelinePipeline_ActionDeclaration) AWSCloudFormationType() string {
	return "AWS::CodePipeline::Pipeline.ActionDeclaration"
}

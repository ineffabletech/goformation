package cloudformation

// AWSCodeBuildProject_ProjectTriggers AWS CloudFormation Resource (AWS::CodeBuild::Project.ProjectTriggers)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-projecttriggers.html
type AWSCodeBuildProject_ProjectTriggers struct {
	dependsOn []string

	// Webhook AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-projecttriggers.html#cfn-codebuild-project-projecttriggers-webhook
	Webhook *Boolean `json:"Webhook,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSCodeBuildProject_ProjectTriggers) AddDependencies(v ...string) *AWSCodeBuildProject_ProjectTriggers {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSCodeBuildProject_ProjectTriggers) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodeBuildProject_ProjectTriggers) AWSCloudFormationType() string {
	return "AWS::CodeBuild::Project.ProjectTriggers"
}

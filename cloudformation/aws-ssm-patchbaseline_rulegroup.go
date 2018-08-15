package cloudformation

// AWSSSMPatchBaseline_RuleGroup AWS CloudFormation Resource (AWS::SSM::PatchBaseline.RuleGroup)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-patchbaseline-rulegroup.html
type AWSSSMPatchBaseline_RuleGroup struct {
	dependsOn []string

	// PatchRules AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-patchbaseline-rulegroup.html#cfn-ssm-patchbaseline-rulegroup-patchrules
	PatchRules []AWSSSMPatchBaseline_Rule `json:"PatchRules,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSSSMPatchBaseline_RuleGroup) AddDependencies(v ...string) *AWSSSMPatchBaseline_RuleGroup {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSSSMPatchBaseline_RuleGroup) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSSMPatchBaseline_RuleGroup) AWSCloudFormationType() string {
	return "AWS::SSM::PatchBaseline.RuleGroup"
}

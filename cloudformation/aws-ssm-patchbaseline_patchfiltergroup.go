package cloudformation

// AWSSSMPatchBaseline_PatchFilterGroup AWS CloudFormation Resource (AWS::SSM::PatchBaseline.PatchFilterGroup)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-patchbaseline-patchfiltergroup.html
type AWSSSMPatchBaseline_PatchFilterGroup struct {
	dependsOn []string

	// PatchFilters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-patchbaseline-patchfiltergroup.html#cfn-ssm-patchbaseline-patchfiltergroup-patchfilters
	PatchFilters []AWSSSMPatchBaseline_PatchFilter `json:"PatchFilters,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSSSMPatchBaseline_PatchFilterGroup) AddDependencies(v ...string) *AWSSSMPatchBaseline_PatchFilterGroup {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSSSMPatchBaseline_PatchFilterGroup) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSSMPatchBaseline_PatchFilterGroup) AWSCloudFormationType() string {
	return "AWS::SSM::PatchBaseline.PatchFilterGroup"
}

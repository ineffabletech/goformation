package cloudformation

// AWSSSMPatchBaseline_PatchFilter AWS CloudFormation Resource (AWS::SSM::PatchBaseline.PatchFilter)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-patchbaseline-patchfilter.html
type AWSSSMPatchBaseline_PatchFilter struct {
	dependsOn []string

	// Key AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-patchbaseline-patchfilter.html#cfn-ssm-patchbaseline-patchfilter-key
	Key *String `json:"Key,omitempty"`

	// Values AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-patchbaseline-patchfilter.html#cfn-ssm-patchbaseline-patchfilter-values
	Values []*String `json:"Values,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSSSMPatchBaseline_PatchFilter) AddDependencies(v ...string) *AWSSSMPatchBaseline_PatchFilter {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSSSMPatchBaseline_PatchFilter) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSSMPatchBaseline_PatchFilter) AWSCloudFormationType() string {
	return "AWS::SSM::PatchBaseline.PatchFilter"
}

package cloudformation

// AWSBatchJobDefinition_VolumesHost AWS CloudFormation Resource (AWS::Batch::JobDefinition.VolumesHost)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-volumeshost.html
type AWSBatchJobDefinition_VolumesHost struct {
	dependsOn []string

	// SourcePath AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-volumeshost.html#cfn-batch-jobdefinition-volumeshost-sourcepath
	SourcePath *String `json:"SourcePath,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSBatchJobDefinition_VolumesHost) AddDependencies(v ...string) *AWSBatchJobDefinition_VolumesHost {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSBatchJobDefinition_VolumesHost) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSBatchJobDefinition_VolumesHost) AWSCloudFormationType() string {
	return "AWS::Batch::JobDefinition.VolumesHost"
}

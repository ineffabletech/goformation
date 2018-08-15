package cloudformation

// AWSDataPipelinePipeline_Field AWS CloudFormation Resource (AWS::DataPipeline::Pipeline.Field)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-pipelineobjects-fields.html
type AWSDataPipelinePipeline_Field struct {
	dependsOn []string

	// Key AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-pipelineobjects-fields.html#cfn-datapipeline-pipeline-pipelineobjects-fields-key
	Key *String `json:"Key,omitempty"`

	// RefValue AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-pipelineobjects-fields.html#cfn-datapipeline-pipeline-pipelineobjects-fields-refvalue
	RefValue *String `json:"RefValue,omitempty"`

	// StringValue AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-pipelineobjects-fields.html#cfn-datapipeline-pipeline-pipelineobjects-fields-stringvalue
	StringValue *String `json:"StringValue,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSDataPipelinePipeline_Field) AddDependencies(v ...string) *AWSDataPipelinePipeline_Field {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSDataPipelinePipeline_Field) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDataPipelinePipeline_Field) AWSCloudFormationType() string {
	return "AWS::DataPipeline::Pipeline.Field"
}

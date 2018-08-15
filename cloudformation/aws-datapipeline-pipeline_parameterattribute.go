package cloudformation

// AWSDataPipelinePipeline_ParameterAttribute AWS CloudFormation Resource (AWS::DataPipeline::Pipeline.ParameterAttribute)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-parameterobjects-attributes.html
type AWSDataPipelinePipeline_ParameterAttribute struct {
	dependsOn []string

	// Key AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-parameterobjects-attributes.html#cfn-datapipeline-pipeline-parameterobjects-attribtues-key
	Key *String `json:"Key,omitempty"`

	// StringValue AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-parameterobjects-attributes.html#cfn-datapipeline-pipeline-parameterobjects-attribtues-stringvalue
	StringValue *String `json:"StringValue,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSDataPipelinePipeline_ParameterAttribute) AddDependencies(v ...string) *AWSDataPipelinePipeline_ParameterAttribute {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSDataPipelinePipeline_ParameterAttribute) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDataPipelinePipeline_ParameterAttribute) AWSCloudFormationType() string {
	return "AWS::DataPipeline::Pipeline.ParameterAttribute"
}

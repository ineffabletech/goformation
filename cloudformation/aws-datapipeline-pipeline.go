package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSDataPipelinePipeline AWS CloudFormation Resource (AWS::DataPipeline::Pipeline)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datapipeline-pipeline.html
type AWSDataPipelinePipeline struct {
	dependsOn []string

	// Activate AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datapipeline-pipeline.html#cfn-datapipeline-pipeline-activate
	Activate *Boolean `json:"Activate,omitempty"`

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datapipeline-pipeline.html#cfn-datapipeline-pipeline-description
	Description *String `json:"Description,omitempty"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datapipeline-pipeline.html#cfn-datapipeline-pipeline-name
	Name *String `json:"Name,omitempty"`

	// ParameterObjects AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datapipeline-pipeline.html#cfn-datapipeline-pipeline-parameterobjects
	ParameterObjects []AWSDataPipelinePipeline_ParameterObject `json:"ParameterObjects,omitempty"`

	// ParameterValues AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datapipeline-pipeline.html#cfn-datapipeline-pipeline-parametervalues
	ParameterValues []AWSDataPipelinePipeline_ParameterValue `json:"ParameterValues,omitempty"`

	// PipelineObjects AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datapipeline-pipeline.html#cfn-datapipeline-pipeline-pipelineobjects
	PipelineObjects []AWSDataPipelinePipeline_PipelineObject `json:"PipelineObjects,omitempty"`

	// PipelineTags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datapipeline-pipeline.html#cfn-datapipeline-pipeline-pipelinetags
	PipelineTags []AWSDataPipelinePipeline_PipelineTag `json:"PipelineTags,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSDataPipelinePipeline) AddDependencies(v ...string) *AWSDataPipelinePipeline {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSDataPipelinePipeline) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDataPipelinePipeline) AWSCloudFormationType() string {
	return "AWS::DataPipeline::Pipeline"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSDataPipelinePipeline) MarshalJSON() ([]byte, error) {
	type Properties AWSDataPipelinePipeline
	return json.Marshal(&struct {
		Type       string
		Properties Properties
		DependsOn  []string `json:"DependsOn,omitempty"`
	}{
		Type:       r.AWSCloudFormationType(),
		Properties: (Properties)(*r),
		DependsOn:  r.dependsOn,
	})
}

// UnmarshalJSON is a custom JSON unmarshalling hook that strips the outer
// AWS CloudFormation resource object, and just keeps the 'Properties' field.
func (r *AWSDataPipelinePipeline) UnmarshalJSON(b []byte) error {
	type Properties AWSDataPipelinePipeline
	res := &struct {
		Type       string
		Properties *Properties
		DependsOn  []string
	}{}
	if err := json.Unmarshal(b, &res); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return err
	}

	// If the resource has no Properties set, it could be nil
	if res.Properties != nil {
		*r = AWSDataPipelinePipeline(*res.Properties)
	}

	if res.DependsOn != nil {
		r.dependsOn = res.DependsOn
	}

	return nil
}

// GetAllAWSDataPipelinePipelineResources retrieves all AWSDataPipelinePipeline items from an AWS CloudFormation template
func (t *Template) GetAllAWSDataPipelinePipelineResources() map[string]AWSDataPipelinePipeline {
	results := map[string]AWSDataPipelinePipeline{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSDataPipelinePipeline:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::DataPipeline::Pipeline" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSDataPipelinePipeline
						if err := json.Unmarshal(b, &result); err == nil {
							results[name] = result
						}
					}
				}
			}
		}
	}
	return results
}

// GetAWSDataPipelinePipelineWithName retrieves all AWSDataPipelinePipeline items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSDataPipelinePipelineWithName(name string) (AWSDataPipelinePipeline, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSDataPipelinePipeline:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::DataPipeline::Pipeline" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSDataPipelinePipeline
						if err := json.Unmarshal(b, &result); err == nil {
							return result, nil
						}
					}
				}
			}
		}
	}
	return AWSDataPipelinePipeline{}, errors.New("resource not found")
}

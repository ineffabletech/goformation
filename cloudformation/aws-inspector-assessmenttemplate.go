package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSInspectorAssessmentTemplate AWS CloudFormation Resource (AWS::Inspector::AssessmentTemplate)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspector-assessmenttemplate.html
type AWSInspectorAssessmentTemplate struct {
	dependsOn []string

	// AssessmentTargetArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspector-assessmenttemplate.html#cfn-inspector-assessmenttemplate-assessmenttargetarn
	AssessmentTargetArn *String `json:"AssessmentTargetArn,omitempty"`

	// AssessmentTemplateName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspector-assessmenttemplate.html#cfn-inspector-assessmenttemplate-assessmenttemplatename
	AssessmentTemplateName *String `json:"AssessmentTemplateName,omitempty"`

	// DurationInSeconds AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspector-assessmenttemplate.html#cfn-inspector-assessmenttemplate-durationinseconds
	DurationInSeconds *Integer `json:"DurationInSeconds,omitempty"`

	// RulesPackageArns AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspector-assessmenttemplate.html#cfn-inspector-assessmenttemplate-rulespackagearns
	RulesPackageArns []*String `json:"RulesPackageArns,omitempty"`

	// UserAttributesForFindings AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspector-assessmenttemplate.html#cfn-inspector-assessmenttemplate-userattributesforfindings
	UserAttributesForFindings []Tag `json:"UserAttributesForFindings,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSInspectorAssessmentTemplate) AddDependencies(v ...string) *AWSInspectorAssessmentTemplate {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSInspectorAssessmentTemplate) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSInspectorAssessmentTemplate) AWSCloudFormationType() string {
	return "AWS::Inspector::AssessmentTemplate"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSInspectorAssessmentTemplate) MarshalJSON() ([]byte, error) {
	type Properties AWSInspectorAssessmentTemplate
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
func (r *AWSInspectorAssessmentTemplate) UnmarshalJSON(b []byte) error {
	type Properties AWSInspectorAssessmentTemplate
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
		*r = AWSInspectorAssessmentTemplate(*res.Properties)
	}

	if res.DependsOn != nil {
		r.dependsOn = res.DependsOn
	}

	return nil
}

// GetAllAWSInspectorAssessmentTemplateResources retrieves all AWSInspectorAssessmentTemplate items from an AWS CloudFormation template
func (t *Template) GetAllAWSInspectorAssessmentTemplateResources() map[string]AWSInspectorAssessmentTemplate {
	results := map[string]AWSInspectorAssessmentTemplate{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSInspectorAssessmentTemplate:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::Inspector::AssessmentTemplate" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSInspectorAssessmentTemplate
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

// GetAWSInspectorAssessmentTemplateWithName retrieves all AWSInspectorAssessmentTemplate items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSInspectorAssessmentTemplateWithName(name string) (AWSInspectorAssessmentTemplate, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSInspectorAssessmentTemplate:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::Inspector::AssessmentTemplate" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSInspectorAssessmentTemplate
						if err := json.Unmarshal(b, &result); err == nil {
							return result, nil
						}
					}
				}
			}
		}
	}
	return AWSInspectorAssessmentTemplate{}, errors.New("resource not found")
}

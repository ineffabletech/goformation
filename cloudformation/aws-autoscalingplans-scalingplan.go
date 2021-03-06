package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSAutoScalingPlansScalingPlan AWS CloudFormation Resource (AWS::AutoScalingPlans::ScalingPlan)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscalingplans-scalingplan.html
type AWSAutoScalingPlansScalingPlan struct {
	dependsOn []string

	// ApplicationSource AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscalingplans-scalingplan.html#cfn-autoscalingplans-scalingplan-applicationsource
	ApplicationSource *AWSAutoScalingPlansScalingPlan_ApplicationSource `json:"ApplicationSource,omitempty"`

	// ScalingInstructions AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscalingplans-scalingplan.html#cfn-autoscalingplans-scalingplan-scalinginstructions
	ScalingInstructions []AWSAutoScalingPlansScalingPlan_ScalingInstruction `json:"ScalingInstructions,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSAutoScalingPlansScalingPlan) AddDependencies(v ...string) *AWSAutoScalingPlansScalingPlan {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSAutoScalingPlansScalingPlan) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSAutoScalingPlansScalingPlan) AWSCloudFormationType() string {
	return "AWS::AutoScalingPlans::ScalingPlan"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSAutoScalingPlansScalingPlan) MarshalJSON() ([]byte, error) {
	type Properties AWSAutoScalingPlansScalingPlan
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
func (r *AWSAutoScalingPlansScalingPlan) UnmarshalJSON(b []byte) error {
	type Properties AWSAutoScalingPlansScalingPlan
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
		*r = AWSAutoScalingPlansScalingPlan(*res.Properties)
	}

	if res.DependsOn != nil {
		r.dependsOn = res.DependsOn
	}

	return nil
}

// GetAllAWSAutoScalingPlansScalingPlanResources retrieves all AWSAutoScalingPlansScalingPlan items from an AWS CloudFormation template
func (t *Template) GetAllAWSAutoScalingPlansScalingPlanResources() map[string]AWSAutoScalingPlansScalingPlan {
	results := map[string]AWSAutoScalingPlansScalingPlan{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSAutoScalingPlansScalingPlan:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::AutoScalingPlans::ScalingPlan" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSAutoScalingPlansScalingPlan
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

// GetAWSAutoScalingPlansScalingPlanWithName retrieves all AWSAutoScalingPlansScalingPlan items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSAutoScalingPlansScalingPlanWithName(name string) (AWSAutoScalingPlansScalingPlan, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSAutoScalingPlansScalingPlan:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::AutoScalingPlans::ScalingPlan" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSAutoScalingPlansScalingPlan
						if err := json.Unmarshal(b, &result); err == nil {
							return result, nil
						}
					}
				}
			}
		}
	}
	return AWSAutoScalingPlansScalingPlan{}, errors.New("resource not found")
}

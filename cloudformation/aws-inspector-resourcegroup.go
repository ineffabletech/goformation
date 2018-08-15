package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSInspectorResourceGroup AWS CloudFormation Resource (AWS::Inspector::ResourceGroup)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspector-resourcegroup.html
type AWSInspectorResourceGroup struct {
	dependsOn []string

	// ResourceGroupTags AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspector-resourcegroup.html#cfn-inspector-resourcegroup-resourcegrouptags
	ResourceGroupTags []Tag `json:"ResourceGroupTags,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSInspectorResourceGroup) AddDependencies(v ...string) *AWSInspectorResourceGroup {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSInspectorResourceGroup) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSInspectorResourceGroup) AWSCloudFormationType() string {
	return "AWS::Inspector::ResourceGroup"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSInspectorResourceGroup) MarshalJSON() ([]byte, error) {
	type Properties AWSInspectorResourceGroup
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
func (r *AWSInspectorResourceGroup) UnmarshalJSON(b []byte) error {
	type Properties AWSInspectorResourceGroup
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
		*r = AWSInspectorResourceGroup(*res.Properties)
	}

	if res.DependsOn != nil {
		r.dependsOn = res.DependsOn
	}

	return nil
}

// GetAllAWSInspectorResourceGroupResources retrieves all AWSInspectorResourceGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSInspectorResourceGroupResources() map[string]AWSInspectorResourceGroup {
	results := map[string]AWSInspectorResourceGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSInspectorResourceGroup:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::Inspector::ResourceGroup" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSInspectorResourceGroup
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

// GetAWSInspectorResourceGroupWithName retrieves all AWSInspectorResourceGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSInspectorResourceGroupWithName(name string) (AWSInspectorResourceGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSInspectorResourceGroup:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::Inspector::ResourceGroup" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSInspectorResourceGroup
						if err := json.Unmarshal(b, &result); err == nil {
							return result, nil
						}
					}
				}
			}
		}
	}
	return AWSInspectorResourceGroup{}, errors.New("resource not found")
}

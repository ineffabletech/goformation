package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSCloudFormationCustomResource AWS CloudFormation Resource (AWS::CloudFormation::CustomResource)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cfn-customresource.html
type AWSCloudFormationCustomResource struct {
	dependsOn []string

	// ServiceToken AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cfn-customresource.html#cfn-customresource-servicetoken
	ServiceToken *String `json:"ServiceToken,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSCloudFormationCustomResource) AddDependencies(v ...string) *AWSCloudFormationCustomResource {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSCloudFormationCustomResource) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCloudFormationCustomResource) AWSCloudFormationType() string {
	return "AWS::CloudFormation::CustomResource"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSCloudFormationCustomResource) MarshalJSON() ([]byte, error) {
	type Properties AWSCloudFormationCustomResource
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
func (r *AWSCloudFormationCustomResource) UnmarshalJSON(b []byte) error {
	type Properties AWSCloudFormationCustomResource
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
		*r = AWSCloudFormationCustomResource(*res.Properties)
	}

	if res.DependsOn != nil {
		r.dependsOn = res.DependsOn
	}

	return nil
}

// GetAllAWSCloudFormationCustomResourceResources retrieves all AWSCloudFormationCustomResource items from an AWS CloudFormation template
func (t *Template) GetAllAWSCloudFormationCustomResourceResources() map[string]AWSCloudFormationCustomResource {
	results := map[string]AWSCloudFormationCustomResource{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSCloudFormationCustomResource:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::CloudFormation::CustomResource" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSCloudFormationCustomResource
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

// GetAWSCloudFormationCustomResourceWithName retrieves all AWSCloudFormationCustomResource items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCloudFormationCustomResourceWithName(name string) (AWSCloudFormationCustomResource, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSCloudFormationCustomResource:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::CloudFormation::CustomResource" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSCloudFormationCustomResource
						if err := json.Unmarshal(b, &result); err == nil {
							return result, nil
						}
					}
				}
			}
		}
	}
	return AWSCloudFormationCustomResource{}, errors.New("resource not found")
}

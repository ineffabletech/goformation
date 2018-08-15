package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSConfigConfigurationRecorder AWS CloudFormation Resource (AWS::Config::ConfigurationRecorder)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configurationrecorder.html
type AWSConfigConfigurationRecorder struct {
	dependsOn []string

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configurationrecorder.html#cfn-config-configurationrecorder-name
	Name *String `json:"Name,omitempty"`

	// RecordingGroup AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configurationrecorder.html#cfn-config-configurationrecorder-recordinggroup
	RecordingGroup *AWSConfigConfigurationRecorder_RecordingGroup `json:"RecordingGroup,omitempty"`

	// RoleARN AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configurationrecorder.html#cfn-config-configurationrecorder-rolearn
	RoleARN *String `json:"RoleARN,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSConfigConfigurationRecorder) AddDependencies(v ...string) *AWSConfigConfigurationRecorder {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSConfigConfigurationRecorder) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSConfigConfigurationRecorder) AWSCloudFormationType() string {
	return "AWS::Config::ConfigurationRecorder"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSConfigConfigurationRecorder) MarshalJSON() ([]byte, error) {
	type Properties AWSConfigConfigurationRecorder
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
func (r *AWSConfigConfigurationRecorder) UnmarshalJSON(b []byte) error {
	type Properties AWSConfigConfigurationRecorder
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
		*r = AWSConfigConfigurationRecorder(*res.Properties)
	}

	if res.DependsOn != nil {
		r.dependsOn = res.DependsOn
	}

	return nil
}

// GetAllAWSConfigConfigurationRecorderResources retrieves all AWSConfigConfigurationRecorder items from an AWS CloudFormation template
func (t *Template) GetAllAWSConfigConfigurationRecorderResources() map[string]AWSConfigConfigurationRecorder {
	results := map[string]AWSConfigConfigurationRecorder{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSConfigConfigurationRecorder:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::Config::ConfigurationRecorder" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSConfigConfigurationRecorder
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

// GetAWSConfigConfigurationRecorderWithName retrieves all AWSConfigConfigurationRecorder items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSConfigConfigurationRecorderWithName(name string) (AWSConfigConfigurationRecorder, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSConfigConfigurationRecorder:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::Config::ConfigurationRecorder" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSConfigConfigurationRecorder
						if err := json.Unmarshal(b, &result); err == nil {
							return result, nil
						}
					}
				}
			}
		}
	}
	return AWSConfigConfigurationRecorder{}, errors.New("resource not found")
}

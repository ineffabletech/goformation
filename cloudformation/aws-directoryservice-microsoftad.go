package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSDirectoryServiceMicrosoftAD AWS CloudFormation Resource (AWS::DirectoryService::MicrosoftAD)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-microsoftad.html
type AWSDirectoryServiceMicrosoftAD struct {
	dependsOn []string

	// CreateAlias AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-microsoftad.html#cfn-directoryservice-microsoftad-createalias
	CreateAlias *Boolean `json:"CreateAlias,omitempty"`

	// Edition AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-microsoftad.html#cfn-directoryservice-microsoftad-edition
	Edition *String `json:"Edition,omitempty"`

	// EnableSso AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-microsoftad.html#cfn-directoryservice-microsoftad-enablesso
	EnableSso *Boolean `json:"EnableSso,omitempty"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-microsoftad.html#cfn-directoryservice-microsoftad-name
	Name *String `json:"Name,omitempty"`

	// Password AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-microsoftad.html#cfn-directoryservice-microsoftad-password
	Password *String `json:"Password,omitempty"`

	// ShortName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-microsoftad.html#cfn-directoryservice-microsoftad-shortname
	ShortName *String `json:"ShortName,omitempty"`

	// VpcSettings AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-microsoftad.html#cfn-directoryservice-microsoftad-vpcsettings
	VpcSettings *AWSDirectoryServiceMicrosoftAD_VpcSettings `json:"VpcSettings,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSDirectoryServiceMicrosoftAD) AddDependencies(v ...string) *AWSDirectoryServiceMicrosoftAD {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSDirectoryServiceMicrosoftAD) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDirectoryServiceMicrosoftAD) AWSCloudFormationType() string {
	return "AWS::DirectoryService::MicrosoftAD"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSDirectoryServiceMicrosoftAD) MarshalJSON() ([]byte, error) {
	type Properties AWSDirectoryServiceMicrosoftAD
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
func (r *AWSDirectoryServiceMicrosoftAD) UnmarshalJSON(b []byte) error {
	type Properties AWSDirectoryServiceMicrosoftAD
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
		*r = AWSDirectoryServiceMicrosoftAD(*res.Properties)
	}

	if res.DependsOn != nil {
		r.dependsOn = res.DependsOn
	}

	return nil
}

// GetAllAWSDirectoryServiceMicrosoftADResources retrieves all AWSDirectoryServiceMicrosoftAD items from an AWS CloudFormation template
func (t *Template) GetAllAWSDirectoryServiceMicrosoftADResources() map[string]AWSDirectoryServiceMicrosoftAD {
	results := map[string]AWSDirectoryServiceMicrosoftAD{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSDirectoryServiceMicrosoftAD:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::DirectoryService::MicrosoftAD" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSDirectoryServiceMicrosoftAD
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

// GetAWSDirectoryServiceMicrosoftADWithName retrieves all AWSDirectoryServiceMicrosoftAD items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSDirectoryServiceMicrosoftADWithName(name string) (AWSDirectoryServiceMicrosoftAD, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSDirectoryServiceMicrosoftAD:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::DirectoryService::MicrosoftAD" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSDirectoryServiceMicrosoftAD
						if err := json.Unmarshal(b, &result); err == nil {
							return result, nil
						}
					}
				}
			}
		}
	}
	return AWSDirectoryServiceMicrosoftAD{}, errors.New("resource not found")
}

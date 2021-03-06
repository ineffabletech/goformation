package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSServiceDiscoveryPublicDnsNamespace AWS CloudFormation Resource (AWS::ServiceDiscovery::PublicDnsNamespace)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicediscovery-publicdnsnamespace.html
type AWSServiceDiscoveryPublicDnsNamespace struct {
	dependsOn []string

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicediscovery-publicdnsnamespace.html#cfn-servicediscovery-publicdnsnamespace-description
	Description *String `json:"Description,omitempty"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicediscovery-publicdnsnamespace.html#cfn-servicediscovery-publicdnsnamespace-name
	Name *String `json:"Name,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSServiceDiscoveryPublicDnsNamespace) AddDependencies(v ...string) *AWSServiceDiscoveryPublicDnsNamespace {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSServiceDiscoveryPublicDnsNamespace) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSServiceDiscoveryPublicDnsNamespace) AWSCloudFormationType() string {
	return "AWS::ServiceDiscovery::PublicDnsNamespace"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSServiceDiscoveryPublicDnsNamespace) MarshalJSON() ([]byte, error) {
	type Properties AWSServiceDiscoveryPublicDnsNamespace
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
func (r *AWSServiceDiscoveryPublicDnsNamespace) UnmarshalJSON(b []byte) error {
	type Properties AWSServiceDiscoveryPublicDnsNamespace
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
		*r = AWSServiceDiscoveryPublicDnsNamespace(*res.Properties)
	}

	if res.DependsOn != nil {
		r.dependsOn = res.DependsOn
	}

	return nil
}

// GetAllAWSServiceDiscoveryPublicDnsNamespaceResources retrieves all AWSServiceDiscoveryPublicDnsNamespace items from an AWS CloudFormation template
func (t *Template) GetAllAWSServiceDiscoveryPublicDnsNamespaceResources() map[string]AWSServiceDiscoveryPublicDnsNamespace {
	results := map[string]AWSServiceDiscoveryPublicDnsNamespace{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSServiceDiscoveryPublicDnsNamespace:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::ServiceDiscovery::PublicDnsNamespace" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSServiceDiscoveryPublicDnsNamespace
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

// GetAWSServiceDiscoveryPublicDnsNamespaceWithName retrieves all AWSServiceDiscoveryPublicDnsNamespace items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSServiceDiscoveryPublicDnsNamespaceWithName(name string) (AWSServiceDiscoveryPublicDnsNamespace, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSServiceDiscoveryPublicDnsNamespace:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::ServiceDiscovery::PublicDnsNamespace" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSServiceDiscoveryPublicDnsNamespace
						if err := json.Unmarshal(b, &result); err == nil {
							return result, nil
						}
					}
				}
			}
		}
	}
	return AWSServiceDiscoveryPublicDnsNamespace{}, errors.New("resource not found")
}

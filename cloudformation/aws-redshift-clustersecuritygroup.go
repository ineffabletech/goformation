package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSRedshiftClusterSecurityGroup AWS CloudFormation Resource (AWS::Redshift::ClusterSecurityGroup)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clustersecuritygroup.html
type AWSRedshiftClusterSecurityGroup struct {
	dependsOn []string

	// Description AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clustersecuritygroup.html#cfn-redshift-clustersecuritygroup-description
	Description *String `json:"Description,omitempty"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clustersecuritygroup.html#cfn-redshift-clustersecuritygroup-tags
	Tags []Tag `json:"Tags,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSRedshiftClusterSecurityGroup) AddDependencies(v ...string) *AWSRedshiftClusterSecurityGroup {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSRedshiftClusterSecurityGroup) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSRedshiftClusterSecurityGroup) AWSCloudFormationType() string {
	return "AWS::Redshift::ClusterSecurityGroup"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSRedshiftClusterSecurityGroup) MarshalJSON() ([]byte, error) {
	type Properties AWSRedshiftClusterSecurityGroup
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
func (r *AWSRedshiftClusterSecurityGroup) UnmarshalJSON(b []byte) error {
	type Properties AWSRedshiftClusterSecurityGroup
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
		*r = AWSRedshiftClusterSecurityGroup(*res.Properties)
	}

	if res.DependsOn != nil {
		r.dependsOn = res.DependsOn
	}

	return nil
}

// GetAllAWSRedshiftClusterSecurityGroupResources retrieves all AWSRedshiftClusterSecurityGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSRedshiftClusterSecurityGroupResources() map[string]AWSRedshiftClusterSecurityGroup {
	results := map[string]AWSRedshiftClusterSecurityGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSRedshiftClusterSecurityGroup:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::Redshift::ClusterSecurityGroup" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSRedshiftClusterSecurityGroup
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

// GetAWSRedshiftClusterSecurityGroupWithName retrieves all AWSRedshiftClusterSecurityGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRedshiftClusterSecurityGroupWithName(name string) (AWSRedshiftClusterSecurityGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSRedshiftClusterSecurityGroup:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::Redshift::ClusterSecurityGroup" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSRedshiftClusterSecurityGroup
						if err := json.Unmarshal(b, &result); err == nil {
							return result, nil
						}
					}
				}
			}
		}
	}
	return AWSRedshiftClusterSecurityGroup{}, errors.New("resource not found")
}

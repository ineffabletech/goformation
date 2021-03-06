package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSAppSyncGraphQLSchema AWS CloudFormation Resource (AWS::AppSync::GraphQLSchema)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlschema.html
type AWSAppSyncGraphQLSchema struct {
	dependsOn []string

	// ApiId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlschema.html#cfn-appsync-graphqlschema-apiid
	ApiId *String `json:"ApiId,omitempty"`

	// Definition AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlschema.html#cfn-appsync-graphqlschema-definition
	Definition *String `json:"Definition,omitempty"`

	// DefinitionS3Location AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlschema.html#cfn-appsync-graphqlschema-definitions3location
	DefinitionS3Location *String `json:"DefinitionS3Location,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSAppSyncGraphQLSchema) AddDependencies(v ...string) *AWSAppSyncGraphQLSchema {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSAppSyncGraphQLSchema) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSAppSyncGraphQLSchema) AWSCloudFormationType() string {
	return "AWS::AppSync::GraphQLSchema"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSAppSyncGraphQLSchema) MarshalJSON() ([]byte, error) {
	type Properties AWSAppSyncGraphQLSchema
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
func (r *AWSAppSyncGraphQLSchema) UnmarshalJSON(b []byte) error {
	type Properties AWSAppSyncGraphQLSchema
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
		*r = AWSAppSyncGraphQLSchema(*res.Properties)
	}

	if res.DependsOn != nil {
		r.dependsOn = res.DependsOn
	}

	return nil
}

// GetAllAWSAppSyncGraphQLSchemaResources retrieves all AWSAppSyncGraphQLSchema items from an AWS CloudFormation template
func (t *Template) GetAllAWSAppSyncGraphQLSchemaResources() map[string]AWSAppSyncGraphQLSchema {
	results := map[string]AWSAppSyncGraphQLSchema{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSAppSyncGraphQLSchema:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::AppSync::GraphQLSchema" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSAppSyncGraphQLSchema
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

// GetAWSAppSyncGraphQLSchemaWithName retrieves all AWSAppSyncGraphQLSchema items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSAppSyncGraphQLSchemaWithName(name string) (AWSAppSyncGraphQLSchema, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSAppSyncGraphQLSchema:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::AppSync::GraphQLSchema" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSAppSyncGraphQLSchema
						if err := json.Unmarshal(b, &result); err == nil {
							return result, nil
						}
					}
				}
			}
		}
	}
	return AWSAppSyncGraphQLSchema{}, errors.New("resource not found")
}

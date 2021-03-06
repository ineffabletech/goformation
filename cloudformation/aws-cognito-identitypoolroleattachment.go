package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSCognitoIdentityPoolRoleAttachment AWS CloudFormation Resource (AWS::Cognito::IdentityPoolRoleAttachment)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypoolroleattachment.html
type AWSCognitoIdentityPoolRoleAttachment struct {
	dependsOn []string

	// IdentityPoolId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypoolroleattachment.html#cfn-cognito-identitypoolroleattachment-identitypoolid
	IdentityPoolId *String `json:"IdentityPoolId,omitempty"`

	// RoleMappings AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypoolroleattachment.html#cfn-cognito-identitypoolroleattachment-rolemappings
	RoleMappings interface{} `json:"RoleMappings,omitempty"`

	// Roles AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypoolroleattachment.html#cfn-cognito-identitypoolroleattachment-roles
	Roles interface{} `json:"Roles,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSCognitoIdentityPoolRoleAttachment) AddDependencies(v ...string) *AWSCognitoIdentityPoolRoleAttachment {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSCognitoIdentityPoolRoleAttachment) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCognitoIdentityPoolRoleAttachment) AWSCloudFormationType() string {
	return "AWS::Cognito::IdentityPoolRoleAttachment"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSCognitoIdentityPoolRoleAttachment) MarshalJSON() ([]byte, error) {
	type Properties AWSCognitoIdentityPoolRoleAttachment
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
func (r *AWSCognitoIdentityPoolRoleAttachment) UnmarshalJSON(b []byte) error {
	type Properties AWSCognitoIdentityPoolRoleAttachment
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
		*r = AWSCognitoIdentityPoolRoleAttachment(*res.Properties)
	}

	if res.DependsOn != nil {
		r.dependsOn = res.DependsOn
	}

	return nil
}

// GetAllAWSCognitoIdentityPoolRoleAttachmentResources retrieves all AWSCognitoIdentityPoolRoleAttachment items from an AWS CloudFormation template
func (t *Template) GetAllAWSCognitoIdentityPoolRoleAttachmentResources() map[string]AWSCognitoIdentityPoolRoleAttachment {
	results := map[string]AWSCognitoIdentityPoolRoleAttachment{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSCognitoIdentityPoolRoleAttachment:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::Cognito::IdentityPoolRoleAttachment" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSCognitoIdentityPoolRoleAttachment
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

// GetAWSCognitoIdentityPoolRoleAttachmentWithName retrieves all AWSCognitoIdentityPoolRoleAttachment items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCognitoIdentityPoolRoleAttachmentWithName(name string) (AWSCognitoIdentityPoolRoleAttachment, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSCognitoIdentityPoolRoleAttachment:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::Cognito::IdentityPoolRoleAttachment" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSCognitoIdentityPoolRoleAttachment
						if err := json.Unmarshal(b, &result); err == nil {
							return result, nil
						}
					}
				}
			}
		}
	}
	return AWSCognitoIdentityPoolRoleAttachment{}, errors.New("resource not found")
}

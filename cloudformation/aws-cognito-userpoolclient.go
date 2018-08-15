package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSCognitoUserPoolClient AWS CloudFormation Resource (AWS::Cognito::UserPoolClient)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclient.html
type AWSCognitoUserPoolClient struct {
	dependsOn []string

	// ClientName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclient.html#cfn-cognito-userpoolclient-clientname
	ClientName *String `json:"ClientName,omitempty"`

	// ExplicitAuthFlows AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclient.html#cfn-cognito-userpoolclient-explicitauthflows
	ExplicitAuthFlows []*String `json:"ExplicitAuthFlows,omitempty"`

	// GenerateSecret AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclient.html#cfn-cognito-userpoolclient-generatesecret
	GenerateSecret *Boolean `json:"GenerateSecret,omitempty"`

	// ReadAttributes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclient.html#cfn-cognito-userpoolclient-readattributes
	ReadAttributes []*String `json:"ReadAttributes,omitempty"`

	// RefreshTokenValidity AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclient.html#cfn-cognito-userpoolclient-refreshtokenvalidity
	RefreshTokenValidity *Double `json:"RefreshTokenValidity,omitempty"`

	// UserPoolId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclient.html#cfn-cognito-userpoolclient-userpoolid
	UserPoolId *String `json:"UserPoolId,omitempty"`

	// WriteAttributes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclient.html#cfn-cognito-userpoolclient-writeattributes
	WriteAttributes []*String `json:"WriteAttributes,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSCognitoUserPoolClient) AddDependencies(v ...string) *AWSCognitoUserPoolClient {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSCognitoUserPoolClient) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCognitoUserPoolClient) AWSCloudFormationType() string {
	return "AWS::Cognito::UserPoolClient"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSCognitoUserPoolClient) MarshalJSON() ([]byte, error) {
	type Properties AWSCognitoUserPoolClient
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
func (r *AWSCognitoUserPoolClient) UnmarshalJSON(b []byte) error {
	type Properties AWSCognitoUserPoolClient
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
		*r = AWSCognitoUserPoolClient(*res.Properties)
	}

	if res.DependsOn != nil {
		r.dependsOn = res.DependsOn
	}

	return nil
}

// GetAllAWSCognitoUserPoolClientResources retrieves all AWSCognitoUserPoolClient items from an AWS CloudFormation template
func (t *Template) GetAllAWSCognitoUserPoolClientResources() map[string]AWSCognitoUserPoolClient {
	results := map[string]AWSCognitoUserPoolClient{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSCognitoUserPoolClient:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::Cognito::UserPoolClient" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSCognitoUserPoolClient
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

// GetAWSCognitoUserPoolClientWithName retrieves all AWSCognitoUserPoolClient items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCognitoUserPoolClientWithName(name string) (AWSCognitoUserPoolClient, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSCognitoUserPoolClient:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::Cognito::UserPoolClient" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSCognitoUserPoolClient
						if err := json.Unmarshal(b, &result); err == nil {
							return result, nil
						}
					}
				}
			}
		}
	}
	return AWSCognitoUserPoolClient{}, errors.New("resource not found")
}

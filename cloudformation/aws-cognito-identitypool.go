package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSCognitoIdentityPool AWS CloudFormation Resource (AWS::Cognito::IdentityPool)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html
type AWSCognitoIdentityPool struct {
	dependsOn []string

	// AllowUnauthenticatedIdentities AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-allowunauthenticatedidentities
	AllowUnauthenticatedIdentities *Boolean `json:"AllowUnauthenticatedIdentities,omitempty"`

	// CognitoEvents AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-cognitoevents
	CognitoEvents interface{} `json:"CognitoEvents,omitempty"`

	// CognitoIdentityProviders AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-cognitoidentityproviders
	CognitoIdentityProviders []AWSCognitoIdentityPool_CognitoIdentityProvider `json:"CognitoIdentityProviders,omitempty"`

	// CognitoStreams AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-cognitostreams
	CognitoStreams *AWSCognitoIdentityPool_CognitoStreams `json:"CognitoStreams,omitempty"`

	// DeveloperProviderName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-developerprovidername
	DeveloperProviderName *String `json:"DeveloperProviderName,omitempty"`

	// IdentityPoolName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-identitypoolname
	IdentityPoolName *String `json:"IdentityPoolName,omitempty"`

	// OpenIdConnectProviderARNs AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-openidconnectproviderarns
	OpenIdConnectProviderARNs []*String `json:"OpenIdConnectProviderARNs,omitempty"`

	// PushSync AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-pushsync
	PushSync *AWSCognitoIdentityPool_PushSync `json:"PushSync,omitempty"`

	// SamlProviderARNs AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-samlproviderarns
	SamlProviderARNs []*String `json:"SamlProviderARNs,omitempty"`

	// SupportedLoginProviders AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-supportedloginproviders
	SupportedLoginProviders interface{} `json:"SupportedLoginProviders,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSCognitoIdentityPool) AddDependencies(v ...string) *AWSCognitoIdentityPool {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSCognitoIdentityPool) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCognitoIdentityPool) AWSCloudFormationType() string {
	return "AWS::Cognito::IdentityPool"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSCognitoIdentityPool) MarshalJSON() ([]byte, error) {
	type Properties AWSCognitoIdentityPool
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
func (r *AWSCognitoIdentityPool) UnmarshalJSON(b []byte) error {
	type Properties AWSCognitoIdentityPool
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
		*r = AWSCognitoIdentityPool(*res.Properties)
	}

	if res.DependsOn != nil {
		r.dependsOn = res.DependsOn
	}

	return nil
}

// GetAllAWSCognitoIdentityPoolResources retrieves all AWSCognitoIdentityPool items from an AWS CloudFormation template
func (t *Template) GetAllAWSCognitoIdentityPoolResources() map[string]AWSCognitoIdentityPool {
	results := map[string]AWSCognitoIdentityPool{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSCognitoIdentityPool:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::Cognito::IdentityPool" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSCognitoIdentityPool
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

// GetAWSCognitoIdentityPoolWithName retrieves all AWSCognitoIdentityPool items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCognitoIdentityPoolWithName(name string) (AWSCognitoIdentityPool, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSCognitoIdentityPool:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::Cognito::IdentityPool" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSCognitoIdentityPool
						if err := json.Unmarshal(b, &result); err == nil {
							return result, nil
						}
					}
				}
			}
		}
	}
	return AWSCognitoIdentityPool{}, errors.New("resource not found")
}

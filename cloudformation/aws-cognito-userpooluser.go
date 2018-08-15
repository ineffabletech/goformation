package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSCognitoUserPoolUser AWS CloudFormation Resource (AWS::Cognito::UserPoolUser)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluser.html
type AWSCognitoUserPoolUser struct {
	dependsOn []string

	// DesiredDeliveryMediums AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluser.html#cfn-cognito-userpooluser-desireddeliverymediums
	DesiredDeliveryMediums []*String `json:"DesiredDeliveryMediums,omitempty"`

	// ForceAliasCreation AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluser.html#cfn-cognito-userpooluser-forcealiascreation
	ForceAliasCreation *Boolean `json:"ForceAliasCreation,omitempty"`

	// MessageAction AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluser.html#cfn-cognito-userpooluser-messageaction
	MessageAction *String `json:"MessageAction,omitempty"`

	// UserAttributes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluser.html#cfn-cognito-userpooluser-userattributes
	UserAttributes []AWSCognitoUserPoolUser_AttributeType `json:"UserAttributes,omitempty"`

	// UserPoolId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluser.html#cfn-cognito-userpooluser-userpoolid
	UserPoolId *String `json:"UserPoolId,omitempty"`

	// Username AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluser.html#cfn-cognito-userpooluser-username
	Username *String `json:"Username,omitempty"`

	// ValidationData AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluser.html#cfn-cognito-userpooluser-validationdata
	ValidationData []AWSCognitoUserPoolUser_AttributeType `json:"ValidationData,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSCognitoUserPoolUser) AddDependencies(v ...string) *AWSCognitoUserPoolUser {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSCognitoUserPoolUser) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCognitoUserPoolUser) AWSCloudFormationType() string {
	return "AWS::Cognito::UserPoolUser"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSCognitoUserPoolUser) MarshalJSON() ([]byte, error) {
	type Properties AWSCognitoUserPoolUser
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
func (r *AWSCognitoUserPoolUser) UnmarshalJSON(b []byte) error {
	type Properties AWSCognitoUserPoolUser
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
		*r = AWSCognitoUserPoolUser(*res.Properties)
	}

	if res.DependsOn != nil {
		r.dependsOn = res.DependsOn
	}

	return nil
}

// GetAllAWSCognitoUserPoolUserResources retrieves all AWSCognitoUserPoolUser items from an AWS CloudFormation template
func (t *Template) GetAllAWSCognitoUserPoolUserResources() map[string]AWSCognitoUserPoolUser {
	results := map[string]AWSCognitoUserPoolUser{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSCognitoUserPoolUser:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::Cognito::UserPoolUser" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSCognitoUserPoolUser
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

// GetAWSCognitoUserPoolUserWithName retrieves all AWSCognitoUserPoolUser items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCognitoUserPoolUserWithName(name string) (AWSCognitoUserPoolUser, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSCognitoUserPoolUser:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::Cognito::UserPoolUser" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSCognitoUserPoolUser
						if err := json.Unmarshal(b, &result); err == nil {
							return result, nil
						}
					}
				}
			}
		}
	}
	return AWSCognitoUserPoolUser{}, errors.New("resource not found")
}

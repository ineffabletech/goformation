package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSServerlessFunction AWS CloudFormation Resource (AWS::Serverless::Function)
// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessfunction
type AWSServerlessFunction struct {
	dependsOn []string

	// CodeUri AWS CloudFormation Property
	// Required: true
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessfunction
	CodeUri *AWSServerlessFunction_CodeUri `json:"CodeUri,omitempty"`

	// DeadLetterQueue AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessfunction
	DeadLetterQueue *AWSServerlessFunction_DeadLetterQueue `json:"DeadLetterQueue,omitempty"`

	// Description AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessfunction
	Description *String `json:"Description,omitempty"`

	// Environment AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessfunction
	Environment *AWSServerlessFunction_FunctionEnvironment `json:"Environment,omitempty"`

	// Events AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessfunction
	Events map[string]AWSServerlessFunction_EventSource `json:"Events,omitempty"`

	// FunctionName AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessfunction
	FunctionName *String `json:"FunctionName,omitempty"`

	// Handler AWS CloudFormation Property
	// Required: true
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessfunction
	Handler *String `json:"Handler,omitempty"`

	// KmsKeyArn AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessfunction
	KmsKeyArn *String `json:"KmsKeyArn,omitempty"`

	// MemorySize AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessfunction
	MemorySize *Integer `json:"MemorySize,omitempty"`

	// Policies AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessfunction
	Policies *AWSServerlessFunction_Policies `json:"Policies,omitempty"`

	// Role AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessfunction
	Role *String `json:"Role,omitempty"`

	// Runtime AWS CloudFormation Property
	// Required: true
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessfunction
	Runtime *String `json:"Runtime,omitempty"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessfunction
	Tags map[string]*String `json:"Tags,omitempty"`

	// Timeout AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessfunction
	Timeout *Integer `json:"Timeout,omitempty"`

	// Tracing AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessfunction
	Tracing *String `json:"Tracing,omitempty"`

	// VpcConfig AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessfunction
	VpcConfig *AWSServerlessFunction_VpcConfig `json:"VpcConfig,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSServerlessFunction) AddDependencies(v ...string) *AWSServerlessFunction {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSServerlessFunction) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSServerlessFunction) AWSCloudFormationType() string {
	return "AWS::Serverless::Function"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSServerlessFunction) MarshalJSON() ([]byte, error) {
	type Properties AWSServerlessFunction
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
func (r *AWSServerlessFunction) UnmarshalJSON(b []byte) error {
	type Properties AWSServerlessFunction
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
		*r = AWSServerlessFunction(*res.Properties)
	}

	if res.DependsOn != nil {
		r.dependsOn = res.DependsOn
	}

	return nil
}

// GetAllAWSServerlessFunctionResources retrieves all AWSServerlessFunction items from an AWS CloudFormation template
func (t *Template) GetAllAWSServerlessFunctionResources() map[string]AWSServerlessFunction {
	results := map[string]AWSServerlessFunction{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSServerlessFunction:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::Serverless::Function" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSServerlessFunction
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

// GetAWSServerlessFunctionWithName retrieves all AWSServerlessFunction items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSServerlessFunctionWithName(name string) (AWSServerlessFunction, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSServerlessFunction:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::Serverless::Function" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSServerlessFunction
						if err := json.Unmarshal(b, &result); err == nil {
							return result, nil
						}
					}
				}
			}
		}
	}
	return AWSServerlessFunction{}, errors.New("resource not found")
}

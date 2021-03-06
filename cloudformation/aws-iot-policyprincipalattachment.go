package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSIoTPolicyPrincipalAttachment AWS CloudFormation Resource (AWS::IoT::PolicyPrincipalAttachment)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-policyprincipalattachment.html
type AWSIoTPolicyPrincipalAttachment struct {
	dependsOn []string

	// PolicyName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-policyprincipalattachment.html#cfn-iot-policyprincipalattachment-policyname
	PolicyName *String `json:"PolicyName,omitempty"`

	// Principal AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-policyprincipalattachment.html#cfn-iot-policyprincipalattachment-principal
	Principal *String `json:"Principal,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSIoTPolicyPrincipalAttachment) AddDependencies(v ...string) *AWSIoTPolicyPrincipalAttachment {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSIoTPolicyPrincipalAttachment) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIoTPolicyPrincipalAttachment) AWSCloudFormationType() string {
	return "AWS::IoT::PolicyPrincipalAttachment"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSIoTPolicyPrincipalAttachment) MarshalJSON() ([]byte, error) {
	type Properties AWSIoTPolicyPrincipalAttachment
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
func (r *AWSIoTPolicyPrincipalAttachment) UnmarshalJSON(b []byte) error {
	type Properties AWSIoTPolicyPrincipalAttachment
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
		*r = AWSIoTPolicyPrincipalAttachment(*res.Properties)
	}

	if res.DependsOn != nil {
		r.dependsOn = res.DependsOn
	}

	return nil
}

// GetAllAWSIoTPolicyPrincipalAttachmentResources retrieves all AWSIoTPolicyPrincipalAttachment items from an AWS CloudFormation template
func (t *Template) GetAllAWSIoTPolicyPrincipalAttachmentResources() map[string]AWSIoTPolicyPrincipalAttachment {
	results := map[string]AWSIoTPolicyPrincipalAttachment{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSIoTPolicyPrincipalAttachment:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::IoT::PolicyPrincipalAttachment" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSIoTPolicyPrincipalAttachment
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

// GetAWSIoTPolicyPrincipalAttachmentWithName retrieves all AWSIoTPolicyPrincipalAttachment items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSIoTPolicyPrincipalAttachmentWithName(name string) (AWSIoTPolicyPrincipalAttachment, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSIoTPolicyPrincipalAttachment:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::IoT::PolicyPrincipalAttachment" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSIoTPolicyPrincipalAttachment
						if err := json.Unmarshal(b, &result); err == nil {
							return result, nil
						}
					}
				}
			}
		}
	}
	return AWSIoTPolicyPrincipalAttachment{}, errors.New("resource not found")
}

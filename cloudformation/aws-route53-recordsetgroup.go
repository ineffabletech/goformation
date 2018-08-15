package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSRoute53RecordSetGroup AWS CloudFormation Resource (AWS::Route53::RecordSetGroup)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-recordsetgroup.html
type AWSRoute53RecordSetGroup struct {
	dependsOn []string

	// Comment AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-recordsetgroup.html#cfn-route53-recordsetgroup-comment
	Comment *String `json:"Comment,omitempty"`

	// HostedZoneId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-recordsetgroup.html#cfn-route53-recordsetgroup-hostedzoneid
	HostedZoneId *String `json:"HostedZoneId,omitempty"`

	// HostedZoneName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-recordsetgroup.html#cfn-route53-recordsetgroup-hostedzonename
	HostedZoneName *String `json:"HostedZoneName,omitempty"`

	// RecordSets AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-recordsetgroup.html#cfn-route53-recordsetgroup-recordsets
	RecordSets []AWSRoute53RecordSetGroup_RecordSet `json:"RecordSets,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSRoute53RecordSetGroup) AddDependencies(v ...string) *AWSRoute53RecordSetGroup {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSRoute53RecordSetGroup) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSRoute53RecordSetGroup) AWSCloudFormationType() string {
	return "AWS::Route53::RecordSetGroup"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSRoute53RecordSetGroup) MarshalJSON() ([]byte, error) {
	type Properties AWSRoute53RecordSetGroup
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
func (r *AWSRoute53RecordSetGroup) UnmarshalJSON(b []byte) error {
	type Properties AWSRoute53RecordSetGroup
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
		*r = AWSRoute53RecordSetGroup(*res.Properties)
	}

	if res.DependsOn != nil {
		r.dependsOn = res.DependsOn
	}

	return nil
}

// GetAllAWSRoute53RecordSetGroupResources retrieves all AWSRoute53RecordSetGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSRoute53RecordSetGroupResources() map[string]AWSRoute53RecordSetGroup {
	results := map[string]AWSRoute53RecordSetGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSRoute53RecordSetGroup:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::Route53::RecordSetGroup" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSRoute53RecordSetGroup
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

// GetAWSRoute53RecordSetGroupWithName retrieves all AWSRoute53RecordSetGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRoute53RecordSetGroupWithName(name string) (AWSRoute53RecordSetGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSRoute53RecordSetGroup:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::Route53::RecordSetGroup" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSRoute53RecordSetGroup
						if err := json.Unmarshal(b, &result); err == nil {
							return result, nil
						}
					}
				}
			}
		}
	}
	return AWSRoute53RecordSetGroup{}, errors.New("resource not found")
}

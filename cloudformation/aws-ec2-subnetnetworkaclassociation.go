package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSEC2SubnetNetworkAclAssociation AWS CloudFormation Resource (AWS::EC2::SubnetNetworkAclAssociation)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet-network-acl-assoc.html
type AWSEC2SubnetNetworkAclAssociation struct {
	dependsOn []string

	// NetworkAclId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet-network-acl-assoc.html#cfn-ec2-subnetnetworkaclassociation-networkaclid
	NetworkAclId *String `json:"NetworkAclId,omitempty"`

	// SubnetId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet-network-acl-assoc.html#cfn-ec2-subnetnetworkaclassociation-associationid
	SubnetId *String `json:"SubnetId,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSEC2SubnetNetworkAclAssociation) AddDependencies(v ...string) *AWSEC2SubnetNetworkAclAssociation {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSEC2SubnetNetworkAclAssociation) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2SubnetNetworkAclAssociation) AWSCloudFormationType() string {
	return "AWS::EC2::SubnetNetworkAclAssociation"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSEC2SubnetNetworkAclAssociation) MarshalJSON() ([]byte, error) {
	type Properties AWSEC2SubnetNetworkAclAssociation
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
func (r *AWSEC2SubnetNetworkAclAssociation) UnmarshalJSON(b []byte) error {
	type Properties AWSEC2SubnetNetworkAclAssociation
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
		*r = AWSEC2SubnetNetworkAclAssociation(*res.Properties)
	}

	if res.DependsOn != nil {
		r.dependsOn = res.DependsOn
	}

	return nil
}

// GetAllAWSEC2SubnetNetworkAclAssociationResources retrieves all AWSEC2SubnetNetworkAclAssociation items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2SubnetNetworkAclAssociationResources() map[string]AWSEC2SubnetNetworkAclAssociation {
	results := map[string]AWSEC2SubnetNetworkAclAssociation{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSEC2SubnetNetworkAclAssociation:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::EC2::SubnetNetworkAclAssociation" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSEC2SubnetNetworkAclAssociation
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

// GetAWSEC2SubnetNetworkAclAssociationWithName retrieves all AWSEC2SubnetNetworkAclAssociation items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2SubnetNetworkAclAssociationWithName(name string) (AWSEC2SubnetNetworkAclAssociation, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSEC2SubnetNetworkAclAssociation:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::EC2::SubnetNetworkAclAssociation" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSEC2SubnetNetworkAclAssociation
						if err := json.Unmarshal(b, &result); err == nil {
							return result, nil
						}
					}
				}
			}
		}
	}
	return AWSEC2SubnetNetworkAclAssociation{}, errors.New("resource not found")
}

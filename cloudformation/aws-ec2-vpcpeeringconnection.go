package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSEC2VPCPeeringConnection AWS CloudFormation Resource (AWS::EC2::VPCPeeringConnection)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcpeeringconnection.html
type AWSEC2VPCPeeringConnection struct {
	dependsOn []string

	// PeerOwnerId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcpeeringconnection.html#cfn-ec2-vpcpeeringconnection-peerownerid
	PeerOwnerId *String `json:"PeerOwnerId,omitempty"`

	// PeerRoleArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcpeeringconnection.html#cfn-ec2-vpcpeeringconnection-peerrolearn
	PeerRoleArn *String `json:"PeerRoleArn,omitempty"`

	// PeerVpcId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcpeeringconnection.html#cfn-ec2-vpcpeeringconnection-peervpcid
	PeerVpcId *String `json:"PeerVpcId,omitempty"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcpeeringconnection.html#cfn-ec2-vpcpeeringconnection-tags
	Tags []Tag `json:"Tags,omitempty"`

	// VpcId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcpeeringconnection.html#cfn-ec2-vpcpeeringconnection-vpcid
	VpcId *String `json:"VpcId,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSEC2VPCPeeringConnection) AddDependencies(v ...string) *AWSEC2VPCPeeringConnection {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSEC2VPCPeeringConnection) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2VPCPeeringConnection) AWSCloudFormationType() string {
	return "AWS::EC2::VPCPeeringConnection"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSEC2VPCPeeringConnection) MarshalJSON() ([]byte, error) {
	type Properties AWSEC2VPCPeeringConnection
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
func (r *AWSEC2VPCPeeringConnection) UnmarshalJSON(b []byte) error {
	type Properties AWSEC2VPCPeeringConnection
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
		*r = AWSEC2VPCPeeringConnection(*res.Properties)
	}

	if res.DependsOn != nil {
		r.dependsOn = res.DependsOn
	}

	return nil
}

// GetAllAWSEC2VPCPeeringConnectionResources retrieves all AWSEC2VPCPeeringConnection items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2VPCPeeringConnectionResources() map[string]AWSEC2VPCPeeringConnection {
	results := map[string]AWSEC2VPCPeeringConnection{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSEC2VPCPeeringConnection:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::EC2::VPCPeeringConnection" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSEC2VPCPeeringConnection
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

// GetAWSEC2VPCPeeringConnectionWithName retrieves all AWSEC2VPCPeeringConnection items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2VPCPeeringConnectionWithName(name string) (AWSEC2VPCPeeringConnection, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSEC2VPCPeeringConnection:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::EC2::VPCPeeringConnection" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSEC2VPCPeeringConnection
						if err := json.Unmarshal(b, &result); err == nil {
							return result, nil
						}
					}
				}
			}
		}
	}
	return AWSEC2VPCPeeringConnection{}, errors.New("resource not found")
}

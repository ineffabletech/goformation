package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSServiceCatalogCloudFormationProduct AWS CloudFormation Resource (AWS::ServiceCatalog::CloudFormationProduct)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationproduct.html
type AWSServiceCatalogCloudFormationProduct struct {
	dependsOn []string

	// AcceptLanguage AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationproduct.html#cfn-servicecatalog-cloudformationproduct-acceptlanguage
	AcceptLanguage *String `json:"AcceptLanguage,omitempty"`

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationproduct.html#cfn-servicecatalog-cloudformationproduct-description
	Description *String `json:"Description,omitempty"`

	// Distributor AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationproduct.html#cfn-servicecatalog-cloudformationproduct-distributor
	Distributor *String `json:"Distributor,omitempty"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationproduct.html#cfn-servicecatalog-cloudformationproduct-name
	Name *String `json:"Name,omitempty"`

	// Owner AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationproduct.html#cfn-servicecatalog-cloudformationproduct-owner
	Owner *String `json:"Owner,omitempty"`

	// ProvisioningArtifactParameters AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationproduct.html#cfn-servicecatalog-cloudformationproduct-provisioningartifactparameters
	ProvisioningArtifactParameters []AWSServiceCatalogCloudFormationProduct_ProvisioningArtifactProperties `json:"ProvisioningArtifactParameters,omitempty"`

	// SupportDescription AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationproduct.html#cfn-servicecatalog-cloudformationproduct-supportdescription
	SupportDescription *String `json:"SupportDescription,omitempty"`

	// SupportEmail AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationproduct.html#cfn-servicecatalog-cloudformationproduct-supportemail
	SupportEmail *String `json:"SupportEmail,omitempty"`

	// SupportUrl AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationproduct.html#cfn-servicecatalog-cloudformationproduct-supporturl
	SupportUrl *String `json:"SupportUrl,omitempty"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationproduct.html#cfn-servicecatalog-cloudformationproduct-tags
	Tags []Tag `json:"Tags,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSServiceCatalogCloudFormationProduct) AddDependencies(v ...string) *AWSServiceCatalogCloudFormationProduct {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSServiceCatalogCloudFormationProduct) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSServiceCatalogCloudFormationProduct) AWSCloudFormationType() string {
	return "AWS::ServiceCatalog::CloudFormationProduct"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSServiceCatalogCloudFormationProduct) MarshalJSON() ([]byte, error) {
	type Properties AWSServiceCatalogCloudFormationProduct
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
func (r *AWSServiceCatalogCloudFormationProduct) UnmarshalJSON(b []byte) error {
	type Properties AWSServiceCatalogCloudFormationProduct
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
		*r = AWSServiceCatalogCloudFormationProduct(*res.Properties)
	}

	if res.DependsOn != nil {
		r.dependsOn = res.DependsOn
	}

	return nil
}

// GetAllAWSServiceCatalogCloudFormationProductResources retrieves all AWSServiceCatalogCloudFormationProduct items from an AWS CloudFormation template
func (t *Template) GetAllAWSServiceCatalogCloudFormationProductResources() map[string]AWSServiceCatalogCloudFormationProduct {
	results := map[string]AWSServiceCatalogCloudFormationProduct{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSServiceCatalogCloudFormationProduct:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::ServiceCatalog::CloudFormationProduct" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSServiceCatalogCloudFormationProduct
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

// GetAWSServiceCatalogCloudFormationProductWithName retrieves all AWSServiceCatalogCloudFormationProduct items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSServiceCatalogCloudFormationProductWithName(name string) (AWSServiceCatalogCloudFormationProduct, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSServiceCatalogCloudFormationProduct:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::ServiceCatalog::CloudFormationProduct" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSServiceCatalogCloudFormationProduct
						if err := json.Unmarshal(b, &result); err == nil {
							return result, nil
						}
					}
				}
			}
		}
	}
	return AWSServiceCatalogCloudFormationProduct{}, errors.New("resource not found")
}

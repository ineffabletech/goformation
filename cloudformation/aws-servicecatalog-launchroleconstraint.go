package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSServiceCatalogLaunchRoleConstraint AWS CloudFormation Resource (AWS::ServiceCatalog::LaunchRoleConstraint)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchroleconstraint.html
type AWSServiceCatalogLaunchRoleConstraint struct {
	dependsOn []string

	// AcceptLanguage AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchroleconstraint.html#cfn-servicecatalog-launchroleconstraint-acceptlanguage
	AcceptLanguage *String `json:"AcceptLanguage,omitempty"`

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchroleconstraint.html#cfn-servicecatalog-launchroleconstraint-description
	Description *String `json:"Description,omitempty"`

	// PortfolioId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchroleconstraint.html#cfn-servicecatalog-launchroleconstraint-portfolioid
	PortfolioId *String `json:"PortfolioId,omitempty"`

	// ProductId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchroleconstraint.html#cfn-servicecatalog-launchroleconstraint-productid
	ProductId *String `json:"ProductId,omitempty"`

	// RoleArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchroleconstraint.html#cfn-servicecatalog-launchroleconstraint-rolearn
	RoleArn *String `json:"RoleArn,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSServiceCatalogLaunchRoleConstraint) AddDependencies(v ...string) *AWSServiceCatalogLaunchRoleConstraint {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSServiceCatalogLaunchRoleConstraint) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSServiceCatalogLaunchRoleConstraint) AWSCloudFormationType() string {
	return "AWS::ServiceCatalog::LaunchRoleConstraint"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSServiceCatalogLaunchRoleConstraint) MarshalJSON() ([]byte, error) {
	type Properties AWSServiceCatalogLaunchRoleConstraint
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
func (r *AWSServiceCatalogLaunchRoleConstraint) UnmarshalJSON(b []byte) error {
	type Properties AWSServiceCatalogLaunchRoleConstraint
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
		*r = AWSServiceCatalogLaunchRoleConstraint(*res.Properties)
	}

	if res.DependsOn != nil {
		r.dependsOn = res.DependsOn
	}

	return nil
}

// GetAllAWSServiceCatalogLaunchRoleConstraintResources retrieves all AWSServiceCatalogLaunchRoleConstraint items from an AWS CloudFormation template
func (t *Template) GetAllAWSServiceCatalogLaunchRoleConstraintResources() map[string]AWSServiceCatalogLaunchRoleConstraint {
	results := map[string]AWSServiceCatalogLaunchRoleConstraint{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSServiceCatalogLaunchRoleConstraint:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::ServiceCatalog::LaunchRoleConstraint" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSServiceCatalogLaunchRoleConstraint
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

// GetAWSServiceCatalogLaunchRoleConstraintWithName retrieves all AWSServiceCatalogLaunchRoleConstraint items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSServiceCatalogLaunchRoleConstraintWithName(name string) (AWSServiceCatalogLaunchRoleConstraint, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSServiceCatalogLaunchRoleConstraint:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::ServiceCatalog::LaunchRoleConstraint" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSServiceCatalogLaunchRoleConstraint
						if err := json.Unmarshal(b, &result); err == nil {
							return result, nil
						}
					}
				}
			}
		}
	}
	return AWSServiceCatalogLaunchRoleConstraint{}, errors.New("resource not found")
}

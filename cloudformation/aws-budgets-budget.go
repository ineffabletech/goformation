package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSBudgetsBudget AWS CloudFormation Resource (AWS::Budgets::Budget)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-budgets-budget.html
type AWSBudgetsBudget struct {
	dependsOn []string

	// Budget AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-budgets-budget.html#cfn-budgets-budget-budget
	Budget *AWSBudgetsBudget_BudgetData `json:"Budget,omitempty"`

	// NotificationsWithSubscribers AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-budgets-budget.html#cfn-budgets-budget-notificationswithsubscribers
	NotificationsWithSubscribers []AWSBudgetsBudget_NotificationWithSubscribers `json:"NotificationsWithSubscribers,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSBudgetsBudget) AddDependencies(v ...string) *AWSBudgetsBudget {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSBudgetsBudget) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSBudgetsBudget) AWSCloudFormationType() string {
	return "AWS::Budgets::Budget"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSBudgetsBudget) MarshalJSON() ([]byte, error) {
	type Properties AWSBudgetsBudget
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
func (r *AWSBudgetsBudget) UnmarshalJSON(b []byte) error {
	type Properties AWSBudgetsBudget
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
		*r = AWSBudgetsBudget(*res.Properties)
	}

	if res.DependsOn != nil {
		r.dependsOn = res.DependsOn
	}

	return nil
}

// GetAllAWSBudgetsBudgetResources retrieves all AWSBudgetsBudget items from an AWS CloudFormation template
func (t *Template) GetAllAWSBudgetsBudgetResources() map[string]AWSBudgetsBudget {
	results := map[string]AWSBudgetsBudget{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSBudgetsBudget:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::Budgets::Budget" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSBudgetsBudget
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

// GetAWSBudgetsBudgetWithName retrieves all AWSBudgetsBudget items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSBudgetsBudgetWithName(name string) (AWSBudgetsBudget, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSBudgetsBudget:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::Budgets::Budget" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSBudgetsBudget
						if err := json.Unmarshal(b, &result); err == nil {
							return result, nil
						}
					}
				}
			}
		}
	}
	return AWSBudgetsBudget{}, errors.New("resource not found")
}

package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSAutoScalingScheduledAction AWS CloudFormation Resource (AWS::AutoScaling::ScheduledAction)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-scheduledaction.html
type AWSAutoScalingScheduledAction struct {
	dependsOn []string

	// AutoScalingGroupName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-scheduledaction.html#cfn-as-scheduledaction-asgname
	AutoScalingGroupName *String `json:"AutoScalingGroupName,omitempty"`

	// DesiredCapacity AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-scheduledaction.html#cfn-as-scheduledaction-desiredcapacity
	DesiredCapacity *Integer `json:"DesiredCapacity,omitempty"`

	// EndTime AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-scheduledaction.html#cfn-as-scheduledaction-endtime
	EndTime *String `json:"EndTime,omitempty"`

	// MaxSize AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-scheduledaction.html#cfn-as-scheduledaction-maxsize
	MaxSize *Integer `json:"MaxSize,omitempty"`

	// MinSize AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-scheduledaction.html#cfn-as-scheduledaction-minsize
	MinSize *Integer `json:"MinSize,omitempty"`

	// Recurrence AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-scheduledaction.html#cfn-as-scheduledaction-recurrence
	Recurrence *String `json:"Recurrence,omitempty"`

	// StartTime AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-scheduledaction.html#cfn-as-scheduledaction-starttime
	StartTime *String `json:"StartTime,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSAutoScalingScheduledAction) AddDependencies(v ...string) *AWSAutoScalingScheduledAction {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSAutoScalingScheduledAction) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSAutoScalingScheduledAction) AWSCloudFormationType() string {
	return "AWS::AutoScaling::ScheduledAction"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSAutoScalingScheduledAction) MarshalJSON() ([]byte, error) {
	type Properties AWSAutoScalingScheduledAction
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
func (r *AWSAutoScalingScheduledAction) UnmarshalJSON(b []byte) error {
	type Properties AWSAutoScalingScheduledAction
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
		*r = AWSAutoScalingScheduledAction(*res.Properties)
	}

	if res.DependsOn != nil {
		r.dependsOn = res.DependsOn
	}

	return nil
}

// GetAllAWSAutoScalingScheduledActionResources retrieves all AWSAutoScalingScheduledAction items from an AWS CloudFormation template
func (t *Template) GetAllAWSAutoScalingScheduledActionResources() map[string]AWSAutoScalingScheduledAction {
	results := map[string]AWSAutoScalingScheduledAction{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSAutoScalingScheduledAction:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::AutoScaling::ScheduledAction" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSAutoScalingScheduledAction
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

// GetAWSAutoScalingScheduledActionWithName retrieves all AWSAutoScalingScheduledAction items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSAutoScalingScheduledActionWithName(name string) (AWSAutoScalingScheduledAction, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSAutoScalingScheduledAction:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::AutoScaling::ScheduledAction" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSAutoScalingScheduledAction
						if err := json.Unmarshal(b, &result); err == nil {
							return result, nil
						}
					}
				}
			}
		}
	}
	return AWSAutoScalingScheduledAction{}, errors.New("resource not found")
}

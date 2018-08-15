package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSOpsWorksLayer AWS CloudFormation Resource (AWS::OpsWorks::Layer)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html
type AWSOpsWorksLayer struct {
	dependsOn []string

	// Attributes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html#cfn-opsworks-layer-attributes
	Attributes map[string]*String `json:"Attributes,omitempty"`

	// AutoAssignElasticIps AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html#cfn-opsworks-layer-autoassignelasticips
	AutoAssignElasticIps *Boolean `json:"AutoAssignElasticIps,omitempty"`

	// AutoAssignPublicIps AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html#cfn-opsworks-layer-autoassignpublicips
	AutoAssignPublicIps *Boolean `json:"AutoAssignPublicIps,omitempty"`

	// CustomInstanceProfileArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html#cfn-opsworks-layer-custominstanceprofilearn
	CustomInstanceProfileArn *String `json:"CustomInstanceProfileArn,omitempty"`

	// CustomJson AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html#cfn-opsworks-layer-customjson
	CustomJson interface{} `json:"CustomJson,omitempty"`

	// CustomRecipes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html#cfn-opsworks-layer-customrecipes
	CustomRecipes *AWSOpsWorksLayer_Recipes `json:"CustomRecipes,omitempty"`

	// CustomSecurityGroupIds AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html#cfn-opsworks-layer-customsecuritygroupids
	CustomSecurityGroupIds []*String `json:"CustomSecurityGroupIds,omitempty"`

	// EnableAutoHealing AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html#cfn-opsworks-layer-enableautohealing
	EnableAutoHealing *Boolean `json:"EnableAutoHealing,omitempty"`

	// InstallUpdatesOnBoot AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html#cfn-opsworks-layer-installupdatesonboot
	InstallUpdatesOnBoot *Boolean `json:"InstallUpdatesOnBoot,omitempty"`

	// LifecycleEventConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html#cfn-opsworks-layer-lifecycleeventconfiguration
	LifecycleEventConfiguration *AWSOpsWorksLayer_LifecycleEventConfiguration `json:"LifecycleEventConfiguration,omitempty"`

	// LoadBasedAutoScaling AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html#cfn-opsworks-layer-loadbasedautoscaling
	LoadBasedAutoScaling *AWSOpsWorksLayer_LoadBasedAutoScaling `json:"LoadBasedAutoScaling,omitempty"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html#cfn-opsworks-layer-name
	Name *String `json:"Name,omitempty"`

	// Packages AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html#cfn-opsworks-layer-packages
	Packages []*String `json:"Packages,omitempty"`

	// Shortname AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html#cfn-opsworks-layer-shortname
	Shortname *String `json:"Shortname,omitempty"`

	// StackId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html#cfn-opsworks-layer-stackid
	StackId *String `json:"StackId,omitempty"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html#cfn-opsworks-layer-tags
	Tags []Tag `json:"Tags,omitempty"`

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html#cfn-opsworks-layer-type
	Type *String `json:"Type,omitempty"`

	// UseEbsOptimizedInstances AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html#cfn-opsworks-layer-useebsoptimizedinstances
	UseEbsOptimizedInstances *Boolean `json:"UseEbsOptimizedInstances,omitempty"`

	// VolumeConfigurations AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html#cfn-opsworks-layer-volumeconfigurations
	VolumeConfigurations []AWSOpsWorksLayer_VolumeConfiguration `json:"VolumeConfigurations,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSOpsWorksLayer) AddDependencies(v ...string) *AWSOpsWorksLayer {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSOpsWorksLayer) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSOpsWorksLayer) AWSCloudFormationType() string {
	return "AWS::OpsWorks::Layer"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSOpsWorksLayer) MarshalJSON() ([]byte, error) {
	type Properties AWSOpsWorksLayer
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
func (r *AWSOpsWorksLayer) UnmarshalJSON(b []byte) error {
	type Properties AWSOpsWorksLayer
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
		*r = AWSOpsWorksLayer(*res.Properties)
	}

	if res.DependsOn != nil {
		r.dependsOn = res.DependsOn
	}

	return nil
}

// GetAllAWSOpsWorksLayerResources retrieves all AWSOpsWorksLayer items from an AWS CloudFormation template
func (t *Template) GetAllAWSOpsWorksLayerResources() map[string]AWSOpsWorksLayer {
	results := map[string]AWSOpsWorksLayer{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSOpsWorksLayer:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::OpsWorks::Layer" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSOpsWorksLayer
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

// GetAWSOpsWorksLayerWithName retrieves all AWSOpsWorksLayer items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSOpsWorksLayerWithName(name string) (AWSOpsWorksLayer, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSOpsWorksLayer:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::OpsWorks::Layer" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSOpsWorksLayer
						if err := json.Unmarshal(b, &result); err == nil {
							return result, nil
						}
					}
				}
			}
		}
	}
	return AWSOpsWorksLayer{}, errors.New("resource not found")
}

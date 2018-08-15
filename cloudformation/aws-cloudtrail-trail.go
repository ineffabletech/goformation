package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSCloudTrailTrail AWS CloudFormation Resource (AWS::CloudTrail::Trail)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html
type AWSCloudTrailTrail struct {
	dependsOn []string

	// CloudWatchLogsLogGroupArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-cloudwatchlogsloggrouparn
	CloudWatchLogsLogGroupArn *String `json:"CloudWatchLogsLogGroupArn,omitempty"`

	// CloudWatchLogsRoleArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-cloudwatchlogsrolearn
	CloudWatchLogsRoleArn *String `json:"CloudWatchLogsRoleArn,omitempty"`

	// EnableLogFileValidation AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-enablelogfilevalidation
	EnableLogFileValidation *Boolean `json:"EnableLogFileValidation,omitempty"`

	// EventSelectors AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-eventselectors
	EventSelectors []AWSCloudTrailTrail_EventSelector `json:"EventSelectors,omitempty"`

	// IncludeGlobalServiceEvents AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-includeglobalserviceevents
	IncludeGlobalServiceEvents *Boolean `json:"IncludeGlobalServiceEvents,omitempty"`

	// IsLogging AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-islogging
	IsLogging *Boolean `json:"IsLogging,omitempty"`

	// IsMultiRegionTrail AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-ismultiregiontrail
	IsMultiRegionTrail *Boolean `json:"IsMultiRegionTrail,omitempty"`

	// KMSKeyId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-kmskeyid
	KMSKeyId *String `json:"KMSKeyId,omitempty"`

	// S3BucketName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-s3bucketname
	S3BucketName *String `json:"S3BucketName,omitempty"`

	// S3KeyPrefix AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-s3keyprefix
	S3KeyPrefix *String `json:"S3KeyPrefix,omitempty"`

	// SnsTopicName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-snstopicname
	SnsTopicName *String `json:"SnsTopicName,omitempty"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-tags
	Tags []Tag `json:"Tags,omitempty"`

	// TrailName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-trailname
	TrailName *String `json:"TrailName,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSCloudTrailTrail) AddDependencies(v ...string) *AWSCloudTrailTrail {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSCloudTrailTrail) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCloudTrailTrail) AWSCloudFormationType() string {
	return "AWS::CloudTrail::Trail"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSCloudTrailTrail) MarshalJSON() ([]byte, error) {
	type Properties AWSCloudTrailTrail
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
func (r *AWSCloudTrailTrail) UnmarshalJSON(b []byte) error {
	type Properties AWSCloudTrailTrail
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
		*r = AWSCloudTrailTrail(*res.Properties)
	}

	if res.DependsOn != nil {
		r.dependsOn = res.DependsOn
	}

	return nil
}

// GetAllAWSCloudTrailTrailResources retrieves all AWSCloudTrailTrail items from an AWS CloudFormation template
func (t *Template) GetAllAWSCloudTrailTrailResources() map[string]AWSCloudTrailTrail {
	results := map[string]AWSCloudTrailTrail{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSCloudTrailTrail:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::CloudTrail::Trail" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSCloudTrailTrail
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

// GetAWSCloudTrailTrailWithName retrieves all AWSCloudTrailTrail items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCloudTrailTrailWithName(name string) (AWSCloudTrailTrail, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSCloudTrailTrail:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::CloudTrail::Trail" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSCloudTrailTrail
						if err := json.Unmarshal(b, &result); err == nil {
							return result, nil
						}
					}
				}
			}
		}
	}
	return AWSCloudTrailTrail{}, errors.New("resource not found")
}

package cloudformation

// AWSCloudTrailTrail_EventSelector AWS CloudFormation Resource (AWS::CloudTrail::Trail.EventSelector)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail-eventselector.html
type AWSCloudTrailTrail_EventSelector struct {
	dependsOn []string

	// DataResources AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail-eventselector.html#cfn-cloudtrail-trail-eventselector-dataresources
	DataResources []AWSCloudTrailTrail_DataResource `json:"DataResources,omitempty"`

	// IncludeManagementEvents AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail-eventselector.html#cfn-cloudtrail-trail-eventselector-includemanagementevents
	IncludeManagementEvents *Boolean `json:"IncludeManagementEvents,omitempty"`

	// ReadWriteType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail-eventselector.html#cfn-cloudtrail-trail-eventselector-readwritetype
	ReadWriteType *String `json:"ReadWriteType,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSCloudTrailTrail_EventSelector) AddDependencies(v ...string) *AWSCloudTrailTrail_EventSelector {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSCloudTrailTrail_EventSelector) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCloudTrailTrail_EventSelector) AWSCloudFormationType() string {
	return "AWS::CloudTrail::Trail.EventSelector"
}

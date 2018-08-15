package cloudformation

// AWSGlueConnection_PhysicalConnectionRequirements AWS CloudFormation Resource (AWS::Glue::Connection.PhysicalConnectionRequirements)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-physicalconnectionrequirements.html
type AWSGlueConnection_PhysicalConnectionRequirements struct {
	dependsOn []string

	// AvailabilityZone AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-physicalconnectionrequirements.html#cfn-glue-connection-physicalconnectionrequirements-availabilityzone
	AvailabilityZone *String `json:"AvailabilityZone,omitempty"`

	// SecurityGroupIdList AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-physicalconnectionrequirements.html#cfn-glue-connection-physicalconnectionrequirements-securitygroupidlist
	SecurityGroupIdList []*String `json:"SecurityGroupIdList,omitempty"`

	// SubnetId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-physicalconnectionrequirements.html#cfn-glue-connection-physicalconnectionrequirements-subnetid
	SubnetId *String `json:"SubnetId,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSGlueConnection_PhysicalConnectionRequirements) AddDependencies(v ...string) *AWSGlueConnection_PhysicalConnectionRequirements {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSGlueConnection_PhysicalConnectionRequirements) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSGlueConnection_PhysicalConnectionRequirements) AWSCloudFormationType() string {
	return "AWS::Glue::Connection.PhysicalConnectionRequirements"
}

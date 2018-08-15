package cloudformation

// AWSEMRInstanceFleetConfig_VolumeSpecification AWS CloudFormation Resource (AWS::EMR::InstanceFleetConfig.VolumeSpecification)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-volumespecification.html
type AWSEMRInstanceFleetConfig_VolumeSpecification struct {
	dependsOn []string

	// Iops AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-volumespecification.html#cfn-elasticmapreduce-instancefleetconfig-volumespecification-iops
	Iops *Integer `json:"Iops,omitempty"`

	// SizeInGB AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-volumespecification.html#cfn-elasticmapreduce-instancefleetconfig-volumespecification-sizeingb
	SizeInGB *Integer `json:"SizeInGB,omitempty"`

	// VolumeType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-volumespecification.html#cfn-elasticmapreduce-instancefleetconfig-volumespecification-volumetype
	VolumeType *String `json:"VolumeType,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSEMRInstanceFleetConfig_VolumeSpecification) AddDependencies(v ...string) *AWSEMRInstanceFleetConfig_VolumeSpecification {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSEMRInstanceFleetConfig_VolumeSpecification) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRInstanceFleetConfig_VolumeSpecification) AWSCloudFormationType() string {
	return "AWS::EMR::InstanceFleetConfig.VolumeSpecification"
}

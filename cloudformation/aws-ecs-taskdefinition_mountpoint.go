package cloudformation

// AWSECSTaskDefinition_MountPoint AWS CloudFormation Resource (AWS::ECS::TaskDefinition.MountPoint)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-mountpoints.html
type AWSECSTaskDefinition_MountPoint struct {
	dependsOn []string

	// ContainerPath AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-mountpoints.html#cfn-ecs-taskdefinition-containerdefinition-mountpoints-containerpath
	ContainerPath *String `json:"ContainerPath,omitempty"`

	// ReadOnly AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-mountpoints.html#cfn-ecs-taskdefinition-containerdefinition-mountpoints-readonly
	ReadOnly *Boolean `json:"ReadOnly,omitempty"`

	// SourceVolume AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-mountpoints.html#cfn-ecs-taskdefinition-containerdefinition-mountpoints-sourcevolume
	SourceVolume *String `json:"SourceVolume,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSECSTaskDefinition_MountPoint) AddDependencies(v ...string) *AWSECSTaskDefinition_MountPoint {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSECSTaskDefinition_MountPoint) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSECSTaskDefinition_MountPoint) AWSCloudFormationType() string {
	return "AWS::ECS::TaskDefinition.MountPoint"
}

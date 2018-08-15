package cloudformation

// AWSECSTaskDefinition_Device AWS CloudFormation Resource (AWS::ECS::TaskDefinition.Device)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-device.html
type AWSECSTaskDefinition_Device struct {
	dependsOn []string

	// ContainerPath AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-device.html#cfn-ecs-taskdefinition-device-containerpath
	ContainerPath *String `json:"ContainerPath,omitempty"`

	// HostPath AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-device.html#cfn-ecs-taskdefinition-device-hostpath
	HostPath *String `json:"HostPath,omitempty"`

	// Permissions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-device.html#cfn-ecs-taskdefinition-device-permissions
	Permissions []*String `json:"Permissions,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSECSTaskDefinition_Device) AddDependencies(v ...string) *AWSECSTaskDefinition_Device {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSECSTaskDefinition_Device) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSECSTaskDefinition_Device) AWSCloudFormationType() string {
	return "AWS::ECS::TaskDefinition.Device"
}

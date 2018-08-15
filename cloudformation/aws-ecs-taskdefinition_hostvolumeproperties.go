package cloudformation

// AWSECSTaskDefinition_HostVolumeProperties AWS CloudFormation Resource (AWS::ECS::TaskDefinition.HostVolumeProperties)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-volumes-host.html
type AWSECSTaskDefinition_HostVolumeProperties struct {
	dependsOn []string

	// SourcePath AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-volumes-host.html#cfn-ecs-taskdefinition-volumes-host-sourcepath
	SourcePath *String `json:"SourcePath,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSECSTaskDefinition_HostVolumeProperties) AddDependencies(v ...string) *AWSECSTaskDefinition_HostVolumeProperties {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSECSTaskDefinition_HostVolumeProperties) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSECSTaskDefinition_HostVolumeProperties) AWSCloudFormationType() string {
	return "AWS::ECS::TaskDefinition.HostVolumeProperties"
}

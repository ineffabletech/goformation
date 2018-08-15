package cloudformation

// AWSOpsWorksLayer_LifecycleEventConfiguration AWS CloudFormation Resource (AWS::OpsWorks::Layer.LifecycleEventConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-lifecycleeventconfiguration.html
type AWSOpsWorksLayer_LifecycleEventConfiguration struct {
	dependsOn []string

	// ShutdownEventConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-lifecycleeventconfiguration.html#cfn-opsworks-layer-lifecycleconfiguration-shutdowneventconfiguration
	ShutdownEventConfiguration *AWSOpsWorksLayer_ShutdownEventConfiguration `json:"ShutdownEventConfiguration,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSOpsWorksLayer_LifecycleEventConfiguration) AddDependencies(v ...string) *AWSOpsWorksLayer_LifecycleEventConfiguration {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSOpsWorksLayer_LifecycleEventConfiguration) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSOpsWorksLayer_LifecycleEventConfiguration) AWSCloudFormationType() string {
	return "AWS::OpsWorks::Layer.LifecycleEventConfiguration"
}

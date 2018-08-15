package cloudformation

// AWSOpsWorksLayer_ShutdownEventConfiguration AWS CloudFormation Resource (AWS::OpsWorks::Layer.ShutdownEventConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-lifecycleeventconfiguration-shutdowneventconfiguration.html
type AWSOpsWorksLayer_ShutdownEventConfiguration struct {
	dependsOn []string

	// DelayUntilElbConnectionsDrained AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-lifecycleeventconfiguration-shutdowneventconfiguration.html#cfn-opsworks-layer-lifecycleconfiguration-shutdowneventconfiguration-delayuntilelbconnectionsdrained
	DelayUntilElbConnectionsDrained *Boolean `json:"DelayUntilElbConnectionsDrained,omitempty"`

	// ExecutionTimeout AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-lifecycleeventconfiguration-shutdowneventconfiguration.html#cfn-opsworks-layer-lifecycleconfiguration-shutdowneventconfiguration-executiontimeout
	ExecutionTimeout *Integer `json:"ExecutionTimeout,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSOpsWorksLayer_ShutdownEventConfiguration) AddDependencies(v ...string) *AWSOpsWorksLayer_ShutdownEventConfiguration {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSOpsWorksLayer_ShutdownEventConfiguration) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSOpsWorksLayer_ShutdownEventConfiguration) AWSCloudFormationType() string {
	return "AWS::OpsWorks::Layer.ShutdownEventConfiguration"
}

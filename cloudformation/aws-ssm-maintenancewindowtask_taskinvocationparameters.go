package cloudformation

// AWSSSMMaintenanceWindowTask_TaskInvocationParameters AWS CloudFormation Resource (AWS::SSM::MaintenanceWindowTask.TaskInvocationParameters)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-taskinvocationparameters.html
type AWSSSMMaintenanceWindowTask_TaskInvocationParameters struct {
	dependsOn []string

	// MaintenanceWindowAutomationParameters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-taskinvocationparameters.html#cfn-ssm-maintenancewindowtask-taskinvocationparameters-maintenancewindowautomationparameters
	MaintenanceWindowAutomationParameters *AWSSSMMaintenanceWindowTask_MaintenanceWindowAutomationParameters `json:"MaintenanceWindowAutomationParameters,omitempty"`

	// MaintenanceWindowLambdaParameters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-taskinvocationparameters.html#cfn-ssm-maintenancewindowtask-taskinvocationparameters-maintenancewindowlambdaparameters
	MaintenanceWindowLambdaParameters *AWSSSMMaintenanceWindowTask_MaintenanceWindowLambdaParameters `json:"MaintenanceWindowLambdaParameters,omitempty"`

	// MaintenanceWindowRunCommandParameters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-taskinvocationparameters.html#cfn-ssm-maintenancewindowtask-taskinvocationparameters-maintenancewindowruncommandparameters
	MaintenanceWindowRunCommandParameters *AWSSSMMaintenanceWindowTask_MaintenanceWindowRunCommandParameters `json:"MaintenanceWindowRunCommandParameters,omitempty"`

	// MaintenanceWindowStepFunctionsParameters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-taskinvocationparameters.html#cfn-ssm-maintenancewindowtask-taskinvocationparameters-maintenancewindowstepfunctionsparameters
	MaintenanceWindowStepFunctionsParameters *AWSSSMMaintenanceWindowTask_MaintenanceWindowStepFunctionsParameters `json:"MaintenanceWindowStepFunctionsParameters,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSSSMMaintenanceWindowTask_TaskInvocationParameters) AddDependencies(v ...string) *AWSSSMMaintenanceWindowTask_TaskInvocationParameters {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSSSMMaintenanceWindowTask_TaskInvocationParameters) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSSMMaintenanceWindowTask_TaskInvocationParameters) AWSCloudFormationType() string {
	return "AWS::SSM::MaintenanceWindowTask.TaskInvocationParameters"
}

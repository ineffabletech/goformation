package cloudformation

// AWSSSMMaintenanceWindowTask_MaintenanceWindowAutomationParameters AWS CloudFormation Resource (AWS::SSM::MaintenanceWindowTask.MaintenanceWindowAutomationParameters)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-maintenancewindowautomationparameters.html
type AWSSSMMaintenanceWindowTask_MaintenanceWindowAutomationParameters struct {
	dependsOn []string

	// DocumentVersion AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-maintenancewindowautomationparameters.html#cfn-ssm-maintenancewindowtask-maintenancewindowautomationparameters-documentversion
	DocumentVersion *String `json:"DocumentVersion,omitempty"`

	// Parameters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-maintenancewindowautomationparameters.html#cfn-ssm-maintenancewindowtask-maintenancewindowautomationparameters-parameters
	Parameters interface{} `json:"Parameters,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSSSMMaintenanceWindowTask_MaintenanceWindowAutomationParameters) AddDependencies(v ...string) *AWSSSMMaintenanceWindowTask_MaintenanceWindowAutomationParameters {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSSSMMaintenanceWindowTask_MaintenanceWindowAutomationParameters) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSSMMaintenanceWindowTask_MaintenanceWindowAutomationParameters) AWSCloudFormationType() string {
	return "AWS::SSM::MaintenanceWindowTask.MaintenanceWindowAutomationParameters"
}

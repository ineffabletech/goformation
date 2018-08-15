package cloudformation

// AWSSSMMaintenanceWindowTask_LoggingInfo AWS CloudFormation Resource (AWS::SSM::MaintenanceWindowTask.LoggingInfo)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-logginginfo.html
type AWSSSMMaintenanceWindowTask_LoggingInfo struct {
	dependsOn []string

	// Region AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-logginginfo.html#cfn-ssm-maintenancewindowtask-logginginfo-region
	Region *String `json:"Region,omitempty"`

	// S3Bucket AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-logginginfo.html#cfn-ssm-maintenancewindowtask-logginginfo-s3bucket
	S3Bucket *String `json:"S3Bucket,omitempty"`

	// S3Prefix AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-logginginfo.html#cfn-ssm-maintenancewindowtask-logginginfo-s3prefix
	S3Prefix *String `json:"S3Prefix,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSSSMMaintenanceWindowTask_LoggingInfo) AddDependencies(v ...string) *AWSSSMMaintenanceWindowTask_LoggingInfo {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSSSMMaintenanceWindowTask_LoggingInfo) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSSMMaintenanceWindowTask_LoggingInfo) AWSCloudFormationType() string {
	return "AWS::SSM::MaintenanceWindowTask.LoggingInfo"
}

package cloudformation

// AWSBatchJobQueue_ComputeEnvironmentOrder AWS CloudFormation Resource (AWS::Batch::JobQueue.ComputeEnvironmentOrder)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobqueue-computeenvironmentorder.html
type AWSBatchJobQueue_ComputeEnvironmentOrder struct {
	dependsOn []string

	// ComputeEnvironment AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobqueue-computeenvironmentorder.html#cfn-batch-jobqueue-computeenvironmentorder-computeenvironment
	ComputeEnvironment *String `json:"ComputeEnvironment,omitempty"`

	// Order AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobqueue-computeenvironmentorder.html#cfn-batch-jobqueue-computeenvironmentorder-order
	Order *Integer `json:"Order,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSBatchJobQueue_ComputeEnvironmentOrder) AddDependencies(v ...string) *AWSBatchJobQueue_ComputeEnvironmentOrder {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSBatchJobQueue_ComputeEnvironmentOrder) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSBatchJobQueue_ComputeEnvironmentOrder) AWSCloudFormationType() string {
	return "AWS::Batch::JobQueue.ComputeEnvironmentOrder"
}

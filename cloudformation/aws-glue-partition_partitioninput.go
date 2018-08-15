package cloudformation

// AWSGluePartition_PartitionInput AWS CloudFormation Resource (AWS::Glue::Partition.PartitionInput)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-partitioninput.html
type AWSGluePartition_PartitionInput struct {
	dependsOn []string

	// Parameters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-partitioninput.html#cfn-glue-partition-partitioninput-parameters
	Parameters interface{} `json:"Parameters,omitempty"`

	// StorageDescriptor AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-partitioninput.html#cfn-glue-partition-partitioninput-storagedescriptor
	StorageDescriptor *AWSGluePartition_StorageDescriptor `json:"StorageDescriptor,omitempty"`

	// Values AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-partitioninput.html#cfn-glue-partition-partitioninput-values
	Values []*String `json:"Values,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSGluePartition_PartitionInput) AddDependencies(v ...string) *AWSGluePartition_PartitionInput {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSGluePartition_PartitionInput) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSGluePartition_PartitionInput) AWSCloudFormationType() string {
	return "AWS::Glue::Partition.PartitionInput"
}

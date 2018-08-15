package cloudformation

// AWSGluePartition_Order AWS CloudFormation Resource (AWS::Glue::Partition.Order)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-order.html
type AWSGluePartition_Order struct {
	dependsOn []string

	// Column AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-order.html#cfn-glue-partition-order-column
	Column *String `json:"Column,omitempty"`

	// SortOrder AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-order.html#cfn-glue-partition-order-sortorder
	SortOrder *Integer `json:"SortOrder,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSGluePartition_Order) AddDependencies(v ...string) *AWSGluePartition_Order {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSGluePartition_Order) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSGluePartition_Order) AWSCloudFormationType() string {
	return "AWS::Glue::Partition.Order"
}

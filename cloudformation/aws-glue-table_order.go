package cloudformation

// AWSGlueTable_Order AWS CloudFormation Resource (AWS::Glue::Table.Order)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-order.html
type AWSGlueTable_Order struct {
	dependsOn []string

	// Column AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-order.html#cfn-glue-table-order-column
	Column *String `json:"Column,omitempty"`

	// SortOrder AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-order.html#cfn-glue-table-order-sortorder
	SortOrder *Integer `json:"SortOrder,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSGlueTable_Order) AddDependencies(v ...string) *AWSGlueTable_Order {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSGlueTable_Order) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSGlueTable_Order) AWSCloudFormationType() string {
	return "AWS::Glue::Table.Order"
}

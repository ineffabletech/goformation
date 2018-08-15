package cloudformation

// AWSDynamoDBTable_GlobalSecondaryIndex AWS CloudFormation Resource (AWS::DynamoDB::Table.GlobalSecondaryIndex)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-gsi.html
type AWSDynamoDBTable_GlobalSecondaryIndex struct {
	dependsOn []string

	// IndexName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-gsi.html#cfn-dynamodb-gsi-indexname
	IndexName *String `json:"IndexName,omitempty"`

	// KeySchema AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-gsi.html#cfn-dynamodb-gsi-keyschema
	KeySchema []AWSDynamoDBTable_KeySchema `json:"KeySchema,omitempty"`

	// Projection AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-gsi.html#cfn-dynamodb-gsi-projection
	Projection *AWSDynamoDBTable_Projection `json:"Projection,omitempty"`

	// ProvisionedThroughput AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-gsi.html#cfn-dynamodb-gsi-provisionedthroughput
	ProvisionedThroughput *AWSDynamoDBTable_ProvisionedThroughput `json:"ProvisionedThroughput,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSDynamoDBTable_GlobalSecondaryIndex) AddDependencies(v ...string) *AWSDynamoDBTable_GlobalSecondaryIndex {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSDynamoDBTable_GlobalSecondaryIndex) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDynamoDBTable_GlobalSecondaryIndex) AWSCloudFormationType() string {
	return "AWS::DynamoDB::Table.GlobalSecondaryIndex"
}

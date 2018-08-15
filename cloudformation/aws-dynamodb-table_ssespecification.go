package cloudformation

// AWSDynamoDBTable_SSESpecification AWS CloudFormation Resource (AWS::DynamoDB::Table.SSESpecification)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-ssespecification.html
type AWSDynamoDBTable_SSESpecification struct {
	dependsOn []string

	// SSEEnabled AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-ssespecification.html#cfn-dynamodb-table-ssespecification-sseenabled
	SSEEnabled *Boolean `json:"SSEEnabled,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSDynamoDBTable_SSESpecification) AddDependencies(v ...string) *AWSDynamoDBTable_SSESpecification {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSDynamoDBTable_SSESpecification) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDynamoDBTable_SSESpecification) AWSCloudFormationType() string {
	return "AWS::DynamoDB::Table.SSESpecification"
}

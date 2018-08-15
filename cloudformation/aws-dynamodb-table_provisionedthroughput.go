package cloudformation

// AWSDynamoDBTable_ProvisionedThroughput AWS CloudFormation Resource (AWS::DynamoDB::Table.ProvisionedThroughput)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-provisionedthroughput.html
type AWSDynamoDBTable_ProvisionedThroughput struct {
	dependsOn []string

	// ReadCapacityUnits AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-provisionedthroughput.html#cfn-dynamodb-provisionedthroughput-readcapacityunits
	ReadCapacityUnits *Long `json:"ReadCapacityUnits,omitempty"`

	// WriteCapacityUnits AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-provisionedthroughput.html#cfn-dynamodb-provisionedthroughput-writecapacityunits
	WriteCapacityUnits *Long `json:"WriteCapacityUnits,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSDynamoDBTable_ProvisionedThroughput) AddDependencies(v ...string) *AWSDynamoDBTable_ProvisionedThroughput {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSDynamoDBTable_ProvisionedThroughput) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDynamoDBTable_ProvisionedThroughput) AWSCloudFormationType() string {
	return "AWS::DynamoDB::Table.ProvisionedThroughput"
}

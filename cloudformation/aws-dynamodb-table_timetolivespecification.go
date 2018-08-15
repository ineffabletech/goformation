package cloudformation

// AWSDynamoDBTable_TimeToLiveSpecification AWS CloudFormation Resource (AWS::DynamoDB::Table.TimeToLiveSpecification)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-timetolivespecification.html
type AWSDynamoDBTable_TimeToLiveSpecification struct {
	dependsOn []string

	// AttributeName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-timetolivespecification.html#cfn-dynamodb-timetolivespecification-attributename
	AttributeName *String `json:"AttributeName,omitempty"`

	// Enabled AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-timetolivespecification.html#cfn-dynamodb-timetolivespecification-enabled
	Enabled *Boolean `json:"Enabled,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSDynamoDBTable_TimeToLiveSpecification) AddDependencies(v ...string) *AWSDynamoDBTable_TimeToLiveSpecification {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSDynamoDBTable_TimeToLiveSpecification) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDynamoDBTable_TimeToLiveSpecification) AWSCloudFormationType() string {
	return "AWS::DynamoDB::Table.TimeToLiveSpecification"
}

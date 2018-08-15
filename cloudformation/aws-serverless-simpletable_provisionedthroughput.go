package cloudformation

// AWSServerlessSimpleTable_ProvisionedThroughput AWS CloudFormation Resource (AWS::Serverless::SimpleTable.ProvisionedThroughput)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-provisionedthroughput.html
type AWSServerlessSimpleTable_ProvisionedThroughput struct {
	dependsOn []string

	// ReadCapacityUnits AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-provisionedthroughput.html
	ReadCapacityUnits *Integer `json:"ReadCapacityUnits,omitempty"`

	// WriteCapacityUnits AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-provisionedthroughput.html
	WriteCapacityUnits *Integer `json:"WriteCapacityUnits,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSServerlessSimpleTable_ProvisionedThroughput) AddDependencies(v ...string) *AWSServerlessSimpleTable_ProvisionedThroughput {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSServerlessSimpleTable_ProvisionedThroughput) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSServerlessSimpleTable_ProvisionedThroughput) AWSCloudFormationType() string {
	return "AWS::Serverless::SimpleTable.ProvisionedThroughput"
}

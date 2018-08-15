package cloudformation

// AWSGlueJob_ConnectionsList AWS CloudFormation Resource (AWS::Glue::Job.ConnectionsList)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-job-connectionslist.html
type AWSGlueJob_ConnectionsList struct {
	dependsOn []string

	// Connections AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-job-connectionslist.html#cfn-glue-job-connectionslist-connections
	Connections []*String `json:"Connections,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSGlueJob_ConnectionsList) AddDependencies(v ...string) *AWSGlueJob_ConnectionsList {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSGlueJob_ConnectionsList) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSGlueJob_ConnectionsList) AWSCloudFormationType() string {
	return "AWS::Glue::Job.ConnectionsList"
}

package cloudformation

// AWSSSMAssociation_InstanceAssociationOutputLocation AWS CloudFormation Resource (AWS::SSM::Association.InstanceAssociationOutputLocation)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-instanceassociationoutputlocation.html
type AWSSSMAssociation_InstanceAssociationOutputLocation struct {
	dependsOn []string

	// S3Location AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-instanceassociationoutputlocation.html#cfn-ssm-association-instanceassociationoutputlocation-s3location
	S3Location *AWSSSMAssociation_S3OutputLocation `json:"S3Location,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSSSMAssociation_InstanceAssociationOutputLocation) AddDependencies(v ...string) *AWSSSMAssociation_InstanceAssociationOutputLocation {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSSSMAssociation_InstanceAssociationOutputLocation) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSSMAssociation_InstanceAssociationOutputLocation) AWSCloudFormationType() string {
	return "AWS::SSM::Association.InstanceAssociationOutputLocation"
}

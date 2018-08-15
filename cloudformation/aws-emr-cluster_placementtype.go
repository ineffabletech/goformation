package cloudformation

// AWSEMRCluster_PlacementType AWS CloudFormation Resource (AWS::EMR::Cluster.PlacementType)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-placementtype.html
type AWSEMRCluster_PlacementType struct {
	dependsOn []string

	// AvailabilityZone AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-placementtype.html#cfn-elasticmapreduce-cluster-placementtype-availabilityzone
	AvailabilityZone *String `json:"AvailabilityZone,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSEMRCluster_PlacementType) AddDependencies(v ...string) *AWSEMRCluster_PlacementType {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSEMRCluster_PlacementType) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRCluster_PlacementType) AWSCloudFormationType() string {
	return "AWS::EMR::Cluster.PlacementType"
}

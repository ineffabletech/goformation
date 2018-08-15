package cloudformation

// AWSElasticsearchDomain_SnapshotOptions AWS CloudFormation Resource (AWS::Elasticsearch::Domain.SnapshotOptions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticsearch-domain-snapshotoptions.html
type AWSElasticsearchDomain_SnapshotOptions struct {
	dependsOn []string

	// AutomatedSnapshotStartHour AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticsearch-domain-snapshotoptions.html#cfn-elasticsearch-domain-snapshotoptions-automatedsnapshotstarthour
	AutomatedSnapshotStartHour *Integer `json:"AutomatedSnapshotStartHour,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSElasticsearchDomain_SnapshotOptions) AddDependencies(v ...string) *AWSElasticsearchDomain_SnapshotOptions {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSElasticsearchDomain_SnapshotOptions) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticsearchDomain_SnapshotOptions) AWSCloudFormationType() string {
	return "AWS::Elasticsearch::Domain.SnapshotOptions"
}

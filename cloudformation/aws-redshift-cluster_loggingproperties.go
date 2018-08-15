package cloudformation

// AWSRedshiftCluster_LoggingProperties AWS CloudFormation Resource (AWS::Redshift::Cluster.LoggingProperties)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-cluster-loggingproperties.html
type AWSRedshiftCluster_LoggingProperties struct {
	dependsOn []string

	// BucketName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-cluster-loggingproperties.html#cfn-redshift-cluster-loggingproperties-bucketname
	BucketName *String `json:"BucketName,omitempty"`

	// S3KeyPrefix AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-cluster-loggingproperties.html#cfn-redshift-cluster-loggingproperties-s3keyprefix
	S3KeyPrefix *String `json:"S3KeyPrefix,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSRedshiftCluster_LoggingProperties) AddDependencies(v ...string) *AWSRedshiftCluster_LoggingProperties {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSRedshiftCluster_LoggingProperties) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSRedshiftCluster_LoggingProperties) AWSCloudFormationType() string {
	return "AWS::Redshift::Cluster.LoggingProperties"
}

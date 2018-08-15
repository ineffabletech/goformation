package cloudformation

// AWSCloudFrontDistribution_Restrictions AWS CloudFormation Resource (AWS::CloudFront::Distribution.Restrictions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-restrictions.html
type AWSCloudFrontDistribution_Restrictions struct {
	dependsOn []string

	// GeoRestriction AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-restrictions.html#cfn-cloudfront-distribution-restrictions-georestriction
	GeoRestriction *AWSCloudFrontDistribution_GeoRestriction `json:"GeoRestriction,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSCloudFrontDistribution_Restrictions) AddDependencies(v ...string) *AWSCloudFrontDistribution_Restrictions {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSCloudFrontDistribution_Restrictions) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCloudFrontDistribution_Restrictions) AWSCloudFormationType() string {
	return "AWS::CloudFront::Distribution.Restrictions"
}

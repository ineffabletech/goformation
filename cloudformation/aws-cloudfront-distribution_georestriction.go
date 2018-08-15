package cloudformation

// AWSCloudFrontDistribution_GeoRestriction AWS CloudFormation Resource (AWS::CloudFront::Distribution.GeoRestriction)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-georestriction.html
type AWSCloudFrontDistribution_GeoRestriction struct {
	dependsOn []string

	// Locations AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-georestriction.html#cfn-cloudfront-distribution-georestriction-locations
	Locations []*String `json:"Locations,omitempty"`

	// RestrictionType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-georestriction.html#cfn-cloudfront-distribution-georestriction-restrictiontype
	RestrictionType *String `json:"RestrictionType,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSCloudFrontDistribution_GeoRestriction) AddDependencies(v ...string) *AWSCloudFrontDistribution_GeoRestriction {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSCloudFrontDistribution_GeoRestriction) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCloudFrontDistribution_GeoRestriction) AWSCloudFormationType() string {
	return "AWS::CloudFront::Distribution.GeoRestriction"
}

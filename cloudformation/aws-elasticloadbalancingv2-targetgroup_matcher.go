package cloudformation

// AWSElasticLoadBalancingV2TargetGroup_Matcher AWS CloudFormation Resource (AWS::ElasticLoadBalancingV2::TargetGroup.Matcher)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-targetgroup-matcher.html
type AWSElasticLoadBalancingV2TargetGroup_Matcher struct {
	dependsOn []string

	// HttpCode AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-targetgroup-matcher.html#cfn-elasticloadbalancingv2-targetgroup-matcher-httpcode
	HttpCode *String `json:"HttpCode,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSElasticLoadBalancingV2TargetGroup_Matcher) AddDependencies(v ...string) *AWSElasticLoadBalancingV2TargetGroup_Matcher {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSElasticLoadBalancingV2TargetGroup_Matcher) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticLoadBalancingV2TargetGroup_Matcher) AWSCloudFormationType() string {
	return "AWS::ElasticLoadBalancingV2::TargetGroup.Matcher"
}

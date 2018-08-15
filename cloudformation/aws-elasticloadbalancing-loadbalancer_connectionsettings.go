package cloudformation

// AWSElasticLoadBalancingLoadBalancer_ConnectionSettings AWS CloudFormation Resource (AWS::ElasticLoadBalancing::LoadBalancer.ConnectionSettings)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-connectionsettings.html
type AWSElasticLoadBalancingLoadBalancer_ConnectionSettings struct {
	dependsOn []string

	// IdleTimeout AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-connectionsettings.html#cfn-elb-connectionsettings-idletimeout
	IdleTimeout *Integer `json:"IdleTimeout,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSElasticLoadBalancingLoadBalancer_ConnectionSettings) AddDependencies(v ...string) *AWSElasticLoadBalancingLoadBalancer_ConnectionSettings {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSElasticLoadBalancingLoadBalancer_ConnectionSettings) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticLoadBalancingLoadBalancer_ConnectionSettings) AWSCloudFormationType() string {
	return "AWS::ElasticLoadBalancing::LoadBalancer.ConnectionSettings"
}

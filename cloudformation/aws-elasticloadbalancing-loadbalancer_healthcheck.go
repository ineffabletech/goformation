package cloudformation

// AWSElasticLoadBalancingLoadBalancer_HealthCheck AWS CloudFormation Resource (AWS::ElasticLoadBalancing::LoadBalancer.HealthCheck)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-health-check.html
type AWSElasticLoadBalancingLoadBalancer_HealthCheck struct {
	dependsOn []string

	// HealthyThreshold AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-health-check.html#cfn-elb-healthcheck-healthythreshold
	HealthyThreshold *String `json:"HealthyThreshold,omitempty"`

	// Interval AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-health-check.html#cfn-elb-healthcheck-interval
	Interval *String `json:"Interval,omitempty"`

	// Target AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-health-check.html#cfn-elb-healthcheck-target
	Target *String `json:"Target,omitempty"`

	// Timeout AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-health-check.html#cfn-elb-healthcheck-timeout
	Timeout *String `json:"Timeout,omitempty"`

	// UnhealthyThreshold AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-health-check.html#cfn-elb-healthcheck-unhealthythreshold
	UnhealthyThreshold *String `json:"UnhealthyThreshold,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSElasticLoadBalancingLoadBalancer_HealthCheck) AddDependencies(v ...string) *AWSElasticLoadBalancingLoadBalancer_HealthCheck {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSElasticLoadBalancingLoadBalancer_HealthCheck) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticLoadBalancingLoadBalancer_HealthCheck) AWSCloudFormationType() string {
	return "AWS::ElasticLoadBalancing::LoadBalancer.HealthCheck"
}

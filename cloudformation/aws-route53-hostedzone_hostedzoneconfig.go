package cloudformation

// AWSRoute53HostedZone_HostedZoneConfig AWS CloudFormation Resource (AWS::Route53::HostedZone.HostedZoneConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-hostedzone-hostedzoneconfig.html
type AWSRoute53HostedZone_HostedZoneConfig struct {
	dependsOn []string

	// Comment AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-hostedzone-hostedzoneconfig.html#cfn-route53-hostedzone-hostedzoneconfig-comment
	Comment *String `json:"Comment,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSRoute53HostedZone_HostedZoneConfig) AddDependencies(v ...string) *AWSRoute53HostedZone_HostedZoneConfig {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSRoute53HostedZone_HostedZoneConfig) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSRoute53HostedZone_HostedZoneConfig) AWSCloudFormationType() string {
	return "AWS::Route53::HostedZone.HostedZoneConfig"
}

package cloudformation

// AWSElasticBeanstalkApplication_ApplicationVersionLifecycleConfig AWS CloudFormation Resource (AWS::ElasticBeanstalk::Application.ApplicationVersionLifecycleConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-application-applicationversionlifecycleconfig.html
type AWSElasticBeanstalkApplication_ApplicationVersionLifecycleConfig struct {
	dependsOn []string

	// MaxAgeRule AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-application-applicationversionlifecycleconfig.html#cfn-elasticbeanstalk-application-applicationversionlifecycleconfig-maxagerule
	MaxAgeRule *AWSElasticBeanstalkApplication_MaxAgeRule `json:"MaxAgeRule,omitempty"`

	// MaxCountRule AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-application-applicationversionlifecycleconfig.html#cfn-elasticbeanstalk-application-applicationversionlifecycleconfig-maxcountrule
	MaxCountRule *AWSElasticBeanstalkApplication_MaxCountRule `json:"MaxCountRule,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSElasticBeanstalkApplication_ApplicationVersionLifecycleConfig) AddDependencies(v ...string) *AWSElasticBeanstalkApplication_ApplicationVersionLifecycleConfig {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSElasticBeanstalkApplication_ApplicationVersionLifecycleConfig) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticBeanstalkApplication_ApplicationVersionLifecycleConfig) AWSCloudFormationType() string {
	return "AWS::ElasticBeanstalk::Application.ApplicationVersionLifecycleConfig"
}

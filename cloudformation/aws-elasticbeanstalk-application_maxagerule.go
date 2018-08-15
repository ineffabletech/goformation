package cloudformation

// AWSElasticBeanstalkApplication_MaxAgeRule AWS CloudFormation Resource (AWS::ElasticBeanstalk::Application.MaxAgeRule)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-application-maxagerule.html
type AWSElasticBeanstalkApplication_MaxAgeRule struct {
	dependsOn []string

	// DeleteSourceFromS3 AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-application-maxagerule.html#cfn-elasticbeanstalk-application-maxagerule-deletesourcefroms3
	DeleteSourceFromS3 *Boolean `json:"DeleteSourceFromS3,omitempty"`

	// Enabled AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-application-maxagerule.html#cfn-elasticbeanstalk-application-maxagerule-enabled
	Enabled *Boolean `json:"Enabled,omitempty"`

	// MaxAgeInDays AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-application-maxagerule.html#cfn-elasticbeanstalk-application-maxagerule-maxageindays
	MaxAgeInDays *Integer `json:"MaxAgeInDays,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSElasticBeanstalkApplication_MaxAgeRule) AddDependencies(v ...string) *AWSElasticBeanstalkApplication_MaxAgeRule {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSElasticBeanstalkApplication_MaxAgeRule) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticBeanstalkApplication_MaxAgeRule) AWSCloudFormationType() string {
	return "AWS::ElasticBeanstalk::Application.MaxAgeRule"
}

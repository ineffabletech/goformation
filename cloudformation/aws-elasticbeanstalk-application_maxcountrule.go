package cloudformation

// AWSElasticBeanstalkApplication_MaxCountRule AWS CloudFormation Resource (AWS::ElasticBeanstalk::Application.MaxCountRule)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-application-maxcountrule.html
type AWSElasticBeanstalkApplication_MaxCountRule struct {
	dependsOn []string

	// DeleteSourceFromS3 AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-application-maxcountrule.html#cfn-elasticbeanstalk-application-maxcountrule-deletesourcefroms3
	DeleteSourceFromS3 *Boolean `json:"DeleteSourceFromS3,omitempty"`

	// Enabled AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-application-maxcountrule.html#cfn-elasticbeanstalk-application-maxcountrule-enabled
	Enabled *Boolean `json:"Enabled,omitempty"`

	// MaxCount AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-application-maxcountrule.html#cfn-elasticbeanstalk-application-maxcountrule-maxcount
	MaxCount *Integer `json:"MaxCount,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSElasticBeanstalkApplication_MaxCountRule) AddDependencies(v ...string) *AWSElasticBeanstalkApplication_MaxCountRule {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSElasticBeanstalkApplication_MaxCountRule) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticBeanstalkApplication_MaxCountRule) AWSCloudFormationType() string {
	return "AWS::ElasticBeanstalk::Application.MaxCountRule"
}

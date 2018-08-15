package cloudformation

// AWSElasticBeanstalkEnvironment_OptionSetting AWS CloudFormation Resource (AWS::ElasticBeanstalk::Environment.OptionSetting)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-option-settings.html
type AWSElasticBeanstalkEnvironment_OptionSetting struct {
	dependsOn []string

	// Namespace AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-option-settings.html#cfn-beanstalk-optionsettings-namespace
	Namespace *String `json:"Namespace,omitempty"`

	// OptionName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-option-settings.html#cfn-beanstalk-optionsettings-optionname
	OptionName *String `json:"OptionName,omitempty"`

	// ResourceName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-option-settings.html#cfn-elasticbeanstalk-environment-optionsetting-resourcename
	ResourceName *String `json:"ResourceName,omitempty"`

	// Value AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-option-settings.html#cfn-beanstalk-optionsettings-value
	Value *String `json:"Value,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSElasticBeanstalkEnvironment_OptionSetting) AddDependencies(v ...string) *AWSElasticBeanstalkEnvironment_OptionSetting {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSElasticBeanstalkEnvironment_OptionSetting) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticBeanstalkEnvironment_OptionSetting) AWSCloudFormationType() string {
	return "AWS::ElasticBeanstalk::Environment.OptionSetting"
}

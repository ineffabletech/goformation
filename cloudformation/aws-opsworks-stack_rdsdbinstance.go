package cloudformation

// AWSOpsWorksStack_RdsDbInstance AWS CloudFormation Resource (AWS::OpsWorks::Stack.RdsDbInstance)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-rdsdbinstance.html
type AWSOpsWorksStack_RdsDbInstance struct {
	dependsOn []string

	// DbPassword AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-rdsdbinstance.html#cfn-opsworks-stack-rdsdbinstance-dbpassword
	DbPassword *String `json:"DbPassword,omitempty"`

	// DbUser AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-rdsdbinstance.html#cfn-opsworks-stack-rdsdbinstance-dbuser
	DbUser *String `json:"DbUser,omitempty"`

	// RdsDbInstanceArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-rdsdbinstance.html#cfn-opsworks-stack-rdsdbinstance-rdsdbinstancearn
	RdsDbInstanceArn *String `json:"RdsDbInstanceArn,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSOpsWorksStack_RdsDbInstance) AddDependencies(v ...string) *AWSOpsWorksStack_RdsDbInstance {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSOpsWorksStack_RdsDbInstance) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSOpsWorksStack_RdsDbInstance) AWSCloudFormationType() string {
	return "AWS::OpsWorks::Stack.RdsDbInstance"
}

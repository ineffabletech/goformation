package cloudformation

// AWSLambdaAlias_AliasRoutingConfiguration AWS CloudFormation Resource (AWS::Lambda::Alias.AliasRoutingConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-alias-aliasroutingconfiguration.html
type AWSLambdaAlias_AliasRoutingConfiguration struct {
	dependsOn []string

	// AdditionalVersionWeights AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-alias-aliasroutingconfiguration.html#cfn-lambda-alias-aliasroutingconfiguration-additionalversionweights
	AdditionalVersionWeights []AWSLambdaAlias_VersionWeight `json:"AdditionalVersionWeights,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSLambdaAlias_AliasRoutingConfiguration) AddDependencies(v ...string) *AWSLambdaAlias_AliasRoutingConfiguration {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSLambdaAlias_AliasRoutingConfiguration) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSLambdaAlias_AliasRoutingConfiguration) AWSCloudFormationType() string {
	return "AWS::Lambda::Alias.AliasRoutingConfiguration"
}

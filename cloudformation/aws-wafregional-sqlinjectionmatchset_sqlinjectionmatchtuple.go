package cloudformation

// AWSWAFRegionalSqlInjectionMatchSet_SqlInjectionMatchTuple AWS CloudFormation Resource (AWS::WAFRegional::SqlInjectionMatchSet.SqlInjectionMatchTuple)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-sqlinjectionmatchset-sqlinjectionmatchtuple.html
type AWSWAFRegionalSqlInjectionMatchSet_SqlInjectionMatchTuple struct {
	dependsOn []string

	// FieldToMatch AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-sqlinjectionmatchset-sqlinjectionmatchtuple.html#cfn-wafregional-sqlinjectionmatchset-sqlinjectionmatchtuple-fieldtomatch
	FieldToMatch *AWSWAFRegionalSqlInjectionMatchSet_FieldToMatch `json:"FieldToMatch,omitempty"`

	// TextTransformation AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-sqlinjectionmatchset-sqlinjectionmatchtuple.html#cfn-wafregional-sqlinjectionmatchset-sqlinjectionmatchtuple-texttransformation
	TextTransformation *String `json:"TextTransformation,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSWAFRegionalSqlInjectionMatchSet_SqlInjectionMatchTuple) AddDependencies(v ...string) *AWSWAFRegionalSqlInjectionMatchSet_SqlInjectionMatchTuple {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSWAFRegionalSqlInjectionMatchSet_SqlInjectionMatchTuple) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFRegionalSqlInjectionMatchSet_SqlInjectionMatchTuple) AWSCloudFormationType() string {
	return "AWS::WAFRegional::SqlInjectionMatchSet.SqlInjectionMatchTuple"
}

package cloudformation

// AWSWAFSqlInjectionMatchSet_SqlInjectionMatchTuple AWS CloudFormation Resource (AWS::WAF::SqlInjectionMatchSet.SqlInjectionMatchTuple)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-sqlinjectionmatchset-sqlinjectionmatchtuples.html
type AWSWAFSqlInjectionMatchSet_SqlInjectionMatchTuple struct {
	dependsOn []string

	// FieldToMatch AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-sqlinjectionmatchset-sqlinjectionmatchtuples.html#cfn-waf-sqlinjectionmatchset-sqlinjectionmatchtuples-fieldtomatch
	FieldToMatch *AWSWAFSqlInjectionMatchSet_FieldToMatch `json:"FieldToMatch,omitempty"`

	// TextTransformation AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-sqlinjectionmatchset-sqlinjectionmatchtuples.html#cfn-waf-sqlinjectionmatchset-sqlinjectionmatchtuples-texttransformation
	TextTransformation *String `json:"TextTransformation,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSWAFSqlInjectionMatchSet_SqlInjectionMatchTuple) AddDependencies(v ...string) *AWSWAFSqlInjectionMatchSet_SqlInjectionMatchTuple {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSWAFSqlInjectionMatchSet_SqlInjectionMatchTuple) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFSqlInjectionMatchSet_SqlInjectionMatchTuple) AWSCloudFormationType() string {
	return "AWS::WAF::SqlInjectionMatchSet.SqlInjectionMatchTuple"
}

package cloudformation

// AWSAppSyncGraphQLApi_OpenIDConnectConfig AWS CloudFormation Resource (AWS::AppSync::GraphQLApi.OpenIDConnectConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-graphqlapi-openidconnectconfig.html
type AWSAppSyncGraphQLApi_OpenIDConnectConfig struct {
	dependsOn []string

	// AuthTTL AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-graphqlapi-openidconnectconfig.html#cfn-appsync-graphqlapi-openidconnectconfig-authttl
	AuthTTL *Double `json:"AuthTTL,omitempty"`

	// ClientId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-graphqlapi-openidconnectconfig.html#cfn-appsync-graphqlapi-openidconnectconfig-clientid
	ClientId *String `json:"ClientId,omitempty"`

	// IatTTL AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-graphqlapi-openidconnectconfig.html#cfn-appsync-graphqlapi-openidconnectconfig-iatttl
	IatTTL *Double `json:"IatTTL,omitempty"`

	// Issuer AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-graphqlapi-openidconnectconfig.html#cfn-appsync-graphqlapi-openidconnectconfig-issuer
	Issuer *String `json:"Issuer,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSAppSyncGraphQLApi_OpenIDConnectConfig) AddDependencies(v ...string) *AWSAppSyncGraphQLApi_OpenIDConnectConfig {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSAppSyncGraphQLApi_OpenIDConnectConfig) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSAppSyncGraphQLApi_OpenIDConnectConfig) AWSCloudFormationType() string {
	return "AWS::AppSync::GraphQLApi.OpenIDConnectConfig"
}

package cloudformation

// AWSServiceCatalogCloudFormationProduct_ProvisioningArtifactProperties AWS CloudFormation Resource (AWS::ServiceCatalog::CloudFormationProduct.ProvisioningArtifactProperties)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicecatalog-cloudformationproduct-provisioningartifactproperties.html
type AWSServiceCatalogCloudFormationProduct_ProvisioningArtifactProperties struct {
	dependsOn []string

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicecatalog-cloudformationproduct-provisioningartifactproperties.html#cfn-servicecatalog-cloudformationproduct-provisioningartifactproperties-description
	Description *String `json:"Description,omitempty"`

	// Info AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicecatalog-cloudformationproduct-provisioningartifactproperties.html#cfn-servicecatalog-cloudformationproduct-provisioningartifactproperties-info
	Info interface{} `json:"Info,omitempty"`

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicecatalog-cloudformationproduct-provisioningartifactproperties.html#cfn-servicecatalog-cloudformationproduct-provisioningartifactproperties-name
	Name *String `json:"Name,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSServiceCatalogCloudFormationProduct_ProvisioningArtifactProperties) AddDependencies(v ...string) *AWSServiceCatalogCloudFormationProduct_ProvisioningArtifactProperties {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSServiceCatalogCloudFormationProduct_ProvisioningArtifactProperties) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSServiceCatalogCloudFormationProduct_ProvisioningArtifactProperties) AWSCloudFormationType() string {
	return "AWS::ServiceCatalog::CloudFormationProduct.ProvisioningArtifactProperties"
}

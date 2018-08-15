package cloudformation

// AWSDirectoryServiceSimpleAD_VpcSettings AWS CloudFormation Resource (AWS::DirectoryService::SimpleAD.VpcSettings)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-directoryservice-simplead-vpcsettings.html
type AWSDirectoryServiceSimpleAD_VpcSettings struct {
	dependsOn []string

	// SubnetIds AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-directoryservice-simplead-vpcsettings.html#cfn-directoryservice-simplead-vpcsettings-subnetids
	SubnetIds []*String `json:"SubnetIds,omitempty"`

	// VpcId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-directoryservice-simplead-vpcsettings.html#cfn-directoryservice-simplead-vpcsettings-vpcid
	VpcId *String `json:"VpcId,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSDirectoryServiceSimpleAD_VpcSettings) AddDependencies(v ...string) *AWSDirectoryServiceSimpleAD_VpcSettings {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSDirectoryServiceSimpleAD_VpcSettings) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDirectoryServiceSimpleAD_VpcSettings) AWSCloudFormationType() string {
	return "AWS::DirectoryService::SimpleAD.VpcSettings"
}

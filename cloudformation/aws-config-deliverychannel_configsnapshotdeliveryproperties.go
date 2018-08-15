package cloudformation

// AWSConfigDeliveryChannel_ConfigSnapshotDeliveryProperties AWS CloudFormation Resource (AWS::Config::DeliveryChannel.ConfigSnapshotDeliveryProperties)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-deliverychannel-configsnapshotdeliveryproperties.html
type AWSConfigDeliveryChannel_ConfigSnapshotDeliveryProperties struct {
	dependsOn []string

	// DeliveryFrequency AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-deliverychannel-configsnapshotdeliveryproperties.html#cfn-config-deliverychannel-configsnapshotdeliveryproperties-deliveryfrequency
	DeliveryFrequency *String `json:"DeliveryFrequency,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSConfigDeliveryChannel_ConfigSnapshotDeliveryProperties) AddDependencies(v ...string) *AWSConfigDeliveryChannel_ConfigSnapshotDeliveryProperties {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSConfigDeliveryChannel_ConfigSnapshotDeliveryProperties) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSConfigDeliveryChannel_ConfigSnapshotDeliveryProperties) AWSCloudFormationType() string {
	return "AWS::Config::DeliveryChannel.ConfigSnapshotDeliveryProperties"
}

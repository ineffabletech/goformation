package cloudformation

// AWSSESTemplate_Template AWS CloudFormation Resource (AWS::SES::Template.Template)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-template-template.html
type AWSSESTemplate_Template struct {
	dependsOn []string

	// HtmlPart AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-template-template.html#cfn-ses-template-template-htmlpart
	HtmlPart *String `json:"HtmlPart,omitempty"`

	// SubjectPart AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-template-template.html#cfn-ses-template-template-subjectpart
	SubjectPart *String `json:"SubjectPart,omitempty"`

	// TemplateName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-template-template.html#cfn-ses-template-template-templatename
	TemplateName *String `json:"TemplateName,omitempty"`

	// TextPart AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-template-template.html#cfn-ses-template-template-textpart
	TextPart *String `json:"TextPart,omitempty"`
}

// AddDependencies allows adding dependencies to the resource.
func (r *AWSSESTemplate_Template) AddDependencies(v ...string) *AWSSESTemplate_Template {
	if r.dependsOn == nil {
		r.dependsOn = []string{}
	}
	r.dependsOn = append(r.dependsOn, v...)
	return r
}

// DependsOn returns the .
func (r *AWSSESTemplate_Template) DependsOn(v ...string) []string {
	if r.dependsOn == nil {
		return []string{}
	} else {
		return r.dependsOn
	}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSESTemplate_Template) AWSCloudFormationType() string {
	return "AWS::SES::Template.Template"
}

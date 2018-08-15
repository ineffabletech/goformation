package cloudformation

type FnGetAtt struct {
	FnGetAtt []*String `json:"Fn::GetAtt"`
}

func NewFnGetAtt(nor, atn interface{}) *FnGetAtt {
	return &FnGetAtt{FnGetAtt: []*String{NewString(nor), NewString(atn)}}
}

func NewStringFnGetAtt(nor, atn interface{}) *String {
	return NewString(NewFnGetAtt(nor, atn))
}

func NewIntegerFnGetAtt(nor, atn interface{}) *Integer {
	return NewInteger(NewFnGetAtt(nor, atn))
}

func NewLongFnGetAtt(nor, atn interface{}) *Long {
	return NewLong(NewFnGetAtt(nor, atn))
}

func NewDoubleFnGetAtt(nor, atn interface{}) *Double {
	return NewDouble(NewFnGetAtt(nor, atn))
}

func NewBooleanFnGetAtt(nor, atn interface{}) *Boolean {
	return NewBoolean(NewFnGetAtt(nor, atn))
}

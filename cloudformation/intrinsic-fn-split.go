package cloudformation

type FnSplit struct {
	FnSplit []*String `json:"Fn::Split"`
}

func NewFnSplit(d, s interface{}) *FnSplit {
	return &FnSplit{FnSplit: []*String{NewString(d), NewString(s)}}
}

func NewStringFnSplit(d, s interface{}) *String {
	return NewString(NewFnSplit(d, s))
}

func NewIntegerFnSplit(d, s interface{}) *Integer {
	return NewInteger(NewFnSplit(d, s))
}

func NewLongFnSplit(d, s interface{}) *Long {
	return NewLong(NewFnSplit(d, s))
}

func NewDoubleFnSplit(d, s interface{}) *Double {
	return NewDouble(NewFnSplit(d, s))
}

func NewBooleanFnSplit(d, s interface{}) *Boolean {
	return NewBoolean(NewFnSplit(d, s))
}

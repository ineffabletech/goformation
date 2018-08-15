package cloudformation

type FnBase64 struct {
	FnBase64 *String `json:"Fn::Base64"`
}

func NewFnBase64(v interface{}) *FnBase64 {
	return &FnBase64{FnBase64: NewString(v)}
}

func NewStringFnBase64(v interface{}) *String {
	return NewString(NewFnBase64(v))
}

func NewIntegerFnBase64(v interface{}) *Integer {
	return NewInteger(NewFnBase64(v))
}

func NewLongFnBase64(v interface{}) *Long {
	return NewLong(NewFnBase64(v))
}

func NewDoubleFnBase64(v interface{}) *Double {
	return NewDouble(NewFnBase64(v))
}

func NewBooleanFnBase64(v interface{}) *Boolean {
	return NewBoolean(NewFnBase64(v))
}

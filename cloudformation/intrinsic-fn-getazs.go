package cloudformation

type FnGetAZs struct {
	FnGetAZs []*String `json:"Fn::GetAZs"`
}

func NewFnGetAZs(r interface{}) *FnGetAZs {
	return &FnGetAZs{FnGetAZs: []*String{NewString(r)}}
}

func NewStringFnGetAZs(v interface{}) *String {
	return NewString(NewFnGetAZs(v))
}

func NewIntegerFnGetAZs(v interface{}) *Integer {
	return NewInteger(NewFnGetAZs(v))
}

func NewLongFnGetAZs(v interface{}) *Long {
	return NewLong(NewFnGetAZs(v))
}

func NewDoubleFnGetAZs(v interface{}) *Double {
	return NewDouble(NewFnGetAZs(v))
}

func NewBooleanFnGetAZs(v interface{}) *Boolean {
	return NewBoolean(NewFnGetAZs(v))
}

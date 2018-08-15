package cloudformation

type FnImportValue struct {
	FnImportValue []*String `json:"Fn::ImportValue"`
}

func NewFnImportValue(sv interface{}) *FnImportValue {
	return &FnImportValue{FnImportValue: []*String{NewString(sv)}}
}

func NewStringFnImportValue(v interface{}) *String {
	return NewString(NewFnImportValue(v))
}

func NewIntegerFnImportValue(v interface{}) *Integer {
	return NewInteger(NewFnImportValue(v))
}

func NewLongFnImportValue(v interface{}) *Long {
	return NewLong(NewFnImportValue(v))
}

func NewDoubleFnImportValue(v interface{}) *Double {
	return NewDouble(NewFnImportValue(v))
}

func NewBooleanFnImportValue(v interface{}) *Boolean {
	return NewBoolean(NewFnImportValue(v))
}

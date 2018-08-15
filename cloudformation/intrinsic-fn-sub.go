package cloudformation

type FnSub struct {
	FnSub []interface{} `json:"Fn::Sub"`
}

func NewFnSub(s interface{}, m map[string]interface{}) *FnSub {
	if m != nil {
		return &FnSub{FnSub: []interface{}{NewString(s), m}}
	} else {
		return &FnSub{FnSub: []interface{}{NewString(s)}}
	}
}

func NewStringFnSub(s interface{}, m map[string]interface{}) *String {
	return NewString(NewFnSub(s, m))
}

func NewIntegerFnSub(s interface{}, m map[string]interface{}) *Integer {
	return NewInteger(NewFnSub(s, m))
}

func NewLongFnSub(s interface{}, m map[string]interface{}) *Long {
	return NewLong(NewFnSub(s, m))
}

func NewDoubleFnSub(s interface{}, m map[string]interface{}) *Double {
	return NewDouble(NewFnSub(s, m))
}

func NewBooleanFnSub(s interface{}, m map[string]interface{}) *Boolean {
	return NewBoolean(NewFnSub(s, m))
}

package cloudformation

type FnSelect struct {
	FnSelect []interface{} `json:"Fn::Select"`
}

func NewFnSelect(i interface{}, vs ...interface{}) *FnSelect {
	res := []interface{}{}
	for _, v := range vs {
		res = append(res, NewString(v))
	}
	return &FnSelect{FnSelect: []interface{}{NewInteger(i), res}}
}

func NewStringFnSelect(i interface{}, vs ...interface{}) *String {
	return NewString(NewFnSelect(i, vs...))
}

func NewIntegerFnSelect(i interface{}, vs ...interface{}) *Integer {
	return NewInteger(NewFnSelect(i, vs...))
}

func NewLongFnSelect(i interface{}, vs ...interface{}) *Long {
	return NewLong(NewFnSelect(i, vs...))
}

func NewDoubleFnSelect(i interface{}, vs ...interface{}) *Double {
	return NewDouble(NewFnSelect(i, vs...))
}

func NewBooleanFnSelect(i interface{}, vs ...interface{}) *Boolean {
	return NewBoolean(NewFnSelect(i, vs...))
}

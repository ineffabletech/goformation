package cloudformation

type FnJoin struct {
	FnJoin []interface{} `json:"Fn::Join"`
}

func NewFnJoin(d interface{}, vs ...interface{}) *FnJoin {
	res := []interface{}{}
	for _, v := range vs {
		res = append(res, NewString(v))
	}
	return &FnJoin{FnJoin: []interface{}{NewString(d), res}}
}

func NewStringFnJoin(d interface{}, vs ...interface{}) *String {
	return NewString(NewFnJoin(d, vs...))
}

func NewIntegerFnJoin(d interface{}, vs ...interface{}) *Integer {
	return NewInteger(NewFnJoin(d, vs...))
}

func NewLongFnJoin(d interface{}, vs ...interface{}) *Long {
	return NewLong(NewFnJoin(d, vs...))
}

func NewDoubleFnJoin(d interface{}, vs ...interface{}) *Double {
	return NewDouble(NewFnJoin(d, vs...))
}

func NewBooleanFnJoin(d interface{}, vs ...interface{}) *Boolean {
	return NewBoolean(NewFnJoin(d, vs...))
}

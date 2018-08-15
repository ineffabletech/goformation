package cloudformation

type Ref struct {
	Ref *String `json:"Ref"`
}

func NewRef(v interface{}) *Ref {
	return &Ref{Ref: NewString(v)}
}

func NewStringRef(v interface{}) *String {
	return NewString(NewRef(v))
}

func NewIntegerRef(v interface{}) *Integer {
	return NewInteger(NewRef(v))
}

func NewLongRef(v interface{}) *Long {
	return NewLong(NewRef(v))
}

func NewDoubleRef(v interface{}) *Double {
	return NewDouble(NewRef(v))
}

func NewBooleanRef(v interface{}) *Boolean {
	return NewBoolean(NewRef(v))
}

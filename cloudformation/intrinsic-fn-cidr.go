package cloudformation

type FnCidr struct {
	FnCidr []interface{} `json:"Fn::Cidr"`
}

func NewFnCidr(ip, cnt, mask interface{}) *FnCidr {
	return &FnCidr{FnCidr: []interface{}{NewString(ip), NewInteger(cnt), NewInteger(mask)}}
}

func NewStringFnCidr(ip, cnt, mask interface{}) *String {
	return NewString(NewFnCidr(ip, cnt, mask))
}

func NewIntegerFnCidr(ip, cnt, mask interface{}) *Integer {
	return NewInteger(NewFnCidr(ip, cnt, mask))
}

func NewLongFnCidr(ip, cnt, mask interface{}) *Long {
	return NewLong(NewFnCidr(ip, cnt, mask))
}

func NewDoubleFnCidr(ip, cnt, mask interface{}) *Double {
	return NewDouble(NewFnCidr(ip, cnt, mask))
}

func NewBooleanFnCidr(ip, cnt, mask interface{}) *Boolean {
	return NewBoolean(NewFnCidr(ip, cnt, mask))
}

package cloudformation

type FnFindInMap struct {
	FnFindInMap []interface{} `json:"Fn::FindInMap"`
}

func NewFnFindInMap(mn, tlk, slk interface{}) *FnFindInMap {
	return &FnFindInMap{FnFindInMap: []interface{}{NewString(mn), NewString(tlk), NewString(slk)}}
}

func NewStringFnFindInMap(mn, tlk, slk interface{}) *String {
	return NewString(NewFnFindInMap(mn, tlk, slk))
}

func NewIntegerFnFindInMap(mn, tlk, slk interface{}) *Integer {
	return NewInteger(NewFnFindInMap(mn, tlk, slk))
}

func NewLongFnFindInMap(mn, tlk, slk interface{}) *Long {
	return NewLong(NewFnFindInMap(mn, tlk, slk))
}

func NewDoubleFnFindInMap(mn, tlk, slk interface{}) *Double {
	return NewDouble(NewFnFindInMap(mn, tlk, slk))
}

func NewBooleanFnFindInMap(mn, tlk, slk interface{}) *Boolean {
	return NewBoolean(NewFnFindInMap(mn, tlk, slk))
}

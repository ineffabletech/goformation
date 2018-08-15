package cloudformation

import (
	"encoding/json"
	"fmt"
)

// String AWS CloudFormation Property (String)
type String struct {
	Value     string
	Intrinsic interface{}
}

func NewString(v interface{}) *String {
	ns := &String{"", nil}
	switch t := v.(type) {
	case string:
		ns.Value = t
	case []byte:
		ns.Value = string(t)
	case *String:
		return t
	default:
		ns.Intrinsic = v
	}
	return ns
}

func (p *String) value() interface{} {
	if p.Intrinsic != nil {
		return p.Intrinsic
	} else {
		return p.Value
	}
}

func (p *String) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.value())
}

func (p *String) UnmarshalJSON(b []byte) error {
	var typecheck interface{}
	if err := json.Unmarshal(b, &typecheck); err != nil {
		return err
	}

	switch val := typecheck.(type) {

	case string:
		p.Value = val

	case map[string]interface{}:
		p.Intrinsic = val

	default:
		return fmt.Errorf("invalid type for String")
	}

	return nil
}

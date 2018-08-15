package cloudformation

import (
	"encoding/json"
	"fmt"
)

// Boolean AWS CloudFormation Property (Boolean)
type Boolean struct {
	Value     bool
	Intrinsic interface{}
}

func NewBoolean(v interface{}) *Boolean {
	ns := &Boolean{false, nil}
	switch t := v.(type) {
	case bool:
		ns.Value = t
	case *Boolean:
		return t
	default:
		ns.Intrinsic = v
	}
	return ns
}

func (p *Boolean) value() interface{} {
	if p.Intrinsic != nil {
		return p.Intrinsic
	} else {
		return p.Value
	}
}

func (p *Boolean) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.value())
}

func (p *Boolean) UnmarshalJSON(b []byte) error {
	var typecheck interface{}
	if err := json.Unmarshal(b, &typecheck); err != nil {
		return err
	}

	switch val := typecheck.(type) {

	case bool:
		p.Value = val

	case map[string]interface{}:
		p.Intrinsic = val

	default:
		return fmt.Errorf("invalid type for Boolean")
	}

	return nil
}

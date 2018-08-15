package cloudformation

import (
	"encoding/json"
	"fmt"
)

// Int AWS CloudFormation Property (Long/Integer)
type Integer struct {
	Value     int
	Intrinsic interface{}
}

func NewInteger(v interface{}) *Integer {
	ns := &Integer{0, nil}
	switch t := v.(type) {
	case int:
		ns.Value = int(t)
	case int8:
		ns.Value = int(t)
	case int16:
		ns.Value = int(t)
	case int32:
		ns.Value = int(t)
	case *Integer:
		return t
	default:
		ns.Intrinsic = v
	}
	return ns
}

func (p *Integer) value() interface{} {
	if p.Intrinsic != nil {
		return p.Intrinsic
	} else {
		return p.Value
	}
}

func (p *Integer) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.value())
}

func (p *Integer) UnmarshalJSON(b []byte) error {
	var typecheck interface{}
	if err := json.Unmarshal(b, &typecheck); err != nil {
		return err
	}

	switch val := typecheck.(type) {

	case float64:
		p.Value = int(val)

	case map[string]interface{}:
		p.Intrinsic = val

	default:
		return fmt.Errorf("invalid type for Integer")
	}

	return nil
}

package cloudformation

import (
	"encoding/json"
	"fmt"
)

// Int AWS CloudFormation Property (Long/Integer)
type Long struct {
	Value     int64
	Intrinsic interface{}
}

func NewLong(v interface{}) *Long {
	ns := &Long{0, nil}
	switch t := v.(type) {
	case int:
		ns.Value = int64(t)
	case int8:
		ns.Value = int64(t)
	case int16:
		ns.Value = int64(t)
	case int32:
		ns.Value = int64(t)
	case int64:
		ns.Value = int64(t)
	case *Long:
		return t
	default:
		ns.Intrinsic = v
	}
	return ns
}

func (p *Long) value() interface{} {
	if p.Intrinsic != nil {
		return p.Intrinsic
	} else {
		return p.Value
	}
}

func (p *Long) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.value())
}

func (p *Long) UnmarshalJSON(b []byte) error {
	var typecheck interface{}
	if err := json.Unmarshal(b, &typecheck); err != nil {
		return err
	}

	switch val := typecheck.(type) {

	case float64:
		p.Value = int64(val)

	case map[string]interface{}:
		p.Intrinsic = val

	default:
		return fmt.Errorf("invalid type for Long")
	}

	return nil
}

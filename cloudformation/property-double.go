package cloudformation

import (
	"encoding/json"
	"fmt"
)

// Int AWS CloudFormation Property (Double)
type Double struct {
	Value     float64
	Intrinsic interface{}
}

func NewDouble(v interface{}) *Double {
	ns := &Double{0, nil}
	switch t := v.(type) {
	case float32:
		ns.Value = float64(t)
	case float64:
		ns.Value = float64(t)
	case *Double:
		return t
	default:
		ns.Intrinsic = v
	}
	return ns
}

func (p *Double) value() interface{} {
	if p.Intrinsic != nil {
		return p.Intrinsic
	} else {
		return p.Value
	}
}

func (p *Double) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.value())
}

func (p *Double) UnmarshalJSON(b []byte) error {
	var typecheck interface{}
	if err := json.Unmarshal(b, &typecheck); err != nil {
		return err
	}

	switch val := typecheck.(type) {

	case float64:
		p.Value = val

	case map[string]interface{}:
		p.Intrinsic = val

	default:
		return fmt.Errorf("invalid type for Double")
	}

	return nil
}

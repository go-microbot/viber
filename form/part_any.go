package form

import (
	"fmt"
	"mime/multipart"
)

// PartAny represents form-data text part (without strong type) model.
type PartAny struct {
	v interface{}
}

// Marshal encodes part text to writer.
// To encode value will be used fmt.Sprintf("%v", ...).
func (p PartAny) Marshal(w *multipart.Writer, partName string, omitempty bool) error {
	var value interface{}
	switch {
	case p.v == nil:
		if omitempty {
			return nil
		}
	default:
		value = p.v
	}

	return w.WriteField(partName, fmt.Sprintf("%v", value))
}

// NewPartAny returns new PartAny instance.
func NewPartAny(v interface{}) PartAny {
	return PartAny{
		v: v,
	}
}

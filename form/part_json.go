package form

import (
	"encoding/json"
	"mime/multipart"
)

// PartJSON represents form-data text part (with JSON format) model.
type PartJSON struct {
	v interface{}
}

// Marshal encodes part text (JSON) to writer.
func (p PartJSON) Marshal(w *multipart.Writer, partName string, omitempty bool) error {
	var value interface{}
	switch {
	case p.v == nil:
		if omitempty {
			return nil
		}
	default:
		value = p.v
	}

	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return w.WriteField(partName, string(data))
}

// NewPartJSON returns new PartJSON instance.
func NewPartJSON(v interface{}) PartJSON {
	return PartJSON{
		v: v,
	}
}

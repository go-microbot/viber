package form

import "mime/multipart"

// PartText represents form-data text part model.
type PartText struct {
	text *string
}

// Marshal encodes part text to writer.
func (p PartText) Marshal(w *multipart.Writer, partName string, omitempty bool) error {
	var value string
	switch {
	case p.text == nil:
		if omitempty {
			return nil
		}
	default:
		value = *p.text
	}

	return w.WriteField(partName, value)
}

// NewPartText returns new PartText instance.
func NewPartText(text string) PartText {
	return PartText{
		text: &text,
	}
}

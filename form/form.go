package form

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"reflect"
	"strings"
)

// Part represents general form part.
type Part interface {
	Marshal(w *multipart.Writer, partName string, omitempty bool) error
}

// Marshal encodes form's data value into a bytes slice.
func Marshal(val interface{}, ct *string) ([]byte, error) {
	typeOf := reflect.TypeOf(val)
	valueOf := reflect.ValueOf(val)
	fPType := reflect.TypeOf((*Part)(nil)).Elem()
	formBody := &bytes.Buffer{}
	writer := multipart.NewWriter(formBody)

	for i := 0; i < typeOf.NumField(); i++ {
		// check if field is implemented `Part` interface.
		field := typeOf.Field(i)
		if !field.Type.Implements(fPType) {
			continue
		}

		var omitempty bool
		partName := field.Name

		// parse form tag and detect part name.
		formTag, ok := field.Tag.Lookup("form")
		if formTag = strings.TrimSpace(formTag); formTag == "-" {
			continue
		}
		if ok {
			partValues := strings.Split(formTag, ",")
			partName = partValues[0]
			omitempty = strings.Contains(formTag, ",omitempty")
		}

		// call `Marshal` method.
		fieldValueOf := valueOf.FieldByName(field.Name)
		if fieldValueOf.CanInterface() {
			v := fieldValueOf.Interface()
			if v == nil {
				continue
			}
		}
		values := fieldValueOf.MethodByName("Marshal").Call([]reflect.Value{
			reflect.ValueOf(writer),
			reflect.ValueOf(partName),
			reflect.ValueOf(omitempty),
		})
		errOut := values[0].Interface()
		if errOut != nil {
			return nil, fmt.Errorf("could not marshal %s value: %v", field.Name, errOut)
		}
	}

	if err := writer.Close(); err != nil {
		return nil, err
	}

	if ct != nil {
		*ct = writer.FormDataContentType()
	}

	return formBody.Bytes(), nil
}

package query

import (
	"reflect"
	"strings"
)

// Param represents general query param.
type Param interface {
	Value() string
	Lookup() (string, bool)
}

// AsMap converts model that contains Param fields
// to the map with string key and string value.
func AsMap(val interface{}) map[string]string {
	queryMap := make(map[string]string)
	typeOf := reflect.TypeOf(val)
	valueOf := reflect.ValueOf(val)
	qPType := reflect.TypeOf((*Param)(nil)).Elem()

	for i := 0; i < typeOf.NumField(); i++ {
		field := typeOf.Field(i)
		if !field.Type.Implements(qPType) {
			continue
		}

		values := valueOf.FieldByName(field.Name).MethodByName("Value").Call(nil)
		if len(values) == 0 {
			continue
		}
		valueResult := values[0].String()

		values = valueOf.FieldByName(field.Name).MethodByName("Lookup").Call(nil)
		if len(values) == 0 {
			continue
		}
		lookupValueResult, lookupOKResult := values[0].String(), values[1].Bool()

		// parse query tag.
		queryTag, ok := field.Tag.Lookup("query")
		if !ok {
			queryMap[field.Name] = valueResult
			continue
		}

		var name, value string
		queryValues := strings.Split(queryTag, ",")
		name = queryValues[0]
		omitempty := strings.Contains(queryTag, ",omitempty")
		switch {
		case omitempty && lookupOKResult:
			value = lookupValueResult
		case omitempty && !lookupOKResult:
			continue
		case !omitempty && lookupOKResult:
			value = lookupValueResult
		default:
			value = valueResult
		}

		queryMap[name] = value
	}

	return queryMap
}

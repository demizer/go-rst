package parse

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"

	"golang.org/x/text/unicode/norm"
)

type equalityCheck struct {
	id                  int
	fieldNum            int
	expectedFieldExists bool
	t                   *testing.T

	parsedValue        reflect.Value
	parsedFieldName    string
	parsedFieldValue   reflect.Value
	expectedValue      reflect.Value
	expectedFieldName  string
	expectedFieldValue reflect.Value
}

func newEqualityCheck(t *testing.T, parsedVal reflect.Value, expectedVal reflect.Value) *equalityCheck {
	return &equalityCheck{
		id:            int(parsedVal.FieldByName("ID").Interface().(ID)),
		parsedValue:   parsedVal,
		expectedValue: expectedVal,
	}
}

func (e *equalityCheck) error() {
	var got, exp string
	switch r := e.parsedValue.Interface().(type) {
	case ID:
		got = e.parsedValue.Interface().(ID).String()
		exp = e.expectedValue.Interface().(ID).String()
	case itemElement:
		got = e.parsedValue.Interface().(itemElement).String()
		exp = e.expectedValue.Interface().(itemElement).String()
	case Line:
		got = e.parsedValue.Interface().(Line).String()
		exp = e.expectedValue.Interface().(Line).String()
	case StartPosition:
		got = e.parsedValue.Interface().(StartPosition).String()
		exp = e.expectedValue.Interface().(StartPosition).String()
	case int:
		got = strconv.Itoa(e.parsedValue.Interface().(int))
		exp = strconv.Itoa(e.expectedValue.Interface().(int))
	case string:
		got = e.parsedValue.Interface().(string)
		exp = e.expectedValue.Interface().(string)
	default:
		panic(fmt.Sprintf("%T is not implemented!", r))
	}
	e.t.Errorf("\n(ID: %d) Got: %s = %q, Expect: %s = %q", e.id, e.parsedFieldName, got, e.expectedFieldName, exp)
}

func (e *equalityCheck) check(fieldNum int) error {
	e.fieldNum = fieldNum
	e.expectedFieldValue = e.expectedValue.Field(e.fieldNum)
	e.expectedFieldName = e.expectedValue.Type().Field(e.fieldNum).Name

	var parsedValueStruct reflect.StructField
	parsedValueStruct, e.expectedFieldExists = e.parsedValue.Type().FieldByName(e.expectedFieldName)
	e.parsedFieldValue = e.parsedValue.FieldByName(e.expectedFieldName)
	e.parsedFieldName = parsedValueStruct.Name

	if !e.expectedFieldExists {
		return fmt.Errorf("ID: %d does not contain field %q", e.id, e.expectedFieldName)
	}

	// Handle special cases when comparing fields
	switch e.expectedFieldName {
	case "Text":
		e.expectedFieldValue = reflect.ValueOf(norm.NFC.String(e.expectedFieldValue.Interface().(string)))
	case "StartPosition":
		if e.parsedFieldValue.Interface().(StartPosition) == 1 {
			// Ignore StartPositions that begin at 1 from the parsed output items. This allows
			// startPosition to be excluded from the expected items tests (*_items.json).
			return nil
		}
	}

	if e.expectedFieldValue.Interface() != e.parsedFieldValue.Interface() {
		e.error()
	}

	return nil
}

package flatstruct

import (
	"errors"
	"reflect"
	"strconv"
)

// converts struct to map[string]string <=> map[{field json tag}]{field string value}
// Note:
// 	- every structure fields must have json tag
// 	- structure field value cannot be type of Array, Map, Slice, Interface
//	- boolean converts to string(uint8) - true and false <=> "1" and "0"
func StructToFlatMap(event interface{}) (map[string]string, error) {
	rv := reflect.ValueOf(event)
	if rv.Kind() == reflect.Ptr {
		return nil, errors.New("must pass a value, not a pointer")
	}

	rt := reflect.TypeOf(event)
	if rt.Kind() != reflect.Struct {
		return nil, errors.New("must pass a structure")
	}

	flatEvent := map[string]string{}
	for i := 0; i < rt.NumField(); i++ {
		fieldName := rt.Field(i).Tag.Get("json")
		if fieldName == "" {
			return nil, errors.New("every struct field must have json tag")
		}

		value, err := PrimitiveTypeToString(rv.Field(i))
		if err != nil {
			return nil, err
		}

		flatEvent[fieldName] = value
	}

	return flatEvent, nil
}

// converts any primitive type to string
func PrimitiveTypeToString(v reflect.Value) (string, error) {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.Interface:
		return "", errors.New("array, map, slice, interface as value are not allowed")
	case reflect.String:
		return v.String(), nil
	case reflect.Bool:
		if v.Bool() {
			return "1", nil
		}
		return "0", nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10), nil
	case reflect.Float32:
		return strconv.FormatFloat(v.Float(), 'f', 6, 64), nil
	case reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', 6, 32), nil
	case reflect.Ptr:
		if v.IsNil() {
			return "", nil
		}
		return PrimitiveTypeToString(v.Elem())
	}

	return "", errors.New("unknown value primitive type given")
}

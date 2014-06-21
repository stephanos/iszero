package iszero

import "reflect"

// Value returns whether the passed-in value is equal to its type's zero value.
func Value(src interface{}) bool {
	if src == nil {
		return true
	}

	switch v := src.(type) {
	case bool:
		return (v == false)
	case string:
		return (v == "")
	case int:
		return (v == int(0))
	case int8:
		return (v == int8(0))
	case int16:
		return (v == int16(0))
	case int32:
		return (v == int32(0))
	case int64:
		return (v == int64(0))
	case uint:
		return (v == uint(0))
	case uint8:
		return (v == uint8(0))
	case uint16:
		return (v == uint16(0))
	case uint32:
		return (v == uint32(0))
	case uint64:
		return (v == uint64(0))
	case float32:
		return (v == float32(0))
	case float64:
		return (v == float64(0))
	case reflect.Value:
		return checkReflectValue(src, v)
	default:
		return checkReflectValue(src, reflect.ValueOf(src))
	}
}

func checkReflectValue(src interface{}, val reflect.Value) bool {
	kind := val.Kind()
	switch kind {
	case reflect.Func, reflect.Map, reflect.Slice:
		return val.IsNil()
	case reflect.Array:
		z := true
		for i := 0; i < val.Len(); i++ {
			z = z && Value(val.Index(i))
		}
		return z
	case reflect.Interface, reflect.Ptr:
		return val.IsNil()
	default:
		zero := reflect.Zero(val.Type()).Interface()
		if kind == reflect.Struct {
			return reflect.DeepEqual(src, zero)
		}
		return val.Interface() == zero
	}
}

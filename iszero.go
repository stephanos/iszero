package iszero

import "reflect"

// Value returns whether the passed-in value is equal to its type's zero value.
func Value(src interface{}) bool {
	if src == nil {
		return true
	}

	if val, ok := src.(reflect.Value); ok {
		return checkReflectValue(nil, val)
	}

	return checkReflectValue(src, reflect.ValueOf(src))
}

func checkReflectValue(src interface{}, val reflect.Value) bool {
	if !val.IsValid() {
		return true
	}

	kind := val.Kind()
	switch kind {
	case reflect.String:
		return val.Len() == 0
	case reflect.Bool:
		return val.Bool() == false
	case reflect.Float32, reflect.Float64:
		return val.Float() == 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return val.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return val.Uint() == 0
	case reflect.Ptr, reflect.Chan, reflect.Func, reflect.Interface, reflect.Slice, reflect.Map:
		return val.IsNil()
	case reflect.Array:
		z := true
		for i := 0; i < val.Len(); i++ {
			z = z && Value(val.Index(i))
		}
		return z
	default:
		if src == nil {
			src = val.Interface()
		}

		zero := reflect.Zero(val.Type()).Interface()
		if kind == reflect.Struct {
			return reflect.DeepEqual(src, zero)
		}
		return src == zero
	}
}

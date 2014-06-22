package iszero

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

type MyString string

type MyStruct struct {
	num  int
	text MyString
}

var (
	result bool

	zeroPtr    *string
	zeroSlice  []int
	zeroFunc   func() string
	zeroMap    map[string]string
	emptyIface interface{}
	zeroIface  fmt.Stringer
	zeroValues = []interface{}{
		nil,

		// bool
		false,

		// int
		0,
		int8(0),
		int16(0),
		int32(0),
		int64(0),
		uint(0),
		uint8(0),
		uint16(0),
		uint32(0),
		uint64(0),

		// float
		0.0,
		float32(0.0),
		float64(0.0),

		// string
		"",

		// alias
		MyString(""),

		// func
		zeroFunc,

		// array / slice
		[0]int{},
		zeroSlice,

		// map
		zeroMap,

		// interface
		emptyIface,
		zeroIface,

		// pointer
		zeroPtr,

		// struct
		MyStruct{},
		time.Time{},
		MyStruct{num: 0},
		MyStruct{text: MyString("")},
	}
	zeroReflectValues = convertToReflectValues(zeroValues)

	nonZeroIface  fmt.Stringer = time.Now()
	nonZeroValues              = []interface{}{
		// bool
		true,

		// int
		1,
		int8(1),
		int16(1),
		int32(1),
		int64(1),
		uint8(1),
		uint16(1),
		uint32(1),
		uint64(1),

		// float
		1.0,
		float32(1.0),
		float64(1.0),

		// string
		"test",

		// alias
		MyString("test"),

		// func
		time.Now,

		// array / slice
		[]int{},
		[]int{42},
		[1]int{42},

		// map
		make(map[string]string, 1),

		// interface
		nonZeroIface,

		// pointer
		&nonZeroIface,

		// struct
		MyStruct{num: 1},
		time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
	}
	nonZeroReflectValues = convertToReflectValues(nonZeroValues)

	allValues        = append(zeroValues, nonZeroValues...)
	allReflectValues = append(zeroReflectValues, nonZeroReflectValues...)
)

func TestIsZeroCheck(t *testing.T) {
	for _, value := range append(zeroValues, zeroReflectValues...) {
		if !Value(value) {
			t.Errorf("expected '%v' (%T) to be recognized as zero value", value, value)
		}
	}

	for _, value := range nonZeroValues {
		if Value(value) {
			t.Errorf("did not expect '%v' (%T) to be recognized as zero value", value, value)
		}
	}
}

func BenchmarkIsZeroCheck(b *testing.B) {
	benchmark(b, allValues)
}

func BenchmarkIsZeroCheckReflect(b *testing.B) {
	benchmark(b, allReflectValues)
}

func benchmark(b *testing.B, fixture []interface{}) {
	fixtureLen := len(fixture)

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		result = Value(fixture[i%fixtureLen])
	}
}

func convertToReflectValues(src []interface{}) []interface{} {
	out := make([]interface{}, len(src))
	for _, val := range src {
		//fmt.Printf("%v: %v\n", i, reflect.ValueOf(val))
		out = append(out, reflect.ValueOf(val))
	}
	return out
}

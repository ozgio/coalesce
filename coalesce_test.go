package coalesce

import (
	"errors"
	"testing"
	"time"
)

func TestAny(t *testing.T) {
	type intType int
	type strType string

	var emptyMap map[string]int
	var emptyArray []int
	var intTypeVal intType
	var strTypeVal strType

	var ptrInt *int
	var ptrMap *(map[string]int)

	tests := []struct {
		params   []interface{}
		expected interface{}
	}{
		{[]interface{}{nil, ptrInt, ptrMap, emptyMap, emptyArray, intTypeVal, strTypeVal, "", "1"}, "1"},
		{[]interface{}{nil, strTypeVal, 0, ""}, nil},
		{[]interface{}{[]int{}, 1}, 1},
		{[]interface{}{[]int{}, 1, 2}, 1},
		// {[]interface{&testError{"1"}, errors.New("2")}, "1"},
		{[]interface{}{nil}, nil},
	}

	for i, test := range tests {
		res := Any(test.params...)
		if res != test.expected {
			t.Errorf("%d: '%v' is not expected result %v", i, res, test.expected)
		}
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		params   []string
		expected string
	}{
		{[]string{"", "1", "2"}, "1"},
		{[]string{"", "", "1"}, "1"},
		{[]string{"1"}, "1"},
		{[]string{"1", "2"}, "1"},
		{[]string{""}, ""},
	}

	for i, test := range tests {
		res := String(test.params...)
		if test.expected != res {
			t.Errorf("%d: '%v' is not expected result %v", i, res, test.expected)
		}
	}
}

func TestDuration(t *testing.T) {
	tests := []struct {
		params   []time.Duration
		expected time.Duration
	}{
		{[]time.Duration{0, time.Second, time.Millisecond}, time.Second},
		{[]time.Duration{0, 0, time.Second}, time.Second},
		{[]time.Duration{time.Second, time.Millisecond}, time.Second},
		{[]time.Duration{0}, 0},
	}

	for i, test := range tests {
		res := Duration(test.params...)
		if test.expected != res {
			t.Errorf("%d: '%v' is not expected result %v", i, res, test.expected)
		}
	}
}

func TestTime(t *testing.T) {
	var zero time.Time
	var d = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		params   []time.Time
		expected time.Time
	}{
		{[]time.Time{zero, d}, d},
		{[]time.Time{zero, zero, d}, d},
		{[]time.Time{zero}, zero},
	}

	for i, test := range tests {
		res := Time(test.params...)
		if test.expected != res {
			t.Errorf("%d: '%v' is not expected result %v", i, res, test.expected)
		}
	}
}

type testError struct {
	message string
}

func (e testError) Error() string {
	return e.message
}

func TestError(t *testing.T) {
	var nilErrType *testError
	tests := []struct {
		params   []error
		expected string
	}{
		{[]error{nil, errors.New("1")}, "1"},
		{[]error{nil, nilErrType, errors.New("1")}, "1"},
		{[]error{errors.New("1"), errors.New("2")}, "1"},
		{[]error{&testError{"1"}, errors.New("2")}, "1"},
		{[]error{nil}, ""},
	}

	for i, test := range tests {
		res := Error(test.params...)
		if test.expected == "" {
			if res != nil {
				t.Errorf("%d: error should be nil", i)
				continue
			}
		} else if res == nil {
			t.Errorf("%d: was expecting nil result", i)
		} else if res.Error() != test.expected {
			t.Errorf("%d: %s is not expected error %s", i, res.Error(), test.expected)
		}
	}
}

func TestByte(t *testing.T) {
	tests := []struct {
		params   []byte
		expected byte
	}{
		{[]byte{0, 1, 2}, 1},
		{[]byte{0, 0, 1}, 1},
		{[]byte{1, 2}, 1},
		{[]byte{0}, 0},
	}

	for i, test := range tests {
		res := Byte(test.params...)
		if test.expected != res {
			t.Errorf("%d: '%v' is not expected result %v", i, res, test.expected)
		}
	}
}

func TestRune(t *testing.T) {
	tests := []struct {
		params   []rune
		expected rune
	}{
		{[]rune{0, 1, 2}, 1},
		{[]rune{0, 0, 1}, 1},
		{[]rune{1, 2}, 1},
		{[]rune{0}, 0},
	}

	for i, test := range tests {
		res := Rune(test.params...)
		if test.expected != res {
			t.Errorf("%d: '%v' is not expected result %v", i, res, test.expected)
		}
	}
}

func TestInt(t *testing.T) {
	tests := []struct {
		params   []int
		expected int
	}{
		{[]int{0, 1, 2}, 1},
		{[]int{0, 0, 1}, 1},
		{[]int{1, 2}, 1},
		{[]int{0}, 0},
	}

	for i, test := range tests {
		res := Int(test.params...)
		if test.expected != res {
			t.Errorf("%d: '%v' is not expected result %v", i, res, test.expected)
		}
	}
}

func TestInt8(t *testing.T) {
	tests := []struct {
		params   []int8
		expected int8
	}{
		{[]int8{0, 1, 2}, 1},
		{[]int8{0, 0, 1}, 1},
		{[]int8{1, 2}, 1},
		{[]int8{0}, 0},
	}

	for i, test := range tests {
		res := Int8(test.params...)
		if test.expected != res {
			t.Errorf("%d: '%v' is not expected result %v", i, res, test.expected)
		}
	}
}

func TestInt16(t *testing.T) {
	tests := []struct {
		params   []int16
		expected int16
	}{
		{[]int16{0, 1, 2}, 1},
		{[]int16{0, 0, 1}, 1},
		{[]int16{1, 2}, 1},
		{[]int16{0}, 0},
	}

	for i, test := range tests {
		res := Int16(test.params...)
		if test.expected != res {
			t.Errorf("%d: '%v' is not expected result %v", i, res, test.expected)
		}
	}
}

func TestInt32(t *testing.T) {
	tests := []struct {
		params   []int32
		expected int32
	}{
		{[]int32{0, 1, 2}, 1},
		{[]int32{0, 0, 1}, 1},
		{[]int32{1, 2}, 1},
		{[]int32{0}, 0},
	}

	for i, test := range tests {
		res := Int32(test.params...)
		if test.expected != res {
			t.Errorf("%d: '%v' is not expected result %v", i, res, test.expected)
		}
	}
}

func TestInt64(t *testing.T) {
	tests := []struct {
		params   []int64
		expected int64
	}{
		{[]int64{0, 1, 2}, 1},
		{[]int64{0, 0, 1}, 1},
		{[]int64{1, 2}, 1},
		{[]int64{0}, 0},
	}

	for i, test := range tests {
		res := Int64(test.params...)
		if test.expected != res {
			t.Errorf("%d: '%v' is not expected result %v", i, res, test.expected)
		}
	}
}

func TestUint(t *testing.T) {
	tests := []struct {
		params   []uint
		expected uint
	}{
		{[]uint{0, 1, 2}, 1},
		{[]uint{0, 0, 1}, 1},
		{[]uint{1, 2}, 1},
		{[]uint{0}, 0},
	}

	for i, test := range tests {
		res := Uint(test.params...)
		if test.expected != res {
			t.Errorf("%d: '%v' is not expected result %v", i, res, test.expected)
		}
	}
}

func TestUint8(t *testing.T) {
	tests := []struct {
		params   []uint8
		expected uint8
	}{
		{[]uint8{0, 1, 2}, 1},
		{[]uint8{0, 0, 1}, 1},
		{[]uint8{1, 2}, 1},
		{[]uint8{0}, 0},
	}

	for i, test := range tests {
		res := Uint8(test.params...)
		if test.expected != res {
			t.Errorf("%d: '%v' is not expected result %v", i, res, test.expected)
		}
	}
}

func TestUint16(t *testing.T) {
	tests := []struct {
		params   []uint16
		expected uint16
	}{
		{[]uint16{0, 1, 2}, 1},
		{[]uint16{0, 0, 1}, 1},
		{[]uint16{1, 2}, 1},
		{[]uint16{0}, 0},
	}

	for i, test := range tests {
		res := Uint16(test.params...)
		if test.expected != res {
			t.Errorf("%d: '%v' is not expected result %v", i, res, test.expected)
		}
	}
}

func TestUint32(t *testing.T) {
	tests := []struct {
		params   []uint32
		expected uint32
	}{
		{[]uint32{0, 1, 2}, 1},
		{[]uint32{0, 0, 1}, 1},
		{[]uint32{1, 2}, 1},
		{[]uint32{0}, 0},
	}

	for i, test := range tests {
		res := Uint32(test.params...)
		if test.expected != res {
			t.Errorf("%d: '%v' is not expected result %v", i, res, test.expected)
		}
	}
}

func TestUint64(t *testing.T) {
	tests := []struct {
		params   []uint64
		expected uint64
	}{
		{[]uint64{0, 1, 2}, 1},
		{[]uint64{0, 0, 1}, 1},
		{[]uint64{1, 2}, 1},
		{[]uint64{0}, 0},
	}

	for i, test := range tests {
		res := Uint64(test.params...)
		if test.expected != res {
			t.Errorf("%d: '%v' is not expected result %v", i, res, test.expected)
		}
	}
}

func TestFloat32(t *testing.T) {
	tests := []struct {
		params   []float32
		expected float32
	}{
		{[]float32{0, 1.2, 2.5}, 1.2},
		{[]float32{0, 0, 0.01}, 0.01},
		{[]float32{1, 2}, 1},
		{[]float32{0}, 0},
		{[]float32{0.00001, 1}, 0.00001},
	}

	for i, test := range tests {
		res := Float32(test.params...)
		if test.expected != res {
			t.Errorf("%d: '%v' is not expected result %v", i, res, test.expected)
		}
	}
}

func TestFloat64(t *testing.T) {
	tests := []struct {
		params   []float64
		expected float64
	}{
		{[]float64{0, 1.2, 2.5}, 1.2},
		{[]float64{0, 0, 0.01}, 0.01},
		{[]float64{1, 2}, 1},
		{[]float64{0}, 0},
		{[]float64{0.00001, 1}, 0.00001},
	}

	for i, test := range tests {
		res := Float64(test.params...)
		if test.expected != res {
			t.Errorf("%d: '%v' is not expected result %v", i, res, test.expected)
		}
	}
}

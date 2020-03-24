package coalesce

import (
	"reflect"
	"time"
)

// Any returns the first non empty value. Nils, zero values, empty arrays/slices/maps/channels
// regarded as empty
// Example:
//  	r := coalesce.Any(nil, 42) //r = 42
//  	r := coalesce.Any([]int{}, 0, "",  42) //r = 42
func Any(vals ...interface{}) interface{} {
	for _, v := range vals {
		if v != nil {
			val := reflect.ValueOf(v)
			switch {
			case (val.Kind() == reflect.Array || val.Kind() == reflect.Chan || val.Kind() == reflect.Map || val.Kind() == reflect.Slice):
				if val.Len() > 0 {
					return v
				}
			case !val.IsZero():
				return v
			case val.Kind() == reflect.Ptr && !val.IsNil():
				return v
			}
		}
	}
	return nil
}

// String returns the first non-empty string
// Example:
//		r := coalesce.String("", "Hello", "", "World") //r = "Hello"
func String(vals ...string) string {
	for _, v := range vals {
		if v != "" {
			return v
		}
	}
	return ""
}

// Duration returns the first non-zero duration
// Example:
//      var zeroDur time.Duration
//		r := coalesce.Duration(zeroDur, time.Second, 0, -1 * time.Second) //r = time.Second
func Duration(vals ...time.Duration) time.Duration {
	for _, v := range vals {
		if v != 0 {
			return v
		}
	}
	return 0
}

// Time returns the first non-zero time
// Example:
//      var zeroTime time.Time
//		var now := time.Now()
//		r := coalesce.Time(zeroTime, time.Now()) //r = now
func Time(vals ...time.Time) time.Time {
	for _, v := range vals {
		if !v.IsZero() {
			return v
		}
	}
	return time.Time{}
}

// Error returns the first non-nil error
// Example:
//		r := coalesce.Error(nil, errors.New("err 1"), errors.New("err 2")) //r.Error() = "err 1"
func Error(vals ...error) error {
	for _, v := range vals {
		if v != nil {
			if val := reflect.ValueOf(v); !val.IsNil() {
				return v
			}
		}
	}
	return nil
}

// Byte returns the first non-zero byte
// Example:
//		r := coalesce.Byte(0, 1, 0, -1) //r = 1
func Byte(vals ...byte) byte {
	for _, v := range vals {
		if v != 0 {
			return v
		}
	}
	return 0
}

// Rune returns the first non-zero rune
// Example:
//		r := coalesce.Rune(0, 1, 0, -1) //r = 1
func Rune(vals ...rune) rune {
	for _, v := range vals {
		if v != 0 {
			return v
		}
	}
	return 0
}

// Int returns the first non-zero int
// Example:
//		r := coalesce.Int(0, 1, 0, -1) //r = 1
func Int(vals ...int) int {
	for _, v := range vals {
		if v != 0 {
			return v
		}
	}
	return 0
}

// Int8 returns the first non-zero int8
// Example:
//		r := coalesce.Int8(0, 1, 0, -1) //r = 1
func Int8(vals ...int8) int8 {
	for _, v := range vals {
		if v != 0 {
			return v
		}
	}
	return 0
}

// Int16 returns the first non-zero int16
// Example:
//		r := coalesce.Int16(0, 1, 0, -1) //r = 1
func Int16(vals ...int16) int16 {
	for _, v := range vals {
		if v != 0 {
			return v
		}
	}
	return 0
}

// Int32  returns the first non-zero int32
// Example:
//		r := coalesce.Int32(0, 1, 0, -1) //r = 1
func Int32(vals ...int32) int32 {
	for _, v := range vals {
		if v != 0 {
			return v
		}
	}
	return 0
}

// Int64 returns the first non-zero int
// Example:
//		r := coalesce.Int64(0, 1, 0, -1) //r = 1
func Int64(vals ...int64) int64 {
	for _, v := range vals {
		if v != 0 {
			return v
		}
	}
	return 0
}

// Uint returns the first non-zero int
// Example:
//		r := coalesce.Uint(0, 1, 0) //r = 1
func Uint(vals ...uint) uint {
	for _, v := range vals {
		if v != 0 {
			return v
		}
	}
	return 0
}

// Uint8 returns the first non-zero uint8
// Example:
//		r := coalesce.Uint8(0, 1, 0) //r = 1
func Uint8(vals ...uint8) uint8 {
	for _, v := range vals {
		if v != 0 {
			return v
		}
	}
	return 0
}

// Uint16 returns the first non-zero uint16
// Example:
//		r := coalesce.Uint16(0, 1, 0) //r = 1
func Uint16(vals ...uint16) uint16 {
	for _, v := range vals {
		if v != 0 {
			return v
		}
	}
	return 0
}

// Uint32  returns the first non-zero uint32
// Example:
//		r := coalesce.Uint32(0, 1, 0) //r = 1
func Uint32(vals ...uint32) uint32 {
	for _, v := range vals {
		if v != 0 {
			return v
		}
	}
	return 0
}

// Uint64 returns the first non-zero uint64
// Example:
//		r := coalesce.Uint64(0, 1, 0) //r = 1
func Uint64(vals ...uint64) uint64 {
	for _, v := range vals {
		if v != 0 {
			return v
		}
	}
	return 0
}

// Float64 returns the first non-zero float64
// Example:
//		r := coalesce.Float64(0, 1, 0, -1) //r = 1
func Float64(vals ...float64) float64 {
	for _, v := range vals {
		if v != 0 {
			return v
		}
	}
	return 0
}

// Float32 returns the first non-zero float32
// Example:
//		r := coalesce.Float32(0, 1, 0, -1) //r = 1
func Float32(vals ...float32) float32 {
	for _, v := range vals {
		if v != 0 {
			return v
		}
	}
	return 0
}

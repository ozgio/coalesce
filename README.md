Coalesce
========

Coalesce functions returns the first non-nil/non-zero argument. 

## Functions

### Any ([Docs](https://godoc.org/github.com/ozgio/coalesce#Any))
Any returns the first non empty value. Nils, zero values, empty arrays/slices/maps/channels regarded
as empty

```go
coalesce.Any(nil, 42) //-> 42 
coalesce.Any([]int{}, 0, "",  42) //-> 42
```

### String ([Docs](https://godoc.org/github.com/ozgio/coalesce#String))
String returns the first non-empty string

Example:
```go
coalesce.String("", "Hello", "", "World") //-> "Hello"
```

### Duration ([Docs](https://godoc.org/github.com/ozgio/coalesce#Duration))
Duration returns the first non-zero duration
Example:
```go
var zeroDur time.Duration
coalesce.Duration(zeroDur, time.Second, 0, -1 * time.Second) //-> time.Second
```

### Time ([Docs](https://godoc.org/github.com/ozgio/coalesce#Time))
Time returns the first non-zero time

Example:
```go
var zeroTime time.Time
var now := time.Now()
coalesce.Time(zeroTime, time.Now()) //-> now
```

### Error ([Docs](https://godoc.org/github.com/ozgio/coalesce#Error))
Error returns the first non-nil error

Example:
```go
coalesce.Error(nil, errors.New("err 1"), errors.New("err 2")) //-> r.Error() = "err 1"
```

### Byte ([Docs](https://godoc.org/github.com/ozgio/coalesce#Byte))
Byte returns the first non-zero byte

Example:
```go
coalesce.Byte(0, 1, 0, -1) //-> 1
```

### Rune ([Docs](https://godoc.org/github.com/ozgio/coalesce#Rune))
Rune returns the first non-zero rune

Example:
```go
coalesce.Rune(0, 1, 0, -1) //-> 1
```

### Int ([Docs](https://godoc.org/github.com/ozgio/coalesce#Int))
Int returns the first non-zero int

Example:
```go
coalesce.Int(0, 1, 0, -1) //-> 1
```

### Int8 ([Docs](https://godoc.org/github.com/ozgio/coalesce#Int8))
Int8 returns the first non-zero int8

Example:
```go
coalesce.Int8(0, 1, 0, -1) //-> 1
```

### Int16 ([Docs](https://godoc.org/github.com/ozgio/coalesce#Int16))
Int16 returns the first non-zero int16

Example:
```go
coalesce.Int16(0, 1, 0, -1) //-> 1
```

### Int32 ([Docs](https://godoc.org/github.com/ozgio/coalesce#Int32))
Int32 returns the first non-zero int32

Example:
```go
coalesce.Int32(0, 1, 0, -1) //-> 1
```

### Int64 ([Docs](https://godoc.org/github.com/ozgio/coalesce#Int64))
Int64 returns the first non-zero int64

Example:
```go
coalesce.Int64(0, 1, 0, -1) //-> 1
```

### Uint ([Docs](https://godoc.org/github.com/ozgio/coalesce#Uint))
Uint returns the first non-zero int

Example:
```go
coalesce.Uint(0, 1, 0) //-> 1
```

### Uint8 ([Docs](https://godoc.org/github.com/ozgio/coalesce#Uint8))
Uint8 returns the first non-zero int8

Example:
```go
coalesce.Uint8(0, 1, 0) //-> 1
```

### Uint16 ([Docs](https://godoc.org/github.com/ozgio/coalesce#Uint16))
Uint16 returns the first non-zero int16

Example:
```go
coalesce.Uint16(0, 1, 0) //-> 1
```

### Uint32 ([Docs](https://godoc.org/github.com/ozgio/coalesce#Uint32))
Uint32 returns the first non-zero int32

Example:
```go
coalesce.Uint32(0, 1, 0) //-> 1
```

### Uint64 ([Docs](https://godoc.org/github.com/ozgio/coalesce#Uint64))
Uint64 returns the first non-zero int64

Example:
```go
coalesce.Uint64(0, 1, 0) //-> 1
```

### Float32 ([Docs](https://godoc.org/github.com/ozgio/coalesce#Float32))
Float32 returns the first non-zero float32

Example:
```go
coalesce.Float32(0, 1, 0, -1) //-> 1
```

### Float64 ([Docs](https://godoc.org/github.com/ozgio/coalesce#Float64))
Float64 returns the first non-zero float64

Example:
```go
coalesce.Float64(0, 1, 0, -1) //-> 1
```
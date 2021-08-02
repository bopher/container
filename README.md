# Container

Central container for managing dependencies and performing dependency injection.

## Create New Container

```go
import "github.com/bopher/container";
cnt := container.NewContainer();
```

## Usage

Container interface contains following methods:

### Register

Add new dependency to container

```go
// Signature:
Register(name string, dep interface{})

// Example:
cnt.Register("app-name", "My App")
```

### Resolve

Get dependency. this function return false as second return value if dependency not exists.

**Caution:** Resolve return data as `interface{}` type. if you need to parse data as special type use helper parser methods described later.

```go
// Signature:
Resolve(name string) (interface{}, bool)

// Example:
appName, exists := cnt.Resolve("app-name") // => "My App", true
```

### Exists

Check if dependency exists.

```go
// Signature:
Exists(name string) bool

// Example:
exists := cnt.Exists("app-name") // => true
```

### Get Value And Parse As Data Type

dependencies can resolved by type. You can pass a fallback value to resolve functions as default value if dependency not exists or has different type. this functions return false as second return value if dependency not exists.

Example:

```go
cnt.Register("app-name", "Bopher App")

name, exists := cnt.String("app-name", "my app") // => "Bopher App", true
name2, exists := cnt.String("non-exists", "golang") // => "golang", false
myBool, exists := cnt.Bool("app-name", false) // false, true
myBool2, exists := cnt.Bool("non-exists", false) // false, false
```

#### Resolve By Type Methods

```go
// Bool parse dependency as boolean
Bool(name string, fallback bool) (bool, bool)
// Int parse dependency as int
Int(name string, fallback int) (int, bool)
// Int8 parse dependency as int8
Int8(name string, fallback int8) (int8, bool)
// Int16 parse dependency as int16
Int16(name string, fallback int16) (int16, bool)
// Int32 parse dependency as int32
Int32(name string, fallback int32) (int32, bool)
// Int64 parse dependency as int64
Int64(name string, fallback int64) (int64, bool)
// UInt parse dependency as uint
UInt(name string, fallback uint) (uint, bool)
// UInt8 parse dependency as uint8
UInt8(name string, fallback uint8) (uint8, bool)
// UInt16 parse dependency as uint16
UInt16(name string, fallback uint16) (uint16, bool)
// UInt32 parse dependency as uint32
UInt32(name string, fallback uint32) (uint32, bool)
// UInt64 parse dependency as uint64
UInt64(name string, fallback uint64) (uint64, bool)
// Float32 parse dependency as float64
Float32(name string, fallback float32) (float32, bool)
// Float64 parse dependency as float64
Float64(name string, fallback float64) (float64, bool)
// String parse dependency as string
String(name string, fallback string) (string, bool)
// Bytes parse dependency as bytes array
Bytes(name string, fallback []byte) ([]byte, bool)
```

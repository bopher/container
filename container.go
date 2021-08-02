package container

// Container dependency manager interface
type Container interface {
	// Register add new dependency to container
	Register(name string, dep interface{})
	// Resolve get a dependency
	Resolve(name string) (interface{}, bool)
	// Exists check if dependency exists
	Exists(name string) bool
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
}

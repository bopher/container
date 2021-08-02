package container

type containerDriver struct {
	dependencies map[string]interface{}
}

// Register add new dependency to container
func (c *containerDriver) Register(name string, dep interface{}) {
	c.dependencies[name] = dep
}

// Resolve get a dependency
func (c *containerDriver) Resolve(name string) (interface{}, bool) {
	dep, exists := c.dependencies[name]
	return dep, exists
}

// Exists check if dependency exists
func (c *containerDriver) Exists(name string) bool {
	_, exists := c.dependencies[name]
	return exists
}

// Bool parse dependency as boolean
func (c *containerDriver) Bool(name string, fallback bool) (bool, bool) {
	var ok bool
	dep, ok := c.dependencies[name]
	if ok {
		res, ok := dep.(bool)
		if ok {
			return res, ok
		} else {
			return fallback, true
		}
	}

	return fallback, false
}

// Int parse dependency as int
func (c *containerDriver) Int(name string, fallback int) (int, bool) {
	var ok bool
	dep, ok := c.dependencies[name]
	if ok {
		res, ok := dep.(int)
		if ok {
			return res, ok
		} else {
			return fallback, true
		}
	}

	return fallback, false
}

// Int8 parse dependency as int8
func (c *containerDriver) Int8(name string, fallback int8) (int8, bool) {
	var ok bool
	dep, ok := c.dependencies[name]
	if ok {
		res, ok := dep.(int8)
		if ok {
			return res, ok
		} else {
			return fallback, true
		}
	}

	return fallback, false
}

// Int16 parse dependency as int16
func (c *containerDriver) Int16(name string, fallback int16) (int16, bool) {
	var ok bool
	dep, ok := c.dependencies[name]
	if ok {
		res, ok := dep.(int16)
		if ok {
			return res, ok
		} else {
			return fallback, true
		}
	}

	return fallback, false
}

// Int32 parse dependency as int32
func (c *containerDriver) Int32(name string, fallback int32) (int32, bool) {
	var ok bool
	dep, ok := c.dependencies[name]
	if ok {
		res, ok := dep.(int32)
		if ok {
			return res, ok
		} else {
			return fallback, true
		}
	}

	return fallback, false
}

// Int64 parse dependency as int64
func (c *containerDriver) Int64(name string, fallback int64) (int64, bool) {
	var ok bool
	dep, ok := c.dependencies[name]
	if ok {
		res, ok := dep.(int64)
		if ok {
			return res, ok
		} else {
			return fallback, true
		}
	}

	return fallback, false
}

// UInt parse dependency as uint
func (c *containerDriver) UInt(name string, fallback uint) (uint, bool) {
	var ok bool
	dep, ok := c.dependencies[name]
	if ok {
		res, ok := dep.(uint)
		if ok {
			return res, ok
		} else {
			return fallback, true
		}
	}

	return fallback, false
}

// UInt8 parse dependency as uint8
func (c *containerDriver) UInt8(name string, fallback uint8) (uint8, bool) {
	var ok bool
	dep, ok := c.dependencies[name]
	if ok {
		res, ok := dep.(uint8)
		if ok {
			return res, ok
		} else {
			return fallback, true
		}
	}

	return fallback, false
}

// UInt16 parse dependency as uint16
func (c *containerDriver) UInt16(name string, fallback uint16) (uint16, bool) {
	var ok bool
	dep, ok := c.dependencies[name]
	if ok {
		res, ok := dep.(uint16)
		if ok {
			return res, ok
		} else {
			return fallback, true
		}
	}

	return fallback, false
}

// UInt32 parse dependency as uint32
func (c *containerDriver) UInt32(name string, fallback uint32) (uint32, bool) {
	var ok bool
	dep, ok := c.dependencies[name]
	if ok {
		res, ok := dep.(uint32)
		if ok {
			return res, ok
		} else {
			return fallback, true
		}
	}

	return fallback, false
}

// UInt64 parse dependency as uint64
func (c *containerDriver) UInt64(name string, fallback uint64) (uint64, bool) {
	var ok bool
	dep, ok := c.dependencies[name]
	if ok {
		res, ok := dep.(uint64)
		if ok {
			return res, ok
		} else {
			return fallback, true
		}
	}

	return fallback, false
}

// Float32 parse dependency as float64
func (c *containerDriver) Float32(name string, fallback float32) (float32, bool) {
	var ok bool
	dep, ok := c.dependencies[name]
	if ok {
		res, ok := dep.(float32)
		if ok {
			return res, ok
		} else {
			return fallback, true
		}
	}

	return fallback, false
}

// Float64 parse dependency as float64
func (c *containerDriver) Float64(name string, fallback float64) (float64, bool) {
	var ok bool
	dep, ok := c.dependencies[name]
	if ok {
		res, ok := dep.(float64)
		if ok {
			return res, ok
		} else {
			return fallback, true
		}
	}

	return fallback, false
}

// String parse dependency as string
func (c *containerDriver) String(name string, fallback string) (string, bool) {
	var ok bool
	dep, ok := c.dependencies[name]
	if ok {
		res, ok := dep.(string)
		if ok {
			return res, ok
		} else {
			return fallback, true
		}
	}

	return fallback, false
}

// Bytes parse dependency as bytes array
func (c *containerDriver) Bytes(name string, fallback []byte) ([]byte, bool) {
	var ok bool
	dep, ok := c.dependencies[name]
	if ok {
		res, ok := dep.([]byte)
		if ok {
			return res, ok
		} else {
			return fallback, true
		}
	}

	return fallback, false
}

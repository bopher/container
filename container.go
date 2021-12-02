package container

import "github.com/bopher/caster"

// Container dependency manager interface
type Container interface {
	// Register add new dependency to container
	Register(name string, dep interface{})
	// Resolve get a dependency
	Resolve(name string) (interface{}, bool)
	// Exists check if dependency exists
	Exists(name string) bool
	// Cast parse dependency as caster
	Cast(name string) caster.Caster
}

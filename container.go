package container

import "github.com/bopher/caster"

// Container dependency manager interface
type Container interface {
	// Register add new dependency to container
	Register(name string, dep any)
	// Resolve get a dependency
	Resolve(name string) (any, bool)
	// Exists check if dependency exists
	Exists(name string) bool
	// Cast parse dependency as caster
	Cast(name string) caster.Caster
}

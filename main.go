package container

// NewContainer create a new container instance
func NewContainer() Container {
	c := new(cDriver)
	c.dependencies = make(map[string]any)
	return c
}

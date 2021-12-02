package container

import "github.com/bopher/caster"

type cDriver struct {
	dependencies map[string]interface{}
}

func (this *cDriver) Register(name string, dep interface{}) {
	this.dependencies[name] = dep
}

func (this cDriver) Resolve(name string) (interface{}, bool) {
	dep, exists := this.dependencies[name]
	return dep, exists
}

func (this cDriver) Exists(name string) bool {
	_, exists := this.dependencies[name]
	return exists
}

func (this cDriver) Cast(name string) caster.Caster {
	return caster.NewCaster(this.dependencies[name])
}

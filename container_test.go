package container_test

import (
	"testing"

	"github.com/bopher/container"
)

func TestContainer(t *testing.T) {
	c := container.NewContainer()
	c.Register("duration", 1000)
	cast := c.Cast("duration")
	_, err := cast.Int8()
	if err == nil {
		t.Errorf("1000 must can't cast to int8")
	}
	i, err := cast.Int()
	if err != nil {
		t.Fatal(err)
	}

	if i != 1000 {
		t.Error("Failed casting!")
	}
}

package main

import (
	"fmt"

	"github.com/bopher/container"
)

func main() {
	c := container.NewContainer()
	c.Register("duration", 1000)
	cast := c.Cast("duration")
	fmt.Println(cast.Int8())
	fmt.Println(cast.Int())
}

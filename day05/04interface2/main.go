package main

import "fmt"

type car interface {
	run()
}

type falali struct {
	name string
}

type baoshijie struct {
	name string
}

func (f falali) run() {
	fmt.Printf("%s 速度1000迈\n", f.name)
}

func (b baoshijie) run() {
	fmt.Printf("%s 速度1000迈\n", b.name)
}

func drive(c car) {
	c.run()
}

func main() {
	var f = falali{
		name: "法拉利",
	}

	var b = baoshijie{
		name: "保时捷",
	}
	drive(f)
	drive(b)
}

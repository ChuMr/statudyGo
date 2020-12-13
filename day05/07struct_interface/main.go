package main

import "fmt"

type animal interface {
	mover
	eater
}
type mover interface {
	move()
}

type eater interface {
	eat()
}

type cat struct {
	name string
}

func (c *cat) move() {
	fmt.Printf("%s move~~~~~~~~~~~~~~\n", c.name)
}

func (c *cat) eat() {
	fmt.Printf("%s eat~~~~~~~~~~~~~~\n", c.name)
}

func main() {
	var a animal
	c := cat{"蓝猫"}
	c.move()
	c.eat()

	a = &c
	fmt.Println(a)
}

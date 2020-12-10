package main

import "fmt"

type speak interface {
	speak()
}

type preson struct {
}

type dog struct {
}

type cat struct {
}

func (p preson) speak() {
	fmt.Println("嘤嘤嘤~")
}

func (d dog) speak() {
	fmt.Println("汪汪汪~")
}

func (c cat) speak() {
	fmt.Println("喵喵喵~")
}

func da(s speak) {
	s.speak()
}

func main() {
	var p preson
	var c cat
	var d dog
	da(p)
	da(c)
	da(d)
}

package main

import "fmt"

//方法,方法是作用于特定类型的函数
//
type dog struct {
	name string
}

func newDog(name string) dog {
	return dog{
		name: name,
	}
}

func (d dog) wang() {
	fmt.Printf("%s:汪汪汪", d.name)
}

func main() {
	d1 := newDog("小狗")
	d1.wang()

}

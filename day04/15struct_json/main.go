package main

import (
	"encoding/json"
	"fmt"
)

//结构体与json

type preson struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	p1 := preson{
		Name: "张三",
		Age:  18,
	}

	j, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("err: %s", err)
		return
	}

	fmt.Printf("%v", string(j))
	// fmt.Println(string(j))

	str := `{"name":"张三","age":12}`
	p2 := preson{}
	json.Unmarshal([]byte(str), &p2)
	fmt.Println(p2)

}

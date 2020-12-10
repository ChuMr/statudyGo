package main

// "calc"

//导入包，包名的导入是从GOPATH下的src目录下开始的,禁止循环导入包
func main() {

	//"zhangqian.com/day05/calc"
	//calc "zhangqian.com/day05/10calc"
	//src/zhangqian.com/day05/10calc/calc

	// ret := calc.Add(10, 20)
	// fmt.Println(ret)

	ret := calc.Add(1,23)
	fmt.Println(ret)

}

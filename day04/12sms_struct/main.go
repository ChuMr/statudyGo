package main

import (
	"fmt"
	"os"
)

//struct版学生管理系统
//查看 新增 删除 学生

type smsStudent struct {
	allStudent map[int]*student
}

type student struct {
	id   int
	name string
	age  int
}

func newStudent(id int, name string, age int) *student {
	return &student{
		id:   id,
		name: name,
		age:  age,
	}
}

func main() {

	// allStudent = make(map[int]*student, 50)
	var sms = smsStudent{
		allStudent: make(map[int]*student, 50),
	}
	// s = sms
	//1 打印菜单
	//2 等待用户选择要做什么
	//3 执行对应的函数
	for {
		fmt.Println(`
		主界面面板:
			1 查看学生
			2 新增学生
			3 删除学生
			4 退出
	  `)

		var choice int64

		fmt.Print("请输入面板序号:")
		fmt.Scan(&choice)
		fmt.Printf("您选择了:%d\n", choice)
		switch choice {
		case 1:
			sms.showAllStudent()
		case 2:
			sms.addStudent()
		case 3:
			sms.deleteStudent()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("无效的序号，请重新输入")
		}
	}

}

func (s smsStudent) showAllStudent() {
	if len(s.allStudent) <= 0 {
		fmt.Println("暂无数据，请新增学生")
	}

	for key, val := range s.allStudent {
		fmt.Printf("学号:%d,姓名:%s 年龄:%d\n", key, val.name, val.age)
	}
}

func (s smsStudent) addStudent() {
	// fmt.Println("新增学生")
	var (
		id   int
		name string
		age  int
	)
	fmt.Print("请输入新增学生id:")
	fmt.Scan(&id)

	fmt.Print("请输入新增学生姓名:")
	fmt.Scan(&name)

	fmt.Print("请输入新增学生年龄:")
	fmt.Scan(&age)

	newStudent := newStudent(id, name, age)
	s.allStudent[id] = newStudent
	fmt.Print("新增成功")
	// fmt.Printf("新增后的学生列表:%v", allStudent)
}

func (s smsStudent) deleteStudent() {
	var id int
	fmt.Print("请输入要删除的学生id:")
	fmt.Scan(&id)
	delete(s.allStudent, id)
}

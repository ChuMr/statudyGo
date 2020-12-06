package main

import (
	"fmt"
)

//struct版学生管理系统
//查看 新增 删除 学生

type studentMgr struct {
	allStudent map[int]student
}

type student struct {
	id   int
	name string
	age  int
}

func newStudent(id int, name string, age int) student {
	return student{
		id:   id,
		name: name,
		age:  age,
	}
}

//查看学生
func (s studentMgr) showAllStudent() {
	if len(s.allStudent) <= 0 {
		fmt.Println("暂无数据，请新增学生")
	}

	for key, val := range s.allStudent {
		fmt.Printf("学号:%d,姓名:%s 年龄:%d\n", key, val.name, val.age)
	}
}

//添加学生
func (s studentMgr) addStudent() {
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

//编辑学生
func (s studentMgr) editStudent() {
	var (
		id   int
		name string
	)
	fmt.Print("请输入要编辑的学生id:")
	fmt.Scan(&id)

	stu, ok := s.allStudent[id]
	if !ok {
		fmt.Println("查无此人")
	}

	fmt.Print("请输入要编辑的学生姓名:")
	fmt.Scan(&name)
	stu.name = name
	s.allStudent[id] = stu
}

//删除学生
func (s studentMgr) deleteStudent() {
	var id int
	fmt.Print("请输入要删除的学生id:")
	fmt.Scan(&id)
	delete(s.allStudent, id)
}

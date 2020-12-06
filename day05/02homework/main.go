package main

import (
	"fmt"
	"os"
)

//结构体版本管理系统
//查看 新增 删除 编辑

var sms studentMgr

func main() {

	// allStudent = make(map[int]*student, 50)

	sms = studentMgr{
		allStudent: make(map[int]student, 50),
	}
	// s = sms
	//1 打印菜单
	//2 等待用户选择要做什么
	//3 执行对应的函数
	for {
		fmt.Println("------------welcome sms------------")
		fmt.Println(`
		主界面面板:
			1 查看学生
			2 新增学生
			3 删除学生
			4 编辑学生
			5 退出
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
			sms.editStudent()
		case 5:
			os.Exit(1)
		default:
			fmt.Println("无效的序号，请重新输入")
		}
	}

}

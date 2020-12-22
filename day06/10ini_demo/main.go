package main

import (
	"fmt"
	"reflect"
)

// MysqlConfig mysql配置结构体
type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"passsword"`
}

type myInt int64

// 通过反射获取类型
func reflectType(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("type:%v kind:%v\n", t.Name(), t.Kind()) //type
}

// 通过反射设置变量的值
func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	// 反射中使用 Elem()方法获取指针对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

func main() {
	// var a *float32 // 指针
	// var b myInt    // 自定义类型
	// var c rune     // 类型别名
	// var d int64 = 1000
	// reflectType(a) // type: kind:ptr
	// reflectType(b) // type:myInt kind:int64
	// reflectType(c) // type:int32 kind:int32
	// fmt.Println(d)
	// reflectSetValue(&d)
	// fmt.Println(d)

}

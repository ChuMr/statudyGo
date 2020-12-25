package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

// type myInt int64

// 通过反射获取类型
// func reflectType(x interface{}) {
// 	t := reflect.TypeOf(x)
// 	fmt.Printf("type:%v kind:%v\n", t.Name(), t.Kind()) //type
// }

// // 通过反射设置变量的值
// func reflectSetValue(x interface{}) {
// 	v := reflect.ValueOf(x)
// 	// 反射中使用 Elem()方法获取指针对应的值
// 	if v.Elem().Kind() == reflect.Int64 {
// 		v.Elem().SetInt(200)
// 	}
// }

// MysqlConfig ...
type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"passsword"`
}

// RedisConfig ...
type RedisConfig struct {
	Host      string `ini:"host"`
	Port      int    `ini:"port"`
	Password  string `ini:"passsword"`
	Databases int    `ini:"databases"`
}

// Config ...
type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func loadIni(filename string, data interface{}) (err error) {

	// 1.1 参数校验，参数必须为指针类型
	t := reflect.TypeOf(data)
	if t.Kind() != reflect.Ptr {
		err = errors.New("data param should be a pointer")
		return
	}

	// 1.2 参数必须为结构体的指针
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("data param should be a struct pointer")
		return
	}

	// 2.1 读文件得到的字节类型
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	// 2.2 将字节类型转换为字符串
	lineSlice := strings.Split(string(b), "\r\n")

	// 2.3 for循环读取数据，过滤注释，空行
	// var structName string
	var structName string
	// var sectionName string
	for _, line := range lineSlice {

		// 3.1 每行过滤空格
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		// 3.2 如果是注释，直接跳过
		if strings.HasPrefix(line, "#") || strings.HasPrefix(line, ";") {
			continue
		}

		//3.3 如果是[]，就表示节,并取出sectionName
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			// 3.5 根据sectionName 找到对应的结构体
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if sectionName == field.Tag.Get("ini") {
					//找到对应的结构体
					structName = field.Name
				}
			}

			// continue
		} else {
			mapVal := strings.Split(line, "=")
			index := strings.Index(line, "=")
			key1 := mapVal[0]
			if index == -1 || key1 == "" {
				return
			}

			key := strings.TrimSpace(line[:index])
			// varlue := strings.TrimSpace(line[index+1:])

			v := reflect.ValueOf(data)
			sValue := v.Elem().FieldByName(structName)
			// sType := sValue.Type()

			var fieldName string
			// var fileType reflect.StructField
			for i := 0; i < sValue.NumField(); i++ {
				filed := sValue.Type().Field(i)
				// fileType = filed
				if filed.Tag.Get("ini") == key {
					fieldName = filed.Name
					break
				}

			}
			fileObj := sValue.FieldByName(fieldName)

			fmt.Println(fileObj)

			// fmt.Println(fileObj, fileType.Type.Kind())

		}
		//3.3.1 根据sectionName 查找对应的结构体

		//3.4 剔除不符合

		// return nil
		// fmt.Println(v)
		//3.5 取出对应的嵌套结构体
	}

	// fmt.Println(lineSlice)
	return nil

}

// Person ...
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// SetValueToStruct ...
func SetValueToStruct(name string, age int) *Person {
	p := &Person{}
	v := reflect.ValueOf(p).Elem()
	v.FieldByName("Name").Set(reflect.ValueOf(name))
	v.FieldByName("Age").Set(reflect.ValueOf(age))
	return p
}

func main() {

	// var cfg Config
	var mc MysqlConfig

	v := reflect.ValueOf(&mc)

	familyPtr := v.FieldByName("address")
	fmt.Println(familyPtr)

	// fmt.Println(v.NumField())

	// for i := 0; i < v.NumField(); i++ {
	// 	fieldType := v.Field(i)
	// 	// filed := v.Field(i)
	// 	// fmt.Println(filed)
	// 	fmt.Println(fieldType)
	// 	// fmt.Printf("name: %v  tag: '%v'\n", fieldType.Name, fieldType.Tag)
	// }

	// fmt.Println(v.FieldByName("address"))
	// sValue := v.Elem().FieldByName("address")
	// fmt.Println(sValue)

	// fmt.Println(v)

	// fmt.Println(v.FieldByName("address"))
	// var cfg Config
	// err := loadIni("./config.ini", &cfg)
	// if err != nil {
	// 	fmt.Printf("load ini failed,err:%v\n", err)
	// 	return
	// }
	// fmt.Println(cfg)

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

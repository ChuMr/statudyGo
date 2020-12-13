package main

import (
	"fmt"
	"time"
)

func main() {
	// time.Time()
	// fmt.Println(time.Now())
	// fmt.Println(time.Now().Date())
	// fmt.Println(time.Now().Year())
	// fmt.Println(time.Now().Month())
	// fmt.Println(time.Now().Day())

	// fmt.Println(time.Now().Hour())
	// fmt.Println(time.Now().Minute())
	// fmt.Println(time.Now().Second())

	// fmt.Println(time.Now().Unix())
	// fmt.Println(time.Now().UnixNano())

	//时间戳 格式化
	// timeLayout := "2006-01-02 15:04:05"
	// t := time.Unix(1607842510, 0).Format(timeLayout)
	// fmt.Println(t)

	//格式化日期转换为时间戳
	// dateTime := "2020-12-13 14:55:10"

	// tm2, _ := time.Parse("01/02/2006", "2020/12/13")
	// tm2, _ := time.Parse("2006-01-02", "2020-12-13")

	// fmt.Println(tm2.Unix())
	// loc, _ := time.LoadLocation("Local")
	// tmp, _ := time.ParseInLocation(timeLayout, dateTime, loc)
	// // t, _ := time.Parse(timeLayout, dateTime)
	// fmt.Println(tmp)

	// fmt.Println(t.Year())
	// fmt.Println(t.Month())
	// fmt.Println(t.Day())
	// fmt.Println(time.Now().Format("2006-01-02 15:04:05")) //格式化时间戳

	//两个日期的差
	now := time.Now()
	// fmt.Println(now)

	// loc, _ := time.LoadLocation("Asia/ShangHai")
	// fmt.Println(loc)

	// timeObj, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-12-15 15:04:05", loc)
	// fmt.Println(timeObj)

	// sub := timeObj.Sub(now)
	// fmt.Println(sub)

	loc, _ := time.LoadLocation("Local") //获取时区
	tmp, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-12-15 15:04:05", loc)
	fmt.Println(tmp)

	sub := tmp.Sub(now)
	fmt.Println(sub)

}

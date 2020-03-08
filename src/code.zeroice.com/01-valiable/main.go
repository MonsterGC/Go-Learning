package main

import "fmt"


func main(){
	//  不指定类型
	name := "小陈"
	// 指定类型
	var age int = 20
	// 多变量声明
	name1 , name2 , name3 := "陈1","陈2","陈3"
	var age1 , age2 , age3 int = 1,2,3
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(name1,name2,name3)
	fmt.Println(age1,age2,age3)

	var (
		name4 string = "测试"
		age4 int = 20
	)
	fmt.Println(name4,age4)

	var test1 int = 123
	var test2 int = test1
	test2 = 1231
	fmt.Println(test1,test2)
}
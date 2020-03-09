package main

import "fmt"

func main()  {
	name := []string{"小陈一号","小陈二号","小陈三号"}
	name2 := [] string{"小陈一号","小陈二号","小陈三号"}
	fmt.Println(name,name2)

	age := make([]int , 10 , 10)
	fmt.Println(age)
	fmt.Println(len(age),cap(age))
	age = append(age,10)
	fmt.Println(age)
	age[0] = 10
	fmt.Println(age)
	fmt.Println(age[:])
	fmt.Println(age[1:])
	fmt.Println(age[1:2])
	fmt.Println(len(age),cap(age))
	for _,value:= range age{
		fmt.Println(value)
	}

	author := make([]string,0,0)
	// var author2 = []string
	var author2 []string
	if author != nil {
		fmt.Println("这个不是一个空切片")
	}

	fmt.Println(author2)
	if author2 == nil {
		fmt.Println("这个就是空切片")
	}

	var doubleAge = make([]int,len(age),cap(age)*2)
	copy(doubleAge,age)
	fmt.Println(doubleAge)
	fmt.Println(len(doubleAge),cap(doubleAge))
}
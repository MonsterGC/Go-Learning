package main

import "fmt"

func main()  {
	var name = [...]string{"小陈一号","小陈二号"}
	var name2 = [5]string{"小陈一号","小陈二号","1","2","3"}
	fmt.Println(name)
	fmt.Println(name2)
	fmt.Println(name2[:2])
	fmt.Println(name2[:])
	fmt.Println(name2[1:])
}
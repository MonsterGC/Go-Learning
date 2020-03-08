package main

import "fmt"

func main()  {
	var x interface{}

	switch i := x.(type) {
		case nil:
			fmt.Printf(" x 的类型 :%T\n",i)
		case int:
			fmt.Println("int")
		case float32:
			fmt.Println("float32")		             
		case func(int) float64:
			fmt.Println("func(int)")
		default:
			fmt.Println("不知道")
	}

	switch{
		case false:
			fmt.Println("false")
			fallthrough
		case true:
			fmt.Println("true")
			fallthrough
		case false:
			fmt.Println("false")
			fallthrough
		case true:
			fmt.Println("true")
		case false:
			fmt.Print("没了")
		default:
			fmt.Println("真没了")
	}
}	
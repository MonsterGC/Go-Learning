package main

import "fmt"

const MAX = 3

func main()  {
	var p *int
	var num int = 10;
	 
	p = &num
	fmt.Println(*p)

	var age *int
	if(p != nil){
		fmt.Println("p不是空指针")
	}
	if(age == nil){
		fmt.Println("空指针")
	}

	a := [...]int{12,3,2}
	var i int
	var per [MAX]*int;

	for i = 0; i < MAX; i++{
		per[i] = &a[i]
	}
	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i , *per[i] )
	 }
}   
package main

import (
	"fmt"
	"math"
)

type person struct {
	name string
}

func add(x int , y int) int  {
	return x + y
}

func max(x int , y int) int  {
	if(x > y){
		x , y = y , x
	}
	return y
}

func swap(x int , y int)(int,int){
	return y,x 
}

func main()  {
	sum := add(10,20)
	fmt.Println(sum)
	max := max(20,10)
	fmt.Printf("最大的值是:%d\n",max)
	x := 10
	y := 20
	
	x , y = swap(x,y)
	fmt.Printf("x=%d,y=%d",x,y)


	fun1 := func (x float64) float64  {
		return math.Sqrt(x)
	}

	fmt.Println(fun1(9))
}
package main

import (
	"fmt"
	"time"
)

func say(s string)  {
	for i := 0; i < 9; i++ {
		time.Sleep(100*time.Millisecond)
		fmt.Printf("你真是个帅哥: %d\n" , s)
	}
}

func sum(list []int , c chan int)	{
	sum := 0
	for _ , v := range list {
		sum += v
	}
	time.Sleep(100*time.Millisecond)
	c <- sum
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < 2*n; i++ {
			c <- x
			x, y = y, x+y
	}
	close(c)
}

func main()  {
	// go say("小陈")
	// say("---------")
	
	s := []int{1,2,3,4,5,6,7,8,9,10}
	c := make(chan int)
	go sum(s , c)
	go sum( s[:len(s)/2] , c )
	x , y := <- c , <- c

	fmt.Printf("两者结果: %d ====> %d",x , y) // x,y 顺序不定

	d := make(chan int, 10)
    go fibonacci(cap(d), d)
    // range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
    // 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
    // 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
    // 会结束，从而在接收第 11 个数据的时候就阻塞了。
    for i := range d {
		fmt.Println(i)
	}        
}
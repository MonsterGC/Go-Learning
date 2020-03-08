package main

import (
	"fmt"
	"unsafe"
)

func main()  {
	// 隐式类型
	const name = "小陈"
	// 显式类型
	const nameTest string = "小陈2"

	const name1 , name2 , name3 = "小陈","小陈2","小陈3"
	fmt.Println(name1,name2,name3,"=>>>>>",name,nameTest)

	const WIDTH int = 10
	const HEIGHT int = 20

	fmt.Println(WIDTH*HEIGHT)


	const (
		a = "abc"
		b = len(a)
		c = unsafe.Sizeof(a)
	)

	fmt.Println(a,b,c)


	// iota
	type Kind uint
	const (
		d = iota
		e
		f
	)
	fmt.Println(d,e,f)


	// 中间插值
	const (
		g = iota
		h
		i
		j = 123
		k
		l = 123
		n = 12
		m
		o
		p
		q = iota
		r
	)
	fmt.Println(g,h,i,j,k,l,n,m,o,p,q,r)

	// 位运算
	const (
		n1 = 1<<iota
		n2 = 4<<iota
		n3
		n4
	)
	fmt.Println(n1,n2,n3,n4)
}
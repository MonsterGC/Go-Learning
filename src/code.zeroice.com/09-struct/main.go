package main

import "fmt"

type Book struct{
	name string
	author string 
	subject string
	id int
}

func testStruct(book Book)  {
	book.author = "小陈"
	fmt.Printf("这是一本: %d\n",book.name)
}


func testPointerStruct(book *Book)  {
	book.author = "小陈"
	fmt.Printf("这是一本: %d\n",book.name)
}

func main()  {
	var book Book
	book.name = "JS权威指南"
	book.author = "匿名者"
	book.subject = "暂无"
	book.id = 123

	book1 := Book{
		name : "暂无",
		author:"没有",
		id:124,
	}

	var book2 = Book{
		name:"没有啊",
		id:123,
	}

	fmt.Println(book)
	fmt.Println(book1)
	fmt.Println(book2)
	testStruct(book)
	fmt.Println(book.author)
	testPointerStruct(&book1)
	fmt.Println(book1.author)
}
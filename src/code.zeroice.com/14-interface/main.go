package main

import "fmt"


type Phone interface{
	call()
}

type Book struct{
	name string
}

type Person struct{
	name string
}

func (book *Book)call()  {
	fmt.Printf("我打算看一本书:%d \n",book.name)
}

func (per *Person)call()  {
	fmt.Printf("我想打电话给一个人:%d \n",per.name)
}

func call(i Phone)  {
	i.call()
}

func main()  {
	// var phone Phone
	// phone = new(Book)
	// phone.call()

	// phone = new(Person)
	// phone.call()
	book := &Book{
		name:"GO权威指南",
	}

	call(book)
}
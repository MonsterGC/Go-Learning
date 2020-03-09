package main

import "fmt"

func main() {
	name := map[string]int{"小陈":1231231,"小伟":123}
	for key,value := range name{
		fmt.Println(key,value)
	}

	book := make(map[string]int)
	book["GO权威指南"] = 1
	book["JS放弃学习"] = 2
	fmt.Println(book)

	_,ok := book["放弃吧"]
	if !ok{
		fmt.Println("放弃吧，没有数据")
	}


	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}

	delete(countryCapitalMap,"France")
	fmt.Println(countryCapitalMap)
}
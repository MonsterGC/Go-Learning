package main

import "fmt"

func main()  {
	name := []string{"小陈一号","小陈二号","小陈三号","小陈四号"}
	for _,value:=range name{
		fmt.Println(value)
	}
	kvs := map[string]string{"小陈":"很帅"}
	for i, c := range "go" {
        fmt.Println(i, c)
	}
	
	for _,value := range kvs{
		fmt.Println(value)
	}
}
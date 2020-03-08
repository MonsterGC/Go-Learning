package main

import "fmt"

func main()  {
	// for {
	// 	fmt.Println("无限寻缘")
	// }
	sum := 0
	for i:= 0;i<= 10;i++{
		sum += i
	}
	fmt.Println(sum)


	for ;sum > 0;{
		sum -= 5
	}
	fmt.Println(sum) 
	for sum < 55{
		sum += 5
	}
	fmt.Println(sum)
	strings := []string{"小陈一号","小陈2号"}
	for key,value := range strings{
		fmt.Printf("key=%d,value=%d",key,value)
	}

	var age int = 0

	for age < 21 {
		if age == 18{
			break
		}
		age++;
		fmt.Println(age)
	}
	fmt.Println("永远十八岁")
}
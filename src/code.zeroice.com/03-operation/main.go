package main

import "fmt"

func main()  {
	var a int = 21
	var b int = 10
	var c int

	// 算术运算符
	c = a + b
	fmt.Printf("相加%d\n",c)
	c = a - b
	fmt.Printf("相减%d\n",c)
	c = a * b
	fmt.Printf("相乘%d\n",c)
	c = a / b
	fmt.Printf("相除%d\n",c)
	a++
	fmt.Printf("输出后自增%d\n",a)
	// ++a
	// fmt.Printf("先自增后输出%d",a) // err
	a--
	fmt.Printf("先输出后自减%d\n",a)
	// --a
	// fmt.Printf("先自减后输出%d",a) // err
	c = a % 2
	fmt.Printf("求余%d\n",c)

	// 关系运算符
	fmt.Printf("等于%d\n",a == b)
	fmt.Printf("不等于%d\n",a != b)
	fmt.Printf("大于%d\n",a > b)
	fmt.Printf("小于%d\n",a < b)
	fmt.Printf("大于或等于%d\n",a >= b)
	fmt.Printf("小于或等于%d\n", a <= b)

	// 逻辑运算符
	var status1 bool  = true
	var status2 bool  = false

	if(status1 && !status2){
		fmt.Printf("两个相反1\n")
	} 
	// if(status1 || status2) fmt.Printf("两个相反2\n")
	if(status1 || status2) {
		fmt.Printf("两个相反2\n")
	}

	// 位运算符
	var d uint = 60      /* 60 = 0011 1100 */  
	var e uint = 13      /* 13 = 0000 1101 */
	var g uint = 0          

	g = d & e       /* 12 = 0000 1100 */ 
	fmt.Printf("第一行 - c 的值为 %d\n", g )

	g = d | e       /* 61 = 0011 1101 */
	fmt.Printf("第二行 - c 的值为 %d\n", g )

	g = d ^ e       /* 49 = 0011 0001 */
	fmt.Printf("第三行 - c 的值为 %d\n", g )

	g = d << 2     /* 240 = 1111 0000 */
	fmt.Printf("第四行 - c 的值为 %d\n", g )

	g = d >> 2     /* 15 = 0000 1111 */
	fmt.Printf("第五行 - c 的值为 %d\n", g )
}
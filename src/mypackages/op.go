package mypackages

import "fmt"

func My_relation_op() {
	var a int = 21
	var b int = 10

	if a == b {
		fmt.Printf("第一行 - a 等于 b\n")
	} else {
		fmt.Printf("第一行 - a 不等于 b\n")
	}
	if a < b {
		fmt.Printf("第二行 - a 小于 b\n")
	} else {
		fmt.Printf("第二行 - a 不小于 b\n")
	}
	if a > b {
		fmt.Printf("第三行 - a 大于 b\n")
	} else {
		fmt.Printf("第三行 - a 不大于 b\n")
	}
	/* Lets change value of a and b */
	a = 5
	b = 20
	if a <= b {
		fmt.Printf("第四行 - a 小于等于  b\n")
	}
	if b >= a {
		fmt.Printf("第五行 - b 大于等于 b\n")
	}
}
func My_logic_op() {
	var a bool = true
	var b bool = false
	if a && b { //逻辑与
		fmt.Printf("第一行 - 条件为 true\n")
	}
	if a || b { //逻辑或
		fmt.Printf("第二行 - 条件为 true\n")
	}
	/* 修改 a 和 b 的值 */
	a = false
	b = true
	if a && b {
		fmt.Printf("第三行 - 条件为 true\n")
	} else {
		fmt.Printf("第三行 - 条件为 false\n")
	}
	if !(a && b) {
		fmt.Printf("第四行 - 条件为 true\n")
	}
}
func My_bit_op() {
	/*
		&, |, 和 ^ 异或
		A = 0011 1100
		B = 0000 1101
		-----------------
		A&B = 0000 1100
		A|B = 0011 1101
		A^B = 0011 0001
		~A  = 1100 0011
	*/
	A := 60
	B := 13
	fmt.Println(A & B)
	fmt.Printf("%.8b\n", A&B) //  00001100
	fmt.Printf("%b\n", A&B)   //  1100
	fmt.Printf("%.8b", A<<2)  //以为，乘以2的平方
	var a uint = 60           /* 60 = 0011 1100 */
	var b uint = 13           /* 13 = 0000 1101 */
	var c uint = 0

	c = a & b /* 12 = 0000 1100 */
	fmt.Printf("第一行 - c 的值为 %d\n", c)

	c = a | b /* 61 = 0011 1101 */
	fmt.Printf("第二行 - c 的值为 %d\n", c)

	c = a ^ b /* 49 = 0011 0001 */
	fmt.Printf("第三行 - c 的值为 %d\n", c)

	c = a << 2 /* 240 = 1111 0000 */
	fmt.Printf("第四行 - c 的值为 %d\n", c)
	c = a >> 2 /* 15 = 0000 1111 */
	fmt.Printf("第五行 - c 的值为 %d\n", c)
}

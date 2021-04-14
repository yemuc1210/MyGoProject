package main

import "fmt"

func main() {
	var a int = 10
	var b float64 = 20.0
	fmt.Printf("变量的地址: %x\n", &a  )
	var ip *int = &a      /* 指向整型*/
	var fp *float64 = &b   /* 指向浮点型 */
	fmt.Printf("a 变量的地址是: %x\n", &a  )
	fmt.Printf("b 变量的地址是: %x\n", &b  )
	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量的存储地址: %x\n", ip )
	fmt.Printf("fp 变量的存储地址: %x\n", fp )
	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip )
	fmt.Printf("*fp 变量的值: %f\n", *fp )


	var  ptr *int

	fmt.Printf("ptr 的值为 : %x\n", ptr  )
	fmt.Println(ptr != nil)     /* ptr 不是空指针 */
	//(ptr == nil)    /* ptr 是空指针 */


}
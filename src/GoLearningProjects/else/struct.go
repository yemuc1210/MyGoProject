package main

import (
	"fmt"
	"time"
)

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func pointer() {

	var a int = 10
	var b float64 = 20.0
	fmt.Printf("变量的地址: %x\n", &a)
	var ip *int = &a     /* 指向整型*/
	var fp *float64 = &b /* 指向浮点型 */
	fmt.Printf("a 变量的地址是: %x\n", &a)
	fmt.Printf("b 变量的地址是: %x\n", &b)
	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量的存储地址: %x\n", ip)
	fmt.Printf("fp 变量的存储地址: %x\n", fp)
	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)
	fmt.Printf("*fp 变量的值: %f\n", *fp)

	var ptr *int

	fmt.Printf("ptr 的值为 : %x\n", ptr)
	fmt.Println(ptr != nil) /* ptr 不是空指针 */
	//(ptr == nil)    /* ptr 是空指针 */

}
func main() {
	var Book1 Books /* 声明 Book1 为 Books 类型 */
	var Book2 Books /* 声明 Book2 为 Books 类型 */

	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "www.w3cschool.cn"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407

	/* book 2 描述 */
	Book2.title = "Python 教程"
	Book2.author = "www.w3cschool.cn"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 6495700

	/* 打印 Book1 信息 */
	fmt.Printf("Book 1 title : %s\n", Book1.title)
	fmt.Printf("Book 1 author : %s\n", Book1.author)
	fmt.Printf("Book 1 subject : %s\n", Book1.subject)
	fmt.Printf("Book 1 book_id : %d\n", Book1.book_id)

	fmt.Println(Book1) //{Go 语言 www.w3cschool.cn Go 语言教程 6495407}
	/* 打印 Book2 信息 */
	fmt.Printf("Book 2 title : %s\n", Book2.title)
	fmt.Printf("Book 2 author : %s\n", Book2.author)
	fmt.Printf("Book 2 subject : %s\n", Book2.subject)
	fmt.Printf("Book 2 book_id : %d\n", Book2.book_id)
	//使用64位整数存储时间
	t := time.Now()
	fmt.Println(t)
}

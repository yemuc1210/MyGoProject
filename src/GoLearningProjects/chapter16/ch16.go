package main

import "fmt"

/*
数组
*/

func test(p [4]string) {
	fmt.Println(p, len(p))
	for i := range p {
		p[i] = "new" + p[i]
	}
	fmt.Println("test 内部", p)
}

// /**
// 实验 chess.go
// */
// func chess(){

// }
func main() {
	fmt.Println("数组章节")

	var planets [8]string
	var planets_1 = [...]string{"`1233,", "asdff", "qwee", "safafa"} //复合字面量初始化方法，
	// var planets_2 = []
	fmt.Println(planets_1, len(planets_1))
	fmt.Println(planets)

	// i := 1
	// a := planets_1[i]   //编译无错，运行出错，出现panic
	// fmt.Println(a)   //panic: runtime error: index out of range [1] with length 1

	//迭代遍历数组
	//(1)  len()   传统for循环可以，可以定制循环过程，设置终止条件、逆序循环等
	for i := 0; i < len(planets_1); i++ {
		a := planets_1[i]
		fmt.Println(i, a)
	}
	//(2)  range  可以避免越界的错误
	for index, a := range planets_1 {
		fmt.Println(index, a)
	}

	//数组作为值传递测试
	test(planets_1)
	fmt.Println("after, main:", planets_1)
	// test(planets)//cannot use planets (variable of type [8]string) as [4]string value in argument to tes
	fmt.Println("结论，数组作为值传递，而不是引用，函数操作的是数组的副本，此外，数组长度必须匹配.数组是值而不是引用")

	//数组复制测试，说明数组是值，而不是引用
	arrCpy := planets_1
	planets_1[3] = "3333333"
	fmt.Println(planets_1)
	fmt.Println(arrCpy)

	//数组的数组，二维数组
	var board [8][8]string
	board[0][0] = "r"
	board[0][7] = "r"

	for row := range board{   // range 返回 index 和值，值可省略
		for col := range board[row]{
			board[row][col] = "p"
		}
	}

	fmt.Print(board)
}

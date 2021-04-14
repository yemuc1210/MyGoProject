package mypackages //非main包，不可定义main

import (
	"fmt"  //提供格式化输入输出
	"math" //数学   Sin  Cos Tan Sqrt
)

func Pkg_func() { //func关键字用于声明函数    main函数是特殊的，是程序入口
	fmt.Println("hello world") //类似java中的Sys.println()
	fmt.Println("%s", "sss")
	fmt.Println("Sqrt(9)-%f", math.Sqrt(9.0))
}

// func main() {
// 	fmt.Println("chapter 1 \n")
// 	Pkg_func()
// }

/**
func main() //func关键字用于声明函数    main函数是特殊的，是程序入口
{      //大括号这种风格是错误的，只允许与main同行
	fmt.Println("hello world")         //类似java中的Sys.println()
	fmt.Println("%s","sss")
	fmt.Println("Sqrt(9)-%f",math.Sqrt(9.0))
}
*/

package mypackages //声明这是一个main包，所以可以定义主函数

import "fmt"

func My_fmt() { //package里面函数首字母需要大写，不然不识别
	/**
	  标识符首字母大写，等于public，可以被外部包访问，称为导出；首字母小写，包外不可见，private
	*/
	fmt.Println("格式化输出")
	fmt.Println("Println函数")
	fmt.Print("Print函数")
	fmt.Printf("Printf函数，My first go project, today is: %s \n", "4-12") //Printf

}

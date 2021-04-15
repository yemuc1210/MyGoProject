package main

/*
在这个路径下，就相当于建立一个新的项目
就不会存在一个项目中一个main包里有多个main主函数的情况
*/
import (
	"fmt"
	"mypackages"
	"time"
)

func main() {
	mypackages.Pkg_func()
	fmt.Println("测试提交，输出", time.Now())
}

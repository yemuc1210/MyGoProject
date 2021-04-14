package main

import (
	"fmt"
	"mypackages"
	"time"
)

func main() {
	mypackages.Pkg_func()
	fmt.Println("测试提交，输出", time.Now())
}

package main

import (
	"fmt"
	"time"
)

func main() {
	//使用64位整数存储时间
	t := time.Now()
	fmt.Println(t)
}

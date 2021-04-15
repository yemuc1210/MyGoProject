package main

import (
	"fmt"
	"mypackages"
	"time"
)


func mytime(){
	// 使用64位整数存储日期
	future := time.Unix(12622780800, 0)
	fmt.Println(future)
}

func main() {
	mypackages.My_time()
	mytime()
}

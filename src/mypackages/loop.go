package mypackages

import "fmt"

func My_loop() {
	var count = 10
	for count = 3; count > 0; count-- {
		fmt.Println(count)
	}
	fmt.Println(count) //3 2 1 0
	count = 10
	for count := 5; count > 0; count-- { //临时声明了一个变量，其值为5
		fmt.Println(count)
	}
	fmt.Println(count) //5 4 3 2 1 10
}

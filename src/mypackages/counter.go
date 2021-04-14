package mypackages

import (
	"fmt"
	"time"
)

func My_counter() {
	count := 10
	for count > 0 {

		time.Sleep(time.Second)
		fmt.Println(count)
		count--
	}
}

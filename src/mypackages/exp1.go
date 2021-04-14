package mypackages

import (
	"fmt"
	"math/rand"
)

func My_exp() {
	var rdm = rand.Intn(10) + 1 //伪随机数  0~9
	fmt.Println(rdm)
	//随机生成一个56 000 000 km到401 000 000 km
	rdm = rand.Intn(401000000-56000000) + 56000000 + 1
	fmt.Println(rdm / 28 / 24)
}

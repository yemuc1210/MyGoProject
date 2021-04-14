package mypackages

import (
	"fmt"
	"math/rand"
)

/**
模拟存钱，随机存入5美分、10美分、25美分
目标20。00美元
1美分=0.01美元
精确打印余额
*/
func My_piggy() {
	coin := []int{5, 10, 25}
	ibalance := 0
	//fbalance := 0.00
	//target := 20.00
	for {
		index := rand.Intn(3)
		ibalance += coin[index]
		//fbalance += ibalance/100.00
		dollars := ibalance / 100
		cents := ibalance % 100
		fmt.Printf("balance:%5.2f    $%d.%02d\n ", float64(ibalance)/100.0, dollars, cents)
		if ibalance >= 2000 {
			break
		}

	}
}

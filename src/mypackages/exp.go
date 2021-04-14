package mypackages

import (
	"fmt"
	"math/rand"
)

func My_guess() {
	var answer int = 75
	for {
		var guess = rand.Intn(100) + 1
		if guess == answer {
			fmt.Println("Success")
			break
		} else if guess > answer {
			fmt.Println(guess, "大于")
		} else {
			fmt.Println(guess, "小于")
		}
	}
}

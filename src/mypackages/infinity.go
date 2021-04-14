package mypackages

import (
	"fmt"
	"math/rand"
)

func My_rnd_infity(){
	degree := 0
	for {
		fmt.Println(degree)
		degree ++
		if degree >= 360 {
			degree = 0
			if rand.Intn(2) == 0 {
				break
			}
		}
	}
}

package mypackages

import (
	"fmt"
	"math/rand"
)

func My_random_dates() {

	for count := 10; count > 0; count-- {
		month := rand.Intn(12) + 1
		daysInMonth := 31
		year := rand.Intn(1020) + 1000 + 1
		var leap = year%400 == 0 || (year%4 == 0 && year%100 != 0)
		switch month {
		case 2:
			daysInMonth = 29
			if leap {
				fmt.Printf("闰年 %v-%v-%v\n", year, month, daysInMonth-1)
			} else {
				fmt.Printf("平年 %v-%v-%v\n", year, month, daysInMonth)
			}
		case 4, 6, 9, 11:
			daysInMonth = 30
			break
		}
		fmt.Printf("平年 %v-%v-%v\n", year, month, daysInMonth)
	}
}

package mypackages

import "fmt"

func My_leap() {
	//判断闰年
	var year = 2021
	var leap = year%400 == 0 || (year%4 == 0 && year%100 != 0)

	if leap {
		fmt.Println(leap, year, "闰年")
	} else {
		fmt.Println(leap, year, "平年")
	}
}

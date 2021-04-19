package main

import "fmt"


//数字类型转换
func numberConvert(){
	a := 41
	b := float64(a)
	fmt.Printf("%d,%[1]T\n%f,%[2]T\n",a,b)
}
func main() {
	fmt.Println(`
第十章，类型转换
学习目标：数值/字符串/布尔值之间的类型转换
	`)
	numberConvert()
}
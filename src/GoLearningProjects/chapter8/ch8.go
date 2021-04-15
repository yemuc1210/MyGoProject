package main

import (
	"fmt"
	"math/big"
	"reflect"
)

func alpha_distance() {
	//求光飞行至半人马alpha星的时间/天
	var distance int64 = 41.3e12
	const speed = 299792 //km/s
	const secondPerDay = 60 * 60 * 24

	days := distance / speed / secondPerDay
	fmt.Printf("半人马座阿尔法星，光飞行需 %v 天\n", days)
}

func andromeda_distance() {
	/*
		进行big包的测试
		三种类型：1. big.Int  存储大整数
		2. big.Float   任意精度的浮点数
		3. big.Rat    诸如1/3的分数

		用起来麻烦，速度慢，但是精度大
	*/
	//NewInt函数
	speed := big.NewInt(299792) //参数为int64，超过取值上限则失败
	//s1 := big.NewInt(10000000000000000000)   //超出上限，失败
	secondPerDay := big.NewInt(86400)
	//存储24艾  仙女座ℹ
	distance := new(big.Int)
	distance.SetString("240000000000000000", 10) //distance对象使用字符串存储数据，第二参数为进制

	seconds := new(big.Int)
	seconds.Div(distance, speed)

	days := new(big.Int)
	days.Div(seconds, secondPerDay)
	fmt.Printf("仙女座星系，光飞行需 %v 天\n", days)

}

func canisMajorDwarf_distance(){
	//大犬座矮星系
	//已知距离太阳之间的距离是236 000 000 000 000 000 km，转化为光年
	lightSpeed := big.NewInt(299792)   //km/s
	//secondPerYear := 60*60*24*365     //365 days   int类型，不等同于big.int
	var c int64 = 60*60*24*365
	secondPerYear := big.NewInt(c)

	distance := new(big.Int)
	distance.SetString("236000000000000000",10)

	seconds := new(big.Int)
	seconds.Div(distance,lightSpeed)

	lightyears := new(big.Int)
	lightyears.Div(seconds, secondPerYear)

	fmt.Println("大犬座矮星系，光年",lightyears)
}
func main() {
	fmt.Println("chapter 8 大数")
	var bignum int64 = 41.3e12
	fmt.Println(bignum)

	alpha_distance()

	//更大的uint64，存储上限为10e18
	var upperBound uint64 = 10e18
	//var errorNum uint64 = 24e18
	fmt.Println(upperBound)
	// fmt.Println(errorNum)  //overflows uint64

	//若不显式的指定类型，则GO推断其为float64
	f64 := 10e28
	fmt.Println(reflect.TypeOf(f64)) //  float64 利用反射机制输出变量类型

	andromeda_distance()

	canisMajorDwarf_distance()
}

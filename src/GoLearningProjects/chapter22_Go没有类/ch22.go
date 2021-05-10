package main

import (
	"fmt"
	"math"
)

//使用度/分/秒格式表示东西南北半球
type coordinate struct{
		d,m,s float64
		h rune
}	//DMS格式，分和秒表示的是位置，60秒"为一分，没60分'为一度   东经137°26'30.1'
type location struct{
	lat, long float64
}
//绑定方法decimal,DMS格式转换为十进制度格式
func (c coordinate) decimal()float64{
	sign := 1.0
	switch c.h {
	case 'S','W','s','w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)   //转化为度
}
func newLocation(lat,long coordinate) location{
	return location{lat.decimal(), long.decimal()}
}
// func newLocationDMS()


/*
构建一个完整的world类
*/
type world struct{
	radius float64   //半径
}
var mars = world{radius: 3389.5}
func rad(deg float64)float64{
	//角度变弧度
	return deg * math.Pi /180
}
func (w world) distance(p1,p2 location) float64{
	s1,c1 := math.Sincos(rad(p1.lat))
	s2,c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long-p2.long))
	return w.radius*math.Acos(s1*s2+c1*c2*clong)
}

func main() {
	fmt.Println("章节22 Go没有类")
	fmt.Println("_________22.1将方法绑定到结构_________________")
	//通过结构及方法实现面向对象涉及到额相关概念

	//南纬4°35'22.2"    东经137°26'30.12"
	lat := coordinate{4,35,22.2,'S'}
	long := coordinate{137,26,30.12,'E'}

	fmt.Println(lat.decimal(),long.decimal())   //-4.5895 137.4417
	fmt.Println("_________22.2 构造函数_________________")
	
	curiosity := location{lat.decimal(),long.decimal()}    //很像构造函数
	fmt.Printf("%+v\n",curiosity)   //{lat:-4.5895 long:137.4417}
	//创建一个函数充当构造函数
	curiosity = newLocation(lat,long)
	fmt.Printf("%+v\n",curiosity)   //{lat:-4.5895 long:137.4417}
	fmt.Println("_________22.3 类的替代品_________________")
	//使用类型替代类
	
	spirit := location{-14.5684,175.462737}
	opporitunity := location{-1.9462,354.4734}

	dist := mars.distance(spirit,opporitunity)
	fmt.Println(dist)  //9669.74072438503

	




}
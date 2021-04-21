package main

import (
	"fmt"
	"math/rand"
	"time"
)
/*一等函数
函数可以赋值给变量，传递给函数，作为函数返回值
用户函数创建函数
*/

type kelvin float64		//声明变量类型

type sensor func() kelvin	//声明函数类型
var offset kelvin = 5
func fakeSensor() kelvin{
	return kelvin(rand.Intn(151)+150)
}
func realSensor() kelvin{
	return 0  //待实现
}
/*
返回一个函数调整误差，相当于对s函数进行了修改
*/
func calibrate(s sensor, offset kelvin) sensor{
	return func() kelvin{
		offset++  //在匿名函数中修改这个变量，则只会改变，因为包含了变量引用？
		return s() + offset
	}
}
func funcPassToVar(){
/*
假定一个场景，获取传感器的数值。但是有时候计算机脱机，则获取传感器数值就困难
所以创意一个fakeSensor，模拟真实值
*/
	sensor:=fakeSensor     //变量是函数本身， var sensor func() kelvin
	fmt.Println(sensor())   //调用函数需要（），所以sensor调用，sensor()
	sensor=realSensor    //能够重新赋值，是因为具有相同数量和相同类型的形参和返回值    func kelvin
	fmt.Println(sensor())
}

func funcPassToFunc(samples int, sensor func() kelvin){
	fmt.Println("函数funcPassToFunc")
	for i:=0;i<samples;i++{
		k:=sensor()
		fmt.Printf("%v K\n",k)
		time.Sleep(time.Second)
	}
}

//声明 函数类型 后改写
func funcPassToFunc_1(samples int, s sensor){
	fmt.Println("函数funcPassToFunc_1")
	for i:=0;i<samples;i++{
		k := s()   //此时s相当于函数，调用的话还是需要加()
		fmt.Printf("%v K\n",k)
		time.Sleep(time.Second)
	}
}

/*
匿名函数，即没有名字的函数，go中称为函数字面量
因为函数字面量需要保留外部作用域的变量引用，所以函数字面量都是闭包的
*/
var f = func(s sensor){  //匿名函数赋值给变量   关键字是var，定义的是变量
	//实际上和上面，将函数赋值给变量一样
	fmt.Println("匿名函数",s())

}

func main(){
	fmt.Println("章节14，一等函数")
	funcPassToVar()
	funcPassToFunc(5,fakeSensor)//函数作为参数传递
	
	s := sensor(fakeSensor)
	funcPassToFunc_1(5,s)   //重写，使用了函数类型之后，提高了复用性

	//匿名函数
	f(s)   //f存的是函数，实例调用需要像函数一样

	//此外，匿名函数可以写在函数内部
	f1 := func(message string){
		fmt.Println(message)
	}

	f1("匿名函数f1,定义在函数中，感觉像python 中 lamda函数的用法")

	//写完即调用
	func(message string){
		fmt.Println(message)
	}("匿名函数,没有函数名，没有赋值给变量，写完即调用")

	
	s1 := calibrate(realSensor, offset)   //offset是全局变量
	fmt.Println(s1())
	fmt.Println("Offset before=",offset)
	//offset ++
	//在匿名函数中修改了offset，但是主函数中的值没有变
	fmt.Println("offset after=",offset)
	//这是传值，匿名函数值并没有包围变量的引用
	fmt.Println(s1())    //值不变，因为offset形参接受的是实参的副本值而不是引用


	//匿名函数会封闭包围作用域中的变量，并且闭包保留的是变量的引用  外部作用域的变量引用
	//闭包=函数+引用环境
	var k kelvin = 294.0
	s2 := func() kelvin{
		return k
	}
	fmt.Println("测试闭包的变量引用，before,",s2())   //294
	k ++
	fmt.Println("测试闭包的变量引用，after,",s2())   //295


	//s1修改，使用fakeSensor，观察是否每次都会有新的随机数
	s3 := calibrate(fakeSensor, offset)
	fmt.Println("匿名函数传入fakeSensor后，第一次调用：",s3())
	fmt.Println("匿名函数传入fakeSensor后，第二次调用：",s3())
	fmt.Println("匿名函数传入fakeSensor后，第三次调用：",s3())
	/*
	结论，每次都会产生新的随机数，因为包含的是函数的引用，所以每次调用都是新的
	*/

}


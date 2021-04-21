package main

import "fmt"

/*
函数首字符大写，可以被导出并对其他包使用
*/
type celsius float64   //定义一个摄氏温度的类型
type fahrenheit float64   //华氏温度类型
type kelvin float64
func typeTest(tt celsius){
	/*使用新类型可以更好地描述值的属性
	虽然都是10，但是10的含义可以是不同的*/
	
	var t1 celsius = 90.0
	var t2 fahrenheit = 90.0

	// t3 := t1+t2          //报错，类型不匹配
	fmt.Printf("%[1]v,%[1]T\t%[2]v,%[2]T\n",t1,t2)
}

/*
go 提供了方法，但是并没有提供类和对象
方法是和某种类型绑定的函数，允许重名
*/
//这是一个函数定义
func fahrenheitToCelsius(f fahrenheit) celsius{
	return celsius(f-67)
}
//这是一个方法  
//格式： func  (接受者名字 类型)  方法名  返回值
func (f fahrenheit) celsius() celsius{
	return celsius((f-32)*5.0/9/0)
}
func (k kelvin) celsius() celsius{
	return celsius( k -273.15)
}
func main(){
	typeTest(12)   //可以接受一个字面量或者是无类型常量
	var tmp fahrenheit = 45
	// typeTest(tmp) //报错，类型不匹配

	//调用方法
	a := tmp.celsius()    
	fmt.Println(a)

	var k kelvin = 90
	b := k.celsius()    //同一个包内允许方法重名
	fmt.Println(b)


}
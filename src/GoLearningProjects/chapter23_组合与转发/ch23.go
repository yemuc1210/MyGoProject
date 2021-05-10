package main

import "fmt"
//结构中的结构
type celsius float64
type temperature struct{
	high,low celsius
}
func (t temperature) average() celsius{
	return (t.high + t.low) / 2
}	

type location struct{
	lat,long float64
}
func (l location) show(){
	fmt.Println(l)
}
type report struct{
	sol int
	temperature temperature
	location location
}
/*
这种转发方式比较繁琐，需要一层一层的编写转发逻辑
*/
func (r report) average() celsius{
	return r.temperature.average()
}
/*使用嵌入实现自动转发*/
type reportEM struct{
	sol int
	temperature
	location     //嵌入，只给定类型，不给字段名
}
/*
命名冲突
*/
func (l location) days(l2 location) int{
	return 5
}
func (r report) days (r2 report) int{
	return 6
}
func main() {
	fmt.Println("章节23 组合与转发")
	/*
	面向对象的组合概念
	嵌入的特殊语言特性实现方法转发
	*/
	fmt.Println("___________23.1 合并结构_______________")
	// type report struct{  //未合并的单一结构   天气报告
	// 	sol int    
	// 	high,low float64
	// 	lat,long float64
	// } //很麻烦，若添加更多的细节信息时，就很不灵活
	
	bradbury := location{-4.5895, 137.417}   //
	t := temperature{high:-1.0,low:-78.0}
	report1 := report{sol: 15, temperature: t,location: bradbury}
	fmt.Printf("%+v\n",report1) //{sol:15 temperature:{high:-1 low:-78} location:{lat:-4.5895 long:137.417}}
	
	//访问温度
	fmt.Printf("average %v\n",report1.temperature.average())
	//另一种方式就是转发，转发至实际的方法，实现封装
	fmt.Printf("average %v\n",report1.average())

	fmt.Println("___________23.2 实现自定转发_______________")
	// type reportEM struct{
	// 	sol int
	// 	temperature
	// 	location     //嵌入，只给定类型，不给字段名
	// }
	reportEM := reportEM{
		sol: 15,
		location: bradbury,
		temperature: t,
	}   
	/*
	会生成与类型同名字段名，这样转发时可以直接根据类型名得到类型名，然后调用对应的类型方法
	实现自动转发
	*/
	fmt.Println(reportEM)
	fmt.Printf("location:")
	reportEM.show() //location:{-4.5895 137.417}
	fmt.Println("___________23.3 命名冲突_______________")
	r2 := reportEM
	// fmt.Println(reportEM.days(r2))   //报错    days（）的形参期望是location类型
	fmt.Println(r2)
	l2 := bradbury
	fmt.Println(bradbury.days(l2))    //正常， 这是命名冲突现象
}
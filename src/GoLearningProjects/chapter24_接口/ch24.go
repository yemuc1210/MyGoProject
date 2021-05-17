package main

import (
	"fmt"
	"strings"
	"time"
)


type talker interface{
	talk() string
 }  //任何类型，只要能够满足接口的要求，就能成为变量t的值。只要它声明talk()方法，不接受实参，返回string

 type s struct{
	a int
	s string
 }
 type intdemo int
 func (i intdemo) talk() string{
	 return strings.Repeat("intdemo ",int(3))
 }

 func (s0 s) talk() string{
	 return s0.s
 }

 func shout(t talker){
	 louder := strings.ToUpper(t.talk())
	 fmt.Println(louder)
 }
 //___________________________________24.2
 
type stardater interface {
	YearDay() int
	Hour() int 
}  //time.Time类型满足这个接口的定义
/*
这种做法在java中无法实现，因为无法让java.time通过显式声明实现接口，它是库函数啊 
*/
//t变量可以指向任意类型变量，只要其类型实现了YearDay()/Hour()
func starDate(t stardater) float64{
	doy := float64(t.YearDay())
	h := float64(t.Hour()) /24.0
	return 1000 + doy + h
}
type sol int
func (s sol) YearDay() int {
	return int(s % 668)
}
func (s sol) Hour() int {
	return 0
}   //现在sol类型也实现了stardate接口
func test24_2(){
	var day stardater = time.Date(2012,8,6,5,17,0,0,time.UTC)
	s := starDate(day)
	fmt.Println(s)
	s1 := sol(1422)
	s2 := starDate(s1)
	fmt.Println(s2)
}


//____________________________________24.3 满足接口
type location struct {
	lat,long float64
}

//实现了Stringer接口，需要实现String函数
func (l location) String() string{
	//类型location的打印函数
	return fmt.Sprintf("(lat:%v,long:%v)\n",l.lat,l.long)
}
func test24_3(){
	curiosity := location{-4.5895,137.4417}
	fmt.Println(curiosity)    //好家伙，直接使用变量名
}
func main() {
	fmt.Println("章节24 接口")

	fmt.Println("______________24.1 接口类型______________________")
	fmt.Println(`
	Go Writer接口，允许通过接口将任意内容写入任意地方，cool
	代码可以通过接口表达抽象的概念，关注物件的行为而不是物件构成

	类型通过方法表达行为，而接口通过列举类型必须满足的一组方法来进行声明
	`)
	// t talker
	t := s{2,"asdf"}
	s := t.talk()
	fmt.Println(t.a)   //输出2 可以当做结构体变量使用
	fmt.Println("s talk:",s)

	//接口的多态

	t1 := intdemo(3)
	fmt.Println(t1.talk())
	/*
	变量t是接口的实例，而t可以指向结构体变量，也可以指向int型变量
	理论上，t可以是任何类型，只有该类型有talk()方法
	此外，结构体s，变量intdemo都不需要显式声明实现一个接口，与Java不同，方便
	为方便复用，将接口声明为类型，加-er后缀
	*/

	shout(t)    //字面来看，t是s结构体变量，但它也是结构体变量
	shout(t1)
	/*
	接口在修改和扩展代码的时候能够淋漓尽致地发挥灵活性，只要声明一个带talk()方法的新类型,shout函数将自动适应

	无论实现何种变化或者新增何种功能，只依赖接口的代码都不用做任何修改
	*/
	
	/*
	嵌入满足接口
	startship嵌入了intdemo，而intdemo实现了talk()方法
	所以starship能够转发intdemo的talk()方法
	当starship需要说话时，intdemo服其劳

	
	*/
	type starship struct {
		intdemo
	}

	s1 := starship{intdemo(3)}
	fmt.Println(s1.talk())
	shout(s1)

	fmt.Println("______________24.2 探索接口______________________")
	/*
	go语言允许在实现代码的过程中随时创建新的接口，任何代码都可以创建接口，包括已经存在的代码
	*/
	test24_2()
	fmt.Println("______________24.3 满足接口______________________")
	/*
	Go标准库导出了很多只有单个方法的接口，人们可以在自己的方法中实现他们
	fmt包声明了
	type Stringer interface{
		String() string
	}
	只要实现了这个方法，就可以通过Println() Sprintf()输出
	*/
	test24_3()
}
package mypackages

import "fmt"

/**
声明变量的语法
var identifier type
v_name := value     不能用于	全局变量的声明和赋值

const identifier [type] = value

type可缺省，编译器根据数值判断
*/
func My_const_var_test() {
	const speed = 300000    //km/s
	var distance = 56000000 //km
	fmt.Printf("已知速度为%v km/s,距离为%v km， 则时间为: %v s \n", speed, distance, distance/speed)
	//常量可以使用len()/cap()/Sizeof()/unsafe()等内置函数计算
	const str = "abcdefghijklmn"
	fmt.Println(str, len(str)) //14
	const (
		a1 = iota //0
		b1        //1
		c1        //2
		d1 = "ha" //独立值，iota += 1
		e1        //"ha"   iota += 1
		f1 = 100  //iota +=1
		g1        //100  iota +=1
		h1 = iota //7,恢复计数
		i         //8
	) //iota是一个特殊的常量
	fmt.Println(a1, b1, c1, i) //0 1 2 8
	//下面申明一个const 中介iota的技术
	const i1 int = 100
	fmt.Println(i, i1) //8 100
	//声明一组变量
	var (
		a = 1
		b = 2
	)
	var c, d = 3, 4
	fmt.Println(a, b, c, d)
	//缺省声明
	e := 5
	fmt.Println(e)

	//增量操作
	a++
	fmt.Println(a) //输出2
	a += 1
	fmt.Println(a)
	//  ++a   不支持，报错
	b = b + 1

	// bool 类型
	var bl bool
	fmt.Println(bl) //false

	//数字类型
	//var int1 uint8 = 256    //无符号8位整型，显然这个数溢出
	var int1 int = 256
	fmt.Println(int1) //variable.go:31:6: constant 256 overflows uint8
	var float1 float32 = 2.3456
	fmt.Println(float1)
	var complex1 complex64 = 2 + 1i
	fmt.Println(complex1) //(2+1i)
	var bt1 byte = '1'
	fmt.Println(bt1)        //输出49
	fmt.Printf("%s\n", bt1) //%!s(uint8=49)
	//内存测试
	var int2 int = int1
	fmt.Println(int2) //256
	int2++
	fmt.Println(int1, int2)   //256 257
	fmt.Println(&int1, &int2) //0xc00000a0d0 0xc00000a0f0     地址

}

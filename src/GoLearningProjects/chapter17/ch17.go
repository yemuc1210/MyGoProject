package main

import (
	"fmt"
	"sort"
	"strings"
)


func testArr(planets [8]string){
	for i := range planets{
		planets[i] = strings.TrimSpace(planets[i])
	}
}

func testSlice(planets []string){
	for i := range planets{
		planets[i] = strings.TrimSpace(planets[i])
	}
}
func main() {
	fmt.Println("章节17 切片：指向数组的窗口")
	//定义太阳系数组
	planets := [...]string{
		" Mercury ",
		" Venus ",
		" Earth ",
		" Mars ",
		" Jupiter ",
		" Saturn ",
		" Uranus ",
		" Neptune ",
	}
	fmt.Println(planets)
	fmt.Printf("%T\n",planets)

	//已知行星可以分为三类：前四个陆地行星 terrestrial 后两个：气态巨行星 gasGiants   z最后两个 冰巨行星 iceGiants
	//使用切片进行分类，得到三个切片  
	terrestrial := planets[0:4]      //这是一个切片slice   不过操作起来像数组
	fmt.Println("terrestrial:",terrestrial)

	gasGiants := planets[4:6]
	fmt.Println("gasGiants:",gasGiants)

	iceGiants := planets[6:] //这是默认切片的手法，等价于  planets[6:8]
	fmt.Println("iceGiants:",iceGiants)

	//下面测试下切片的一些操作，如遍历
	for index,e := range terrestrial{
		fmt.Println(index,e)
	}   //像数组一样

	//创建切片的切片
	giants := planets[4:]
	gas := giants[0:2]    //这里的切片下标就是  以需要切片为基准
	ice := giants[2:]
	fmt.Println(giants,gas,ice)

	//切片的实质是数组的一个视图，对切片的操作会影响到数组
	gas[0] = "gasGiant-No.1"
	fmt.Println(giants,gasGiants,planets) 
	//[gasGiant No.1 Saturn Uranus Neptune] [gasGiant No.1 Saturn] [Mercury Venus Earth Mars gasGiant No.1 Saturn Uranus Neptune]
	
	//默认切片   切片索引不能是负数
	planets_1 := planets[:]
	terrestrial_1 := planets_1[:4]
	fmt.Println(planets_1)
	fmt.Println(terrestrial_1)


	//对字符串的切片
	str := "Neptune"
	tune := str[3:]
	fmt.Println(tune)
	//切片字符串将会创建另一个字符串，所以字符串的切片不是一个视图
	str = "Poseidon"
	fmt.Println(str)   //Poseidon
	fmt.Println(tune)  //tune
	//切分字符串时，索引代表的是字节号码的下标，而不是符文下标
	str = "中国"    //汉字是三？？
	fmt.Println(str[:3])   //打印出：中

	//Go语言的函数倾向于使用切片作为输入，而不是整个数组，即便是使用数组，也使用[:]切片
	//直接声明切片
	slice := []string{"ceres","Pluto","Haumea","Makemake","Eris"}
	fmt.Println(slice)

	//数组作为参数传递，传的是值，函数内操作的是数组的一个副本，不会对数组造成影响
	//但是传入切片，就相当于传递引用
	fmt.Println(planets)
	testArr(planets)
	fmt.Println(planets)
	testSlice(planets[:])   //传递一个切片
	fmt.Println(planets)
	/*
	[ Mercury   Venus   Earth   Mars  gasGiant-No.1  Saturn   Uranus   Neptune ]     
	[ Mercury   Venus   Earth   Mars  gasGiant-No.1  Saturn   Uranus   Neptune ]
	[Mercury Venus Earth Mars gasGiant-No.1 Saturn Uranus Neptune]      //切片传递，对原数组进行操作了，nb
	*/ 
	/*
	此外，数组传递时，需要保证形参和实参的大小长度是一样，这个比较困难
	*/

	//————————————————————————17.4  带有方法的切片
	fmt.Println("___________17.4____________")
	giants_1 := strings.Join(giants,"-")  //切片连接，通过Join函数，将切片中多个元素连接，创建出字符串
	fmt.Printf("%[1]T\t%[1]v\n",giants_1)  // output:  string  gasGiant-No.1-Saturn-Uranus-Neptune
	//在Go语言中声明类型，并绑定相应的方法，很通用
	/*
	type StringSlice []string    sort库中声明
	对应的绑定方法 func(p StrinSlice) Sort()
	*/
	planets_2 := planets_1  //负责一个副本
	sort.StringSlice(planets_1).Sort()
	fmt.Println(planets_1)  //[Earth Mars Mercury Neptune Saturn Uranus Venus gasGiant-No.1]

	sort.Strings(planets_2)
	fmt.Println(planets_2)


}
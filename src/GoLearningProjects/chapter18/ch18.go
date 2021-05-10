package main

import "fmt"

//dump函数会打印出切片的长度、容量和内容
func dump(label string, slice []string){
	fmt.Printf("%v:length %v, capacity %v %v\n",label,len(slice),cap(slice),slice)
}
func terraform(prefix string, worlds ...string) []string{
	newWorlds := make([]string, len(worlds))    //新切片

	for i := range worlds{
		newWorlds[i] = prefix +" " + worlds[i]    //覆盖表
	}
	return newWorlds
}

func exp_cap(){
	s := []string{}
	lastCap := cap(s)

	for i:=0;i<20;i++{
		s = append(s, "add one")
		if cap(s) != lastCap{
			fmt.Println("full:",cap(s))
			lastCap = cap(s)
		}

	}

}
func main() {
	fmt.Println(
		"章节18  更大的切片")
	fmt.Println("_________________18.1 append函数______________________")
	//在实际中，需要可变长度的数组，通过切片和内置的append实现
	dwarfs := []string{
		"ceres","Pluto","Haumea","Makemake","Eris",
	}      //这是一个切片

	dwarfs = append(dwarfs, "NewOrcus")

	fmt.Println(dwarfs)     //实现了可变长度
	fmt.Println("_________________18.2 长度和容量______________________")
	//切片长度是数组可见的部分长度，切片底层的数组若比切片大，则称切片还有容量可增长
	dump("dwarfs",dwarfs)  //dwarfs:length 6, capacity 10 [ceres Pluto Haumea Makemake Eris NewOrcus]
	/*
	因为上面使用了append函数，但实际上原数组没有位置添加NewOrcus，所以元素是先复制一个新数组，再添加
	新数组大小是原数组的两倍
	*/
	dump("dwarfs[1:2]",dwarfs[1:2])   //容量比上面的小1，因为切片从1开始
	dump("dwarfs[3:]",dwarfs[3:])    //容量为7
	fmt.Println("_________________18.3 详解append函数______________________")
	fmt.Println("因为上面使用了append函数，但实际上原数组没有位置添加NewOrcus，所以元素是先复制一个新数组，再添加,新数组大小是原数组的两倍")

	fmt.Println("_________________18.4 三索引切分操作______________________")
	planets := []string{
		"a","b","c","d","e","f",
	}
	slice := planets[0:4:4]
	slice1 := append(slice,"add")
	dump("slice",slice)
	dump("slice1",slice1)
	dump("planets",planets)
/**
slice:length 4, capacity 4 [a b c d]
slice1:length 5, capacity 8 [a b c d add]
planets:length 6, capacity 6 [a b c d e f]
说明，使用三索引切片，没有对原数组进行扩充，而是为切片新申请了一个数组
*/
	fmt.Println("_________________18.5 使用make对切片预分配______________________")
	/*
	当切片容量不足时，需要对底层数组进行扩容：新建、复制；这是一个耗时的操作
	可以通过make函数对切片进行预分配，避免额外的内存分配和数组复制
	*/
	slice_make := make([]string,0,10)   //长度0，容量10
	slice_make_1 := make([]string,2,10)
	dump("slice_make",slice_make)
	dump("slice_make_1",slice_make_1)
	/*
	slice_make:length 0, capacity 10 []
	slice_make_1:length 2, capacity 10 [ ]    显然，初值是两个空字符（字符类型的零值）
	*/
	slice_make = append(slice_make,"first")
	dump("slice_make",slice_make)	//slice_make:length 1, capacity 10 [first]
	slice_make_1 = append(slice_make_1, "first")
	dump("slice_make_1",slice_make_1)  //slice_make_1:length 3, capacity 10 [  first]
	//这说明零值不等于空值，零值是占空间的，所以从第三个地方插入

	fmt.Println("_________________18.6 声明可变参数函数______________________")
	/*
	像Printf和append函数，可以接受可变数量的实参，这种叫做可变参数函数
	方法：在最后一个形参后面加省略号...   编译器捕捉多个参数为切片
	*/

	twoworlds := terraform("new", "Mars", "Earth")
	dump("twoWorlds",twoworlds)  //twoWorlds:length 2, capacity 2 [new Mars new Earth]
	
	//使用省略号可以展开切片中的多个元素
	newPlanets := terraform("new", planets...)
	dump("newPlanets",newPlanets)  //newPlanets:length 6, capacity 6 [new a new b new c new d new e new f]


	exp_cap()


}	
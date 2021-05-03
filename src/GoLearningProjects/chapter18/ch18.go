package main

import "fmt"

//dump函数会打印出切片的长度、容量和内容
func dump(label string, slice []string){
	fmt.Printf("%v:length %v, capacity %v %v\n",label,len(slice),cap(slice),slice)
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



}
package main

import (
	"fmt"
	"go-ethereum/common/math"
)

/*
多语言文本章节，字符串
*/

func stringTest() {
	var black string
	fmt.Println("空字符串：", black)

	//转义字符串
	fmt.Println(`换行符\n`)
	fmt.Println(`
跨越多行，lin1
lin2
	带有一个制表符lin3
	`)


	//字符串用v%表示,%T显示类型，[1]位置
	fmt.Printf("%v is a %[1]T, second input %v is a %[2]T\n\n","a string",12)

	//单引号包围的叫 原始字符串字面量， 双引号的叫字符串字面量，区别主要在于里面的转义字符

}

func unicode(){
	//字符和代码点
	fmt.Println(65)
	//符文类型rune,用于表示单个统一码代码点，int32的别名
	//起别名type,如type byte = uint8

	var a rune = 65
	fmt.Printf("%c\n",a)   //转成字符输出
	var b rune = 128515
	fmt.Printf("%c\n",b)

	//为了避免记忆代码点，go有字符字面量句法
	grade := 'A'   //单引号包围，取得代码点
	fmt.Printf("%[1]c ,代码点 %[1]v, 类型%[1]T\n",grade)

}

func stringOp(){
	//字符串的操作
	//初始化
	s := "asdfg"
	//赋值
	s = "qwert"
	//但是，无法对字符串本身进行修改
	c := s[0]
	fmt.Printf("字符串 %s, 第一个字符为%c\n",s,c)
	// s[0] = 'a'  //报错，无法赋值，所以字符串是个常量
	s = "shalom"
	for i:=0;i<len(s);i++{   //遍历字符串
		fmt.Printf("%c,%[1]T\n",s[i])
	}


}

func kaixa(){
	//使用凯撒加密法加密字符串
	//对每个字符进行移位
	//step := rand.Uint32()+1  //移位1~10
	var step uint8 = 3   //凯撒加密 step = 3
	source := "asdfggkwqqqeq"
	// var dest [len(source)] byte  //需要什么元素个数，常量值
	var dest [20] uint8
	var c byte = 'a'
	fmt.Printf("\n%c,%[1]T\n",c+3)
	fmt.Println("source is ",source)
	for i:=0;i<len(source);i++{
		// source[i] += step   //报错，因为字符串不可修改
		re := (source[i] + step)
		if re > 'z'{
			dest[i] = re - 26
		} else {
			dest[i] = re
		}
		fmt.Printf("%c",dest[i])
	}
	fmt.Printf("\n\n\n\n")
}

func rotate13(){
	//凯撒加密的变体，变量为13
	var step uint8 = 13
	source := "uv 2332  vagareangvbany fcnpr fgngvba"
	//length := len(source)
	var result[math.MaxInt8] byte   //定义一个超大的数组

	for i:=0;i<len(source);i++{
		//只加密英文字符
		c := source[i]
		if c>='a' && c<='z'{
			c += step
			if c>'z'{
				c -= 26
			}
		}
		result[i] = c
	}
	fmt.Println(source)
	fmt.Printf("%s\n\n\n\n",result)

	//以上，只能处理英文字符
	//utf8包提供RuneCountInString函数，能够以符文而不字节为单位返回字符串长度
	//DecodeRuneInStrin解码字符串首个字符并返回解码后符文所占字节数
}

func exp(){
	//章节实验
	source := "L fdph,L vdz,L frqtxhuhg"
	var step uint8 = 3
	for i:=0;i<len(source);i++{
		c := source[i]
		if (c>='a'&&c<='z') || (c>='A'&&c<='Z') {
			if (c>='a'&&c<='c'){
				c = 'z' - (c-'c')
			}else if (c>='A'&&c<='C'){
				c = 'Z' - (c-'C')
			}else{
				c -=step
			}
			// c -= step
			fmt.Printf("%c",c)
		}else{
			fmt.Printf("%c",c)
		}
	}
	fmt.Printf("\n\n")
	fmt.Println("使用range遍历字符串")
	for index,c := range source{
		fmt.Printf("%d %c\n",index,c)
	}
}

func main() {
	// Hello()
	stringTest()
	unicode()
	stringOp()
	kaixa()
	rotate13()
	exp()
}

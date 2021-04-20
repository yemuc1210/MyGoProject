package main

import (
	"fmt"
	"math"
	"strconv"
)

//数字类型转换
func numberConvert(){
	a := 41
	b := float64(a)   //go  只有显式转换
	fmt.Printf("%d,%[1]T\n%f,%[2]T\n",a,b)

	//类型转换的危险
	var bh float64 = 32767
	var h = int16(bh)
	fmt.Println(bh,h)
	bh += 1
	h = int16(bh)
	fmt.Println("+1后",bh,h)       //出现上溢

	//math包中存储最小、最大常量
	fmt.Println("math.MaxInt16=",math.MaxInt16,"math.MinInt16=",math.MinInt16)

	//数字转换成字符串
	cnt := 10000121212
	fmt.Println("cnt:"+strconv.Itoa(cnt))     //数字转化成ASCII字符
	s := fmt.Sprintf("cnt:%d\n",cnt)    //SPrintf返回格式化字符串
	fmt.Println("s:",s)
	//字符串转数字 Atoi ACSII to Integer
	count,err := strconv.Atoi("233333")
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println("233333",count)
	fmt.Println(strconv.FormatInt(int64(cnt),2))   //进制转换
	var bt  = []byte{1,2,3,4}
	fmt.Println(strconv.AppendInt(bt,0,10))  //[1 2 3 4 48]    字符0的ASCII

}

func booleanConvert(){
	fmt.Println("bool convert.......")
	var launch bool
	launchText := fmt.Sprintf("%v",launch)
	fmt.Printf("%v,%[1]T\n",launchText)

	//字符串转换为布尔值
	ans := "no"
	launch = (ans != "no")
	fmt.Printf("%v,%[1]T\n",launch)
	//bool值没有对应的字符值和数值，所以string(false)/int(false)报错

}

func exp(){
	trueText:= "true"
	// falseText,yesText,noText,oneText,zeroText := "false","yes","no","1","0"
	fmt.Println("exp..........")
	T := (trueText == "true")
	// F := (falseText != "false")
	// yes := (yesText == "yes")
	// no := (noText != "no")
	// one := (oneText == "1")
	// zero := (zeroText != "0")
	fmt.Printf("%v,%[1]T\n",T)
}
func main() {
	fmt.Println(`
第十章，类型转换
学习目标：数值/字符串/布尔值之间的类型转换
	`)
	numberConvert()
	booleanConvert()
	exp()
}
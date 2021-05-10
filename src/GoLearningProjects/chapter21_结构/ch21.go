/*
第五单元开始：状态与行为
Go中，值表示状态
函数和方法表示行为


Go实现面向对象
*/

package main

import (
	"encoding/json"
	"fmt"
	"os"
)
/**
以JSON格式打印出3个着陆点位置，需要包含名称
*/
func Landing(names []string,lats []float64,longs []float64){
	type newLocation struct{
		Lat float64 `json:"latitude"`
		Long float64 `json:"longitude"`
		Name string `json:"name"`
	}
	locations := []newLocation{
		{Name:"1",Lat:lats[0],Long:longs[0]},
		{Name:"2",Lat:lats[1],Long:longs[1]},
		{Name:"3",Lat:lats[2],Long:longs[2]},
	}
	for index := range locations{
		types,err := json.Marshal(locations[index])
		if err != nil{
			fmt.Println(err)
			os.Exit(1)
		}
		jsIndet,_ := json.MarshalIndent(locations[index],"","\t")
		fmt.Println(string(types),"\njsIndent:\n",string(jsIndet))
	}

}

/*
结构  C语言中的struct
*/
func main(){
	fmt.Println("章节21 结构")
	fmt.Println("________21.1 声明结构__________")
	var curiosity struct {	// var 变量名 变量类型weidu
		lat float64			//纬度
		long float64
	}   //经纬度 结构

	curiosity.lat = -4.595
	curiosity.long = 137.4417
	//输出
	fmt.Println(curiosity)   //{-4.595 137.4417}

	fmt.Println("________21.2 通过类型复用结构__________")
	/*
	应用场景。多个结构中可能使用同一个字段
	*/
	type location struct{
		lat float64
		long float64
	}
	var spirit location
	spirit.lat = -14.5684
	spirit.long = 175.472636
	fmt.Println(spirit)  //{-14.5684 175.472636}

	fmt.Println("________21.3 通过复合字面量初始化结构__________")
	type location3D struct{
		lat float64
		long float64
		attitude float64
	}
	opportunity := location3D{1,2,3}  //若不给出对应的字段，则需要给出所有的值
	fmt.Println(opportunity)   //{0 0 0}
	spirit1 := location3D{lat:-1.9462, long:354.4734}  //可以单独给某些字段赋值，没有赋值的默认零值
	fmt.Println(spirit1)  //{-1.9462 354.4734 0}
	fmt.Printf("%+v\n",spirit1)   //{lat:-1.9462 long:354.4734 attitude:0}  打印结构
	fmt.Println("________21.4 结构被复制__________")
	//复制 时创建的是副本
	bardbury := location{-4.5895,137.417}
	curiosity1 := bardbury

	curiosity1.long += 0.0106   //向东移动
	fmt.Println("bardbury:",bardbury)
	fmt.Println("curiosity1:",curiosity1)
	/*
	bardbury: {-4.5895 137.417}
	curiosity1: {-4.5895 137.4276}
	所以复制是复制值，并不是指向同一个底层数据
	*/
	fmt.Println("________21.5 由结构组成的切片__________")
	//[] struct 表示由结构组成的切片，切片的每个值都是一个切片
	lats := []float64{-4.5895,-14.5684,-1.9462}
	longs := []float64{137.4417,175.472636,354.4734}
	names := []string{"1","2","3"}
	//上面表示三组坐标，很麻烦

	type newLocation struct{
		name string
		lat float64
		long float64
	}
	
	locations := []newLocation{
		{name:"1",lat:lats[0],long:longs[0]},
		{name:"2",lat:lats[1],long:longs[1]},
		{name:"3",lat:lats[2],long:longs[2]},
	}
	for index := range locations{
		fmt.Printf("%+v\n",locations[index])
		/*
		{name:1 lat:-4.5895 long:137.4417}
		{name:2 lat:-14.5684 long:175.472636}
		{name:3 lat:-1.9462 long:354.4734}
		*/
	}
	fmt.Println("________21.6 将结构编码成JSON__________")
	/*
	JSON: JavaScript 对象表示法
	JSON和结构体
	*/
	type jsonLocation struct{
		Lat,Long float64   //字段首字母必须大写
	}

	curiosity2 := jsonLocation{lats[0],longs[0]}
	bytes, err := json.Marshal(curiosity2)  
	/**
	Marshal只会对结构中被导出的字段编码，
	首字母大写的会被导出，和函数的定义一样
	*/

	if err != nil{    //打印错误信息并退出    做成函数比较好
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(bytes))   //{"Lat":-4.5895,"Long":137.4417}
	type jsonLocation1 struct{
		lat,Long float64   //字段首字母必须大写   这里lat不大写
	}
	curiosity3 := jsonLocation1{lats[0],longs[0]}
	bytes, err = json.Marshal(curiosity3)
	if err != nil{    //打印错误信息并退出    做成函数比较好
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(bytes))   //{"Long":137.4417}    只有首字母大写的Long能够被导出

	fmt.Println("________21.7 将结构标签定制JSON__________")
	/*
	Go中的json包对字段的命名要求严格：
	（1）首字母必须大写
	（2）包含多个单词的字段必须使用类似 CamelCase的驼峰命名规则
	若想使用snake_case这种命名规则，就需要定制了，因为Python中使用snake_case，交互时可能用到
	定制方法：
	对 结构中的字段打标签tag，使得json包按照我们的意愿修改字段名称
	标签是与字段关联的   字符串 ，所以可以用双引号，不过麻烦，因为key:"value"中包含双引号，需要转义
	*/
	type tagLocation struct{
		Lat float64 `json:"latitude"`
		Long float64 `json:"longitude"`
	}

	curiosity4 := tagLocation{lats[0],longs[0]}

	bytes,err = json.Marshal(curiosity4)
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	//实现定制json，打标签
	fmt.Println(string(bytes))    //{"latitude":-4.5895,"longitude":137.4417}



	Landing(names,lats,longs)
	/*
	{"latitude":-1.9462,"longitude":354.4734,"name":"3"}
	jsIndent:
 	{
        "latitude": -1.9462,
        "longitude": 354.4734,
        "name": "3"
	}
	*/

 


}
package main

import (
	"fmt"
	"math"
	"strings"
)

func words() {
	text := `As far as eye could reach he saw nothing but t he stems of  
	the  great plants about him receding in the violet shade, 
	and far ovrehead the multiple transparency  of  huge  leaves filtering  the sunshine  to  the  solemn  
	splendour  of twilight in which he  walked.  
	Whenever  he  felt able  he ran again;
	the  ground  continued soft and springy, covered with the same resilient weed which was the first thing 
	his hands had touched in Malacandra. 
	Once or twice a small red creature scuttled across his  path, 
	but otherwise there seemed to  be no life stirring in the wood; 
	nothing to fear-except the fact of wandering unprovisioned and alone in a forest of unknown 
	vegetation  thousands or millions of miles beyond the reach or knowledge of man.`
	//统计各个单词出现的频率
	words := strings.Fields(strings.ToLower(text))   //
	frequency := make(map[string]int,len(words))     //统计嘛，用映射
	for _,val := range words{
		val = strings.Trim(val, `.,"- `)
		frequency[val]++
	}
	for word,count := range frequency{
		if(count > 1){
			fmt.Printf("%v:%d\n",word,count)
		}
	}


}

func main() {
	fmt.Println("章节19   无所不能的映射")
	/*
	映射，是一种收集器，类似于google map，可以将键映射到值，键可以是任何类型的
	python 字典；
	映射适用于在运行期间就确定了键的非结构化数据

	*/
	fmt.Println("___________________19.1 声明映射___________")
	tmp := map[string]int{
		"A":1,
		"B":2, 
	}    //声明键类型string,值类型int
	fmt.Println("tmp[\"A\"]:",tmp["A"])
	fmt.Println("tmp[\"C\"]:",tmp["C"])  //tmp["C"]: 0   返回int类型的零值，但这没法和实际值0区分

	//逗号，ok句法   ok可用其他变量名替换，实际上就是个Boolean值罢了
	// and,ok := tmp["C"]
	if ans,ok :=  tmp["C"];ok {
		fmt.Println(ans)
	}else{
		fmt.Println("not ok")
	}

	fmt.Println("___________________19.2 映射不会被复制___________")
	//数组 赋值或作为参数传递时，会创建副本   映射不会
	planets := map[string]string{
		"Earth":"Sector ZZ9",
		"Mars":"Sector ZZ9",
	}

	planetsMark2 := planets    //赋值传递

	planets["Earth"] = "whoops"   //修改原始数据
	//指向相同的底层数据
	fmt.Println(planets)   //map[Earth:whoops Mars:Sector ZZ9]
	fmt.Println(planetsMark2)  //map[Earth:whoops Mars:Sector ZZ9]

	delete(planets,"Earth")     //内置函数   映射
	fmt.Println(planetsMark2)  //map[Mars:Sector ZZ9]

	fmt.Println("___________________19.3 使用make函数对映射进行预分配___________")
	temperature := make(map[float64]int,6)   //预分配6割空间
	fmt.Println(temperature)
	fmt.Println("___________________19.4 使用映射计数___________")
	//比如需要统计某些非int型数字的出现频率，用数组是不可行，因为数组的键就是索引（是int）
	
	temperature1 := []float64{
		-28.0,32.0,-31.0,-29.0,-23.0,-29.0,-28.0,-33.0,
	}
	frequency := make(map[float64]int)
	for _,t := range temperature1{
		frequency[t] ++
	}
	fmt.Println(frequency)  //map[-33:1 -31:1 -29:2 -28:2 -23:1 32:1]
	for key,val := range frequency{
		fmt.Printf("%+.2f: %d\n",key,val)
		/*
		-28.00: 2
		+32.00: 1
		*/
	}
	fmt.Println("___________________19.5 使用映射和切片实现数据分组___________")
	groups := make(map[float64][]float64)   //值是切片

	for _,t := range temperature1{
		level := math.Trunc(t/10)*10   //得到t的级别
		groups[level] = append(groups[level],t)     //加到对应切片中

	}
	for level,val := range groups{
		fmt.Printf("%v:%v\n",level,val)
		/*
		-30:[-31 -33]
		-20:[-28 -29 -23 -29 -28]
		30:[32]
		*/
	}

	fmt.Println("___________________19.6 映射用作集合___________")
	/*
	集合作为收集器，和数组类似，不过集合是无序的，每个元素只能出现一次
	Go语言并没有直接提供集合收集器
	映射用作集合时，键de值不重要，设置为true即可
	*/
	set := make(map[float64]bool)
	for _,val := range temperature1{
		set[val] = true    //将元素值作为键，利用键唯一的属性
	}
	fmt.Println(set)    //一般而言，键是无序的，若要有序，需要将键转换成切片排序
	//此时输出的是映射，将键转化成切片，就完成了到集合的映射
	unique := make([]float64,0,len(set))
	for t := range set{
		unique = append(unique, t)
	}
	fmt.Println(unique)    //这就是我们一般认识的集合了，没有重复元素

	words()

}
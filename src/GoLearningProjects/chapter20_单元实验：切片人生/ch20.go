package main

import (
	"fmt"
	"math/rand"
	"time"
)
/*
康威生命游戏，单元实验
*/

const (
	width  = 80
	height = 15
)                      //网格的尺寸
type Universe [][]bool //二维细胞网格，用true表示细胞的存货
func NewUniverse() Universe {
	//返回
	u := make(Universe, height) //行矩阵
	for index := range u {
		u[index] = make([]bool, width) //列矩阵
	}
	return u
}
//为Universe类型绑定一个方法，首字大写说明方法是public的
func (u Universe) Print()string{
	var b byte
	buf := make([]byte,0,(width+1)*height)
	for y:=0;y<height;y++{
		for x:=0;x<width;x++{
			b = ' '
			if u[y][x] {
				b = '*'   //细胞活着
			}
			buf = append(buf, b)
		}
		buf = append(buf, '\n')
	}
	return string(buf)
}
func (u Universe) Show() {
	fmt.Print("\x0c", u.Print())   //清屏，输出
	fmt.Println()
}
//激活细胞
func (u Universe) Set(x, y int, val bool) {
	u[y][x] = val
}
func (u Universe) Seed() {
	for i := 0; i < (width * height / 4); i++ {
		u.Set(rand.Intn(width), rand.Intn(height), true) //随机激活
	}
}

func (u Universe) Alive(x,y int) bool{
	//需要处理越界问题
	x = (x+width) % width
	y = (y+height) % height
	return u[y][x]
}
//统计临近细胞的存活数量
func (u Universe) Neighbor(x,y int) int{
	cnt := 0
	for i:=-1;i<=1;i++{
		for j:=-1;j<=1;j++{
			//遍历8角位置，跳过自己的位置  0，0
			if(!(i==0 && j==0) && u.Alive(x+j,y+i)){
				cnt++
			}
		}
	}
	return cnt
}

//游戏规则 next
func (u Universe) Next(x,y int)bool{
	//根据附近细胞数，实现规则定义
	cnt := u.Neighbor(x,y)
	return u.Alive(x,y) && cnt == 2 || cnt ==3
} 
//平行世界  将当前世界a的状态更新，保存到b中
func Step(a,b Universe){
	for y:=0;y<height;y++{
		for x:=0;x<width;x++{
			b.Set(x,y,a.Next(x,y))
		}
	}
}
func main(){
	a,b := NewUniverse(),NewUniverse()
	a.Seed()
	for i:=0;i<20;i++{
		Step(a,b)
		a.Show()
		time.Sleep(time.Second/30)
		a,b = b,a   //交换
	}

}
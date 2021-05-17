// package main

// import (

// 	"testing"
// 	"time"
// )



// type stardater interface {
// 	YearDay() int
// 	Hour() int 
// }  //time.Time类型满足这个接口的定义
// /*
// 这种做法在java中无法实现，因为无法让java.time通过显式声明实现接口，它是库函数啊 
// */
// //t变量可以指向任意类型变量，只要其类型实现了YearDay()/Hour()
// func starDate(t stardater) float64{
// 	doy := float64(t.YearDay())
// 	h := float64(t.Hour()) /24.0
// 	return 1000 + doy + h
// }
// type sol int
// func (s sol) YearDay() int {
// 	return int(s % 668)
// }
// func (s sol) Hour() int {
// 	return 0
// }   //现在sol类型也实现了stardate接口
// /*
// 使用单元测试的方法，姑且解决多个main函数的问题
// 书中一个单元，每个实例都是一个main函数，显然每次创建新文件，新库不合适
// */
// func TestStarDate(t *testing.T){
	
// 	var day stardater = time.Date(2012,8,6,5,17,0,0,time.UTC)
// 	s := starDate(day)
// 	if s != 0 {
// 		t.Errorf("day is %s",day)
// 	}
	
// 	s1 := sol(1422)
// 	s2 := starDate(s1)
// 	if s2 != 0 {
// 		t.Errorf("sol day is %v\n",s2)
// 	}
// 	// t.Log(day)
// 	t.Failed()
// }

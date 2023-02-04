package main

import "fmt"

// ================== 各个子系统类 =====================
type SubSystemA struct{}

func (sa *SubSystemA) MethodA() {
	fmt.Println("子系统方法A")
}

type SubSystemB struct{}

func (sb *SubSystemB) MethodB() {
	fmt.Println("子系统方法B")
}

type SubSystemC struct{}

func (sc *SubSystemC) MethodC() {
	fmt.Println("子系统方法C")
}

type SubSystemD struct{}

func (sd *SubSystemD) MethodD() {
	fmt.Println("子系统方法D")
}

// 外观模式，提供了一个简单的接口供外部使用
type Facade struct {
	sa *SubSystemA
	sb *SubSystemB
	sc *SubSystemC
	sd *SubSystemD
}

func (f *Facade) MethodOne() {
	f.sa.MethodA()
	f.sb.MethodB()
}

func (f *Facade) MethodTwo() {
	f.sc.MethodC()
	f.sd.MethodD()
}

func NewFacade() *Facade {
	return &Facade{
		sa: &SubSystemA{},
		sb: &SubSystemB{},
		sc: &SubSystemC{},
		sd: &SubSystemD{},
	}
}

// 业务代码
func main() {
	facade := NewFacade()
	facade.MethodOne()
	facade.MethodTwo()
}

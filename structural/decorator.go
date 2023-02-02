/*
感觉和代理模式很相似，不同之处在于：代理模式增加的额外功能是在代理内部添加的，对客户端是透明的；而装饰模式需要在客户端自行添加。
*/

package main

import (
	"fmt"
)

// 抽象的构件
type Phone interface {
	Show()
}

// 具体的构件
type HuaWei struct{}

func (h *HuaWei) Show() {
	fmt.Println("华为手机")
}

type XiaoMi struct{}

func (x *XiaoMi) Show() {
	fmt.Println("小米手机")
}

// 抽象的装饰器
// golang 1.9 之前的接口不支持成员属性
type Decorator struct {
	phone Phone
}

// 具体的装饰器类
type MoDecorator struct {
	Decorator
}

func (md *MoDecorator) Show() {
	md.phone.Show()
	fmt.Println("手机贴膜")
}

type KeDecorator struct {
	Decorator
}

func (kd *KeDecorator) Show() {
	kd.phone.Show()
	fmt.Println("手机加壳")
}

func NewMoDecorator(ph Phone) Phone {
	return &MoDecorator{Decorator{ph}}
}

func NewKeDecorator(ph Phone) Phone {
	return &KeDecorator{Decorator{ph}}
}

func main() {
	hw := &HuaWei{}
	hw.Show()
	xm := &XiaoMi{}
	xm.Show()

	fmt.Println("=======================")

	mhw := NewMoDecorator(hw)
	mhw.Show()

	fmt.Println("=======================")

	khw := NewKeDecorator(hw)
	khw.Show()

	fmt.Println("=======================")
	kmhw := NewMoDecorator(khw)
	kmhw.Show()
}

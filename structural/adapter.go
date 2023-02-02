/*
适配器模式
适配目标 <-----> 适配器 <-----> 适配者
客户端 <-----> 中间层 <-----> 最上层
*/

package main

import "fmt"

// 抽象的适配目标
type V5 interface {
	Use5V()
}

// +
type V12 interface {
	Use12V()
}

type Phone struct {
	v V5
}

type PhoneHigh struct {
	v V12
}

func NewPhone(v V5) *Phone {
	return &Phone{v}
}

// +
func NewPhoneHigh(v V12) *PhoneHigh {
	return &PhoneHigh{v}
}

func (p *Phone) Charge() {
	fmt.Println("Phone正在充电")
	p.v.Use5V()
}

func (ph *PhoneHigh) Charge() {
	fmt.Println("Phone正在充电")
	ph.v.Use12V()
}

// 适配者
type V220 struct{}

func (v *V220) Use220V() {
	fmt.Println("使用220v电压充电")
}

// 新增加的适配者
type V300 struct{}

func (v *V300) Use300V() {
	fmt.Println("使用300电压充电")
}

// 适配器
type Adapter struct {
	v220 *V220
}

func (a *Adapter) Use5V() {
	fmt.Println("使用5v充电")
	a.v220.Use220V()
}

// +适配器
type AdapterHigh struct {
	v300 *V300
}

func (ah *AdapterHigh) Use12V() {
	fmt.Println("使用12v充电")
	ah.v300.Use300V()
}

func NewAdapter(v220 *V220) *Adapter {
	return &Adapter{v220}
}

func NewAdapterHigh(v300 *V300) *AdapterHigh {
	return &AdapterHigh{v300}
}

// 业务层
func main() {
	phone := NewPhone(NewAdapter(&V220{}))
	phone.Charge()
	fmt.Println("-------------")
	p := NewPhoneHigh(NewAdapterHigh(&V300{}))
	p.Charge()
}

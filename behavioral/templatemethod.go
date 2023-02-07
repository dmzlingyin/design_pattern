package main

import "fmt"

// Beverage 抽象类
type Beverage interface {
	BoilWater()
	Brew()
	PourInCup()
	AddThings()

	WantAddThings() bool
}

// 封装流程模板
type template struct {
	b Beverage
}

func (t *template) MakeBeverage() {
	if t == nil {
		return
	}

	t.b.BoilWater()
	t.b.Brew()
	t.b.PourInCup()

	if t.b.WantAddThings() {
		t.b.AddThings()
	}
}

// 具体的模板子类
type MakeCaffee struct {
	template
}

// 自己给自己赋值？
func NewMakeCaffee() *MakeCaffee {
	mk := &MakeCaffee{}
	mk.b = mk
	return mk
}

func (mc *MakeCaffee) BoilWater() {
	fmt.Println("将水烧到100摄氏度")
}

func (mc *MakeCaffee) Brew() {
	fmt.Println("用水煮咖啡豆")
}

func (mc *MakeCaffee) PourInCup() {
	fmt.Println("将煮好的咖啡豆放入陶瓷杯中")
}

func (mc *MakeCaffee) AddThings() {
	fmt.Println("添加奶和糖")
}

func (mc *MakeCaffee) WantAddThings() bool {
	return true
}

// 另一个具体的模板子类
type MakeTea struct {
	template
}

func NewMakeTea() *MakeTea {
	mt := &MakeTea{}
	mt.b = mt
	return mt
}

func (mt *MakeTea) BoilWater() {
	fmt.Println("将水烧到80摄氏度")
}

func (mt *MakeTea) Brew() {
	fmt.Println("用水煮茶叶")
}

func (mt *MakeTea) PourInCup() {
	fmt.Println("将煮好的茶倒入茶壶中")
}

func (mt *MakeTea) AddThings() {
	fmt.Println("添加柠檬")
}

func (mt *MakeTea) WantAddThings() bool {
	return false
}

func main() {
	mc := NewMakeCaffee()
	mc.MakeBeverage()
	fmt.Println("---------------")
	mt := NewMakeTea()
	mt.MakeBeverage()
}

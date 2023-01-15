/*
优点：
1.  拥有工厂方法模式的优点
2. 当一个产品族中的多个对象被设计成一起工作时，它能够保证客户端始终只使用同一个产品族中的对象。
3   增加新的产品族很方便，无须修改已有系统，符合“开闭原则”。

缺点：
1. 增加新的产品等级结构麻烦，需要对原有系统进行较大的修改，甚至需要修改抽象层代码，这显然会带来较大的不便，违背了“开闭原则”。

3.3.5 适用场景
(1) 系统中有多于一个的产品族。而每次只使用其中某一产品族。可以通过配置文件等方式来使得用户可以动态改变产品族，也可以很方便地增加新的产品族。
(2) 产品等级结构稳定。设计完成之后，不会向系统中增加新的产品等级结构或者删除已有的产品等级结构。
*/
package main

import "fmt"

// 抽象层
type AbstractBanana interface {
	ShowBanana()
}

type AbstractApple interface {
	ShowApple()
}

type AbstractPear interface {
	ShowPear()
}

type AbstractFactory interface {
	CreateApple() AbstractApple
	CreateBanana() AbstractBanana
	CreatePear() AbstractPear
}

// 实现层
// 中国产品族
type ChineseApple struct{}

func (ca *ChineseApple) ShowApple() {
	fmt.Println("中国苹果")
}

type ChineseBanana struct{}

func (cb *ChineseBanana) ShowBanana() {
	fmt.Println("中国香蕉")
}

type ChinesePear struct{}

func (cp *ChinesePear) ShowPear() {
	fmt.Println("中国梨")
}

type ChineseFactory struct{}

func (cf *ChineseFactory) CreateApple() AbstractApple {
	return &ChineseApple{}
}

func (cf *ChineseFactory) CreateBanana() AbstractBanana {
	return &ChineseBanana{}
}

func (cf *ChineseFactory) CreatePear() AbstractPear {
	return &ChinesePear{}
}

// 美国产品族
type AmericanApple struct{}

func (aa *AmericanApple) ShowApple() {
	fmt.Println("美国苹果")
}

type AmericanBanana struct{}

func (an *AmericanBanana) ShowBanana() {
	fmt.Println("美国香蕉")
}

type AmericanPear struct{}

func (ap *AmericanPear) ShowPear() {
	fmt.Println("美国梨")
}

type AmericanFactory struct{}

func (af *AmericanFactory) CreateApple() AbstractApple {
	return &AmericanApple{}
}

func (af *AmericanFactory) CreateBanana() AbstractBanana {
	return &AmericanBanana{}
}

func (af *AmericanFactory) CreatePear() AbstractPear {
	return &AmericanPear{}
}

// 业务逻辑层
func main() {
	cf := &ChineseFactory{}
	cf.CreateApple().ShowApple()
	cf.CreateBanana().ShowBanana()
	cf.CreatePear().ShowPear()

	af := &AmericanFactory{}
	af.CreateApple().ShowApple()
	af.CreateBanana().ShowBanana()
	af.CreatePear().ShowPear()
}

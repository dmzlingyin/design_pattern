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

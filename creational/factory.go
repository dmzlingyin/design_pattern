/*
工厂模式组成: 抽象工厂 <- 工厂 <-> 抽象产品 <- 具体产品
工厂方法 = 简单工厂 + “开闭原则”

优点：
1. 不需要记住具体类名，甚至连具体参数都不用记忆。
2. 实现了对象创建和使用的分离。
3. 系统的可扩展性也就变得非常好，无需修改接口和原类。
4.对于新产品的创建，符合开闭原则。

缺点：
1. 增加系统中类的个数，复杂度和理解度增加。
2. 增加了系统的抽象性和理解难度。

适用场景：
1. 客户端不知道它所需要的对象的类。
2. 抽象工厂类通过其子类来指定创建哪个对象。
*/
package main

import "fmt"

// 抽象层
type Fruit interface {
	Show()
}

type AbstractFactory interface {
	CreateFruit() Fruit
}

// 基础模块
type Apple struct{}

func (a *Apple) Show() {
	fmt.Println("apple")
}

type Pear struct{}

func (p *Pear) Show() {
	fmt.Println("pear")
}

// 增加新的水果
type Banana struct{}

func (b *Banana) Show() {
	fmt.Println("新增加的banana")
}

// 工厂模块
type AppleFactory struct{}

func (a *AppleFactory) CreateFruit() Fruit {
	return &Apple{}
}

type PearFatory struct{}

func (p *PearFatory) CreateFruit() Fruit {
	return &Pear{}
}

type BananaFactory struct{}

func (b *BananaFactory) CreateFruit() Fruit {
	return &Banana{}
}

// 业务逻辑层
func main() {
	af := &AppleFactory{}
	apple := af.CreateFruit()
	apple.Show()

	pf := &PearFatory{}
	pear := pf.CreateFruit()
	pear.Show()

	bf := &BananaFactory{}
	banana := bf.CreateFruit()
	banana.Show()
}

package main

import "fmt"

// ===================== 抽象的动物类 =====================
type Animal interface {
	Eat()
}

// ===================== 具体的动物类 =====================
type Cat struct{}

func (c *Cat) Eat() {
	fmt.Println("小猫在吃饭")
}

type Dog struct{}

func (d *Dog) Eat() {
	fmt.Println("小狗在吃饭")
}

func NewAnimal(name string) Animal {
	if name == "cat" {
		return &Cat{}
	} else if name == "dog" {
		return &Dog{}
	} else if name == "bird" {
		return &Bird{}
	}
	return nil
}

// 新增加一个具体的动物类
type Bird struct{}

func (b *Bird) Eat() {
	fmt.Println("小鸟在吃饭")
}

// ===================== 装饰器类 =====================
type Decorator struct {
	animal Animal
}

type CatDecorator struct {
	Decorator
}

func (cd *CatDecorator) Eat() {
	fmt.Println("吃之前，先舔一舔")
	cd.animal.Eat()
}

type DogDecorator struct {
	Decorator
}

func (dd *DogDecorator) Eat() {
	fmt.Println("吃之前先叫一叫")
	dd.animal.Eat()
}

type BirdDecorator struct {
	Decorator
}

func (bd *BirdDecorator) Eat() {
	fmt.Println("吃之前先跳一跳")
	bd.animal.Eat()
}

func NewCatDecorator(animal Animal) Animal {
	return &CatDecorator{
		Decorator{animal},
	}
}

func NewDogDecorator(animal Animal) Animal {
	return &DogDecorator{
		Decorator{animal},
	}
}

func NewBirdDecorator(animal Animal) Animal {
	return &BirdDecorator{
		Decorator{animal},
	}
}

// 业务层代码
func main() {
	cat := NewAnimal("cat")
	dog := NewAnimal("dog")
	bird := NewAnimal("bird")
	cat.Eat()
	dog.Eat()
	bird.Eat()

	fmt.Println("=============================")

	cd := NewCatDecorator(cat)
	cd.Eat()
	fmt.Println("-----")
	dd := NewDogDecorator(dog)
	dd.Eat()
	fmt.Println("-----")
	bd := NewBirdDecorator(bird)
	bd.Eat()
}

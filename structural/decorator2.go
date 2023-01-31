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
	}
	return nil
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

func main() {
	cat := NewAnimal("cat")
	dog := NewAnimal("dog")
	cat.Eat()
	dog.Eat()

	fmt.Println("=============================")

	cd := NewCatDecorator(cat)
	cd.Eat()
	fmt.Println("-----")
	dd := NewDogDecorator(dog)
	dd.Eat()
}

package main

import "fmt"

/*
依赖倒转原则
面向接口编程

业务逻辑层<---->抽象层<---->实现层
*/

/*------------ 抽象层 -------------*/
type Car interface {
	Run()
}

type Driver interface {
	Drive(car Car)
}

/*------------实现层 -------------*/
type Toyoto struct{}

func (t *Toyoto) Run() {
	fmt.Println("toyoto running")
}

type Handi struct{}

func (h *Handi) Run() {
	fmt.Println("Handi running")
}

type ZS struct{}

func (z *ZS) Drive(car Car) {
	fmt.Println("zhang san")
	car.Run()
}

type LS struct{}

func (l *LS) Drive(car Car) {
	fmt.Println("li si")
	car.Run()
}

// 业务逻辑层
func main() {
	t := &Toyoto{}
	h := &Handi{}

	zs := &ZS{}
	ls := &LS{}

	zs.Drive(t)
	zs.Drive(h)
	ls.Drive(h)
	ls.Drive(t)
}

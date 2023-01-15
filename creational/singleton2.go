package main

import "fmt"

type singleton struct{}

var instance *singleton

// "懒汉模式"
func GetInstance() *singleton {
	if instance == nil {
		return &singleton{}
	}
	return instance
}

func (s *singleton) Show() {
	fmt.Println("懒汉模式")
}

func main() {
	ins := GetInstance()
	ins.Show()
}

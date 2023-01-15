package main

import (
	"fmt"
	"sync"
)

var lock sync.Mutex

type singleton struct{}

var instance *singleton

func GetInstance() *singleton {
	lock.Lock()
	defer lock.Unlock()

	if instance == nil {
		return &singleton{}
	}
	return instance
}

func (s *singleton) Show() {
	fmt.Println("线程安全的懒汉模式")
}

func main() {
	ins := GetInstance()
	ins.Show()
}

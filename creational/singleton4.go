package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var initialized int32
var lock sync.Mutex

type singleton struct{}

var instance *singleton

func GetInstance() *singleton {
	// 已经实例化
	if atomic.LoadInt32(&initialized) == 1 {
		return instance
	}

	lock.Lock()
	defer lock.Unlock()

	if initialized == 0 {
		instance = &singleton{}
		atomic.StoreInt32(&initialized, 1)
	}
	return instance
}

func (s *singleton) Show() {
	fmt.Println("引入automic的懒汉模式")
}

func main() {
	ins := GetInstance()
	ins.Show()
}

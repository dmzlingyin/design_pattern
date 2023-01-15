/*
优点：
(1) 单例模式提供了对唯一实例的受控访问。
(2) 节约系统资源。由于在系统内存中只存在一个对象。

缺点：
(1) 扩展略难。单例模式中没有抽象层。
(2) 单例类的职责过重。

3.4.5 适用场景
(1) 系统只需要一个实例对象，如系统要求提供一个唯一的序列号生成器或资源管理器，或者需要考虑资源消耗太大而只允许创建一个对象。
(2) 客户调用类的单个实例只允许使用一个公共访问点，除了该公共访问点，不能通过其他途径访问该实例。
*/
package main

import (
	"fmt"
	"sync"
)

var once sync.Once

// 该类只能有一个实例, 并且提供可以全局访问的方法
type singleton struct{}

var instance *singleton

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

func (s *singleton) Show() {
	fmt.Println("使用sync.Once实现懒汉模式")
}

func main() {
	ins := GetInstance()
	ins.Show()
}

// 单例模式：保证一个类只能有一个对象，且该对象的功能依然能被其他模块使用
package main

import "fmt"

// 保证这个类非公有化
type singleton struct{}

// 保证指针指向这个唯一对象，指针不能改变方向
var instance = &singleton{}

// “饿汉模式”
func GetInstance() *singleton {
	return instance
}

func (s *singleton) Show() {
	fmt.Println("饿汉模式")
}

func main() {
	s := GetInstance()
	s.Show()
}

package main

import "fmt"

// 原型模式适用于创建复杂对象，需要将这些对象创建成原型，并通过复制这些原型来创建新的对象。

// 抽象的产品
type IResume interface {
	Print()
	Clone() IResume
}

// 具体的产品
type Resume struct {
	name string
}

func (r *Resume) Print() {
	fmt.Println(r.name)
}

func (r *Resume) Clone() IResume {
	return &Resume{name: r.name + "(cloned)"}
}

func main() {
	r := &Resume{name: "hello"}
	r.Print()
	r.Clone().Print()
	r.Clone().Clone().Print()
}

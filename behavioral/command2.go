package main

import (
	"fmt"
)

// 实际的计算核心
type Worker struct{}

func (w *Worker) Sheep() {
	fmt.Println("烤羊肉")
}

func (w *Worker) Chicken() {
	fmt.Println("烤鸡肉")
}

// 抽象的命令
type Command interface {
	BBQ()
}

// 实际的命令
type BBQSheep struct {
	w *Worker
}

func (b *BBQSheep) BBQ() {
	b.w.Sheep()
}

type BBQChicken struct {
	w *Worker
}

func (b *BBQChicken) BBQ() {
	b.w.Chicken()
}

// 命令的调用者
type Waiter struct {
	Cmds []Command
}

func (w *Waiter) Notify() {
	if w.Cmds == nil {
		return
	}
	for _, cmd := range w.Cmds {
		cmd.BBQ()
	}
}

// 业务代码
func main() {
	sheep := &BBQSheep{}
	chicken := &BBQChicken{}

	// 负责命令的调用
	w := &Waiter{}
	w.Cmds = append(w.Cmds, sheep)
	w.Cmds = append(w.Cmds, chicken)
	w.Notify()
}

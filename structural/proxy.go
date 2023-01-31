/*
代理在帮助客户端实现基本需求的同时，还可以做一些其他的额外功能。感觉类似于简单工厂的思想，屏蔽了一些内部的细节。
加一个中间层！！！！！！
*/

package main

import (
	"fmt"
)

type Goods struct {
	Kind string // 商品种类
	Fact bool   // 商品真伪
}

// ====================== 抽象层(代理方需要帮客户端实现的抽象方法) ======================
type Shopping interface {
	Buy(goods *Goods)
}

// ====================== 实现层 ======================
type KoreaShopping struct{}

func (ks *KoreaShopping) Buy(goods *Goods) {
	fmt.Printf("韩国代购, 买了%s\n", goods.Kind)
}

type AmericanShopping struct{}

func (as *AmericanShopping) Buy(goods *Goods) {
	fmt.Printf("美国代购, 买了%s\n", goods.Kind)
}

type JapanShoppint struct{}

func (js *JapanShoppint) Buy(goods *Goods) {
	fmt.Printf("日本代购, 买了%s\n", goods.Kind)
}

// 具体的海外代理
type OverseaProxy struct {
	shopping Shopping
}

// 海外代理要实现的方法
func (op *OverseaProxy) Buy(goods *Goods) {
	if op.distinguish(goods) {
		op.shopping.Buy(goods)
		op.check(goods)
	}
}

func (op *OverseaProxy) distinguish(goods *Goods) bool {
	fmt.Println("进行货物鉴定")
	return goods.Fact
}

func (op *OverseaProxy) check(goods *Goods) {
	fmt.Println("进行海关检查")
}

// 创建一个代理，代理对应的货物
func NewProxy(sp Shopping) Shopping {
	return &OverseaProxy{sp}
}

// 业务实现层
func main() {
	// 如果不是用代理模式
	sumsang := &Goods{"phone", true}
	shopping := &KoreaShopping{}
	// 先验货
	if sumsang.Fact {
		// 自己去韩国购买
		shopping.Buy(sumsang)
		fmt.Println("自己去海关安检")
	}

	// 使用代理模式
	overseaProxy := NewProxy(shopping)
	overseaProxy.Buy(sumsang)
}

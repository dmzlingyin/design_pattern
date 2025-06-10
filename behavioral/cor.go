package main

import "fmt"

// 责任链是一种行为设计模式， 允许你将请求沿着处理者链进行发送， 直至其中一个处理者对其进行处理

type Patient struct {
	name             string
	registrationDone bool
	doctorDone       bool
	medicineDone     bool
	paymentDone      bool
}

// 抽象类, 每个处理者都必须实现
type Department interface {
	execute(*Patient)
	setNext(Department)
}

// 具体处理者
type Reception struct {
	next Department
}

func (r *Reception) execute(p *Patient) {
	if p.registrationDone {
		fmt.Println(fmt.Sprintf("%s已挂号", p.name))
		r.next.execute(p)
		return
	}
	fmt.Println(fmt.Sprintf("%s挂号", p.name))
	p.registrationDone = true
	if r.next != nil {
		r.next.execute(p)
	}
}

func (r *Reception) setNext(next Department) {
	r.next = next
}

type Doctor1 struct {
	next Department
}

func (d *Doctor1) execute(p *Patient) {
	if p.doctorDone {
		fmt.Println(fmt.Sprintf("%s已看完👩🏻‍⚕️", p.name))
		d.next.execute(p)
		return
	}
	fmt.Println(fmt.Sprintf("%s看👩🏻‍⚕️", p.name))
	p.doctorDone = true
	if d.next != nil {
		d.next.execute(p)
	}
}

func (d *Doctor1) setNext(next Department) {
	d.next = next
}

type Medical struct {
	next Department
}

func (m *Medical) execute(p *Patient) {
	if p.medicineDone {
		fmt.Println(fmt.Sprintf("%s已拿完💊", p.name))
		m.next.execute(p)
		return
	}
	fmt.Println(fmt.Sprintf("%s拿💊", p.name))
	p.medicineDone = true
	if m.next != nil {
		m.next.execute(p)
	}
}

func (m *Medical) setNext(next Department) {
	m.next = next
}

type Cashier struct {
	next Department
}

func (c *Cashier) execute(p *Patient) {
	if p.paymentDone {
		fmt.Println(fmt.Sprintf("%s已付钱", p.name))
		c.next.execute(p)
		return
	}
	fmt.Println(fmt.Sprintf("%s付钱", p.name))
	p.paymentDone = true
	if c.next != nil {
		c.next.execute(p)
	}
}

func (c *Cashier) setNext(next Department) {
	c.next = next
}

func main() {
	cashier := &Cashier{}
	medical := &Medical{}
	medical.setNext(cashier)
	doctor := &Doctor1{}
	doctor.setNext(medical)
	reception := &Reception{}
	reception.setNext(doctor)

	reception.execute(&Patient{
		name: "小明",
	})
}

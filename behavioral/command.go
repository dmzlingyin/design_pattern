package main

import "fmt"

// 命令接收者
type Doctor struct {
}

func (d *Doctor) treatEye() {
	fmt.Println("医生治疗眼睛")
}

func (d *Doctor) treatNose() {
	fmt.Println("医生治疗鼻子")
}

// 抽象的命令;
type Command interface {
	Treat()
}

// 治疗眼睛的订单
type CommandTreatEye struct {
	d *Doctor
}

func (cmd *CommandTreatEye) Treat() {
	cmd.d.treatEye()
}

// 治疗鼻子的订单
type CommandTreatNose struct {
	d *Doctor
}

// 调用命令者
type Nurse struct {
	Cmds []Command
}

func (n *Nurse) Notify() {
	if n.Cmds == nil {
		return
	}
	for _, cmd := range n.Cmds {
		cmd.Treat()
	}
}

func (cmd *CommandTreatNose) Treat() {
	cmd.d.treatNose()
}

func main() {
	eye := &CommandTreatEye{}
	nose := &CommandTreatNose{}

	nurse := &Nurse{}
	nurse.Cmds = append(nurse.Cmds, eye)
	nurse.Cmds = append(nurse.Cmds, nose)
	nurse.Notify()
}

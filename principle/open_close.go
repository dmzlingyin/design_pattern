package main

import "fmt"

// 绘制矩形和正方形的类
type Drawer struct {
	x int
	y int
}

func (d *Drawer) Rect(w, h int) {
	fmt.Printf("起点x: %d, 起点y:%d, 长: %d, 宽: %d\n", d.x, d.y, w, h)
}

func (d *Drawer) Square(length int) {
	fmt.Printf("起点x: %d, 起点y:%d, 边长: %d\n", d.x, d.y, length)
}

func (d *Drawer) Circle(radius int) {
	fmt.Printf("起点x: %d, 起点y:%d, 半径: %d\n", d.x, d.y, radius)
}

func main() {
	d := &Drawer{0, 0}
	d.Rect(8, 6)
	d.Square(6)
	d.Circle(3)
}

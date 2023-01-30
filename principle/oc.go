package main

import "fmt"

type Drawer interface {
	Draw()
}

type RectDrawer struct {
	x int
	y int
	w int
	h int
}

func (rd *RectDrawer) Draw() {
	fmt.Printf("起点x: %d, 起点y:%d, 长: %d, 宽: %d\n", rd.x, rd.y, rd.w, rd.h)
}

type SquareDrawer struct {
	x      int
	y      int
	length int
}

func (sd *SquareDrawer) Draw() {
	fmt.Printf("起点x: %d, 起点y:%d, 边长: %d\n", sd.x, sd.y, sd.length)
}

type CircleDrawer struct {
	x      int
	y      int
	radius int
}

func (cd *CircleDrawer) Draw() {
	fmt.Printf("起点x: %d, 起点y:%d, 半径: %d\n", cd.x, cd.y, cd.radius)
}

func DrawShape(d Drawer) {
	d.Draw()
}

func main() {
	rd := &RectDrawer{0, 0, 8, 6}
	DrawShape(rd)

	sd := &SquareDrawer{0, 0, 6}
	DrawShape(sd)

	cd := &CircleDrawer{0, 0, 3}
	DrawShape(cd)
}

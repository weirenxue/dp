package main

import "fmt"

type Renderer interface {
	DrawCircle(radius int)
}

type VectorRenderer struct{}

func (v *VectorRenderer) DrawCircle(radius int) {
	fmt.Println("Drawing a circle of radius", radius)
}

type RasterRenderer struct {
	Dpi int
}

func (r *RasterRenderer) DrawCircle(radius int) {
	fmt.Println("Drawing pixels for circle of radius", radius)
}

type Circle struct {
	renderer Renderer
	radius   int
}

func NewCircle(renderer Renderer, radius int) *Circle {
	return &Circle{renderer: renderer, radius: radius}
}

func (c *Circle) Draw() {
	c.renderer.DrawCircle(c.radius)
}

func (c *Circle) Resize(factor int) {
	c.radius *= factor
}

func main() {
	// raster := RasterRenderer{}
	vector := VectorRenderer{}
	c := NewCircle(&vector, 5)
	c.Draw()
	c.Resize(2)
	c.Draw()
}

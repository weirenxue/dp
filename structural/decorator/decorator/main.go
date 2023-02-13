package main

import "fmt"

type Shape interface {
	Render() string
}

type Circle struct {
	Radius float32
}

func (c *Circle) Render() string {
	return fmt.Sprintf("Circle of radius %f", c.Radius)
}

func (c *Circle) Resize(factor float32) {
	c.Radius *= factor
}

type Square struct {
	Side float32
}

func (s *Square) Render() string {
	return fmt.Sprintf("Square of side %f", s.Side)
}

type ColoredShape struct {
	Shape Shape
	Color string
}

func (c *ColoredShape) Render() string {
	return fmt.Sprintf("%s has color %s", c.Shape.Render(), c.Color)
}

type TransparencyShape struct {
	Shape        Shape
	Transparency float32
}

func (t *TransparencyShape) Render() string {
	return fmt.Sprintf("%s has %f%% transparency", t.Shape.Render(), t.Transparency)
}

func main() {
	circle := Circle{Radius: 10}
	fmt.Println(circle.Render())

	coloredCircle := ColoredShape{Shape: &circle, Color: "Red"}
	fmt.Println(coloredCircle.Render())

	rhsCircle := TransparencyShape{Shape: &coloredCircle, Transparency: 0.5}
	fmt.Println(rhsCircle.Render())

}

package main

import (
	"fmt"
	"strings"
)

type GraphicObject struct {
	Name, Color string
	Children    []GraphicObject
}

func (g *GraphicObject) String() string {
	sb := strings.Builder{}
	g.print(&sb, 0)
	return sb.String()
}

func (g *GraphicObject) print(sb *strings.Builder, depth int) {
	sb.WriteString(strings.Repeat("*", depth))
	if len(g.Color) > 0 {
		sb.WriteString(g.Color)
		sb.WriteString(" ")
	}
	sb.WriteString(g.Name)
	sb.WriteString("\n")
	for _, c := range g.Children {
		c.print(sb, depth+1)
	}
}

func NewCircle(color string) *GraphicObject {
	return &GraphicObject{
		Name:  "Circle",
		Color: color,
	}
}

func NewRectangle(color string) *GraphicObject {
	return &GraphicObject{
		Name:  "Rectangle",
		Color: color,
	}
}

func main() {
	myGraphic := GraphicObject{Name: "Grapich 1", Color: "", Children: nil}
	myGraphic.Children = append(myGraphic.Children, *NewCircle("Red"))
	myGraphic.Children = append(myGraphic.Children, *NewRectangle("Yellow"))

	group1 := GraphicObject{Name: "Group 1", Color: "", Children: nil}
	group1.Children = append(group1.Children, *NewCircle("Blue"))
	group1.Children = append(group1.Children, *NewRectangle("Blue"))

	myGraphic.Children = append(myGraphic.Children, group1)

	fmt.Println(myGraphic.String())
}

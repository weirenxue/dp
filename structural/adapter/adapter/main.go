package main

import (
	"fmt"
	"strings"
)

type Line struct {
	X1, Y1, X2, Y2 int
}

type VectorImage struct {
	Lines []Line
}

func NewRectangle(width, height int) VectorImage {
	width--
	height--
	return VectorImage{
		Lines: []Line{
			{X1: 0, Y1: 0, X2: width, Y2: 0},
			{X1: width, Y1: 0, X2: width, Y2: height},
			{X1: 0, Y1: height, X2: width, Y2: height},
			{X1: 0, Y1: 0, X2: 0, Y2: height},
		},
	}
}

// ↑↑↑ the interface we're given

type vectorToRasterAdapter struct {
	points []Point
}

func (v *vectorToRasterAdapter) addLine(l Line) {
	dx := l.X2 - l.X1
	if dx == 0 {
		for i := l.Y1; i <= l.Y2; i++ {
			v.points = append(v.points, Point{X: l.X1, Y: i})
		}
	} else {
		for i := l.X1; i <= l.X2; i++ {
			v.points = append(v.points, Point{X: i, Y: l.Y1})
		}
	}

	fmt.Println("generated", len(v.points), "points")
}

func (v *vectorToRasterAdapter) GetPoints() []Point {
	return v.points
}

func VectorToRaster(v VectorImage) RasterImage {
	adapter := vectorToRasterAdapter{}
	for _, l := range v.Lines {
		adapter.addLine(l)
	}
	return &adapter
}

// ↓↓↓ the interface we have

type Point struct {
	X, Y int
}

type RasterImage interface {
	GetPoints() []Point
}

func DrawPoints(ri RasterImage) string {
	points := ri.GetPoints()
	xmax := 0
	ymax := 0
	for _, p := range points {
		if p.X > xmax {
			xmax = p.X
		}
		if p.Y > ymax {
			ymax = p.Y
		}
	}

	xmax++
	ymax++

	s := make([][]rune, ymax)
	for i := range s {
		s[i] = make([]rune, xmax)
		for j := range s[i] {
			s[i][j] = rune(' ')
		}
	}

	for _, p := range points {
		s[p.Y][p.X] = rune('*')
	}

	sb := strings.Builder{}
	for i := range s {
		sb.WriteString(string(s[i]))
		sb.WriteString("\n")
	}
	return sb.String()
}

func main() {
	r := NewRectangle(6, 4)
	a := VectorToRaster(r)
	fmt.Println(DrawPoints(a))
}

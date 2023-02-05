package main

import "fmt"

type Sized interface {
	Width() int
	SetWidth(width int)
	Height() int
	SetHeight(height int)
}

type Rectangle struct {
	width, height int
}

func (r *Rectangle) Width() int {
	return r.width
}

func (r *Rectangle) SetWidth(width int) {
	r.SetWidth(width)
}

func (r *Rectangle) Height() int {
	return r.height
}

func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

// break lsp
type Square struct {
	Rectangle
}

func NewSquare(size int) *Square {
	s := Square{}
	s.width = size
	s.height = size
	return &s
}

func (s *Square) SetWidth(width int) {
	s.width = width
	s.height = width
}

func (s *Square) SetHeight(height int) {
	s.width = height
	s.height = height
}

type Squre2 struct {
	size int
}

func (s *Squre2) Rectangle() Rectangle {
	return Rectangle{width: s.size, height: s.size}
}

func UseIt(sized Sized) {
	width := sized.Width()
	sized.SetHeight(10)
	expectedArea := 10 * width
	actualArea := sized.Width() * sized.Height()
	fmt.Printf("Expected area of %d, but get %d\n", expectedArea, actualArea)
}

func main() {
	r := Rectangle{width: 2, height: 5}
	UseIt(&r)

	s := NewSquare(5)
	UseIt(s)

	s2 := Squre2{size: 5}
	r = s2.Rectangle()
	UseIt(&r)
}

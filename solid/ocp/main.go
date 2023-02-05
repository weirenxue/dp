package main

import "fmt"

// open for extension, close for modification
// Specification

type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct{}

func (f *Filter) FilterByColor(products []Product, color Color) []*Product {
	var result []*Product
	for i, p := range products {
		if p.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

func (f *Filter) FilterBySize(products []Product, size Size) []*Product {
	var result []*Product
	for i, p := range products {
		if p.size == size {
			result = append(result, &products[i])
		}
	}
	return result
}

func (f *Filter) FilterByColorAndSize(products []Product, color Color, size Size) []*Product {
	var result []*Product
	for i, p := range products {
		if p.color == color && p.size == size {
			result = append(result, &products[i])
		}
	}
	return result
}

type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecfication struct {
	color Color
}

func (c *ColorSpecfication) IsSatisfied(p *Product) bool {
	return p.color == c.color
}

type SizeSpecification struct {
	size Size
}

func (s *SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == s.size
}

type AndSpecification struct {
	first, second Specification
}

func (a *AndSpecification) IsSatisfied(p *Product) bool {
	return a.first.IsSatisfied(p) && a.second.IsSatisfied(p)
}

type OrSpecification struct {
	first, second Specification
}

func (o *OrSpecification) IsSatisfied(p *Product) bool {
	return o.first.IsSatisfied(p) || o.second.IsSatisfied(p)
}

type BetterFilter struct{}

func (b *BetterFilter) Filter(products []Product, spec Specification) []*Product {
	var result []*Product
	for i, p := range products {
		if spec.IsSatisfied(&p) {
			result = append(result, &products[i])
		}
	}
	return result
}

func main() {
	apple := Product{name: "Apple", color: green, size: small}
	tree := Product{name: "Tree", color: green, size: large}
	house := Product{name: "House", color: blue, size: large}
	products := []Product{apple, tree, house}

	f := Filter{}
	fmt.Println("Green products (old):")
	for _, v := range f.FilterByColor(products, green) {
		fmt.Printf(" - %s\n", v.name)
	}

	bf := BetterFilter{}
	fmt.Println("Green product (new):")
	greenSpec := ColorSpecfication{color: green}
	for _, v := range bf.Filter(products, &greenSpec) {
		fmt.Printf(" - %s\n", v.name)
	}

	largeSpec := SizeSpecification{size: large}
	largeAndGreenSpec := AndSpecification{first: &largeSpec, second: &greenSpec}
	fmt.Println("Green large product")
	for _, v := range bf.Filter(products, &largeAndGreenSpec) {
		fmt.Printf(" - %s\n", v.name)
	}

	redSpec := ColorSpecfication{color: red}
	redOrGreenSpec := OrSpecification{first: &redSpec, second: &greenSpec}
	fmt.Println("Red or green product")
	for _, v := range bf.Filter(products, &redOrGreenSpec) {
		fmt.Printf(" - %s\n", v.name)
	}
}

package main

import (
	"fmt"
	"strings"
)

const (
	indentSize = 2
)

type HtmlElement struct {
	name, text string
	elements   []HtmlElement
}

func (h *HtmlElement) String() string {
	return h.string(0)
}

func (h *HtmlElement) string(indent int) string {
	sb := strings.Builder{}
	i := strings.Repeat(" ", indentSize*indent)
	sb.WriteString(fmt.Sprintf("%s<%s>\n", i, h.name))

	if len(h.text) > 0 {
		i := strings.Repeat(" ", indentSize*(indent+1))
		sb.WriteString(fmt.Sprintf("%s%s\n", i, h.text))
	}

	for _, e := range h.elements {
		sb.WriteString(e.string(indent + 1))
	}

	sb.WriteString(fmt.Sprintf("%s</%s>\n", i, h.name))
	return sb.String()
}

type HtmlBuilder struct {
	rootName string
	root     HtmlElement
}

func NewHtmlBuilder(rootName string) *HtmlBuilder {
	return &HtmlBuilder{rootName: rootName, root: HtmlElement{name: rootName}}
}

func (h *HtmlBuilder) AddChild(name, text string) {
	h.root.elements = append(h.root.elements, HtmlElement{name: name, text: text})
}

func (h *HtmlBuilder) AddChildFluent(name, text string) *HtmlBuilder {
	h.root.elements = append(h.root.elements, HtmlElement{name: name, text: text})
	return h
}

func (h *HtmlBuilder) String() string {
	return h.root.String()
}

func main() {
	hello := "hello"
	sb := strings.Builder{}
	sb.WriteString("<p>")
	sb.WriteString(hello)
	sb.WriteString("</p>")
	fmt.Println(sb.String())

	words := []string{"hello", "world"}
	sb.Reset()
	sb.WriteString("<ul>")
	for _, v := range words {
		sb.WriteString("<li>")
		sb.WriteString(v)
		sb.WriteString("</li>")
	}
	sb.WriteString("</ul>")
	fmt.Println(sb.String())

	b := NewHtmlBuilder("ul")
	b.AddChildFluent("li", "hello").
		AddChild("li", "world")
	fmt.Println(b.String())
}

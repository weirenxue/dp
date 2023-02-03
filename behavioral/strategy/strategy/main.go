package main

import (
	"fmt"
	"strings"
)

type OutputFormat int

const (
	Markdown OutputFormat = iota
	Html
)

type ListStrategy interface {
	Start(builder *strings.Builder)
	End(builder *strings.Builder)
	AddItem(builder *strings.Builder, item string)
}

type MarkdownListStrategy struct{}

func (m *MarkdownListStrategy) Start(builder *strings.Builder) {}
func (m *MarkdownListStrategy) End(builder *strings.Builder)   {}
func (m *MarkdownListStrategy) AddItem(builder *strings.Builder, item string) {
	builder.WriteString(" * " + item + "\n")
}

type HtmlListStrategy struct{}

func (m *HtmlListStrategy) Start(builder *strings.Builder) {
	builder.WriteString("<ul>\n")
}
func (m *HtmlListStrategy) End(builder *strings.Builder) {
	builder.WriteString("</li>\n")
}
func (m *HtmlListStrategy) AddItem(builder *strings.Builder, item string) {
	builder.WriteString(" <li>" + item + "</li>\n")
}

type TextProcessor struct {
	builder      strings.Builder
	listStrategy ListStrategy
}

func NewTextProcessor(l ListStrategy) *TextProcessor {
	return &TextProcessor{builder: strings.Builder{}, listStrategy: l}
}

func (t *TextProcessor) AddItems(items []string) {
	t.listStrategy.Start(&t.builder)
	for _, item := range items {
		t.listStrategy.AddItem(&t.builder, item)
	}
	t.listStrategy.End(&t.builder)
}

func (t *TextProcessor) Reset() {
	t.builder.Reset()
}

func (t *TextProcessor) SetOutputFormat(o OutputFormat) {
	switch o {
	case Html:
		t.listStrategy = &HtmlListStrategy{}
	case Markdown:
		t.listStrategy = &MarkdownListStrategy{}
	}
}

func (t *TextProcessor) String() string {
	return t.builder.String()
}

func main() {
	tp := NewTextProcessor(&MarkdownListStrategy{})
	tp.AddItems([]string{"foo", "bar", "baz"})
	fmt.Println(tp)

	tp.Reset()
	tp.SetOutputFormat(Html)
	tp.AddItems([]string{"foo", "bar", "baz"})
	fmt.Print(tp)
}

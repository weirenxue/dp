package main

import (
	"fmt"
	"strings"
	"unicode"
)

type FormattedText struct {
	plainText  string
	capatalize []bool
}

func NewFormattedText(text string) *FormattedText {
	return &FormattedText{plainText: text, capatalize: make([]bool, len(text))}
}

func (f *FormattedText) String() string {
	sb := strings.Builder{}
	for i := range f.plainText {
		c := f.plainText[i]
		if f.capatalize[i] {
			c = byte(unicode.ToUpper(rune(c)))
		}
		sb.WriteByte(c)
	}
	return sb.String()
}

func (f *FormattedText) Capatalize(start, end int) {
	for i := start; i <= end; i++ {
		f.capatalize[i] = true
	}
}

type TextRange struct {
	Start, End               int
	Capatalize, Bold, Italic bool
}

func (t *TextRange) Covers(index int) bool {
	return t.Start <= index && index <= t.End
}

type BetterFormattedText struct {
	plainText  string
	formatting []*TextRange
}

func NewBetterFormattedText(text string) *BetterFormattedText {
	return &BetterFormattedText{plainText: text}
}

func (b *BetterFormattedText) String() string {
	sb := strings.Builder{}
	for i := range b.plainText {
		c := b.plainText[i]
		for _, f := range b.formatting {
			if f.Covers(i) && f.Capatalize {
				c = byte(unicode.ToUpper(rune(c)))
			}
		}
		sb.WriteByte(c)
	}
	return sb.String()
}

func (b *BetterFormattedText) Range(start, end int) *TextRange {
	r := &TextRange{Start: start, End: end}
	b.formatting = append(b.formatting, r)
	return r
}

func main() {
	text := "This is a brave new world"

	ft := NewFormattedText(text)
	ft.Capatalize(10, 15)
	fmt.Println(ft.String())

	bft := NewBetterFormattedText(text)
	bft.Range(16, 19).Capatalize = true
	fmt.Println(bft.String())
}

package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

var entryCount int

type Journal struct {
	entries []string
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) Remove(index int) {
	// ...
}

// seperation of concerns

func (j *Journal) Save(filename string) {
	os.WriteFile(filename, []byte(j.String()), 0644)
}

func (j *Journal) Load(filename string) {}

func (j *Journal) LoadFromWeb(url *url.URL) {}

// In other package

var LineSeperator = "\n"

func SaveToFile(j *Journal, filename string) {
	os.WriteFile(filename, []byte(strings.Join(j.entries, LineSeperator)), 0644)
}

// Other struct

type Persistence struct {
	lineSeperator string
}

func (p *Persistence) SaveToFile(j *Journal, filename string) {
	os.WriteFile(filename, []byte(strings.Join(j.entries, p.lineSeperator)), 0644)
}

func main() {
	j := &Journal{}
	j.AddEntry("Washing dog")
	j.AddEntry("On Udemy")

	// break srp
	// j.Save("journal0.txt")

	// seperation of concerns
	SaveToFile(j, "journal1.txt")

	p := Persistence{lineSeperator: "\n"}
	p.SaveToFile(j, "journal2.txt")
}

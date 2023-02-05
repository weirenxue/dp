package main

import "fmt"

// HLM should not depend on LLM
// Both should depend on abstractions

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	name string
}

type Info struct {
	from         *Person
	relationship Relationship
	to           *Person
}

// LLM, low-level module
type Relationships struct {
	relations []Info
}

func (r *Relationships) AddParentAndChild(parent, child *Person) {
	r.relations = append(r.relations,
		Info{from: parent, relationship: Parent, to: child},
		Info{from: child, relationship: Child, to: parent},
	)
}

// HLM, hight-level module
type Research struct {
	relationships *Relationships
}

func (r *Research) Investigate() {
	relations := r.relationships.relations
	for _, rel := range relations {
		if rel.from.name == "John" && rel.relationship == Parent {
			fmt.Println("John has a child called", rel.to.name)
		}
	}
}

// follow DIP

type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

type BetterRelationships struct {
	relations []Info
}

func (b *BetterRelationships) AddParentAndChild(parent, child *Person) {
	b.relations = append(b.relations,
		Info{from: parent, relationship: Parent, to: child},
		Info{from: child, relationship: Child, to: parent},
	)
}

func (b *BetterRelationships) FindAllChildrenOf(name string) []*Person {
	results := make([]*Person, 0)
	for _, rel := range b.relations {
		if rel.from.name == name && rel.relationship == Parent {
			results = append(results, rel.to)
		}
	}
	return results
}

type BetterResearch struct {
	browser RelationshipBrowser
}

func (b *BetterResearch) Investigate() {
	for _, p := range b.browser.FindAllChildrenOf("John") {
		fmt.Println("John has a child called", p.name)
	}
}

func main() {
	parent := Person{name: "John"}
	child1 := Person{name: "Chris"}
	child2 := Person{name: "Matt"}

	relationships := &Relationships{}
	relationships.AddParentAndChild(&parent, &child1)
	relationships.AddParentAndChild(&parent, &child2)

	r := &Research{relationships: relationships}
	r.Investigate()

	fmt.Println("Apply dependency inversion principle...")
	brel := &BetterRelationships{}
	brel.AddParentAndChild(&parent, &child1)
	brel.AddParentAndChild(&parent, &child2)

	brea := &BetterResearch{browser: brel}
	brea.Investigate()
}

package main

import "fmt"

type Creature struct {
	Name            string
	Attack, Defense int
}

func (c *Creature) String() string {
	return fmt.Sprintf("%s (%d/%d)",
		c.Name, c.Attack, c.Defense)
}

func NewCreature(name string, attack, defense int) *Creature {
	return &Creature{Name: name, Attack: attack, Defense: defense}
}

// Each element in the linked list should implement the modifier interface.
type Modifier interface {
	Add(m Modifier)
	Handle()
}

// The basic implementer in the linked list.
type CreatureModifier struct {
	creature *Creature
	next     Modifier
}

func NewCreatureModifier(c *Creature) *CreatureModifier {
	return &CreatureModifier{creature: c}
}

// Add m to the linked list tail by recursion.
func (c *CreatureModifier) Add(m Modifier) {
	if c.next != nil {
		c.next.Add(m)
	} else {
		c.next = m
	}
}

func (c *CreatureModifier) Handle() {
	if c.next != nil {
		c.next.Handle()
	}
}

type DoubleAttackModifier struct {
	CreatureModifier
}

func NewDoubleAttackModifier(c *Creature) *DoubleAttackModifier {
	return &DoubleAttackModifier{CreatureModifier{creature: c}}
}

// Override the Handle function of the CreatureModifier
func (d *DoubleAttackModifier) Handle() {
	fmt.Println("Doubling", d.creature.Name, "\b's attack...")
	d.creature.Attack *= 2
	d.CreatureModifier.Handle()
}

type IncreaseDefenseModifier struct {
	CreatureModifier
}

func NewIncreaseDefenseModifier(c *Creature) *IncreaseDefenseModifier {
	return &IncreaseDefenseModifier{CreatureModifier{creature: c}}
}

// Override the Handle function of the CreatureModifier
func (i *IncreaseDefenseModifier) Handle() {
	if i.creature.Attack <= 2 {
		fmt.Println("Increasing", i.creature.Name, "\b's defense...")
		i.creature.Defense += 1
	}
	i.CreatureModifier.Handle()
}

type NoBonusesModifier struct {
	CreatureModifier
}

func NewNoBonusesModifier(c *Creature) *NoBonusesModifier {
	return &NoBonusesModifier{CreatureModifier{creature: c}}
}

// Override the Handle function of the CreatureModifier
func (n *NoBonusesModifier) Handle() {
	// nothing
}

func main() {
	goblin := NewCreature("Goblin", 1, 1)
	fmt.Println(goblin)

	root := NewCreatureModifier(goblin)

	root.Add(NewIncreaseDefenseModifier(goblin))
	root.Add(NewDoubleAttackModifier(goblin))
	root.Add(NewIncreaseDefenseModifier(goblin))
	root.Add(NewDoubleAttackModifier(goblin))
	root.Add(NewIncreaseDefenseModifier(goblin)) // No effect, attack power greater than 2.
	root.Add(NewNoBonusesModifier(goblin))       // Stop the method chain.
	root.Add(NewDoubleAttackModifier(goblin))

	root.Handle()
	fmt.Println(goblin)
}

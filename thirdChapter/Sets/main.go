package main

import "fmt"

// crate set class
type Set struct {
	intMap map[int]bool
}

// init sets
func (set *Set) New() {
	set.intMap = make(map[int]bool)
}

// AddElement
func (set *Set) AddElement(element int) {
	set.intMap[element] = true
}

// DeleteElement
func (set *Set) DeleteElement(element int) {
	delete(set.intMap, element)
}

// ContainsElement
func (set *Set) ContainsElement(element int) bool {
	_, exists := set.intMap[element]
	return exists
}

// InterSect
func (set *Set) InterSect(another *Set) *Set {
	var interSect *Set = &Set{}
	interSect.New()
	var value int
	for value, _ = range set.intMap {
		if another.ContainsElement(value) {
			interSect.AddElement(value)
		}
	}
	return interSect
}

// Union
func (set *Set) Union(another *Set) *Set {
	var unionSect *Set = &Set{}
	unionSect.New()
	var value int
	for value, _ = range set.intMap {
		unionSect.AddElement(value)
	}

	for value, _ = range another.intMap {
		unionSect.AddElement(value)
	}
	return unionSect
}

func main() {
	set := &Set{}
	set.New()
	set.AddElement(1)
	set.AddElement(2)
	fmt.Println(set.ContainsElement(1))
	fmt.Println(set.ContainsElement(3))
	anotherSet := &Set{}
	anotherSet.New()
	anotherSet.AddElement(2)
	anotherSet.AddElement(4)
	anotherSet.AddElement(5)
	fmt.Println(set.InterSect(anotherSet))
	fmt.Println(set.Union(anotherSet))

}

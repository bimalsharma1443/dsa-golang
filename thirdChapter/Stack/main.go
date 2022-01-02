package main

import "fmt"

type Element struct {
	elementValue int
}

type Stack struct {
	elements     []*Element
	elementCount int
}

func (stack *Stack) New() {
	stack.elements = make([]*Element, 0)
}

func (stack *Stack) Push(element *Element) {
	stack.elements = append(stack.elements, element)
	stack.elementCount++
}

func (stack *Stack) Pop() *Element {
	if stack.elementCount == 0 {
		return nil
	}
	stack.elementCount--
	var element *Element = stack.elements[stack.elementCount]
	if stack.elementCount > 1 {
		stack.elements = stack.elements[:stack.elementCount]
	} else {
		stack.elements = stack.elements[0:]
	}

	return element
}

func main() {
	var stack *Stack = &Stack{}
	stack.New()
	var element1 *Element = &Element{3}
	var element2 *Element = &Element{5}
	var element3 *Element = &Element{7}
	var element4 *Element = &Element{9}
	stack.Push(element1)
	stack.Push(element2)
	stack.Push(element3)
	stack.Push(element4)
	fmt.Println(stack.Pop(), stack.Pop(), stack.Pop(), stack.Pop())
}

// A package to implement a generic stack in Go

package genstack

import "fmt"

type Stack[T any] struct {
	vals []interface{}
}

func (s *Stack[T]) Push(val interface{}) {
	s.vals = append(s.vals, val)
}

func (s *Stack[T]) isEmpty() bool {
	return len(s.vals) == 0
}

func (s *Stack[T]) Pop() (val interface{}, err error) {
	if s.isEmpty() {
		var zero T
		return zero, fmt.Errorf("Stack is empty")
	}
	val = s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	return val, nil
}

func (s *Stack[T]) Top() (val interface{}, err error) {
	if s.isEmpty() {
		var zero T
		return zero, fmt.Errorf("stack is empty")
	}
	return s.vals[len(s.vals)-1], nil
}

// Fill the stack from a slice
func (s *Stack[T]) CopyFromSlice(slice []interface{}) {
	for _, val := range slice {
		s.Push(val)
	}
}

// Pops the stack contents into a slice
func (s *Stack[T]) CopyToSlice() []interface{} {
	var slice []interface{}
	for !s.isEmpty() {
		val, err := s.Pop()
		if err != nil {
			break
		}
		slice = append(slice, val)
	}
	return slice
}

func main() {
	fmt.Println("Stacks")
	var intStack Stack[int]
	fmt.Println(intStack)
	intStack.Push(15)
	intStack.Push("dog")
	intStack.Push(25)
	fmt.Println(intStack)
	fmt.Println(intStack.isEmpty())

	// Copy stack contents to a slice
	slice := intStack.CopyToSlice()
	fmt.Println("Slice:", slice)
	fmt.Println("Stack after CopyToSlice:", intStack)
	intStack.CopyFromSlice(slice)
	fmt.Println("Stack after CopyFromSlice:", intStack)
}

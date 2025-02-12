// Running unit tests for genstack package

package genstack

import (
	"testing"
)

func TestPushPop(t *testing.T) {
	stack := Stack[int]{}

	stack.Push(10)
	stack.Push(20)

	val, err := stack.Pop()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if val != 20 {
		t.Errorf("Expected 20, got %v", val)
	}

	val2, err2 := stack.Pop()
	if err2 != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if val2 != 10 {
		t.Errorf("Expected 10, got %v", val2)
	}

	_, err = stack.Pop()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestIsEmpty(t *testing.T) {
	stack := Stack[int]{}

	if !stack.isEmpty() {
		t.Errorf("Expected stack to be empty")
	}

	stack.Push(10)
	if stack.isEmpty() {
		t.Errorf("Expected stack to be non-empty")
	}

	stack.Pop()
	if !stack.isEmpty() {
		t.Errorf("Expected stack to be empty")
	}
}

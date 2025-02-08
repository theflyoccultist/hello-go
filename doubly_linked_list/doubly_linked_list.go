// Doubly linked list to store a sequence of characters and determine if it is a palindrome

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ListElement struct {
	data rune         // data of the element
	next *ListElement // pointer to the next element
	prev *ListElement // pointer to the previous element
}

func createListElement(data rune, ptr *ListElement) *ListElement {
	var element ListElement
	element.data = data
	element.next = ptr
	if ptr != nil {
		element.prev = ptr.prev
	}
	return &element
}

func (h *ListElement) PrintList() {
	if h == nil {
		fmt.Println("List is empty")
		return
	}
	fmt.Printf("%c -> ", h.data)
	h.next.PrintList()
}

func AddToFront(dataSlice []rune, h **ListElement) {
	head := *h
	for _, v := range dataSlice {
		head = createListElement(v, head)
	}
	*h = head
}

func AddToRear(dataSlice []rune, h **ListElement) {
	head := *h
	for _, v := range dataSlice {
		newElement := createListElement(v, nil)
		if head != nil {
			head.next = newElement
		}
		head = newElement
	}
}

func DeleteFront(h **ListElement) {
	head := *h
	if head == nil {
		return
	}
	*h = head.next
	if head.next != nil {
		head.next.prev = nil
	}
}

func DeleteRear(h **ListElement) {
	head := *h
	if head == nil {
		return
	}
	for head.next != nil {
		head = head.next
	}
	if head.prev != nil {
		head.prev.next = nil
	} else {
		*h = nil
	}
}

func FindValue(value rune, h *ListElement) *ListElement {
	if h == nil {
		return nil
	}
	if h.data == value {
		return h
	}
	return FindValue(value, h.next)
}

func DeleteValue(value rune, h **ListElement) {
	head := *h
	if head == nil {
		return
	}
	if head.data == value {
		DeleteFront(h)
		return
	}
	for head.next != nil {
		if head.next.data == value {
			head.next = head.next.next
			if head.next != nil {
				head.next.prev = head
			}
			return
		}
		head = head.next
	}
}

func IsEmpty(h *ListElement) bool {
	if h == nil {
		return true
	}
	return false
}

func FindLength(h *ListElement) int {
	if h == nil {
		return 0
	}
	return 1 + FindLength(h.next)
}

func InsertPosition(value rune, position int, h **ListElement) {
	head := *h
	if position < 0 {
		return
	}
	if position == 0 {
		*h = createListElement(value, head)
		return
	}
	for i := 0; i < position-1; i++ {
		if head == nil {
			return
		}
		head = head.next
	}
	if head == nil {
		return
	}
	head.next = createListElement(value, head.next)
}

func DeletePosition(position int, h **ListElement) {
	head := *h
	if position < 0 {
		return
	}
	if position == 0 {
		DeleteFront(h)
		return
	}
	for i := 0; i < position-1; i++ {
		if head == nil {
			return
		}
		head = head.next
	}
	if head == nil {
		return
	}
	head.next = head.next.next
	if head.next != nil {
		head.next.prev = head
	}
}

func IsPalindrome(h *ListElement) bool {
	if h == nil {
		return false
	}

	// Find the tail of the List
	tail := h
	for tail.next != nil {
		tail = tail.next
	}

	// Iterate from both ends towards the middle
	for h != nil && tail != nil && h != tail && h.prev != tail {
		if h.data != tail.data {
			return false
		}
		h = h.next
		tail = tail.prev
	}

	return true
}

func main() {
	var head *ListElement

	fmt.Print("Type a word into the terminal to check if it is a palindrome: \n")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input")
		return
	}
	input = strings.TrimSpace(input)

	// Convert the input string to a slice of runes
	dataslice := ([]rune)(input)

	// Add the input to the front of the doubly linked list
	AddToFront(dataslice, &head)
	fmt.Println("Added to front")
	head.PrintList()
	fmt.Println()

	// Test if the input is a palindrome
	fmt.Println("Is the input a palindrome? ")
	fmt.Println(IsPalindrome(head))
	fmt.Println()

	// Test the other doubly linked list functions
	fmt.Println("Testing the doubly linked list functions")

	AddToRear(dataslice, &head)
	fmt.Println("Added to rear")
	head.PrintList()
	fmt.Println()

	FindValue('a', head)
	if FindValue('a', head) != nil {
		fmt.Println("Value 'a' found")
	} else {
		fmt.Println("Value 'a' not found")
	}
	head.PrintList()
	fmt.Println()

	if FindValue('a', head) != nil {
		fmt.Println("Deleted value 'a'")
		DeleteValue('a', &head)
	} else {
		fmt.Println("Value 'a' not found and not deleted")
	}

	IsEmpty(head)
	if IsEmpty(head) {
		fmt.Println("List is empty")
	} else {
		fmt.Println("List is not empty")
	}

	fmt.Println()
	FindLength(head)
	fmt.Println("Length of the list is: ", FindLength(head))

	InsertPosition('a', 0, &head)
	fmt.Println("Inserted 'a' at position 0")
	head.PrintList()
	fmt.Println()

	DeletePosition(0, &head)
	fmt.Println("Deleted position 0")
	head.PrintList()
	fmt.Println()

	DeleteFront(&head)
	fmt.Println("Deleted front element")
	head.PrintList()
	fmt.Println()

	DeleteRear(&head)
	fmt.Println("Deleted rear element")
	head.PrintList()
	fmt.Println()
}

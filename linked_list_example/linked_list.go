package main

import "fmt"

type ListElement struct {
	data int          // can be any other type
	next *ListElement // pointer to the next element
}

func createListElement(data int, ptr *ListElement) *ListElement {
	var element ListElement
	element.data = data
	element.next = ptr
	return &element
}

func (h *ListElement) PrintList() {
	if h == nil {
		fmt.Println("List is empty")
		return
	}
	fmt.Println(h.data, "->")
	h.next.PrintList()
}

func FillList(dataSlice []int, h **ListElement) {
	currentElement := ListElement{dataSlice[0], nil}
	fmt.Println("currentElement: ", currentElement)
	*h = &currentElement
	ptrCurrentElement := &currentElement

	for i := 1; i < len(dataSlice); i++ {
		nextElement := ListElement{dataSlice[i], nil}
		fmt.Println("nextElement: ", nextElement)
		ptrCurrentElement.next = &nextElement
		fmt.Println("CurrentElement", currentElement)
		ptrCurrentElement = &nextElement
	}
}

func main() {
	var head *ListElement
	var e, f ListElement
	directFill := []int{4, 8, 16, 32, 64, 128, 256, 512}
	fmt.Println("Create List from", directFill)

	head = &e
	e.data = 5
	e.next = &f
	f.data = 10
	fmt.Println("Element e is: ", e, "Head: ", head)
	fmt.Println("Element f is: ", f, "Head: ", head)
	f.next = new(ListElement)
	f.next.data = 15
	fmt.Println("Element f.next is: ", f.next, "Head: ", head)
	head.PrintList()
	FillList(directFill, &head)
	fmt.Println("Head points to: ", head)
	head.PrintList()
}

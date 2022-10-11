package main

import "fmt"

type NodeSinglyLinked struct {
	value int
	next  *NodeSinglyLinked
}

type SinglyLinkedList struct {
	head   *NodeSinglyLinked
	length int
}

func (L *SinglyLinkedList) Insert(newNode NodeSinglyLinked, index int) {
	if index <= L.length {
		if L.head == nil {
			L.head = &newNode
		} else if index == 0 {
			newNode.next = L.head
			L.head = &newNode
		} else {
			current := L.head
			for i := 0; i < index-1; i++ {
				current = current.next
			}
			previousNextNode := current.next
			newNode.next = previousNextNode
			current.next = &newNode
		}
		L.length += 1
	} else {
		fmt.Printf("Cannot insert node at index %v since the List's current length is %v", index, L.length)
	}
}

func (L *SinglyLinkedList) Traverse() []int {
	var current *NodeSinglyLinked = L.head
	var resultArray []int
	for current != nil {
		resultArray = append(resultArray, current.value)
		current = current.next
	}
	return resultArray
}

// Remove Element at Given Position within List
func (L *SinglyLinkedList) RemoveELement(index int) *SinglyLinkedList {
	current := L.head
	var prev *NodeSinglyLinked = nil
	for i := 0; i < index && current != nil; i++ {
		prev = current
		current = current.next
	}
	if L.head == current {
		L.head = current.next
	} else {
		prev.next = current.next
	}
	L.length -= 1

	return L
}

// Get value of item given an index

func (L *SinglyLinkedList) Get(index int) *NodeSinglyLinked {
	return nil
}

// Get index of item within a node by given value.
func (L *SinglyLinkedList) Search(value int) int {
	return -1
}

func main() {
	fmt.Println("Linked List Lesson")
	var mySinglyLinkedList SinglyLinkedList

	mySinglyLinkedList.Insert(NodeSinglyLinked{0, nil}, 0)
	mySinglyLinkedList.Insert(NodeSinglyLinked{1, nil}, 1)
	mySinglyLinkedList.Insert(NodeSinglyLinked{2, nil}, 2)
	fmt.Println(mySinglyLinkedList.length)
	mySinglyLinkedList.Traverse()
	mySinglyLinkedList.Insert(NodeSinglyLinked{5, nil}, 0)
	fmt.Println(mySinglyLinkedList.length)
	fmt.Println(mySinglyLinkedList.Traverse())
	mySinglyLinkedList.RemoveELement(2)
	fmt.Println(mySinglyLinkedList.length)
	fmt.Println(mySinglyLinkedList.Traverse())
	// Use this method to play with your linked list implementation
}

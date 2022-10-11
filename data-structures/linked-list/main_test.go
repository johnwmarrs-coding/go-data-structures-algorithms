package main

import "testing"
import "reflect"

func TestTester(t *testing.T) {
	valueFailed := false
	if valueFailed {
		t.Errorf("The testing test has failed.")
	}
}

func TestInsert(t *testing.T) {
	var mySinglyLinkedList SinglyLinkedList
	// Insert to empty list works
	mySinglyLinkedList.Insert(NodeSinglyLinked{0, nil}, 0)

	result := []int{0}
	if !reflect.DeepEqual(mySinglyLinkedList.Traverse(), result) {
		t.Errorf("Inserting at beginning failed")
	}

	// Insert to end of list works
	result = []int{0, 1}
	mySinglyLinkedList.Insert(NodeSinglyLinked{1, nil}, 1)
	if !reflect.DeepEqual(mySinglyLinkedList.Traverse(), result) {
		t.Errorf("Inserting at end failed")
	}

	// Insert into middle of list works
	result = []int{0, 2, 1}
	mySinglyLinkedList.Insert(NodeSinglyLinked{2, nil}, 1)
	if !reflect.DeepEqual(mySinglyLinkedList.Traverse(), result) {
		t.Errorf("Inserting into middle failed")
	}
}

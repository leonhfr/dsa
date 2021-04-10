/*
Package linkedlist creates a Linked List data structure for the Item type.
Use this generic implementation to generate type-specific sets.

Generate a `IntLinkedList` linked list of `int` values:
	genny -in linkedlist.go -out linkedlist-int.go gen "Item=int"
Generate a `StringLinkedList` linked list of `string` values:
	genny -in linkedlist.go -out linkedlist-string.go gen "Item=string"

A linked list provides a data structure similar to an array, but with the advantage
that inserting an element in the middle of the list is very cheap, compared to
doing so in an array, where we need to shift all elements after the current position,

While arrays keep all the elements in the same block of memory, next to each other,
linked lists can contain elements scattered around memory, by storing a pointer to
the next element.

A disadvantage over arrays on the other hand is that if we want to pick an element
in the middle of the list, we don't know its address, so we need to start at the
beginning of the list and iterate on the list until we find it.
*/
package linkedlist

import (
	"fmt"
	"sync"

	"github.com/cheekybits/genny/generic"
)

// Item the type of the linked list
type Item generic.Type

// Node a single node that composes the list
type Node struct {
	content Item
	next    *Node
}

// ItemLinkedList the linked list of Items
type ItemLinkedList struct {
	head *Node
	size int
	lock sync.RWMutex
}

// Append adds an Item to the end of the linked lost
func (ll *ItemLinkedList) Append(t Item) {
	ll.lock.Lock()
	defer ll.lock.Unlock()
	ll.size++
	node := Node{t, nil}
	if ll.head == nil {
		ll.head = &node
		return
	}
	last := ll.head
	for {
		if last.next == nil {
			break
		}
		last = last.next
	}
	last.next = &node
}

// Insert adds an Item at position i
func (ll *ItemLinkedList) Insert(i int, t Item) error {
	ll.lock.Lock()
	defer ll.lock.Unlock()
	if i < 0 || i > ll.size {
		return fmt.Errorf("index out of bounds")
	}
	ll.size++
	addNode := Node{t, nil}
	if i == 0 {
		addNode.next = ll.head
		ll.head = &addNode
		return nil
	}
	node := ll.head
	for j := 0; j < i-2; j++ {
		node = node.next
	}
	addNode.next = node.next
	node.next = &addNode
	return nil
}

// RemoveAt removes an Item at position i
func (ll *ItemLinkedList) RemoveAt(i int) (*Item, error) {
	ll.lock.Lock()
	defer ll.lock.Unlock()
	if i < 0 || i > ll.size {
		return nil, fmt.Errorf("index out of bounds")
	}
	node := ll.head
	for j := 0; j < i-1; j++ {
		node = node.next
	}
	remove := node.next
	node.next = remove.next
	ll.size--
	return &remove.content, nil
}

// IndexOf returns the position of the Item t
func (ll *ItemLinkedList) IndexOf(t Item) int {
	ll.lock.RLock()
	defer ll.lock.RUnlock()
	node := ll.head
	index := 0
	for {
		if node.content == t {
			return index
		}
		if node.next == nil {
			return -1
		}
		node = node.next
		index++
	}
}

// IsEmpty returns true if the list is empty
func (ll *ItemLinkedList) IsEmpty() bool {
	ll.lock.RLock()
	defer ll.lock.RUnlock()
	return ll.head == nil
}

// Size returns the linked list size
func (ll *ItemLinkedList) Size() int {
	ll.lock.RLock()
	defer ll.lock.RUnlock()
	size := 1
	last := ll.head
	for {
		if last == nil || last.next == nil {
			break
		}
		last = last.next
		size++
	}
	return size
}

// Insert adds an Item at position i
func (ll *ItemLinkedList) String() {
	ll.lock.RLock()
	defer ll.lock.RUnlock()
	node := ll.head
	j := 0
	for {
		if node == nil {
			break
		}
		j++
		fmt.Print(node.content)
		fmt.Print(" ")
		node = node.next
	}
	fmt.Println()
}

// Head returns a pointer to the first node of the list
func (ll *ItemLinkedList) Head() *Node {
	ll.lock.RLock()
	defer ll.lock.RUnlock()
	return ll.head
}

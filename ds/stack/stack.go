/*
Package stack creates a Stack data structure for the Item type.
Use this generic implementation to generate type-specific stacks.

Generate a `IntStack` stack of `int` values:
	genny -in stack.go -out stack-int.go gen "Item=int"
Generate a `StringStack` stack of `string` values:
	genny -in stack.go -out stack-string.go gen "Item=string"

A stack is an ordered collection of items following the last in, first out (LIFO) principle.
We add and remove items from the same part of the stack, called the top.
We cannot remove items from the base, just like a pile of books.

Internally, the stack is represented by a slice type.
*/
package stack

import (
	"sync"

	"github.com/cheekybits/genny/generic"
)

// Item the type of the stack
type Item generic.Type

// ItemStack the stack of Items
type ItemStack struct {
	items []Item
	lock  sync.RWMutex
}

// New creates a new ItemStack
func New() *ItemStack {
	return &ItemStack{items: []Item{}}
}

// Push inserts an Item at the top of the ItemStack
func (s *ItemStack) Push(t Item) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.items = append(s.items, t)
}

// Pop removes an Item at the top of the ItemStack
func (s *ItemStack) Pop() *Item {
	s.lock.Lock()
	defer s.lock.Unlock()
	item := s.items[len(s.items)-1]
	s.items = s.items[0 : len(s.items)-1]
	return &item
}

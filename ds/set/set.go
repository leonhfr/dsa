/*
Package set creates a Set data structure for the Item type.
Use this generic implementation to generate type-specific sets.

Generate a `IntSet` set of `int` values:
	genny -in set.go -out set-int.go gen "Item=int"
Generate a `StringSet` set of `string` values:
	genny -in set.go -out set-string.go gen "Item=string"

A set is collection of values.
We can add new values, iterate over those values, remove values,
clear the set, get the set size, and check if the set contains an item.
A value in the set is only stored once, duplicates are not possible.

Internally, the set is represented by a boolean map.
*/
package set

import (
	"sync"

	"github.com/cheekybits/genny/generic"
)

// Item the type of the set
type Item generic.Type

// ItemSet the set of Items
type ItemSet struct {
	items map[Item]bool
	lock  sync.RWMutex
}

// New creates a new ItemSet
func New() *ItemSet {
	return &ItemSet{items: make(map[Item]bool)}
}

// Add adds a new element to the set and returns a pointer to the set
func (s *ItemSet) Add(t Item) *ItemSet {
	s.lock.Lock()
	defer s.lock.Unlock()
	_, ok := s.items[t]
	if !ok {
		s.items[t] = true
	}
	return s
}

// Clear removes all elements from the set
func (s *ItemSet) Clear() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.items = make(map[Item]bool)
}

// Delete removes the Item from the set and returns Has(Item)
func (s *ItemSet) Delete(t Item) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	_, ok := s.items[t]
	if ok {
		delete(s.items, t)
	}
	return ok
}

// Has returns true if the set contains the item
func (s *ItemSet) Has(t Item) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	_, ok := s.items[t]
	return ok
}

// Items returns the items stored
func (s *ItemSet) Items() []Item {
	s.lock.RLock()
	defer s.lock.RUnlock()
	items := []Item{}
	for i := range s.items {
		items = append(items, i)
	}
	return items
}

// Size returnsthe size of the set
func (s *ItemSet) Size() int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.items)
}

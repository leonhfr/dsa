/*
Package queue creates a Queue data structure for the Item type.
Use this generic implementation to generate type-specific queues.

Generate a `IntQueue` queue of `int` values:
	genny -in queue.go -out queue-int.go gen "Item=int"
Generate a `StringQueue` queue of `string` values:
	genny -in queue.go -out queue-string.go gen "Item=string"

A queue is an ordered collection of items following the first in, first out (FIFO) principle.
We add items from one end of the queue, and we retrieve items from the other end.

Internally, the queue is represented by a slice type.
*/
package queue

import (
	"sync"

	"github.com/cheekybits/genny/generic"
)

// Item the type of the queue
type Item generic.Type

// ItemQueue the queue of Items
type ItemQueue struct {
	items []Item
	lock  sync.RWMutex
}

// New creates a new ItemQueue
func New() *ItemQueue {
	return &ItemQueue{items: []Item{}}
}

// Enqueue inserts an Item at the back of the ItemQueue
func (q *ItemQueue) Enqueue(t Item) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.items = append(q.items, t)
}

// Dequeue removes an Item from the front of the ItemQueue
func (q *ItemQueue) Dequeue() *Item {
	q.lock.Lock()
	defer q.lock.Unlock()
	item := q.items[0]
	q.items = q.items[1:len(q.items)]
	return &item
}

// Front returns the next Item in the ItemQueue, without removing it
func (q *ItemQueue) Front() *Item {
	q.lock.RLock()
	defer q.lock.RUnlock()
	item := q.items[0]
	return &item
}

// IsEmpty returns true if the ItemQueue is empty
func (q *ItemQueue) IsEmpty() bool {
	return len(q.items) == 0
}

// Size returns the number of Items in the ItemQueue
func (q *ItemQueue) Size() int {
	return len(q.items)
}

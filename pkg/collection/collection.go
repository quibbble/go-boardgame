package collection

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	ErrCollectionEmpty           = fmt.Errorf("collection empty")
	ErrCollectionIndexOutOfBound = fmt.Errorf("collection index out of bounds")
)

// Collection is an abstraction on a deck, hand, etc and can be used in their place
type Collection[T any] struct {
	items []T
	seed  int64
	r     *rand.Rand
}

func NewCollection[T any](seed int64) *Collection[T] {
	return &Collection[T]{
		items: make([]T, 0),
		seed:  seed,
		r:     rand.New(rand.NewSource(seed)),
	}
}

// Add adds any numver of items to the end of the collection
func (c *Collection[T]) Add(items ...T) {
	c.items = append(c.items, items...)
}

// Draw removes an item from the start of the collection
func (c *Collection[T]) Draw() (*T, error) {
	if len(c.items) == 0 {
		return nil, ErrCollectionEmpty
	}
	draw := c.items[0]
	if err := c.Remove(0); err != nil {
		return nil, err
	}
	return &draw, nil
}

// Remove removes an item by index
func (c *Collection[T]) Remove(index int) error {
	if index < 0 || index >= len(c.items) {
		return ErrCollectionIndexOutOfBound
	}
	c.items = append(c.items[:index], c.items[index+1:]...)
	return nil
}

// Replace replaces the item at the given index with a replacement
func (c *Collection[T]) Replace(index int, replacement T) error {
	if index < 0 || index >= len(c.items) {
		return ErrCollectionIndexOutOfBound
	}
	c.items[index] = replacement
	return nil
}

// Shuffle randomly shuffles the collection
func (c *Collection[T]) Shuffle() {
	if c.r == nil {
		// if random source was never set then set it manually
		seed := time.Now().Unix()
		c.seed = seed
		c.r = rand.New(rand.NewSource(seed))
	}
	for i := 0; i < len(c.items); i++ {
		r := c.r.Intn(len(c.items))
		if i != r {
			c.items[r], c.items[i] = c.items[i], c.items[r]
		}
	}
}

// GetItem returns the item at 'index' in the collection
func (c *Collection[T]) GetItem(index int) (*T, error) {
	if c.GetSize() <= index {
		return nil, ErrCollectionIndexOutOfBound
	}
	return &c.items[index], nil
}

// Contains returns whether or not the item is in the collection
func (c *Collection[T]) Contains(item T, equals func(a, b T) bool) bool {
	for _, it := range c.items {
		if equals(item, it) {
			return true
		}
	}
	return false
}

// IndexOf returns the index of the item in the collection and returns -1 if not found
func (c *Collection[T]) IndexOf(item T, equals func(a, b T) bool) int {
	for idx, it := range c.items {
		if equals(item, it) {
			return idx
		}
	}
	return -1
}

// GetItems returns the items in the collection
func (c *Collection[T]) GetItems() []T {
	return c.items
}

// GetSize returns the number of items in the collection
func (c *Collection[T]) GetSize() int {
	return len(c.items)
}

func (c *Collection[T]) GetSeed() int64 {
	return c.seed
}

// Copy returns a new copy of the collection
func (c *Collection[T]) Clone() *Collection[T] {
	// because maps and slices are references they do not get copied even when passing by value
	// copy creates new slices to mimic passing the slice by value
	return &Collection[T]{
		items: append(make([]T, 0), c.items...),
		seed:  c.seed,
	}
}

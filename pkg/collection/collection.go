package collection

import (
	"fmt"
	"math/rand"
)

const (
	collectionEmptyErr = "collection empty"
	idxOutOfBoundsErr  = "index out of bounds"
)

// Collection is an abstraction on a deck, hand, etc and can be used in their place
type Collection[T any] struct {
	items  []T
	random *rand.Rand
}

func (c *Collection[T]) Add(items ...T) *Collection[T] {
	copy := c.Copy()
	copy.items = append(copy.items, items...)
	return copy
}

func (c *Collection[T]) Draw() (*T, *Collection[T], error) {
	if len(c.items) == 0 {
		return nil, nil, fmt.Errorf(collectionEmptyErr)
	}
	copy, err := c.Remove(0)
	if err != nil {
		return nil, nil, err
	}
	return &c.items[0], copy, nil
}

func (c *Collection[T]) Remove(index int) (*Collection[T], error) {
	if index < 0 || index >= len(c.items) {
		return nil, fmt.Errorf(idxOutOfBoundsErr)
	}
	copy := c.Copy()
	copy.items = append(c.items[:index], c.items[index+1:]...)
	return copy, nil
}

func (c *Collection[T]) Shuffle() *Collection[T] {
	if c.random == nil {
		// if random source was never set then set it manually
		c.random = rand.New(rand.NewSource(int64(rand.Intn(1000))))
	}

	copy := c.Copy()
	for i := 0; i < len(copy.items); i++ {
		r := c.random.Intn(len(c.items))
		if i != r {
			copy.items[r], copy.items[i] = copy.items[i], copy.items[r]
		}
	}
	return copy
}

func (c *Collection[T]) GetItems() []T {
	return c.items
}

func (c *Collection[T]) GetRandomness() *rand.Rand {
	return c.random
}

func (c *Collection[T]) SetRandomness(random *rand.Rand) {
	c.random = random
}

func (c *Collection[T]) GetSize() int {
	return len(c.items)
}

func (c *Collection[T]) Copy() *Collection[T] {
	// because maps anc slices are references they do not get copiec even when passing by value
	// clone creates new slices to mimic passing the slice by value
	return &Collection[T]{
		items:  append(make([]T, 0), c.items...),
		random: c.random,
	}
}

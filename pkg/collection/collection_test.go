package collection

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_Collection(t *testing.T) {
	type card struct {
		suit string
		rank string
	}
	suits := []string{"Diamonds", "Hearts", "Clubs", "Spades"}
	ranks := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	deck := NewCollection[card](0)
	for _, suit := range suits {
		for _, rank := range ranks {
			deck.Add(card{suit: suit, rank: rank})
		}
	}

	assert.True(t, deck.GetSize() == 52, "cards are missing")
	assert.True(t, deck.items[0].suit == "Diamonds" && deck.items[0].rank == "A", "cards in incorrect order")

	deck.Shuffle()

	assert.True(t, !(deck.items[0].suit == "Diamonds" && deck.items[0].rank == "A"), "failed to shuffle")

	c, err := deck.Draw()
	if err != nil {
		t.Fatal(err)
		t.FailNow()
	}

	assert.NotNil(t, c, "card is nil")
	assert.True(t, len(deck.GetItems()) == 51, "cards are missing")

	if err = deck.Remove(0); err != nil {
		t.Fatal(err)
		t.FailNow()
	}
	assert.True(t, deck.GetSize() == 50, "cards are missing")
}

func Test_CollectionRemove(t *testing.T) {
	testCases := []struct {
		name        string
		collection  Collection[int]
		index       int
		shouldError bool
	}{
		{
			name:        "empty collection should error",
			collection:  Collection[int]{},
			index:       0,
			shouldError: true,
		},
		{
			name:        "negative index should error",
			collection:  Collection[int]{items: []int{1, 2, 3}},
			index:       -1,
			shouldError: true,
		},
		{
			name:        "index larger than collection should error",
			collection:  Collection[int]{items: []int{1, 2, 3}},
			index:       10,
			shouldError: true,
		},
		{
			name:        "index in range should not error",
			collection:  Collection[int]{items: []int{1, 2, 3}},
			index:       2,
			shouldError: false,
		},
	}
	for _, test := range testCases {
		if err := test.collection.Remove(test.index); (err != nil) != test.shouldError {
			t.Fatalf(test.name)
			t.FailNow()
		}
	}
}

func Test_CollectionDraw(t *testing.T) {
	testCases := []struct {
		name        string
		collection  Collection[string]
		draw        string
		shouldError bool
	}{
		{
			name:        "draw on empty collection should error",
			collection:  Collection[string]{},
			shouldError: true,
		},
		{
			name:        "draw on non-empty collection should not error",
			collection:  Collection[string]{items: []string{"A", "B", "C"}},
			draw:        "A",
			shouldError: false,
		},
	}
	for _, test := range testCases {
		draw, err := test.collection.Draw()
		if (err != nil) != test.shouldError {
			t.Fatalf(test.name)
			t.FailNow()
		}
		if !test.shouldError {
			assert.Equal(t, test.draw, *draw)
		}
	}
}

func Test_GetItem(t *testing.T) {
	testCases := []struct {
		name        string
		collection  Collection[string]
		index       int
		item        string
		shouldError bool
	}{
		{
			name:        "get item on empty collection should error",
			collection:  Collection[string]{},
			index:       0,
			shouldError: true,
		},
		{
			name:        "get item on non-empty collection should not error",
			collection:  Collection[string]{items: []string{"A", "B", "C"}},
			index:       1,
			item:        "B",
			shouldError: false,
		},
	}
	for _, test := range testCases {
		item, err := test.collection.GetItem(test.index)
		if (err != nil) != test.shouldError {
			t.Fatalf(test.name)
			t.FailNow()
		}
		if !test.shouldError {
			assert.Equal(t, test.item, *item)
		}
	}
}

func Test_CollectionCloning(t *testing.T) {
	c1 := NewCollection[string](0)
	c1.Add("A", "B", "C")

	c2 := c1.Clone()

	if _, err := c1.Draw(); err != nil {
		fmt.Println(err)
	}
	assert.True(t, c1.GetSize() != c2.GetSize())
}

func Test_CollectionRandomness(t *testing.T) {
	seed := time.Now().Unix()
	c1 := NewCollection[int](seed)
	c1.Add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15)

	c2 := NewCollection[int](seed)
	c2.Add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15)

	c1.Shuffle()
	c2.Shuffle()

	for i := 0; i < c1.GetSize(); i++ {
		it1, _ := c1.GetItem(i)
		it2, _ := c1.GetItem(i)
		assert.Equal(t, it1, it2)
	}
}

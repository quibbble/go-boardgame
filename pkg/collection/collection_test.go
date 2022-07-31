package collection

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Collection(t *testing.T) {
	type card struct {
		suit string
		rank string
	}
	suits := []string{"Diamonds", "Hearts", "Clubs", "Spades"}
	ranks := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	deck := &Collection[card]{}
	deck.SetRandomness(rand.New(rand.NewSource(0)))
	for _, suit := range suits {
		for _, rank := range ranks {
			deck = deck.Add(card{suit: suit, rank: rank})
		}
	}

	assert.True(t, deck.GetSize() == 52, "cards are missing")
	assert.True(t, deck.items[0].suit == "Diamonds" && deck.items[0].rank == "A", "cards in incorrect order")

	deck = deck.Shuffle()

	assert.True(t, !(deck.items[0].suit == "Diamonds" && deck.items[0].rank == "A"), "failed to shuffle")

	c, deck, err := deck.Draw()
	if err != nil {
		t.Fatal(err)
		t.FailNow()
	}

	assert.NotNil(t, c, "card is nil")
	assert.True(t, len(deck.GetItems()) == 51, "cards are missing")

	deck, err = deck.Remove(0)
	if err != nil {
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
		_, err := test.collection.Remove(test.index)
		if (err != nil) != test.shouldError {
			t.Fatalf(test.name)
			t.FailNow()
		}
	}
}

func Test_CollectionDraw(t *testing.T) {
	testCases := []struct {
		name        string
		collection  Collection[string]
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
			shouldError: false,
		},
	}
	for _, test := range testCases {
		_, _, err := test.collection.Draw()
		if (err != nil) != test.shouldError {
			t.Fatalf(test.name)
			t.FailNow()
		}
	}
}

func Test_CollectionCopying(t *testing.T) {
	c1 := &Collection[string]{}
	c1 = c1.Add("A", "B", "C")
	c1.SetRandomness(rand.New(rand.NewSource(0)))

	_, c2, err := c1.Draw()
	if err != nil {
		fmt.Println(err)
	}
	assert.True(t, c1.GetSize() != c2.GetSize())

	c3, _ := c1.Remove(1)
	assert.True(t, c1.GetSize() != c3.GetSize())

	c4 := c1.Add("D")
	assert.True(t, c1.GetSize() != c4.GetSize())

	c5 := c1.Shuffle()
	assert.True(t, c1.GetItems()[0] != c5.GetItems()[0])
}

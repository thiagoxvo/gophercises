//go:generate stringer -type=Suit,Rank

package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Suit represents the card suit
type Suit uint8

// All card suits plus the Joker
const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

var suits = [...]Suit{Spade, Diamond, Club, Heart}

// Rank represents the card rank
type Rank uint8

// All card ranks
const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

const (
	minRank = Ace
	maxRank = King
)

// Card contains a suit and a rank which represent a game card
type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank, c.Suit)
}

// New creates a new card deck
func New(opts ...func([]Card) []Card) []Card {
	var cards []Card
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
		for _, opt := range opts {
			cards = opt(cards)
		}
	}
	return cards
}

// Sort is the function which can receive a less function to sort cards
func Sort(less func(cards []Card) func(i, j int) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

// DefaultSort sorts the deck of cards
func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, less(cards))
	return cards
}

func less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}

func absRank(c Card) int {
	return int(c.Suit)*int(maxRank) + int(c.Rank)
}

// Shuffle shuffles the deck of cards
func Shuffle(cards []Card) []Card {
	ret := make([]Card, len(cards))
	r := rand.New(rand.NewSource(time.Now().Unix()))
	perm := r.Perm(len(cards))
	for i, j := range perm {
		ret[i] = cards[j]
	}
	return ret
}

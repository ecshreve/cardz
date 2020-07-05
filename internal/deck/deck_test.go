package deck_test

import (
	"fmt"
	"testing"

	"github.com/samsarahq/go/snapshotter"
	"github.com/stretchr/testify/assert"

	"github.com/ecshreve/cardz/internal/deck"
)

func TestNewDeck(t *testing.T) {
	snap := snapshotter.New(t)
	defer snap.Verify()

	testDeck := deck.NewDeck()
	assert.Equal(t, 52, len(testDeck.Cards))
	assert.Equal(t, 0, len(testDeck.Dealt))
	snap.Snapshot("basic new deck", testDeck)
}

func TestCards(t *testing.T) {
	snap := snapshotter.New(t)
	defer snap.Verify()

	testDeck := deck.NewDeck()
	var cardStrs []string
	for _, c := range testDeck.Cards {
		cardStrs = append(cardStrs, fmt.Sprint(c))
	}

	snap.Snapshot("card strings", cardStrs)
}

func TestDealOne(t *testing.T) {
	testDeck := deck.NewDeck()

	// Deal one Card and verify expected behavior.
	card, err := testDeck.DealOne()
	assert.NotNil(t, card)
	assert.NoError(t, err)
	assert.Equal(t, 51, len(testDeck.Cards))
	assert.Equal(t, 1, len(testDeck.Dealt))

	// Deal the remaining 51 Cards and verify no error.
	for i := 1; i <= 51; i++ {
		card, err := testDeck.DealOne()
		assert.NotNil(t, card)
		assert.NoError(t, err)
	}

	// Try to deal a card with an empty Deck and verify error.
	card, err = testDeck.DealOne()
	assert.Nil(t, card)
	assert.Error(t, err)
	assert.Equal(t, 0, len(testDeck.Cards))
	assert.Equal(t, 52, len(testDeck.Dealt))
}

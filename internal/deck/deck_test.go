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

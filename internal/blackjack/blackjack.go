package blackjack

import (
	"fmt"

	"github.com/ecshreve/cardz/internal/deck"
)

// Play contains the main game loop for this silly blackjack game.
func Play() {
	d := deck.NewDeck()
	d.ShuffleRemaining()

	player := &Player{IsDealer: false}
	dealer := &Player{IsDealer: true}
	continueGame := true

	// Loop until the player decides they don't want to play anymore.
	for continueGame {
		if len(d.Cards) < 26 {
			d.ReShuffle()
		}

		deal(player, dealer, d)
		fmt.Println("dealer: ")
		fmt.Println(dealer.Hand)
		fmt.Println("---")
		fmt.Println("player: ")
		fmt.Println(player.Hand)

		player.takeTurn(d)
		if !player.Bust {
			dealer.takeTurn(d)
		}

		winner := getWinner(player, dealer)
		continueGame = winner.handleHandEnd()
	}

	fmt.Println("thanks for playing!")
}

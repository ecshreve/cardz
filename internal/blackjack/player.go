package blackjack

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ecshreve/cardz/internal/deck"
)

// Player represents either a human player or the CPU dealer.
type Player struct {
	Hand
	IsDealer bool
}

// dealerTurn takes care of the logic for the Dealer's hitting/standing.
func (p *Player) dealerTurn(d *deck.Deck) {
	for p.Hand.Total < 17 {
		c, _ := d.DealOne()
		p.addCard(*c)
		fmt.Printf("---\ndealer:\n")
		fmt.Println(p.Hand)
	}
	if p.Bust {
		fmt.Println("dealer busted :)")
	}
}

// takeTurn contains the main hit/stand loop for the player.
func (p *Player) takeTurn(d *deck.Deck) {
	if p.IsDealer {
		p.dealerTurn(d)
		return
	}

	for {
		fmt.Println("press 'h' to hit, or 's' to stand; followed by 'enter'")
		reader := bufio.NewReader(os.Stdin)
		char, _, err := reader.ReadRune()
		if err != nil {
			fmt.Println(err)
		}

		switch char {
		case 'h':
			c, _ := d.DealOne()
			p.addCard(*c)
			fmt.Printf("---\nplayer:\n")
			fmt.Println(p.Hand)
			if p.Bust {
				fmt.Println("busted :(")
				return
			}
			break
		case 's':
			fmt.Printf("player stands at %d\n", p.Hand.Total)
			return
		}
	}
}

// handleHandEnd prints a game over message and checks if the player wants to continue.
func (p Player) handleHandEnd() bool {
	retStr := "yay you won!"
	if p.IsDealer {
		retStr = "better luck next time!"
	}
	fmt.Println(retStr)
	return continueGame()
}

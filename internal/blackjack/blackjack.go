package blackjack

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ecshreve/cardz/internal/deck"
)

// Hand store the Player's Cards and some other useful data.
type Hand struct {
	Cards []deck.Card
	Total int
	//HasAce    bool
	//Blackjack bool
	Bust bool
}

// Implement the Stringer interface for the Hand type.
func (h Hand) String() string {
	return fmt.Sprintf("total: %d -- bust: %v\ncards:\n%+v\n", h.Total, h.Bust, h.Cards)
}

// addCard adds a Card to the Player's hand and recalculates the Hand total, also
// updates the Bust field depending on the Total.
func (h *Hand) addCard(c deck.Card) {
	h.Cards = append(h.Cards, c)
	h.Total += c.Value
	h.Bust = h.Total > 21
}

// Player represents either a human player or the CPU dealer.
type Player struct {
	Hand
	IsDealer bool
}

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
		continueGame = handleHandEnd(winner)
	}

	fmt.Println("thanks for playing!")
}

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

func (p *Player) takeTurn(d *deck.Deck) {
	if p.IsDealer {
		p.dealerTurn(d)
		return
	}

	// This is the Hit/Stand loop for the player.
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

// deal clears the player and dealer Hands and deals the first two Cards.
func deal(player, dealer *Player, d *deck.Deck) {
	player.Hand = Hand{}
	dealer.Hand = Hand{}

	c, _ := d.DealOne()
	player.Hand.addCard(*c)

	c, _ = d.DealOne()
	dealer.Hand.addCard(*c)

	c, _ = d.DealOne()
	player.Hand.addCard(*c)

	c, _ = d.DealOne()
	dealer.Hand.addCard(*c)
}

// getWinner returns the winning Player, or nil in the case of a push.
func getWinner(player, dealer *Player) *Player {
	if player.Bust {
		return dealer
	}

	if dealer.Bust {
		return player
	}

	if player.Total > dealer.Total {
		return player
	}

	if dealer.Total > player.Total {
		return dealer
	}

	return nil
}

// handleHandEnd prints a game over message and checks if the player wants to continue.
func handleHandEnd(p *Player) bool {
	retStr := "yay you won!"
	if p.IsDealer {
		retStr = "better luck next time!"
	}
	fmt.Println(retStr)
	return continueGame()
}

// continueGame prompts the user to see if they want to play another hand.
func continueGame() bool {
	fmt.Println("do you want to play another hand? (y/n) followed by 'enter'")

	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err)
	}

	return char == 'y'
}

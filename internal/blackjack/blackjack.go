package blackjack

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ecshreve/cardz/internal/deck"
)

type Hand struct {
	Cards []deck.Card
	Total int
	//HasAce    bool
	//Blackjack bool
	Bust bool
}

func (h Hand) String() string {
	return fmt.Sprintf("total: %d -- bust: %v\ncards:\n%+v\n", h.Total, h.Bust, h.Cards)
}

func (h *Hand) AddCard(c deck.Card) {
	h.Cards = append(h.Cards, c)
	h.Total += c.Value
	h.Bust = h.Total > 21
}

type Player struct {
	Hand
	IsDealer bool
}

func Play() {
	d := deck.NewDeck()
	d.ShuffleRemaining()

	player := &Player{}
	dealer := &Player{IsDealer: true}

	continueGame := true
	for continueGame {
		if len(d.Cards) < 26 {
			d.ReShuffle()
		}
		player.Hand = Hand{}
		dealer.Hand = Hand{}

		// Deal out the initial 2 cards.
		c, _ := d.DealOne()
		player.Hand.AddCard(*c)

		c, _ = d.DealOne()
		dealer.Hand.AddCard(*c)

		c, _ = d.DealOne()
		player.Hand.AddCard(*c)

		c, _ = d.DealOne()
		dealer.Hand.AddCard(*c)

		fmt.Println("player")
		fmt.Println(player.Hand)
		fmt.Println("dealer")
		fmt.Println(dealer.Hand)

		// player's turn
		player.TakeTurn(d)
		if player.Bust {
			continueGame = ContinueGame()
			continue
		}

		// dealer's turn
		dealer.TakeTurn(d)
		if dealer.Bust {
			continueGame = ContinueGame()
			continue
		}

		if player.Total > dealer.Total {
			fmt.Println("yay you won!")
		} else {
			fmt.Println("you lost :(")
		}

		continueGame = ContinueGame()
	}
}

func (p *Player) TakeTurn(d *deck.Deck) {
	if p.IsDealer {
		for p.Hand.Total < 17 {
			c, _ := d.DealOne()
			fmt.Println(c)
			p.AddCard(*c)
			fmt.Println(p.Hand)
		}
		if p.Bust {
			fmt.Println("dealer busted :)")
		}
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
			fmt.Println(c)
			p.AddCard(*c)
			fmt.Println(p.Hand)
			if p.Bust {
				fmt.Println("player busted :(")
				return
			}
			break
		case 's':
			fmt.Printf("player stands at %d\n", p.Hand.Total)
			return
		}
	}
}

func ContinueGame() bool {
	fmt.Println("do you want to play another hand? (y/n) followed by 'enter'")
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()

	if err != nil {
		fmt.Println(err)
	}

	return char == 'y'
}

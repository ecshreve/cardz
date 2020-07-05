package main

import (
	"fmt"

	"github.com/ecshreve/cardz/internal/deck"
)

func main() {
	d := deck.NewDeck()
	c := d.Cards[11]
	fmt.Println(c.PrettyPrint())
}

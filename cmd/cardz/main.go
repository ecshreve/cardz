package main

import (
	"github.com/kr/pretty"

	"github.com/ecshreve/cardz/internal/deck"
)

func main() {
	d := deck.NewDeck()
	c, _ := d.DealOne()
	pretty.Print(c)
}

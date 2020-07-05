package main

import (
	"fmt"

	"github.com/ecshreve/cardz/internal/deck"
)

func main() {
	d := deck.NewDeck()
	c, _ := d.DealOne()
	fmt.Println(c)
}

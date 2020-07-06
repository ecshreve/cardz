package main

import (
	"fmt"

	"github.com/kr/pretty"

	"github.com/ecshreve/cardz/internal/blackjack"
)

//"github.com/ecshreve/cardz/internal/blackjack"

func main() {
	//blackjack.StartGame()
	st := &blackjack.Stats{}
	st.Total = 69
	fmt.Println(st)
	st.Save()
	st2 := &blackjack.Stats{}
	st2.Load()
	pretty.Print(st2)

}

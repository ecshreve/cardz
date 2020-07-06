package blackjack

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Stats struct {
	Win   int
	Loss  int
	Tie   int
	Total int
}

func (st Stats) String() string {
	return fmt.Sprintf(`W:   %d
	 					L:   %d
	 					T:   %d
	 					TOT: %d`, st.Win, st.Loss, st.Tie, st.Total)
}

func LoadStats() *Stats {
	f, err := os.Open("persist.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var st Stats
	byteValue, _ := ioutil.ReadAll(f)
	json.Unmarshal(byteValue, &st)
	return &st
}

func (st *Stats) Save() {
	jsonStats, err := json.Marshal(st)
	if err != nil {
		log.Fatal(err)
	}

	dir, _ := os.Getwd()
	fname := dir + "/persist.json"

	err = ioutil.WriteFile(fname, jsonStats, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func (bg *Game) UpdateStats() {
	bg.Stats.Total++
	if bg.Winner == nil {
		bg.Stats.Tie++
		return
	}
	if bg.Winner.IsDealer {
		bg.Stats.Loss++
		return
	}
	bg.Stats.Win++
}

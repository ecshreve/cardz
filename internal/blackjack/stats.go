package blackjack

import (
	"encoding/json"
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

func (st *Stats) Load() {
	f, err := os.Open("persist.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	byteValue, _ := ioutil.ReadAll(f)
	json.Unmarshal(byteValue, &st)
}

func (st *Stats) Save() {
	jsonStats, err := json.Marshal(st)
	if err != nil {
		log.Fatal(err)
	}

	dir, _ := os.Getwd()
	fname := dir + "/persist.json"

	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile(fname, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write(append(jsonStats, []byte("\n")...)); err != nil {
		f.Close() // ignore error; Write error takes precedence
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
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

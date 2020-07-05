package main

// import (
// 	"github.com/ecshreve/cardz/internal/deck"
// 	"github.com/gdamore/tcell"
// 	"github.com/rivo/tview"
// )

// var customTheme = tview.Theme{
// 	PrimitiveBackgroundColor:    tcell.Color(272727),
// 	ContrastBackgroundColor:     tcell.Color(448488),
// 	MoreContrastBackgroundColor: tcell.ColorGreen,
// 	BorderColor:                 tcell.ColorWhite,
// 	TitleColor:                  tcell.ColorWhite,
// 	GraphicsColor:               tcell.ColorWhite,
// 	PrimaryTextColor:            tcell.ColorWhite,
// 	SecondaryTextColor:          tcell.ColorYellow,
// 	TertiaryTextColor:           tcell.ColorGreen,
// 	InverseTextColor:            tcell.Color(448488),
// 	ContrastSecondaryTextColor:  tcell.ColorDarkCyan,
// }

// func main() {
// 	d := deck.NewDeck()
// 	app := tview.NewApplication()
// 	tview.Styles = customTheme
// 	statsArea := tview.NewTextView().SetText("test").SetTextAlign(1)
// 	dealerArea := tview.NewTextView().SetWrap(false).SetText(d.Cards[7].PrettyPrint()).SetTextAlign(1)
// 	playerArea := tview.NewTextView().SetText(d.Cards[17].PrettyPrint()).SetTextAlign(1)

// 	statsArea.SetBorder(true).SetTitle("stats").SetBorderPadding(0, 1, 1, 1)
// 	dealerArea.SetBorder(true).SetTitle("dealer").SetBorderPadding(0, 1, 1, 1)
// 	playerArea.SetBorder(true).SetTitle("player").SetBorderPadding(1, 1, 1, 1)

// 	flex := tview.NewFlex().
// 		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
// 			AddItem(statsArea, 0, 1, false).
// 			AddItem(dealerArea, 0, 2, false).
// 			AddItem(playerArea, 0, 2, false), 0, 5, false).
// 		AddItem(tview.NewBox().SetBorder(true).SetTitle("History"), 20, 1, false)
// 	flex.SetBackgroundColor(tview.Styles.PrimitiveBackgroundColor)

// 	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
// 		panic(err)
// 	}
// }

import (
	"github.com/ecshreve/cardz/internal/blackjack"
)

func main() {
	blackjack.Play()
}

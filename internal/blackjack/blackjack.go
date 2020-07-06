package blackjack

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"

	"github.com/ecshreve/cardz/internal/deck"
)

// These are declared here because they need to be accessible by both goroutines
// and it was the quickes and easiest way to achieve that.
var (
	app        *tview.Application
	dealerFlex *tview.Flex
	playerFlex *tview.Flex
	statusFlex *tview.Flex
	infoFlex   *tview.Flex
	hitButton  *tview.Button
	stayButton *tview.Button
	yesButton  *tview.Button
	noButton   *tview.Button
)

// These three channels are used to coordinate exeuction between the two goroutines.
var (
	continuePlaying = make(chan bool)
	playerHit       = make(chan bool)
	hasUpdate       = make(chan bool)
)

// Player represents either a human player or the CPU dealer.
type Player struct {
	Hand
	IsDealer bool
}

// Game holds data and configuration variables for the game.
type Game struct {
	Player          *Player
	Dealer          *Player
	Deck            *deck.Deck
	PlayerTurn      bool
	HandComplete    bool
	ContinuePlaying bool
}

func update(bg *Game) {
	for {
		<-hasUpdate
		app.QueueUpdateDraw(func() {
			// Any time we receive a message on the hasUpdate channel we update
			// the flex containers that hold the Dealer and Player Hands.
			playerFlex.Clear()
			playerScore := tview.NewTextView().SetText(strconv.Itoa(bg.Player.Hand.Total)).SetTextAlign(1)
			playerScore.SetBorder(true).SetBorderPadding(2, 2, 0, 0)
			playerScore.SetBorderAttributes(tcell.AttrBlink).SetBorderColor(tcell.ColorLightGreen)

			dealerFlex.Clear()
			dealerScore := tview.NewTextView().SetText(strconv.Itoa(bg.Dealer.Hand.Total)).SetTextAlign(1)
			dealerScore.SetBorder(true).SetBorderPadding(2, 2, 0, 0)

			// Add the score box and Cards to the Player and Dealer flex containers.
			playerFlex.AddItem(playerScore, 10, 0, false)
			for _, card := range bg.Player.Hand.Cards {
				playerArea := tview.NewTextView().SetText(card.PrettyPrint()).SetTextAlign(1)
				playerFlex.AddItem(playerArea, 0, 1, false)
			}

			dealerFlex.AddItem(dealerScore, 10, 1, false)
			for _, card := range bg.Dealer.Hand.Cards {
				dealerArea := tview.NewTextView().SetText(card.PrettyPrint()).SetTextAlign(1)
				dealerFlex.AddItem(dealerArea, 0, 1, false)
			}

			// Update the status box on the Player's turn so the Player can
			// choose to hit or stay.
			if bg.PlayerTurn {
				statusFlex.Clear()
				hitButton = tview.NewButton("HIT")
				hitButton.SetBackgroundColorActivated(tcell.ColorRed)
				hitButton.SetBorder(true)
				hitButton.SetSelectedFunc(func() {
					playerHit <- true
				})

				stayButton = tview.NewButton("STAY")
				stayButton.SetBackgroundColorActivated(tcell.ColorGreen)
				stayButton.SetBorder(true)
				stayButton.SetSelectedFunc(func() {
					playerHit <- false
				})

				statusFlex.AddItem(hitButton, 0, 1, false)
				statusFlex.AddItem(stayButton, 0, 1, true)
				app.SetFocus(statusFlex)
			}

			// Update the status box when the Hand is completed so the Player can
			// choose to play another Hand or not.
			if bg.HandComplete {
				winner := bg.getWinner()

				verb := "won"
				if winner == nil {
					verb = "tied"
				} else if winner.IsDealer {
					verb = "lost"
				}

				resultStr := "you " + strings.ToUpper(verb) + " do you want to play another hand?"
				gameOver := tview.NewTextView().SetText(resultStr).SetTextAlign(1)
				gameOver.SetBorderPadding(1, 1, 2, 2)
				gameOver.SetWrap(true).SetWordWrap(true)

				yesButton = tview.NewButton("yes")
				yesButton.SetBackgroundColorActivated(tcell.ColorGreen)
				yesButton.SetBorder(true)
				yesButton.SetSelectedFunc(func() {
					continuePlaying <- true
				})

				noButton = tview.NewButton("no")
				noButton.SetBackgroundColorActivated(tcell.ColorRed)
				noButton.SetBorder(true)
				noButton.SetSelectedFunc(func() {
					continuePlaying <- false
					app.Stop()
				})

				statusFlex.Clear()
				statusFlex.AddItem(gameOver, 0, 2, false)
				statusFlex.AddItem(yesButton, 0, 1, true)
				statusFlex.AddItem(noButton, 0, 1, false)
				app.SetFocus(statusFlex)
			}
		})
	}
}

// Play contains the main game loop for this silly blackjack game.
func Play(bg *Game) {
	// Make sure we shuffle the deck before starting this Hand.
	bg.Deck.ShuffleRemaining()

	for bg.ContinuePlaying {
		// If there's less than half the cards left in the Deck then reshuffle.
		if len(bg.Deck.Cards) < 26 {
			bg.Deck.ReShuffle()
		}

		// Deal the initial Hands to the Player and Dealer and send a message
		// to the channel so we update the UI.
		bg.deal()
		bg.PlayerTurn = true
		hasUpdate <- true

		// Process the Player's turn.
		bg.PlayerTurn = <-playerHit
		for bg.PlayerTurn {
			// Deal a Card to the Player and update the UI.
			c, err := bg.Deck.DealOne()
			if err != nil {
				fmt.Println(err)
			}
			bg.Player.addCard(*c)
			hasUpdate <- true

			// If the Player busted then we can early exit and skip the Dealer loop.
			if bg.Player.Bust {
				bg.PlayerTurn = false
				break
			}

			// Wait here until we get a response from the Player.
			bg.PlayerTurn = <-playerHit
		}

		// If the Player busted then we don't need to process the Dealer's turn.
		if !bg.Player.Bust {
			for bg.Dealer.Hand.Total < 17 {
				// Deal a Card to the Dealer, update the UI and wait 3 seconds
				// before Dealer makes their next move.
				c, err := bg.Deck.DealOne()
				if err != nil {
					fmt.Println(err)
				}
				bg.Dealer.addCard(*c)
				hasUpdate <- true
			}
		}

		// Set the HandComplete Field on the Game and send a message on the hasUpdate
		// channel so we can update the status box can update to let the Player
		// choose to play another hand or not.
		bg.HandComplete = true
		hasUpdate <- true

		// Set the Game's ContinuePlaying field based on which button the Player
		// selected when asked to play another hand.
		bg.ContinuePlaying = <-continuePlaying
	}
}

// StartGame initializes a Game, sets up the tview app, and calls the
// update and play functions in goroutines.
func StartGame() {
	bg := &Game{
		Player: &Player{
			Hand:     Hand{},
			IsDealer: false,
		},
		Dealer: &Player{
			Hand:     Hand{},
			IsDealer: true,
		},
		Deck:            deck.NewDeck(),
		PlayerTurn:      false,
		HandComplete:    false,
		ContinuePlaying: true,
	}

	app = tview.NewApplication()
	tview.Styles = customCliTheme

	// Initialize the flex boxes for our content.
	dealerFlex = tview.NewFlex().SetDirection(tview.FlexColumn)
	playerFlex = tview.NewFlex().SetDirection(tview.FlexColumn)
	statusFlex = tview.NewFlex().SetDirection(tview.FlexColumn)
	infoFlex = tview.NewFlex().SetDirection(tview.FlexRow)

	dealerFlex.SetBorder(true).SetTitle(" dealer ").SetBorderPadding(1, 1, 1, 1)
	playerFlex.SetBorder(true).SetTitle(" player ").SetBorderPadding(1, 1, 1, 1)
	statusFlex.SetBorder(true)
	infoFlex.SetBorder(true).SetBorderPadding(0, 0, 1, 1)

	statsText := tview.NewTextView().SetText("these will be stats").SetTextAlign(0).SetWordWrap(true)
	statsText.SetBorder(true).SetBorderPadding(1, 1, 1, 1).SetTitle(" stats ").SetBorderAttributes(tcell.AttrBlink)
	historyText := tview.NewTextView().SetText("this will be history").SetTextAlign(0).SetWordWrap(true)
	historyText.SetBorder(true).SetBorderPadding(1, 1, 1, 1).SetTitle(" history ")

	infoFlex.AddItem(statsText, 0, 2, false)
	infoFlex.AddItem(historyText, 0, 2, false)

	// Attach our flex boxes to the outer flex container.
	flex := tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(dealerFlex, 0, 2, false).
			AddItem(playerFlex, 0, 2, false).
			AddItem(statusFlex, 0, 1, false), 0, 5, false).
		AddItem(infoFlex, 25, 1, false)
	flex.SetBackgroundColor(tview.Styles.PrimitiveBackgroundColor)

	// Explicitly configure input handling for these specific Keys. tview has some
	// quirks when it comes to setting Focus on child elements on flex boxes, and
	// this is the workaround.
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyTAB || event.Key() == tcell.KeyLeft || event.Key() == tcell.KeyRight {
			switch {
			case hitButton.HasFocus():
				app.SetFocus(stayButton)
			case stayButton.HasFocus():
				app.SetFocus(hitButton)
			case yesButton.HasFocus():
				app.SetFocus(noButton)
			case noButton.HasFocus():
				app.SetFocus(yesButton)
			}
		}

		if event.Key() == tcell.KeyRune {
			switch event.Rune() {
			case 'q':
				app.Stop()
				os.Exit(0)
			}
		}
		return event
	})

	// Call these functions in goroutines so we can asynchronously update the UI
	// whenever needed.
	go update(bg)
	go Play(bg)

	if err := app.SetRoot(flex, true).SetFocus(statusFlex).Run(); err != nil {
		panic(err)
	}
}

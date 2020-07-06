package deck

var symbolMap = map[Suit]string{
	Hearts:   "♥",
	Diamonds: "♦",
	Clubs:    "♣",
	Spades:   "♠",
}

// CardCodes are all the possible single char codes.
var CardCodes = []string{
	"A",
	"2",
	"3",
	"4",
	"5",
	"6",
	"7",
	"8",
	"9",
	"T",
	"J",
	"Q",
	"K",
}

// CardCodeToValueMap maps a Card.Code to a Card's Value.
var CardCodeToValueMap = map[string]int{
	"A": 11,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 10,
	"Q": 10,
	"K": 10,
}

// CardCodeToNameMap maps a Card.code to a Card's string name.
var CardCodeToNameMap = map[string]string{
	"A": "ace",
	"2": "two",
	"3": "three",
	"4": "four",
	"5": "five",
	"6": "six",
	"7": "seven",
	"8": "eight",
	"9": "nine",
	"T": "ten",
	"J": "jack",
	"Q": "queen",
	"K": "king",
}

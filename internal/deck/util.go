package deck

import "fmt"

// PrettyPrint returns a string form of the card that looks like this:
// ┌─────────┐
// │ A       │
// │         │
// │    ♠    │
// │         │
// │       A │
// └─────────┘
func (c Card) PrettyPrint() string {
	r1 := "┌─────────┐\n"
	r2 := fmt.Sprintf("│ %s       │\n", c.Code)
	r3 := "│         │\n"
	r4 := fmt.Sprintf("│    %s    │\n", symbolMap[c.Suit])
	r5 := "│         │\n"
	r6 := fmt.Sprintf("│       %s │\n", c.Code)
	r7 := "└─────────┘"

	// Special handling for the 10 because it's the only one that's 2 chars.
	if c.Code == "T" {
		r2 = "│ 10      │\n"
		r6 = "│      10 │\n"
	}
	return r1 + r2 + r3 + r4 + r5 + r6 + r7
}

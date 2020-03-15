package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Two, Suit: Spade})
	//Output:
	//Ace of Hearts
	//Two of Spades
}

func TestNew(t *testing.T) {
	cards := New()
	//13 ranks * Suits
	if len(cards) != 13*4 {
		t.Error("Wrong number of cards in the deck.")
	}
}

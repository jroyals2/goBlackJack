package deck

import "fmt"
import "math/rand"

var suits = [4]string{"Diamonds", "Spades", "Hearts", "Clubs"}
var values = [13]string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

// Deck is thge main function exporting my deck making ability
func Deck() {

	var startDeck []string
	var pHand []string

	for i := 0; i < len(suits); i++ {
		for j := 0; j < len(values); j++ {
			card := values[j] + " of " + suits[i] + "\n"
			startDeck = append(startDeck, card)
		}
	}

	pHand = append(pHand, startDeck[rand.Intn(51)])
	pHand = append(pHand, startDeck[rand.Intn(51)])
	fmt.Println(startDeck)
	fmt.Println(pHand)

}

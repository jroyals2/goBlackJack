package deck

import "fmt"
import "crypto/rand"
import "math/big"

var suits = [4]string{"Diamonds", "Spades", "Hearts", "Clubs"}
var values = [13]string{ "Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

// Deck is thge main function exporting my deck making ability
func Deck() {

	var startDeck []string
	var pHand []string
	var dHand []string

	for i := 0; i < len(suits); i++ {
		for j := 0; j < len(values); j++ {
			card := values[j] + " of " + suits[i] + "\n"
			startDeck = append(startDeck, card)
		}
	}

	pHand = append(pHand, startDeck[getSecureRandInt(51).Int64()])
	pHand = append(pHand, startDeck[getSecureRandInt(51).Int64()])
	dHand = append(dHand, startDeck[getSecureRandInt(51).Int64()])
	dHand = append(dHand, startDeck[getSecureRandInt(51).Int64()])
	//pHand = append(pHand, startDeck[rand.Intn(51)])
	//fmt.Println(startDeck)
	fmt.Println("The Player has", pHand)
	fmt.Println("The Dealer has ", dHand)
}
func getSecureRandInt(max int64) *big.Int {
	nBig, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		panic(err)
	}
	return nBig
}

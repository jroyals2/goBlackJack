package deck

import "fmt"
import "crypto/rand"
import "math/big"

var suits = [4]string{"Diamonds", "Spades", "Hearts", "Clubs"}
var values = [13]string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
var value = [13]int{11, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10}

//var startDeck = map[string]*decks{}

var startDeck []string
var pHand []string
var dHand []string

type decks struct {
	suit      string
	psuedoVal string
	val       int
}

// Deck is thge main function exporting my deck making ability
func Deck() {

	for i := 0; i < len(suits); i++ {
		for j := 0; j < len(values); j++ {
			newCard := decks{
				suit:      suits[i],
				psuedoVal: values[j],
				val:       value[j],
			}
			//card := values[j] + " of " + suits[i] + "\n"
			startDeck = append(startDeck, newCard.psuedoVal+" of "+newCard.suit+string(newCard.val))
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
	playerTurn()

}
func getSecureRandInt(max int64) *big.Int {
	nBig, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		panic(err)
	}
	return nBig
}
func playerTurn() {
	var option int
	fmt.Println("What would you like to do? \n 1.Hit   2.Stand")
	fmt.Scanln(&option)
	if option == 1 {
		fmt.Println("You chose to hit")
		pHand = append(pHand, startDeck[getSecureRandInt(51).Int64()])
		fmt.Println("You added a card. Your new hand is ", pHand)
		playerTurn()
	} else if option == 2 {
		fmt.Println("You chose to stand")
		dealerTurn()
	} else {
		fmt.Println("You done messed up A-ARON pick again")
		playerTurn()
	}

}

func dealerTurn() {
	for len(dHand) < 4 {
		dHand = append(dHand, startDeck[getSecureRandInt(51).Int64()])
		fmt.Println("The Dealer has \n", dHand)
	}
}

func bust() {
	//var totalPoints int
	for i := 0; i < len(pHand); i++ {

		println(pHand[i])
	}
}

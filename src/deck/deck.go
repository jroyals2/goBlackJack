package deck


type Card struct {
	Suit string
	Value int
	Type string
}

func (card Card) CreateCard(s string, v int, t string){
	

}


// Deck is thge main function exporting my deck making ability
func Deck(message string) {
	println(message)
}
package main

func newCard() string {
	return "Five of Diamonds"
}

func main() {

	cards := deck{"Ace of Diamonds", newCard()}

	// Append does not modify existing DS but instead creates a new DS and adds the value to it
	cards = append(cards, "Six of Spades")

	//fmt.Println(card)

	cards.print()
}

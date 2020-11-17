package main

func main() {

	//cards := newDeckFromFile("my")
	//fmt.Println(cards.toString())
	//cards.saveToFile("my_cards")

	cards := newDeck()
	cards.shuffle()
	cards.print()

}

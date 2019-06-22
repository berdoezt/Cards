package main

import "log"

func main() {
	cards := newDeck()

	hand, remainingCards := deal(cards, 3)
	hand.print()
	remainingCards.print()

	save := cards.saveToFile("file3")

	if save == nil {
		log.Println("sukses")
	} else {
		log.Println(save)
	}
}

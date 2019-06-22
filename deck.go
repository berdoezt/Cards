package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}

	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, value := range cardValues {
		for _, suite := range cardSuites {
			cards = append(cards, value+" of "+suite)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0644)
}

func newDeckFromFile(fileName string) deck {
	byteData, err := ioutil.ReadFile(fileName)

	if err != nil {
		log.Fatal(err)
	}

	stringData := string(byteData)
	splitString := strings.Split(stringData, ",")

	return deck(splitString)
}

func (d deck) shuffle() deck {
	time := time.Now()
	source := rand.NewSource(time.UnixNano())
	r := rand.New(source)

	for i := range d {
		num := r.Intn(len(d) - 1)
		d[i], d[num] = d[num], d[i]
	}

	return d
}

package main

// Create a new type of deck
// which is a slice of string
import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string
type deck1 []int

func newDeck() deck {

	cards := deck{}

	cardSuits := []string{"Spade", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {

		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards

}

// below is the example of func with receiver
func (d deck) print() {

	for _, card := range d {
		fmt.Println(card)
	}
}

func (d deck1) print() {
	fmt.Println("inside deck1 ")
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//we have to declare the receiver just after the func if want to receive any argument

func deal(d deck, handSize int) (deck, deck) {

	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {

	return os.WriteFile(filename, []byte(d.toString()), 0666)

}

func newDeckFromFile(fileName string) deck {
	bs, err := os.ReadFile(fileName)

	if err != nil {
		//Option 1 - log the error and return a call  to newDeck
		//Option 2 - log the error and entirely quit the program
		fmt.Println("Error::", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]

	}
}

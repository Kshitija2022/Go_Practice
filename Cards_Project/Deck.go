package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardsuits := []string{"Diamonds", "Heart", "Spades", "Club"}
	cardvalues := []string{"Ace", "Two", "Three", "Four"}

	for _, suits := range cardsuits {
		for _, values := range cardvalues {
			cards = append(cards, values+" of "+suits)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
func deal(d deck, HandSize int) (deck, deck) {
	return d[:HandSize], d[HandSize:]
}
func (d deck) toString() string {
	return strings.Join([]string(d), ",")

}
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
func newDeckfromfile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:=", err)
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

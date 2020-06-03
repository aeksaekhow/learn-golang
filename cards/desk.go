package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	d := deck{}

	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuites {
		for _, value := range cardValues {
			d = append(d, value+" of "+suit)
		}
	}

	return d
}

func newDeckFromFile(filename string) (deck, error) {
	contentBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	contentStr := string(contentBytes)
	d := deck(strings.Split(contentStr, ","))
	return d, nil
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) saveToFile(filename string) error {
	contentStr := strings.Join(d, ",")
	contentBytes := []byte(contentStr)
	err := ioutil.WriteFile(filename, contentBytes, 0644)
	return err
}

func (d deck) shuffle() {
	seed := time.Now().Unix()
	rand.Seed(seed)
	for i := range d {
		j := rand.Intn(len(d))
		d[i], d[j] = d[j], d[i]
	}
}

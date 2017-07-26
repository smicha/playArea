package main

import (
"fmt"
"math/rand"
"time"
)

const (
	numberOfCardsPerSuit = 13
    numberOfSuits = 4
)


var  (
	suit = [numberOfSuits]string{"Hearts", "Diamonds", "Clubs", "Spades"}

    card = [numberOfCardsPerSuit]string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
)

func shuffle(deck []string) []string {

	shuffled := make([]string, len(deck))

	rand.Seed(time.Now().UTC().UnixNano())
	perm := rand.Perm(len(deck))

	for i, v := range perm {
		shuffled[v] = deck[i]
	}

	return shuffled
}


func main() {

	var initDeck []string

	//initialize   deck of cards
	//inefficient bubble  sorting :)
	for r := 0; r <= numberOfCardsPerSuit-1; r++ {
		for c := 0; c <= numberOfSuits-1; c++ {
			tmp := card[r] + " of " + suit[c]
			initDeck = append(initDeck, tmp)

		}
	}
	//dump sorted deck
 for k, v := range initDeck {
		fmt.Printf("%d %s\n", k, v)
	}

	fmt.Println("-----shuffling deck------")
	shuffledDeck := shuffle(initDeck)

	//  dump shuffled deck!
	for k, v := range shuffledDeck {
		fmt.Printf("%d %s\n", k, v)
	}


}

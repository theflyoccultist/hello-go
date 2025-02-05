package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Suit int
type Pip int

const (
	club Suit = iota
	diamond
	heart
	spade
)

type Card struct {
	s   Suit
	pip Pip
}

func shuffle(d *[]Card) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(*d), func(i, j int) {
		(*d)[i], (*d)[j] = (*d)[j], (*d)[i]
	})
}

func isStraightFlush(h []Card) bool {
	var ccount, dcount, hcount, scount int
	var sameSuitCards []Card

	for _, v := range h {
		switch v.s {
		case club:
			ccount++
		case diamond:
			dcount++
		case heart:
			hcount++
		case spade:
			scount++
		}
	}

	// Step 1 : Check if all cards are of the same suit
	if ccount >= 5 || dcount >= 5 || hcount >= 5 || scount >= 5 {
		// Collect all cards of the same suit
		for _, v := range h {
			if (ccount >= 5 && v.s == club) ||
				(dcount >= 5 && v.s == diamond) ||
				(hcount >= 5 && v.s == heart) ||
				(scount >= 5 && v.s == spade) {
				sameSuitCards = append(sameSuitCards, v)
			}
		}

		// Step 2 : Sort the cards by pip value
		sort.Slice(sameSuitCards, func(i, j int) bool {
			return sameSuitCards[i].pip < sameSuitCards[j].pip
		})

		// Step 3 : Check if all cards are in sequence
		consecutive := 1
		for i := 1; i < len(sameSuitCards); i++ {
			if sameSuitCards[i].pip == sameSuitCards[i-1].pip+1 {
				consecutive++
				if consecutive == 5 {
					return true
				}
			} else if sameSuitCards[i].pip == sameSuitCards[i-1].pip {
				consecutive = 1
			}
		}
	}

	return false
}

func main() {
	deck := make([]Card, 52)
	var sfcount int // number of straight flushes
	var totcnt int  // Number of trials

	fmt.Print("Enter the number of trials: ")
	_, err := fmt.Scanln(&totcnt)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return
	}

	// Initialize the deck
	for i := 0; i < 13; i++ {
		deck[i] = Card{club, Pip(i + 1)}
		deck[i+13] = Card{diamond, Pip(i + 1)}
		deck[i+26] = Card{heart, Pip(i + 1)}
		deck[i+39] = Card{spade, Pip(i + 1)}
	}

	// Run the trials
	for i := 0; i < totcnt; i++ {
		shuffle(&deck)
		hand := deck[:7]

		if isStraightFlush(hand) {
			sfcount++
		}
	}

	fmt.Printf("\nStraight flushes for %d trials: %d \n", totcnt, sfcount)
	fmt.Printf("Probability of straight flush: %.8f\n", float64(sfcount)/float64(totcnt))
}

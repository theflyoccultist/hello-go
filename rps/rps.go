package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	var move, machineMove, prevMove int
	const (
		rock     = 0
		paper    = 1
		scissors = 2
	)
	const (
		cRock     = 'R'
		cPaper    = 'P'
		cScissors = 'S'
	)
	var cMove string
	var draws, wins, machineWins int
	var rounds int

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("How many rounds do you want to play? ")
	fmt.Scanf("%d", &rounds)
	reader.ReadString('\n') // Clear the newline character from the input buffer

	var rockCounter, scissorCounter, paperCounter int

	// Initialize prevMove to an invalid value
	prevMove = -1

	for i := 0; i < rounds; i++ {

		// Player move
		fmt.Println("\nRound ", i+1, ": Choose either R, P or S")
		cMove, _ = reader.ReadString('\n')
		cMove = strings.TrimSpace(cMove)

		if cMove == "R" {
			move = rock
			rockCounter++
		} else if cMove == "P" {
			move = paper
			paperCounter++
		} else if cMove == "S" {
			move = scissors
			scissorCounter++
		} else {
			fmt.Println("-> Illegal move")
			i--
			continue // Go back to the top of the loop
		}

		// Reset counter if player changes their move
		if prevMove != -1 {
			if move != prevMove {
				// fmt.Println("-> You played a different move than the previous round")
				rockCounter = 0
				scissorCounter = 0
				paperCounter = 0
			}
		}

		// Set machine move based on counters
		if rockCounter >= 10 {
			machineMove = paper
		} else if scissorCounter >= 10 {
			machineMove = rock
		} else if paperCounter >= 10 {
			machineMove = scissors
		} else {
			// Random Move
			source := rand.NewSource(time.Now().UnixNano())
			rng := rand.New(source)
			machineMove = rng.Intn(3)
		}

		// Determine the result using switch
		switch move {
		case rock:
			if machineMove == rock {
				fmt.Println("-> draw")
				draws++
			} else if machineMove == paper {
				fmt.Println("-> machine wins")
				machineWins++
			} else {
				fmt.Println("-> you win")
				wins++
			}
		case paper:
			if machineMove == rock {
				fmt.Println("-> you win")
				wins++
			} else if machineMove == paper {
				fmt.Println("-> draw")
				draws++
			} else {
				fmt.Println("-> machine wins")
				machineWins++
			}
		case scissors:
			if machineMove == rock {
				fmt.Println("-> machine wins")
				machineWins++
			} else if machineMove == paper {
				fmt.Println("-> you win")
				wins++
			} else {
				fmt.Println("-> draw")
				draws++
			}
		}

		// Update previous move
		prevMove = move
	}
	fmt.Println("\nAfter", rounds, "rounds:\n",
		"you win: ", wins,
		" machine wins ", machineWins,
		", with ", draws, "draws")
}

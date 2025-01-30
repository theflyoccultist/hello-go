package main

import "fmt"

func birthdayProbability(n int) float64 {
	probability := 1.0

	for i := 0; i < n; i++ {
		probability *= float64(365-i) / 365
	}

	return 1 - probability
}

func main() {
	const iterations = 10000
	fmt.Printf("People in the room \t Probability of 2 (or more) having the same birthday\n")

	for n := 10; n <= 100; n += 10 {
		totalProbability := 0.0

		for i := 0; i < iterations; i++ {
			totalProbability += birthdayProbability(n)
		}

		averageProbability := totalProbability / iterations
		fmt.Printf("\t%d\t\t\t%f\n", n, averageProbability)
	}

	fmt.Printf("Total number of iterations : %d\n", iterations)
}

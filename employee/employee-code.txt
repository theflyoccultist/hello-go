package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Open the file for reading
	f, err := os.Open("employee.txt")
	check(err)
	defer f.Close()

	// Read the file line by line
	scanner := bufio.NewScanner(f)

	// Declare slices
	var fullNames []string
	var salaries []uint32

	// Add data to the corresponding slice
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		if len(parts) == 3 {
			fullName := parts[0] + "." + parts[1]
			fullNames = append(fullNames, fullName)
			salary, err := strconv.ParseUint(parts[2], 10, 32)
			check(err)
			salaries = append(salaries, uint32(salary))
		} else {
			fmt.Println("-> Invalid line format")
		}
	}

	// Error handling
	if err := scanner.Err(); err != nil {
		check(err)
	}

	//Find the employee with the smallest salary
	minSalary := salaries[0]
	minIndex := 0
	for i, salary := range salaries {
		if salary < minSalary {
			minSalary = salary
			minIndex = i
		}
	}

	//Find the employee with the largest salary
	maxSalary := salaries[0]
	maxIndex := 0
	for i, salary := range salaries {
		if salary > maxSalary {
			maxSalary = salary
			maxIndex = i
		}
	}

	//Find the average salary
	var totalSalary uint32
	for _, salary := range salaries {
		totalSalary += salary
	}
	averageSalary := float64(totalSalary) / float64(len(salaries))

	fmt.Printf("Company's smallest salary: %s, with: %v\n", fullNames[minIndex], minSalary)
	fmt.Printf("Company's largest salary: %s, with: %v\n", fullNames[maxIndex], maxSalary)
	fmt.Printf("Company's average salary: %.2f", averageSalary)
}

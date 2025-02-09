package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Solid interface {
	Volume() float64
}

type Sphere struct {
	radius float64
}

type Cube struct {
	length float64
}

type Pyramid struct {
	base   float64
	height float64
}

func (s Sphere) Volume() float64 {
	return 4 * math.Pi * math.Pow(s.radius, 3) / 3
}

func (l Cube) Volume() float64 {
	return math.Pow(l.length, 3)
}

func (p Pyramid) Volume() float64 {
	return math.Pow(p.base, 2) * p.height / 3
}

func main() {
	fmt.Println("Reading data.txt")
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	var solids []Solid

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) != 3 {
			fmt.Println("Invalid line format: ", line)
			continue
		}

		shapeType := parts[0]
		dimension1, err1 := strconv.ParseFloat(parts[1], 64)
		dimension2, err2 := strconv.ParseFloat(parts[2], 64)
		if err1 != nil || err2 != nil {
			fmt.Println("Invalid number format in line:", line)
			continue
		}

		switch shapeType {
		case "S":
			solids = append(solids, Sphere{radius: dimension1})
		case "C":
			solids = append(solids, Cube{length: dimension1})
		case "P":
			solids = append(solids, Pyramid{base: dimension1, height: dimension2})
		default:
			fmt.Println("Unknown shape type in line:", line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file", err)
	}

	for _, solid := range solids {
		fmt.Printf("Volume: %.2f\n", solid.Volume())
	}
}

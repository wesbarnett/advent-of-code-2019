package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func readWires(file string) ([]string, []string) {
	content, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(strings.Trim(string(content), "\n"), "\n")
	wire1 := strings.Split(strings.Trim(input[0], "\n"), ",")
	wire2 := strings.Split(strings.Trim(input[1], "\n"), ",")
	return wire1, wire2
}

func convertToCoords(wire []string) ([]int, []int) {
	x := make([]int, 1)
	y := make([]int, 1)

	for i := 1; i < len(wire); i++ {
		dir := string(wire[i-1][0])
		amount, err := strconv.Atoi(string(wire[i-1][1:]))
		if err != nil {
			log.Fatal(err)
		}
		switch dir {
		case "R":
			for j := 1; j <= amount; j++ {
				x = append(x, x[len(x)-1]+1)
				y = append(y, y[len(y)-1])
			}
		case "L":
			for j := 1; j <= amount; j++ {
				x = append(x, x[len(x)-1]-1)
				y = append(y, y[len(y)-1])
			}
		case "U":
			for j := 1; j <= amount; j++ {
				x = append(x, x[len(x)-1])
				y = append(y, y[len(y)-1]+1)
			}
		case "D":
			for j := 1; j <= amount; j++ {
				x = append(x, x[len(x)-1])
				y = append(y, y[len(y)-1]-1)
			}
		}
	}

	return x, y
}

func findIntersections(wire1_x []int, wire1_y []int, wire2_x []int, wire2_y []int) ([]int, []int) {
	result_x := make([]int, 0)
	result_y := make([]int, 0)
	for i, _ := range wire1_x {
		for j, _ := range wire2_x {
			if wire1_x[i] == wire2_x[j] && wire1_y[i] == wire2_y[j] && (i != 0 && j != 0) {
				result_x = append(result_x, wire1_x[i])
				result_y = append(result_y, wire1_y[i])
			}
		}
	}
	return result_x, result_y
}

func calcMinDist(x []int, y []int) int {
	minDist := math.MaxInt
	for i, _ := range x {
		dist := x[i] + y[i]
		if dist < minDist {
			minDist = dist
		}
	}
	return minDist
}

func main() {
	var file string
	flag.StringVar(&file, "infile", "input", "Input file")
	flag.Parse()

	wire1, wire2 := readWires(file)
	wire1_x, wire1_y := convertToCoords(wire1)
	wire2_x, wire2_y := convertToCoords(wire2)
	intersection_x, intersection_y := findIntersections(wire1_x, wire1_y, wire2_x, wire2_y)
	minDist := calcMinDist(intersection_x, intersection_y)
	fmt.Println(minDist)
}

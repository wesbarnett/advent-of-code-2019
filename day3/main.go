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

func findIntersections(wire1X []int, wire1Y []int, wire2X []int, wire2Y []int) ([]int, []int) {
	resultX := make([]int, 0)
	resultY := make([]int, 0)
	for i, _ := range wire1X {
		for j, _ := range wire2X {
			if wire1X[i] == wire2X[j] && wire1Y[i] == wire2Y[j] && (i != 0 && j != 0) {
				resultX = append(resultX, wire1X[i])
				resultY = append(resultY, wire1Y[i])
			}
		}
	}
	return resultX, resultY
}

func calcMinDist(x []int, y []int) int {
	minDist := math.MaxInt
	for i, _ := range x {
		if x[i] < 0 {
			x[i] *= -1
		}
		if y[i] < 0 {
			y[i] *= -1
		}
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
	wire1X, wire1Y := convertToCoords(wire1)
	wire2X, wire2Y := convertToCoords(wire2)
	intersectionX, intersectionY := findIntersections(wire1X, wire1Y, wire2X, wire2Y)
	minDist := calcMinDist(intersectionX, intersectionY)
	fmt.Println(minDist)
}

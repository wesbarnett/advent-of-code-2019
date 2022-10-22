package main

import (
	"flag"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func calc_fuel(mass int) int {
	return int(math.Max(math.Floor(float64(mass)/3)-2, 0))
}

func main() {
	var file string
	var masses = []int{}
	flag.StringVar(&file, "infile", "input", "Input file")
	flag.Parse()

	content, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	lines = lines[:len(lines)-1]
	for _, line := range lines {
		mass, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		masses = append(masses, mass)
	}

	fuel := 0
	for _, mass := range masses {
		fuel = fuel + calc_fuel(mass)
	}
	log.Printf("%d", fuel)

	total_fuel := 0
	for _, mass := range masses {
		for mass > 0 {
			fuel = calc_fuel(mass)
			total_fuel = total_fuel + fuel
			mass = fuel
		}
		total_fuel = total_fuel + fuel
	}
	log.Printf("%d", total_fuel)
}

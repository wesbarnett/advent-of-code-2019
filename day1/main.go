package main

import (
	"flag"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var file string
	flag.StringVar(&file, "infile", "input", "Input file")
	flag.Parse()

	content, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	fuel := 0
	for _, line := range lines {
		if len(line) == 0 {
			break
		}
		mass, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		fuel = fuel + int(math.Floor(float64(mass)/3)-2)
	}

	log.Printf("%d", fuel)
}

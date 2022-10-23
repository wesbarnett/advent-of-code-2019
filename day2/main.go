package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func read_program(file string) []int {
	content, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(strings.Trim(string(content), "\n"), ",")
	program := make([]int, len(input))
	for i := range program {
		program[i], err = strconv.Atoi(input[i])
		if err != nil {
			log.Fatal(err)
		}
	}
	return program
}

func run_intcode_program(program []int) int {
	i := 0
	for {
		v := program[i]
		if v == 1 {
			program[program[i+3]] = program[program[i+1]] + program[program[i+2]]
		} else if v == 2 {
			program[program[i+3]] = program[program[i+1]] * program[program[i+2]]
		} else if v == 99 {
			break
		}
		i = i + 4
	}
	return program[0]
}

func main() {
	var file string
	flag.StringVar(&file, "infile", "input", "Input file")
	flag.Parse()

	program := read_program(file)
	result := run_intcode_program(program)
	fmt.Println(result)
}

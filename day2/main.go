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

func find_noun_and_verb(program []int, target int) (int, int) {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			program[1] = noun
			program[2] = verb
			result := run_intcode_program(program)
			if result == target {
				return noun, verb
			}

		}
	}
	return -1, -1
}

func run_intcode_program(program []int) int {
	i := 0
	for {
		v := program[i]
		target := program[i+3]
		if target > len(program) || target < 0 {
			return -1
		}
		if v == 1 {
			program[target] = program[program[i+1]] + program[program[i+2]]
		} else if v == 2 {
			program[target] = program[program[i+1]] * program[program[i+2]]
		} else if v == 99 {
			return program[0]
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

	noun, verb := find_noun_and_verb(program, 19690720)
	fmt.Println(noun, verb)
}

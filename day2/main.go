package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readProgram(file string) []int {
	content, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(strings.Trim(string(content), "\n"), ",")
	mem := make([]int, len(input))
	for i := range mem {
		mem[i], err = strconv.Atoi(input[i])
		if err != nil {
			log.Fatal(err)
		}
	}
	return mem
}

func findNounAndVerb(mem []int, target int) (int, int) {
	mem_copy := make([]int, len(mem))
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			copy(mem_copy, mem)
			result := runIntcodeProgram(mem_copy, noun, verb)
			if result == target {
				return noun, verb
			}
		}
	}
	return -1, -1
}

func runIntcodeProgram(mem []int, noun int, verb int) int {
	mem[1] = noun
	mem[2] = verb
	var addr int
	for {
		switch mem[addr] {
		case 1:
			mem[mem[addr+3]] = mem[mem[addr+1]] + mem[mem[addr+2]]
		case 2:
			mem[mem[addr+3]] = mem[mem[addr+1]] * mem[mem[addr+2]]
		case 99:
			return mem[0]
		}
		addr += 4
	}
	return mem[0]
}

func main() {
	var file string
	flag.StringVar(&file, "infile", "input", "Input file")
	flag.Parse()

	mem := readProgram(file)
	result := runIntcodeProgram(mem, 12, 2)
	fmt.Println(result)

	mem = readProgram(file)
	noun, verb := findNounAndVerb(mem, 19690720)
	fmt.Println(100*noun + verb)
}

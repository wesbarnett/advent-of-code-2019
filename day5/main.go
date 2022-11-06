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

func runIntcodeProgram(mem []int, input int) {
	var output, param1, param2 int
	addr := 0
	for {
		x := mem[addr]
		opcode := x % 100
		switch opcode {
		case 1:
			switch x / 100 % 10 {
			case 0:
				param1 = mem[mem[addr+1]]
			case 1:
				param1 = mem[addr+1]
			}

			switch x / 1000 % 10 {
			case 0:
				param2 = mem[mem[addr+2]]
			case 1:
				param2 = mem[addr+2]
			}

			mem[mem[addr+3]] = param1 + param2

			addr += 4

		case 2:
			switch x / 100 % 10 {
			case 0:
				param1 = mem[mem[addr+1]]
			case 1:
				param1 = mem[addr+1]
			}

			switch x / 1000 % 10 {
			case 0:
				param2 = mem[mem[addr+2]]
			case 1:
				param2 = mem[addr+2]
			}

			mem[mem[addr+3]] = param1 * param2

			addr += 4
		case 3:
			mem[mem[addr+1]] = input
			addr += 2
		case 4:
			switch x / 100 % 10 {
			case 0:
				param1 = mem[mem[addr+1]]
			case 1:
				param1 = mem[addr+1]
			}
			output = param1
			fmt.Println(output)
			addr += 2
		case 99:
			return
		default:
			addr += 1
		}
	}
	log.Fatal("Missing code 99")
	return -1
}

func main() {
	var file string
	flag.StringVar(&file, "infile", "input", "Input file")
	flag.Parse()

	mem := readProgram(file)
	runIntcodeProgram(mem, 1)

}

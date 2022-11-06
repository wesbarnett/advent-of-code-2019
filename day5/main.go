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

func getParams(mem []int, addr int, num int) []int {
	var params []int
	var param int
	divisor := 100
	for i := 1; i <= num; i++ {
		switch mem[addr] / divisor % 10 {
		case 0:
			param = mem[mem[addr+i]]
		case 1:
			param = mem[addr+i]
		}
		params = append(params, param)
		divisor *= 10
	}
	return params
}

func runIntcodeProgram(mem []int, input int) {
	var output int
	addr := 0
	for {
		switch mem[addr] % 100 {
		case 1:
			params := getParams(mem, addr, 3)
			mem[mem[addr+3]] = params[0] + params[1]
			addr += 4
		case 2:
			params := getParams(mem, addr, 3)
			mem[mem[addr+3]] = params[0] * params[1]
			addr += 4
		case 3:
			mem[mem[addr+1]] = input
			addr += 2
		case 4:
			params := getParams(mem, addr, 1)
			output = params[0]
			fmt.Println(output)
			addr += 2
		case 99:
			return
		default:
			addr += 1
		}
	}
	log.Fatal("Missing code 99")
	return
}

func main() {
	var file string
	flag.StringVar(&file, "infile", "input", "Input file")
	flag.Parse()

	mem := readProgram(file)
	runIntcodeProgram(mem, 1)

}

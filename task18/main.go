package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var input = flag.String("input file", "input.txt", "input testing file")

func main() {
	flag.Parse()

	f, err := os.Open(*input)
	if err != nil {
		log.Fatalf("error opening input file %s", *input)
	}
	defer f.Close()

	p1Sum := 0
	p2Sum := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		exp := strings.TrimSpace(scanner.Text())

		var precedence = make(map[Expression]int)
		// Set precedence level for t1: + and * have the same
		precedence["+"] = 1 //+ has higher precedence in T.18.2
		precedence["*"] = 1
		// TASK 1
		q1, err := ShuntingYard(exp, precedence)
		if err != nil {
			log.Fatal(err)
		}
		t1 := RPN(q1)
		p1Sum += t1

		// TASK 2
		// Set precedence level for aoc-Task18.2 (only * and + operators)
		precedence["+"] = 2 //+ has higher precedence in T.18.2
		precedence["*"] = 1

		q2, err := ShuntingYard(exp, precedence)
		if err != nil {
			log.Fatal(err)
		}
		t2 := RPN(q2)
		p2Sum += t2
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("RESULTANT SUM task 1: ", p1Sum)
	fmt.Println("RESULTANT SUM task 2: ", p2Sum)
}

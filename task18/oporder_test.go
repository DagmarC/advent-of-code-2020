package main

import (
	"testing"

	"github.com/DagmarC/introtoalgo/chapter10/queue"
)

// Every test can be set with differrent precedence of operators
var precedence = make(map[Expression]int)

func TestEvalPart1(t *testing.T) {
	precedence["+"] = 1 // + has the same precedence as * in part 1
	precedence["*"] = 1

	tests := []struct {
		exp  string
		want int
	}{
		{exp: "22+120", want: 142},
		{exp: "(1+2)*(2+3)", want: 15},
		{exp: "1 + (2 * 3) + (4 * (5 + 6))", want: 51},
		{exp: "1 + 2 * 3 + 4 * 5 + 6", want: 71},
		{exp: "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", want: 12240},
		{exp: "5 + (8 * 3 + 9 + 3 * 4 * 3)", want: 437},
		{exp: "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", want: 13632},
	}

	for _, tc := range tests {
		q, err := ShuntingYard(tc.exp, precedence) // infix to postfix
		if err != nil {
			t.Errorf(err.Error())
		}

		if got := RPN(q); got != tc.want {
			t.Errorf("expected %d, got %d", tc.want, got)
		}
	}
}

func TestEvalPart2(t *testing.T) {
	precedence["+"] = 2 //+ has higher precedence in T.18.2
	precedence["*"] = 1

	tests := []struct {
		exp  string
		want int
	}{
		{exp: "2*3+5", want: 16},
		{exp: "(1+2)*(2+3)", want: 15},
		{exp: "1 + (2 * 3) + (4 * (5 + 6))", want: 51},
		{exp: "1 + 2 * 3 + 4 * 5 + 6", want: 231},
		{exp: "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", want: 669060},
		{exp: "5 + (8 * 3 + 9 + 3 * 4 * 3)", want: 1445},
		{exp: "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", want: 23340},
		{exp: "((2*(1*2) + (3+1)) + 2) * 2 + 4 + 2", want: 112},
	}
	for _, tc := range tests {
		q, err := ShuntingYard(tc.exp, precedence) // infix to postfix
		if err != nil {
			t.Errorf(err.Error())
		}

		if got := RPN(q); got != tc.want {
			t.Errorf("expected %d, got %d", tc.want, got)
		}
	}
}

func TestShuntingYardAlgo(t *testing.T) {
	var resQ *queue.Queue[Expression]
	var err error
	precedence["+"] = 1
	precedence["*"] = 2

	q1 := queue.NewQueue[Expression]() // postfix notation (+ has higher precedence than * here)
	q1.Enqueue("31")
	q1.Enqueue("22")
	q1.Enqueue("1")
	q1.Enqueue("*")
	q1.Enqueue("+")
	q1.Enqueue("11")
	q1.Enqueue("+")

	tests := []struct {
		exp string
		res *queue.Queue[Expression]
	}{
		{exp: "31 + 22 * 1 + 11", res: q1},
	}
	for _, tc := range tests {
		resQ, err = ShuntingYard(tc.exp, precedence)
		if err != nil {
			t.Errorf(err.Error())
		}

		for {
			e1, err := q1.Dequeue(Expression("empty"))    // expected
			e2, err2 := resQ.Dequeue(Expression("empty")) // result
			if err != nil {
				if err2 == nil {
					t.Errorf("not same queues exp: %v, result: %v ", q1, resQ)
				}
				break
			}
			if e1.String() != e2.String() {
				t.Errorf("expected %s, got %s", e1.String(), e2.String())
			}
		}
	}
}

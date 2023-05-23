package main

import (
	"fmt"
	"strconv"

	"github.com/DagmarC/introtoalgo/chapter10/queue"
	"github.com/DagmarC/introtoalgo/chapter10/stack"
)

type Expression string

func (e Expression) String() string {
	return string(e)
}

// To calculate the expression use RPN - Reverse Polish Notation, where 2*3+1 is converted to 23*1+

// EVALUATION OF 23*1+
// RPN pushes numbers to stack from left [3-Top, 2]
// then encounters sign * so pops two elements out of stack 3*2=6
// and pushes result 6 back to stack [6]
// then pushes 1 to stack [1-Top, 6] and then encounters + sign,
// so pops 1+6=7 and pushes result back to stack [7]
// pops element - empty stack so the RESULT=7

// PROCESS of RPN:
// 3 + 4 × (2 − 1) will be 3 4 2 1 - * + (since * has higher precedence than +) (not in this task though)
// [1, 2, 4, 3] -> - so pop 1 and 2 and substract then push result back [1, 4, 3]
// -> * so pop 1 and 4 a push back result [4, 3] -> + so pop 4 and 3
// and push result back [7] = 7

// TASK 2 NOTE: + has higher precedence than *, so RPN notation is differrent:
//
//	3 + 4 × (2 − 1) will be 3 4 + 2 1 - *
//
// STACK [4, 3] -> [7] -> [1, 2, 7] -> [1, 7] -> [7] = 7
type MyInt int

func (mi MyInt) String() string {
	return strconv.Itoa(int(mi))
}

// RPN (Reverse Polish Notation) will take the queue (expression in postfix notation) and calculates the result
func RPN(postfixQ *queue.Queue[Expression]) int {
	var intStack stack.Stack[MyInt]
	fmt.Println("Postfix Q:", postfixQ)

	for el, err := postfixQ.Dequeue(Expression("empty")); err == nil; el, err = postfixQ.Dequeue(Expression("empty")) {
		// if int conversion succeeds - push to stack, otherwise it is an operator
		if num, err := strconv.Atoi(string(el)); err == nil {
			intStack.Push(MyInt(num))
		} else {
			x, err := intStack.Pull(MyInt(-1))
			if err != nil {
				break // empty stack
			}
			y, err := intStack.Pull(MyInt(-1))
			if err != nil {
				break // empty stack
			}
			// Decide the operation and push back to stack
			if el == "+" {
				intStack.Push(MyInt(y + x))
			} else if el == "*" {
				intStack.Push(MyInt(y * x))
			} else {
				fmt.Println("undefined operation", el)
				break // undefined operation
			}
		}
	}
	result, err := intStack.Pull(MyInt(-1))
	if err != nil {
		fmt.Println("ERROR RESULT")
		return int(result)
	}
	return int(result)
}

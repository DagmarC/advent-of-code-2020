package main

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	"github.com/DagmarC/introtoalgo/chapter10/queue"
	"github.com/DagmarC/introtoalgo/chapter10/stack"
)

// SHUNTING YARD ALGORITHM by Edsger Dijkstra
// Converting INFIX to RPN - use SHUNTING YARD ALGORITHM by Edsger Dijkstra
// method for parsing arithmetical and logical expressions
// can produce either RPN or AST (abstract syntax tree)
// algorithm is stack based

// SIMPLE CONVERSION
// Input: 3 + 4
// Push 3 to the output queue (whenever a number is read it is pushed to the output)
// Push + (or its ID) onto the operator stack
// Push 4 to the output queue
// After reading the expression, pop the operators off the stack and add them to the output.
// In this case there is only one, "+".
// Output: 3 4 +

// PRECEDENCE:
// LOAD  operator o1 from the expression one by one
// if o1 is a NUMBER:
// ....... PUSH to the output QUEUE
// 1. if o1 has higher precedence than the operator on the top of the stack o2:
// ...... PUSH to STACK
// 2. WHILE o1 has the SAME prec. as op on top of stack o2 OR o2 has greater precedence:
// ...... POP top of the stack to QUEUE until lower precedence is on the top of the stack, then PUSH operator o1 to STACK
// 3. If op is a LEFT paranthesis:
// ...... PUSH to STACK
// 4. If o1 and LEFT paranthesis is on top STACK:
// ...... PUSH op to STACK
// 5, If op is RIGHT paranthesis:
// ...... POP from stack to QUEUE until ( is found and then discard ( (if not found - paranthesis mismatch)
// Finally POP the entire Stack to the output QUEUE

// This already shows a couple of rules:
// All numbers are pushed to the output when they are read.
// At the end of reading the expression, pop all operators off the stack onto the output.

// ShuntingYard algorithm will convert infix notation to RPN (Reverse Polish notation)
// result is an output queue, precedence is a map of operators with a given operators needed
func ShuntingYard(exp string, precedence map[Expression]int) (*queue.Queue[Expression], error) {
	var opStack stack.Stack[Expression]
	var numQueue queue.Queue[Expression]
	var num string

	exp = strings.ReplaceAll(exp, " ", "")
	fmt.Println("EXPRESSION", exp)

	for i := 0; i <= len(exp)-1; i++ {
		r := rune(exp[i])

		if isNum(r) {
			num, i = tokenizeNum(exp, i)
			numQueue.Enqueue(Expression(num)) // put the n. into the output queue

		} else if isOperator(r) {
			o1 := Expression(string(r))

			for o2, err := opStack.Top(Expression("empty")); err == nil && o2 != "(" && (precedence[o2] >= precedence[o1]); o2, err = opStack.Top(Expression("empty")) {
				// pop o2 from the top of the stack (until condition is fullfilled - lower precedence on the top of the stack) then
				o2, err = opStack.Pull(Expression("empty"))
				if err != nil {
					break // empty stack
				}
				numQueue.Enqueue(o2) // enque operator to the output queue
			}
			// after 'while loop' ended push o1 on the top of the stack
			opStack.Push(o1)

		} else if isLeftBracket(r) {
			opStack.Push(Expression(string(r)))

		} else if isRightBracket(r) {
			for o2, _ := opStack.Top(Expression("empty")); o2 != "("; o2, _ = opStack.Top(Expression("empty")) {
				// pop the operator Stack into the Queue
				o2, err := opStack.Pull(Expression("empty"))
				if err != nil {
					return &numQueue, errors.New("mismached paranthesis")
				}
				numQueue.Enqueue(o2) // enque operator to the output queue
			}
			o2, _ := opStack.Top(Expression("empty"))
			if o2 == "(" {
				// pop the left parenthesis from the operator stack and discard it
				opStack.Pull(Expression("empty"))
			}
		}
	}
	// Pop the rest of the entire stack onto the Queue
	for o2, err := opStack.Pull(Expression("empty")); err == nil; o2, err = opStack.Pull(Expression("empty")) {
		/* If the operator token on the top of the stack is a parenthesis, then there are mismatched parentheses. */
		if o2 == "(" {
			return &numQueue, errors.New("mismached paranthesis")
		}
		numQueue.Enqueue(o2) // enque operator to the output queue
	}

	return &numQueue, nil
}

func tokenizeNum(exp string, i int) (string, int) {
	var res bytes.Buffer
	for ; i <= len(exp)-1; i++ {
		r := rune(exp[i])
		if !isNum(r) {
			return res.String(), i - 1 // index of the last digit read
		}
		res.WriteRune(r)
	}
	return res.String(), i - 1 // index of the last digit read
}

func isOperator(r rune) bool {
	return r == '+' || r == '*' || r == '-' || r == '/'
}

func isLeftBracket(r rune) bool {
	return r == '('
}

func isRightBracket(r rune) bool {
	return r == ')'
}

func isNum(r rune) bool {
	return r >= '0' && r <= '9'
}

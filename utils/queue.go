package utils

import (
	"errors"
	"fmt"
	"sort"
)

type Queue struct {
	elements []int
}

func (q *Queue) Elements() []int {
	return q.elements
}

func (q *Queue) Enqueue(el int) {
	q.elements = append(q.elements, el)
}

func (q *Queue) Dequeue() (int, error) {
	el, err := q.Peek()
	if err != nil {
		return -1, err
	}
	q.elements = q.elements[1:]
	return el, nil
}

func (q *Queue) Peek() (int, error) {
	if q.IsEmpty() {
		return -1, errors.New("queue is empty")
	}
	return q.elements[0], nil
}

func (q *Queue) Length() int {
	return len(q.elements)
}

func (q *Queue) IsEmpty() bool {
	return q.Length() == 0
}

func (q *Queue) Print() {
	fmt.Println(q.elements)
}

func (q *Queue) Get(x int) int {
	if x < 0 || x >= q.Length() {
		return 0
	}
	return q.elements[x]
}

func (q *Queue) Max() int {
	if !q.IsEmpty() {
		sort.Ints(q.elements)
		return q.elements[len(q.elements)-1]
	}
	return 0
}

func (q *Queue) Min() int {
	if !q.IsEmpty() {
		sort.Ints(q.elements)
		return q.elements[0]
	}
	return 0
}

package main

import "errors"

type integersQueue []int

func (q *integersQueue) put(v int) {
	*q = append(*q, v)
}

func (q *integersQueue) get() (int, error) {
	if len(*q) == 0 {
		return 0, errors.New("queue is empty")
	}
	r := (*q)[0]
	*q = (*q)[1:]
	return r, nil
}

func (q *integersQueue) isEmpty() bool {
	if len(*q) == 0 {
		return true
	}
	return false
}

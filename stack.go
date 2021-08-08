package main

import "errors"

type integersStack []int

func (s *integersStack) push(v int) {
	*s = append(*s, v)
}

func (s *integersStack) pop() (int, error) {
	if len(*s) == 0 {
		return 0, errors.New("stack is empty")
	}

	ret := (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]
	return ret, nil
}

func (s *integersStack) isEmpty() bool {
	if len(*s) == 0 {
		return true
	}
	return false
}

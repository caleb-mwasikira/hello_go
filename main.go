package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	ErrStackFull        error = errors.New("stack already full")
	ErrStackEmpty       error = errors.New("stack is empty")
	ErrInvalidStackSize error = errors.New("stack size must be greater than one")
)

type Stack struct {
	items []int
	top   int
	size  int
}

func NewStack(size int, items ...int) (Stack, error) {
	if size <= 0 {
		return Stack{}, ErrInvalidStackSize
	}

	return Stack{
		items: make([]int, size),
		top:   -1,
		size:  size,
	}, nil
}

func (s *Stack) push(items ...int) (bool, error) {
	for _, item := range items {
		if s.top == s.size-1 {
			return false, ErrStackFull
		}

		s.top++
		s.items[s.top] = item
	}
	return true, nil
}

func (s *Stack) pop() (int, error) {
	if s.top == -1 {
		return 0, ErrStackEmpty
	}

	item := s.items[s.top]
	s.top--
	return item, nil
}

func (s Stack) getItems() []int {
	items := s.items[0 : s.top+1]
	return items
}

func main() {
	stack, err := NewStack(10)
	if err != nil {
		log.Fatal(err)
	}

	for i := 5; i < 20; i++ {
		_, err = stack.push(i)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("stack: %#v\n", stack.getItems())
	}

	for i := 0; i < 2; i++ {
		val, err := stack.pop()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("popped value: %v\n", val)
		fmt.Printf("stack: %#v\n", stack.getItems())
	}
}

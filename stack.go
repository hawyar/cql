package cql

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

type node struct {
	val  antlr.RuleContext
	next *node
}

type stack struct {
	curr  *node
	count int
}

func newStack() *stack {
	return &stack{nil, 0}
}

func (s *stack) size() int {
	return s.count
}

func (s *stack) peek() antlr.RuleContext {
	if s.curr == nil {
		return nil
	}
	return s.curr.val
}

func (s *stack) push(val antlr.RuleContext) {
	if s.curr == nil {
		s.curr = &node{val, nil}
		s.count++
		return
	}

	s.curr = &node{val, s.curr}
	s.count++
}

func (s *stack) pop() antlr.RuleContext {
	if s.curr == nil {
		return nil
	}

	val := s.curr.val
	s.curr = s.curr.next
	s.count--
	return val
}

func (s *stack) isEmpty() bool {
	return s.curr == nil
}

func (s *stack) clear() {
	s.curr = nil
	s.count = 0
}

func (s *stack) print() {
	for curr := s.curr; curr != nil; curr = curr.next {
		fmt.Println(curr.val.GetRuleIndex())
	}
	fmt.Println("------ \n size: ", s.count)
}

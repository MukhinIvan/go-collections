package stack

import (
	"container/list"
	"errors"
	"sync"
)

type StackInterface[T any] interface {
	Push(T)
	Pop() (T, error)
	Peek() (T, error)
	IsEmpty() bool
	Size() int
}

type Stack[T any] struct {
	mutex sync.RWMutex
	stack *list.List
}

const errorString = "stack is empty"

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{mutex: sync.RWMutex{}, stack: list.New()}
}

func (s *Stack[T]) Push(value T) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.stack.PushBack(value)
}

func (s *Stack[T]) Pop() (T, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.IsEmpty() {
		var t T
		return t, errors.New(errorString)
	}
	result := s.stack.Back().Value.(T)
	s.stack.Remove(s.stack.Back())
	return result, nil
}

func (s *Stack[T]) Peek() (T, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.IsEmpty() {
		var t T
		return t, errors.New(errorString)
	}
	return s.stack.Back().Value.(T), nil
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Stack[T]) Size() int {
	return s.stack.Len()
}

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

// Возвращает инициализированный контейнер stack
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{mutex: sync.RWMutex{}, stack: list.New()}
}

// Добавляет элемент в стек
func (s *Stack[T]) Push(value T) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.stack.PushBack(value)
}

// Возвращает самый верхний элемент стека (и удаляет его из стека)
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

// Возвращает самый верхний элемент стека (не удаляя его из стека)
func (s *Stack[T]) Peek() (T, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.IsEmpty() {
		var t T
		return t, errors.New(errorString)
	}
	return s.stack.Back().Value.(T), nil
}

// Проверяет стек на наличие элементов (если стек пуст, то возвращается true)
func (s *Stack[T]) IsEmpty() bool {
	return s.Size() == 0
}

// Возвращает количество элементов в стеке
func (s *Stack[T]) Size() int {
	return s.stack.Len()
}

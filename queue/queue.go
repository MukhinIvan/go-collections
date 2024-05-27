package queue

import (
	"container/list"
	"errors"
	"sync"
)

type QueueInterface[T any] interface {
	Push(T)
	Front() (T, error)
	Back() (T, error)
	Pop()
	IsEmpty() bool
	Size() int
}

const errorString = "queue is empty"

type Queue[T any] struct {
	mutex sync.RWMutex
	queue *list.List
}

// Возвращает инициализированный контейнер queue
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{mutex: sync.RWMutex{}, queue: list.New()}
}

// Добавляет элемент в очередь
func (q *Queue[T]) Push(value T) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	q.queue.PushBack(value)
}

// Возвращает самый первый элемент очереди (не удаляя его из очереди)
func (q *Queue[T]) Front() (T, error) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if q.IsEmpty() {
		var t T
		return t, errors.New(errorString)
	}
	return q.queue.Front().Value.(T), nil
}

// Возвращает самый последний элемент очереди (не удаляя его из очереди)
func (q *Queue[T]) Back() (T, error) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if q.IsEmpty() {
		var t T
		return t, errors.New(errorString)
	}
	return q.queue.Back().Value.(T), nil
}

// Удаляет элемент из начала очереди
func (q *Queue[T]) Pop() error {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if q.IsEmpty() {
		return errors.New(errorString)
	}
	q.queue.Remove(q.queue.Front())
	return nil
}

// Проверяет очередь на наличие элементов (если очередь пуста, то возвращается true)
func (q *Queue[T]) IsEmpty() bool {
	return q.Size() == 0
}

// Возвращает количество элементов в очереди
func (q *Queue[T]) Size() int {
	return q.queue.Len()
}

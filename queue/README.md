# Queue (очередь)

Контейнер `queue` (очередь) работает по принципу FIFO (first-in first-out, "первый вошел -- первым вышел"), то есть первым извлекается тот элемент, который добавляется первым.

## Типы и функции

```go
type QueueInterface[T any] interface {
	Push(T)
	Front() (T, error)
	Back() (T, error)
	Pop()
	IsEmpty() bool
	Size() int
}

type Queue[T any] struct

const errorString = "queue is empty"

func NewQueue[T any]() *Queue[T]
func (q *Queue[T]) Push(value T)
func (q *Queue[T]) Front() (T, error)
func (q *Queue[T]) Back() (T, error)
func (q *Queue[T]) Pop() error
func (q *Queue[T]) IsEmpty() bool
func (q *Queue[T]) Size() int
```

### `func NewQueue[T any]() *Queue[T]`

```go
func NewQueue[T any]() *Queue[T]
```

Возвращает инициализированный контейнер `queue`.

### `func (q *Queue[T]) Push(value T)`

```go
func (q *Queue[T]) Push(value T)
```

Добавляет элемент в очередь.

### `func (q *Queue[T]) Front() (T, error)`

```go
func (q *Queue[T]) Front() (T, error)
```

Возвращает самый первый элемент очереди (не удаляя его из очереди).

### `func (q *Queue[T]) Back() (T, error)`

```go
func (q *Queue[T]) Back() (T, error)
```

Возвращает самый последний элемент очереди (не удаляя его из очереди).

### `func (q *Queue[T]) Pop() error`

```go
func (q *Queue[T]) Pop() error
```

Удаляет элемент из начала очереди.

### `func (q *Queue[T]) IsEmpty() bool`

```go
func (q *Queue[T]) IsEmpty() bool
```

Проверяет очередь на наличие элементов (если очередь пуста, то возвращается `true`).

### `func (q *Queue[T]) Size() int`

```go
func (q *Queue[T]) Size() int
```

Возвращает количество элементов в очереди.

## Пример

```bash
go install github.com/MukhinIvan/go-collections/queue
```

```go
package main

import (
	"fmt"

	"github.com/MukhinIvan/go-collections/queue"
)

func main() {
	myQueue := queue.NewQueue[int]()

	myQueue.Push(1) // добавим элемент в очередь
	myQueue.Push(2)
	myQueue.Push(3)

	element, err := myQueue.Front() // получим первый элемент очереди
	myQueue.Pop()                   // удалим первый элемент очереди
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(element)

	element, err = myQueue.Front()
	myQueue.Pop()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(element)

	element, err = myQueue.Front()
	myQueue.Pop()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(element)
}
```

Результат:

```bash
1
2
3
```

# Queue (очередь)

Контейнер `stack` (стек) работает по принципу LIFO (last-in first-out или "последний вошел -- первым вышел"), то есть первым извлекается тот элемент, который добавляется последним


## Типы и функции

```go
type StackInterface[T any] interface {
	Push(T)
	Pop() (T, error)
	Peek() (T, error)
	IsEmpty() bool
	Size() int
}

type Stack[T any] struct

const errorString = "stack is empty"

func NewStack[T any]() *Stack[T]
func (s *Stack[T]) Push(value T)
func (s *Stack[T]) Pop() (T, error)
func (s *Stack[T]) Peek() (T, error)
func (s *Stack[T]) IsEmpty() bool
func (s *Stack[T]) Size() int
```

### `func NewStack[T any]() *Stack[T]`

```go
func NewStack[T any]() *Stack[T]
```

Возвращает инициализированный контейнер `stack`.

### `func (s *Stack[T]) Push(value T)`

```go
func (s *Stack[T]) Push(value T)
```

Добавляет элемент в стек.

### `func (s *Stack[T]) Pop() (T, error)`

```go
func (s *Stack[T]) Pop() (T, error)
```

Возвращает самый верхний элемент стека (и удаляет его из стека).

### `func (s *Stack[T]) Peek() (T, error)`

```go
func (s *Stack[T]) Peek() (T, error)
```

Возвращает самый верхний элемент стека (не удаляя его из стека).

### `func (s *Stack[T]) IsEmpty() bool`

```go
func (s *Stack[T]) IsEmpty() bool
```

Проверяет стек на наличие элементов (если стек пуст, то возвращается `true`).

### `func (s *Stack[T]) Size() int`

```go
func (s *Stack[T]) Size() int
```

Возвращает количество элементов в стеке.

## Пример

```bash
go install github.com/MukhinIvan/go-collections/stack
```

```go
package main

import (
	"fmt"

	"github.com/MukhinIvan/go-collections/stack"
)

func main() {
	myStack := stack.NewStack[int]()

	myStack.Push(1)
	myStack.Push(2)
	myStack.Push(3)

	element, err := myStack.Pop()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(element)

	element, err = myStack.Pop()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(element)

	element, err = myStack.Pop()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(element)
}
```
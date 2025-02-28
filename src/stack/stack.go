package stack

import (
    "errors"
	"container/list"
)

type Stack[T comparable] struct {
    data *list.List
	size uint64
}

func New[T comparable]() *Stack[T] {
    return &Stack[T]{data: list.New()}
}

func (s *Stack[T]) Size() uint64 {
	return s.size
}

func (s *Stack[T]) IsEmpty() bool {
    return s.data.Len() == 0
}

func (s *Stack[T]) Push(data T) {
    s.data.PushFront(data)
    s.size++
}

func (s *Stack[T]) Pop() (T, error) {
	var zero T
	if s.IsEmpty() {
        return zero, errors.New("stack is empty")
    }

    data := s.data.Front().Value.(T)
    s.data.Remove(s.data.Front())
    s.size--

    return data, nil
}

func (s *Stack[T]) Peek() (T, error) {
	var zero T
    if s.IsEmpty() {
        return zero, errors.New("stack is empty")
    }

    return s.data.Front().Value.(T), nil
}
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


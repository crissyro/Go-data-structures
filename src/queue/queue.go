package queue

import (
    "errors"
    "container/list"
)

type Queue[T comparable] struct {
    data *list.List
    size uint64
}

func New[T comparable]() *Queue[T] {
    return &Queue[T]{data: list.New()}
}


package dequeue

import (
    "errors"
    "container/list"
)

type Dequeue[T comparable] struct {
    data *list.List
    size uint64
}

func New[T comparable]() *Dequeue[T] {
    return &Dequeue[T]{data: list.New()}
}

func (d *Dequeue[T]) Size() uint64 {
    return d.size
}


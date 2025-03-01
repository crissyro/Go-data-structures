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

func (q *Queue[T]) Size() uint64 {
    return q.size
}

func (q *Queue[T]) IsEmpty() bool {
    return q.data.Len() == 0
}

func (q *Queue[T]) Enqueue(data T) {
    q.data.PushBack(data)
    q.size++
}

func (q *Queue[T]) Dequeue() (T, error) {
	var zero T
    if q.IsEmpty() {
        return zero, errors.New("queue is empty")
    }

    data := q.data.Front().Value.(T)
    q.data.Remove(q.data.Front())
    q.size--

    return data, nil
}
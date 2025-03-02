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

func (d *Dequeue[T]) IsEmpty() bool {
    return d.data.Len() == 0
}

func (d *Dequeue[T]) Front() (T, error) {
	var zero T
    if d.IsEmpty() {
        return zero, errors.New("dequeue is empty")
    }

    return d.data.Front().Value.(T), nil
}

func (d *Dequeue[T]) Rear() (T, error) {
	var zero T
    if d.IsEmpty() {
        return zero, errors.New("dequeue is empty")
    }

    return d.data.Back().Value.(T), nil
}

func (d *Dequeue[T]) EnqueueFront(data T) {
    d.data.PushFront(data)
    d.size++
}

func (d *Dequeue[T]) EnqueueRear(data T) {
    d.data.PushBack(data)
    d.size++
}

func (d *Dequeue[T]) DequeueFront() (T, error) {
    var zero T
    if d.IsEmpty() {
        return zero, errors.New("dequeue is empty")
    }

    data := d.data.Front().Value.(T)
    d.data.Remove(d.data.Front())
    d.size--

    return data, nil
}

func (d *Dequeue[T]) DequeueRear() (T, error) {
    var zero T
    if d.IsEmpty() {
        return zero, errors.New("dequeue is empty")
    }

    data := d.data.Back().Value.(T)
    d.data.Remove(d.data.Back())
    d.size--

    return data, nil
}

func (d *Dequeue[T]) Rotate(k int) error {
    if d.IsEmpty() || k <= 0 || k > int(d.size) {
        return errors.New("invalid rotation value")
    }

    for i := 0; i < k; i++ {
        d.data.MoveToBack(d.data.Front())
    }

    return nil
}
package dequeue

import (
    "errors"
    "container/list"
)

type Dequeue[T comparable] struct {
    data *list.List
    size uint64
}


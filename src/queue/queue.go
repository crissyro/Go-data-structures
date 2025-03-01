package queue

import (
    "errors"
    "container/list"
)

type Queue[T comparable] struct {
    data *list.List
    size uint64
}

package linkedlist

import (
	"errors"
	"fmt"
)

type Node struct {
    data interface{}
    next *Node
}

type LinkedList struct {
    head *Node
    size uint64
}

func New() *LinkedList {
	return &LinkedList{}
}
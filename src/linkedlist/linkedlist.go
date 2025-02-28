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

func new() *LinkedList {
	return &LinkedList{}
}

func (list *LinkedList) append(data interface{}) {
	newNode := &Node{data: data}
	
    if list.head == nil {
        list.head = newNode
    } else {
        current := list.head
        for current.next != nil {
            current = current.next
        }
        current.next = newNode
    }
    list.size++
}
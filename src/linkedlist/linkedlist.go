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

func (list *LinkedList) print() {
	current := list.head
    fmt.Print("List: ")
    for current != nil {
        fmt.Printf("%v -> ", current.data)
        current = current.next
    }
    fmt.Println("nil")
}

func (list *LinkedList) prepend(data interface{}) {
	newNode := &Node{data: data, next: list.head}
    list.head = newNode
    list.size++
}

func (list *LinkedList) delete(data interface{}) error {
	if list.head == nil {
        return errors.New("List is empty")
    }

    if list.head.data == data {
        list.head = list.head.next
        list.size--
        return nil
    }

    current := list.head
    for current.next != nil && current.next.data != data {
        current = current.next
    }

    if current.next == nil {
        return errors.New("Element not found")
    }

    current.next = current.next.next
    list.size--
    return nil
}
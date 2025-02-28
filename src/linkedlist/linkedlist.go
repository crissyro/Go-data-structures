package linkedlist

import (
    "errors"
    "fmt"
)

type Node[T comparable] struct {
    data T
    next *Node[T]
}

type LinkedList[T comparable] struct {
    head *Node[T]
    size uint64
}

func New[T comparable]() *LinkedList[T] {
    return &LinkedList[T]{}
}

func (list *LinkedList[T]) Size() uint64 {
    return list.size
}

func (list *LinkedList[T]) Append(data T) {
    newNode := &Node[T]{data: data}

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

func (list *LinkedList[T]) Print() {
    current := list.head
    fmt.Print("List: ")

    for current != nil {
        fmt.Printf("%v -> ", current.data)
        current = current.next
    }

    fmt.Println("nil")
}

func (list *LinkedList[T]) IsEmpty() bool {
    return list.head == nil
}

func (list *LinkedList[T]) Prepend(data T) {
    newNode := &Node[T]{data: data, next: list.head}
    list.head = newNode
    list.size++
}

func (list *LinkedList[T]) Delete(data T) error {
    if list.IsEmpty() {
        return errors.New("list is empty")
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
        return errors.New("element not found")
    }

    current.next = current.next.next
    list.size--
    return nil
}

func (list *LinkedList[T]) Get(index int) (T, error) {
    var zero T
    if index < 0 || index >= int(list.size) {
        return zero, errors.New("index out of range")
    }

    current := list.head
    for i := 0; i < index; i++ {
        current = current.next
    }

    return current.data, nil
}

func (list *LinkedList[T]) Reverse() {
    var prev *Node[T]
    current := list.head

    for current != nil {
        next := current.next
        current.next = prev
        prev = current
        current = next
    }

    list.head = prev
}

func (list *LinkedList[T]) Contains(data T) bool {
    current := list.head
    for current != nil {
        if current.data == data {
            return true
        }
        current = current.next
    }

    return false
}
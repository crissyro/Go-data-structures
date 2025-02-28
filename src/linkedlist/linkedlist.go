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

func (list *LinkedList) Size() uint64 {
	return list.size
}

func (list *LinkedList) Append(data interface{}) {
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

func (list *LinkedList) Print() {
	current := list.head
    fmt.Print("List: ")

    for current != nil {
        fmt.Printf("%v -> ", current.data)
        current = current.next
    }

    fmt.Println("nil")
}

func (list *LinkedList) IsEmpty() bool {
	return list.head == nil
}

func (list *LinkedList) Prepend(data interface{}) {
	newNode := &Node{data: data, next: list.head}
    list.head = newNode
    list.size++
}

func (list *LinkedList) Delete(data interface{}) error {
	if list.isEmpty() {
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

func (list *LinkedList) Get(index int) (interface{}, error) {
	if index < 0 || index >= int(list.size) {
        return nil, errors.New("Index out of range")
    }

    current := list.head
    for i := 0; i < index; i++ {
        current = current.next
    }

    return current.data, nil
}


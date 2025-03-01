package bintree

import (
	"errors"
)

type Node[T comparable] struct {
	data T
	left, right *Node[T]
}

type BinaryTree[T comparable] struct {
    root *Node[T]
}

func New[T comparable]() *BinaryTree[T] {
    return &BinaryTree[T]{}
}

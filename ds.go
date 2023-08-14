package main

import (
	"github.com/adelapazborrero/ds/queue"
	"github.com/adelapazborrero/ds/stack"
)

func NewQueue[T any]() *queue.Queue[T] {
	return queue.NewQueue[T]()
}

func NewStack[T any]() *stack.Stack[T] {
	return stack.NewStack[T]()
}

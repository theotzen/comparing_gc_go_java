package utils

import (
	"fmt"
	"go_gc_app/internal/metrics"
)

type Node struct {
	value int
	next  *Node
}

func GenerateList(size int) *Node {
	fmt.Printf("Generating a %v long LinkedList \n", size)

	metrics.ListsCreated.Inc()

	if size == 0 {
		return nil
	}
	head := &Node{value: size}
	current := head
	for i := 1; i < size; i++ {
		current.next = &Node{value: size - i}
		current = current.next
	}
	return head
}

func generateListRec(size int) *Node {
	if size == 0 {
		return nil
	}
	return &Node{
		value: size,
		next:  generateListRec(size - 1),
	}
}

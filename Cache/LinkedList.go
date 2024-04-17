package cache

import (
	"fmt"
	"sync"
	"time"
)

type LinkedList struct {
	head     *Node
	Tail     *Node
	size     int
	Capacity int
	mu       sync.Mutex
}

func NewLinkedList(capacity int) *LinkedList {
	ll := &LinkedList{
		head:     &Node{},
		Tail:     &Node{},
		size:     0,
		Capacity: capacity,
	}

	ll.head.Next = ll.Tail
	ll.Tail.Prev = ll.head

	return ll
}

func (ll *LinkedList) remove(node *Node) {
	ll.mu.Lock()
	defer ll.mu.Unlock()
	prev := node.Prev
	next := node.Next

	prev.Next = next
	next.Prev = prev

	node.Next = nil
	node.Prev = nil

	ll.size--

}

func (ll *LinkedList) add(key string, value interface{}) {
	ll.mu.Lock()
	defer ll.mu.Unlock()

	next := ll.head.Next
	node := &Node{Key: key, Value: value, TimeAccessed: time.Now()}
	next.Prev = node
	node.Next = next
	node.Prev = ll.head
	ll.head.Next = node
	ll.size++
}

func (ll *LinkedList) swap(node *Node) {
	value := node.Value
	key := node.Key

	ll.remove(node)
	ll.add(key, value)
}

func (ll *LinkedList) printList() {
	temp := ll.head.Next
	fmt.Println()
	for temp.Next != nil {
		fmt.Printf("%s:%v -> ", temp.Key, temp.Value)
		temp = temp.Next
	}
	fmt.Println()
}

package cache

import "time"

type Node struct {
	Prev         *Node
	Next         *Node
	Key          string
	Value        interface{}
	TimeAccessed time.Time
}

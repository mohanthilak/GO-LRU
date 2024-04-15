package cache

type Node struct {
	Prev  *Node
	Next  *Node
	Key   int
	Value int
}

package avl

// Tree struct represents a AVL BinarySearch Tree (BST)
type Tree[K Ordered, V any] struct {
	RootNode *Node[K, V] //The root node of the Tree
}

// NewTree() return an empty new Tree
func NewTree[K Ordered, V any]() *Tree[K, V] {
	return &Tree[K, V]{}
}

// Size() returns the size (number of Nodes) of the Tree
// Basically, it delegates the Size to its RootNode (or returns 0)
func (t *Tree[K, V]) Size() int {
	if t.RootNode == nil {
		return 0
	}
	return t.RootNode.Size()

}

// Depth() returns the depth of the Tree (the maximum iteration for searching a Node)
// Basically, it delegates the Size to its RootNode (or returns 0)
func (t *Tree[K, V]) Depth() int {
	if t.RootNode == nil {
		return 0
	}
	return t.RootNode.Depth()
}

// AddOne() add one element in the Tree
func (t *Tree[K, V]) AddOne(key K, value V) bool {

	if t.RootNode == nil {
		t.RootNode = &Node[K, V]{Key: key, Value: value}
	} else {
		t.RootNode = t.RootNode.Add(key, value)
	}
	return true
}

// Add() adds elements to the Node.
func (t *Tree[K, V]) Add(items ...struct {
	key   K
	value V
}) (addedItems int) {

	for _, item := range items {
		if ok := t.AddOne(item.key, item.value); ok {
			addedItems++
		}
	}

	return
}

// Get() returns the value present in the tree for the key
func (t *Tree[K, V]) Get(key K) (value V, ok bool) {
	if t.RootNode == nil {
		return
	}
	if foundNode := t.RootNode.Get(key); foundNode != nil {
		return t.RootNode.Get(key).Value, true
	} else {
		return
	}
}

// Remove() remove a key in the tree
func (t *Tree[K, V]) RemoveOne(key K) bool {

	if t.RootNode == nil {
		return false
	}

	if foundNode := t.RootNode.Get(key); foundNode != nil {
		newRoot := foundNode.Remove()
		t.RootNode = newRoot
		return true
	} else {
		return false
	}

}

func (t *Tree[K, V]) Remove(keys ...K) (removedItems int) {
	for _, key := range keys {
		if ok := t.RemoveOne(key); ok {
			removedItems++
		}
	}
	return
}

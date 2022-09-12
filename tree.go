package avl

import (
	"sync"
)

// Tree struct represents a AVL BinarySearch Tree (BST)
type Tree[K Ordered, V any] struct {
	sync.RWMutex             //RWMutex for preventing concurrent writing operations
	RootNode     *Node[K, V] //The root node of the Tree
	deletedNodes int
}

// NewTree() return an empty new Tree
func NewTree[K Ordered, V any]() *Tree[K, V] {
	return &Tree[K, V]{}
}

// Size() returns the size (number of Nodes) of the Tree
// Basically, it delegates the Size to its RootNode (or returns 0)
func (t *Tree[K, V]) Size() int {
	t.RLock()
	defer t.RUnlock()
	if t.RootNode == nil {
		return 0
	}
	return t.RootNode.Size()

}

// Depth() returns the depth of the Tree (the maximum iteration for searching a Node)
// Basically, it delegates the Size to its RootNode (or returns 0)
func (t *Tree[K, V]) Depth() int {
	t.RLock()
	defer t.RUnlock()

	if t.RootNode == nil {
		return 0
	}
	return t.RootNode.Depth()
}

// Print() returns the ordered array of non-deleted nodes in the tree
// depth represents the depth in which print the elements (0 for all depths)
func (t *Tree[K, V]) Print(depth uint) (nodes []*Node[K, V]) {
	t.RLock()
	defer t.RUnlock()

	if t.RootNode == nil || depth > uint(t.Depth()) {
		return nodes
	}
	return t.RootNode.Print(depth, 1)
}

// PrintKeys() act like Print but returns only the ordered array of non-deleted keys in the tree
func (t *Tree[K, V]) PrintKeys(depth uint) (keys []K) {
	nodes := t.Print(depth)
	for _, n := range nodes {
		keys = append(keys, n.Key)
	}
	return keys
}

// PrintValues() act like Print but returns only the ordered array of non-deleted values in the tree
func (t *Tree[K, V]) PrintValues(depth uint) (values []V) {
	nodes := t.Print(depth)

	for _, n := range nodes {
		values = append(values, n.Value)
	}
	return values
}

// NeedFlush() says if the tree needs to be flushed or not
func (t *Tree[K, V]) NeedFlush() bool {
	return t.deletedNodes != 0
}

// Flush() will rebase the tree, removing the deleted nodes
func (t *Tree[K, V]) Flush() bool {

	if !t.NeedFlush() {
		return false
	}

	var newRoot *Node[K, V]
	t.RLock()
	for depth := 1; depth <= t.Depth(); depth++ {
		for _, node := range t.Print(uint(depth)) {

			if newRoot == nil {
				newRoot = &Node[K, V]{Key: node.Key, Value: node.Value}
			} else {
				newRoot, _ = newRoot.Put(node.Key, node.Value)
			}
		}
	}
	t.RUnlock()

	t.Lock()
	defer t.Unlock()

	t.RootNode = newRoot
	t.deletedNodes = 0
	return true
}

// AddOne() add one element in the Tree[K,V]. It returns true if succeded
// If the key K is already present, its value is replaced
// Because adding an element can produce a re-balance of the tree, AddOne() will LOCK the tree
func (t *Tree[K, V]) PutOne(key K, value V) bool {
	t.Lock()
	defer t.Unlock()

	if t.RootNode == nil {
		t.RootNode = &Node[K, V]{Key: key, Value: value}
	} else {
		newRoot, unDeleteNode := t.RootNode.Put(key, value)
		t.RootNode = newRoot
		if unDeleteNode {
			t.deletedNodes--
		}
	}
	return true
}

// Add() adds elements `items` to the Node in a concurrent way
// It calls AddOne in goroutines. As AddOne `Lock()` the tree before adding something, Add() is safe to use
func (t *Tree[K, V]) Put(items ...struct {
	key   K
	value V
}) (addedItems int) {

	count := make(chan bool)

	for _, item := range items {
		go func(item struct {
			key   K
			value V
		}) {
			if ok := t.PutOne(item.key, item.value); ok {
				count <- true
			} else {
				count <- false
			}
		}(item)
	}

	for i := 0; i < len(items); i++ {
		result := <-count
		if result {
			addedItems++
		}
	}

	return
}

// Get() returns the value present in the tree for the key
func (t *Tree[K, V]) Get(key K) (value V, ok bool) {
	t.RLock()
	defer t.RUnlock()
	if t.RootNode == nil {
		return
	}
	if foundNode := t.RootNode.Get(key); foundNode != nil {
		return t.RootNode.Get(key).Value, true
	} else {
		return
	}
}

// Delete() marks as "deleted" the nodes corresponding to the passed keys
// returns false if all keys aren't present in the tree
func (t *Tree[K, V]) Delete(keys ...K) bool {

	foundNodes := make([]*Node[K, V], 0)

	for _, k := range keys {
		if foundNode := t.RootNode.Get(k); foundNode == nil {
			return false
		} else {
			foundNodes = append(foundNodes, foundNode)
		}
	}

	t.Lock()
	defer t.Unlock()

	for _, n := range foundNodes {
		if !n.Deleted {
			n.Deleted = true
			t.deletedNodes++
		}
	}

	return true
}

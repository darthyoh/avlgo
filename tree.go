package avlgo

import (
	"encoding/gob"
	"fmt"
	"os"
	"sync"
)

// Tree struct represents a AVL BinarySearch Tree (BST)
type Tree[K Ordered, V any] struct {
	rwMutex  sync.RWMutex //RWMutex for preventing concurrent writing operations
	RootNode *Node[K, V]  //The root node of the Tree
}

// NewTree() return an empty new Tree
func NewTree[K Ordered, V any]() *Tree[K, V] {
	return &Tree[K, V]{}
}

// Encode() serialize the tree in gob format
func (t *Tree[K, V]) Encode(output string) error {
	//simply encode the tree structure
	file, err := os.Create(output)
	if err != nil {
		return fmt.Errorf("unable to create the output file : %s", err)
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	if err = encoder.Encode(t); err != nil {
		return fmt.Errorf("unable to encode tree : %s", err)
	}
	return nil
}

// Decode() deserialize a tree from an input file
func Decode[K Ordered, V any](input string) (*Tree[K, V], error) {
	//first, decode the tree
	tree := NewTree[K, V]()
	file, err := os.Open(input)
	if err != nil {
		return nil, fmt.Errorf("unable to open the input file : %s", err)
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)

	if err = decoder.Decode(&tree); err != nil {
		return nil, fmt.Errorf("unable to decode tree : %s", err)
	}

	//then, re-build the "parent" field of each Node (parent field is private, so not encoded by the Encode() method to prevent infinite loop while encoding)
	if tree.RootNode != nil {
		if !tree.RootNode.affectParentToChildren() {
			return nil, fmt.Errorf("unable to decode tree : %s", err)
		}
	}

	return tree, nil
}

// Size() returns the size (number of Nodes) of the Tree
// Basically, it delegates the Size to its RootNode (or returns 0)
func (t *Tree[K, V]) Size() int {
	t.rwMutex.RLock()
	defer t.rwMutex.RUnlock()
	if t.RootNode == nil {
		return 0
	}
	return t.RootNode.Size()

}

// Depth() returns the depth of the Tree (the maximum iteration for searching a Node)
// Basically, it delegates the Size to its RootNode (or returns 0)
func (t *Tree[K, V]) Depth() int {
	t.rwMutex.RLock()
	defer t.rwMutex.RUnlock()

	if t.RootNode == nil {
		return 0
	}
	return t.RootNode.Depth()
}

// Print() returns the ordered nodes in the tree
// depth represents the depth in which print the elements (0 for all depths)
func (t *Tree[K, V]) Print(depth uint) (nodes []*Node[K, V]) {
	t.rwMutex.RLock()
	defer t.rwMutex.RUnlock()

	if t.RootNode == nil || depth > uint(t.Depth()) {
		return nodes
	}
	return t.RootNode.Print(depth, 1)
}

// PrintKeys() act like Print but returns only the ordered array if keys in the tree
func (t *Tree[K, V]) PrintKeys(depth uint) (keys []K) {
	nodes := t.Print(depth)
	for _, n := range nodes {
		keys = append(keys, n.Key)
	}
	return keys
}

// PrintValues() act like Print but returns only the ordered array of values in the tree
func (t *Tree[K, V]) PrintValues(depth uint) (values []V) {
	nodes := t.Print(depth)

	for _, n := range nodes {
		values = append(values, n.Value)
	}
	return values
}

// AddOne() add one element in the Tree[K,V]. It returns true if succeded
// If the key K is already present, its value is replaced
// Because adding an element can produce a re-balance of the tree, AddOne() will LOCK the tree
func (t *Tree[K, V]) PutOne(key K, value V) bool {
	t.rwMutex.Lock()
	defer t.rwMutex.Unlock()

	if t.RootNode == nil {
		t.RootNode = &Node[K, V]{Key: key, Value: value}
	} else {
		newRoot := t.RootNode.Put(key, value)
		t.RootNode = newRoot
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

// GetFromTo() return an ordered slice of values for keys found between from and to (including bounds or not)
func (t *Tree[K, V]) GetFromTo(from, to K, boundsIncluded bool) (values []V) {
	t.rwMutex.RLock()
	defer t.rwMutex.RUnlock()

	if t.RootNode == nil {
		return
	}

	for _, node := range t.RootNode.GetFromTo(from, to, boundsIncluded) {
		values = append(values, node.Value)
	}

	return
}

// Get() returns the value present in the tree for the key
func (t *Tree[K, V]) Get(key K) (value V, ok bool) {
	t.rwMutex.RLock()
	defer t.rwMutex.RUnlock()
	if t.RootNode == nil {
		return
	}
	if foundNode := t.RootNode.Get(key); foundNode != nil {
		return t.RootNode.Get(key).Value, true
	} else {
		return
	}
}

// Delete() will remove the nodes corresponding to the passed keys
// and returns the number of nodes deleted
func (t *Tree[K, V]) Delete(keys ...K) int {

	deleted := 0

	for _, k := range keys {
		if foundNode := t.RootNode.Get(k); foundNode != nil {
			t.rwMutex.Lock()
			t.RootNode = foundNode.Delete()
			deleted++
			t.rwMutex.Unlock()
		}
	}

	return deleted
}

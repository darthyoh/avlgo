package avl

import "sync"

// Node is one element of a Tree
type Node[K Ordered, V any] struct {
	sync.Mutex
	Key                    K           // Key of the Node must be ordered
	Value                  V           // Value of the Node can be anything
	Parent, Previous, Next *Node[K, V] // Parent, Previous and Next are references to other Node in the Tree
}

// Size() returns the Size of the node + its children
// it returns 1 + recursive size of its children
func (n *Node[K, V]) Size() int {

	size := 1
	if n.Previous != nil {
		size += n.Previous.Size()
	}
	if n.Next != nil {
		size += n.Next.Size()
	}
	return size
}

// Depth() returns the depth of the tree from this node
// it returns 1 + the biggest depth of its children
func (n *Node[K, V]) Depth() int {

	previousDepth := 0
	if n.Previous != nil {
		previousDepth = n.Previous.Depth()
	}
	nextDepth := 0
	if n.Next != nil {
		nextDepth = n.Next.Depth()
	}

	switch {
	case previousDepth > nextDepth:
		return 1 + previousDepth
	default:
		return 1 + nextDepth
	}

}

// Add() add a new Node in the tree, preserving the order and the balance of the Tree
func (n *Node[K, V]) Add(key K, value V) *Node[K, V] {
	switch {
	case key > n.Key: //key is bigger than the n.Key
		if n.Next != nil { //delegates to its Next (if exist)
			return n.Next.Add(key, value)
		}
		//otherwise : create a new Node and affect to its next
		n.Next = &Node[K, V]{Key: key, Parent: n, Value: value}
		return n.balance()

	case key < n.Key: //key is smaller than the n.Key
		if n.Previous != nil { //delegates to its Previous (if exist)
			return n.Previous.Add(key, value)
		}
		//otherwise : create a new Node and affect to its previiys
		n.Previous = &Node[K, V]{Key: key, Parent: n, Value: value}
		return n.balance()

	default: //key is the same than the n.Key so replace the Value
		n.Value = value
		return n.RootNode()
	}
}

// RootNode returns the root node of the tree
// (recursive call to the node which has no parent)
func (n *Node[K, V]) RootNode() *Node[K, V] {
	if n.Parent != nil {
		return n.Parent.RootNode()
	}
	return n
}

// getBalance() returns the difference between next depth and previous depth
// A node will be balanced if this difference is -1, 0 or +1
func (n *Node[K, V]) getBalance() int {

	previousBalance := 0
	nextBalance := 0
	if n.Previous != nil {
		previousBalance = n.Previous.Depth()
	}
	if n.Next != nil {
		nextBalance = n.Next.Depth()
	}

	return nextBalance - previousBalance
}

// balance() balance a node. If the node is unbalanced, it will perform one (or two) rotation
// and returns the new root node
func (n *Node[K, V]) balance() *Node[K, V] {

	balance := n.getBalance()

	//case of balanced node : recursive call to balance() to its parent
	if balance >= -1 && balance <= 1 {
		if n.Parent == nil {
			return n
		}
		return n.Parent.balance()
	}
	if balance > 1 { //unbalanced node with deeper Next
		if n.Next.getBalance() < 0 { //double rotation (to avoir infinite rotation)
			n.Next.rotateRight()
		}
		n.rotateLeft()
	} else if balance < -1 { //unbalanced node with deeper Previous
		if n.Previous.getBalance() > 0 { //double rotation (to avoir infinite rotation)
			n.Previous.rotateLeft()
		}
		n.rotateRight()
	}

	//recursive balance on parent
	if n.Parent == nil {
		return n
	}
	return n.Parent.balance()

}

// rotateRight() rotates the node to the right
func (n *Node[K, V]) rotateRight() {
	if n.Parent == nil {
		n.Previous.Parent = nil
	} else {
		n.Previous.Parent = n.Parent

		if n.Parent.Next == n {
			n.Parent.Next = n.Previous
		} else {
			n.Parent.Previous = n.Previous
		}
	}

	n.Parent = n.Previous

	if n.Previous.Next != nil {
		n.Previous = n.Previous.Next
		n.Previous.Parent = n
	} else {
		n.Previous = nil
	}
	n.Parent.Next = n
}

// rotateRight() rotates the node to the left
func (n *Node[K, V]) rotateLeft() {
	if n.Parent == nil {
		n.Next.Parent = nil
	} else {
		n.Next.Parent = n.Parent

		if n.Parent.Next == n {
			n.Parent.Next = n.Next
		} else {
			n.Parent.Previous = n.Next
		}
	}
	n.Parent = n.Next

	if n.Next.Previous != nil {
		n.Next = n.Next.Previous
		n.Next.Parent = n
	} else {
		n.Next = nil
	}
	n.Parent.Previous = n
}

// Get() search in the node the value of the key and returns it if present
func (n *Node[K, V]) Get(key K) *Node[K, V] {
	switch {
	case key > n.Key: //key is bigger than the n.Key
		if n.Next != nil { //delegates to its Next (if exist)
			return n.Next.Get(key)
		}
		//otherwise : the key isn't present in the Tree
		return nil

	case key < n.Key: //key is smaller than the n.Key
		if n.Previous != nil { //delegates to its Previous (if exist)
			return n.Previous.Get(key)
		}
		//otherwise : the key isn't present in the Tree
		return nil

	default: //This is the key !
		return n
	}
}

func (n *Node[K, V]) biggestChild() *Node[K, V] {
	if n.Next == nil {
		return n
	}
	return n.Next.biggestChild()
}

func (n *Node[K, V]) smallestChild() *Node[K, V] {
	if n.Previous == nil {
		return n
	}
	return n.Previous.biggestChild()
}

// Remove() removes a node with the key
func (n *Node[K, V]) Remove() *Node[K, V] {

	//the node is a leaf, simply delete it
	if n.Next == nil && n.Previous == nil {
		if n.Parent == nil {
			return nil
		} else {
			if n.Parent.Next == n {
				n.Parent.Next = nil
			} else {
				n.Parent.Previous = nil
			}
			newRoot := n.Parent.balance()
			n.Parent = nil
			return newRoot
		}
	}

	//the node has an only child
	if (n.Next == nil && n.Previous != nil) || (n.Next != nil && n.Previous == nil) {
		child := n.Next
		if child == nil {
			child = n.Previous
		}

		n.Next, n.Previous = nil, nil

		if n.Parent == nil {
			child.Parent = nil
			return child
		} else {
			if n.Parent.Next == n {
				n.Parent.Next = child
			} else {
				n.Parent.Previous = child
			}
			child.Parent = n.Parent
			newRoot := n.Parent.balance()
			n.Parent = nil
			return newRoot
		}
	}

	//the node has a previous and a next
	//check the deeper child

	previousDepth, nextDepth := n.Previous.Depth(), n.Next.Depth()

	successor, attachment, location := n.Previous, n.Next, n.Previous.biggestChild()
	if nextDepth > previousDepth {
		successor, attachment, location = n.Next, n.Previous, n.Next.smallestChild()
	}

	n.Previous = nil
	n.Next = nil
	attachment.Parent = location
	if previousDepth >= nextDepth {
		location.Next = attachment
	} else {
		location.Previous = attachment
	}

	successor.Parent = n.Parent
	if n.Parent != nil {
		if n.Parent.Next == n {
			n.Parent.Next = successor
		} else {
			n.Parent.Previous = successor
		}
	}

	n.Parent = nil

	return location.balance()
}

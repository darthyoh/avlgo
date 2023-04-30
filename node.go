package avlgo

import (
	"encoding/json"
	"fmt"
	"sync"
)

// jsonNode struct is a wrapper struct used for Marshalling and Unmarshalling operations
type jsonNode[K Ordered, V any] struct {
	ID         string `json:"ID"`
	Key        K      `json:"key"`
	Value      V      `json:"value"`
	ParentID   string `json:"parentId,omitempty"`
	PreviousID string `json:"previousId,omitempty"`
	NextID     string `json:"nextId,omitempty"`
}

// Node is one element of a Tree
type Node[K Ordered, V any] struct {
	sync.Mutex
	Key                    K           // Key of the Node must be ordered
	Value                  V           // Value of the Node can be anything
	Parent, Previous, Next *Node[K, V] // Parent, Previous and Next are references to other Node in the Tree
}

// MarshalJSON **flats** the node replacing pointers to adresses as ID
func (n *Node[K, V]) MarshalJSON() ([]byte, error) {

	marshalNode := jsonNode[K, V]{
		ID:    fmt.Sprintf("%p", n),
		Key:   n.Key,
		Value: n.Value,
	}
	if n.Parent != nil {
		marshalNode.ParentID = fmt.Sprintf("%p", n.Parent)
	}
	if n.Previous != nil {
		marshalNode.PreviousID = fmt.Sprintf("%p", n.Previous)
	}
	if n.Next != nil {
		marshalNode.NextID = fmt.Sprintf("%p", n.Next)
	}
	return json.Marshal(marshalNode)
}

func (n *Node[K, V]) UnmarshalJSON(data []byte) error {
	marshalNode := &struct {
		ID         string `json:"ID"`
		Key        K      `json:"key"`
		Value      V      `json:"value"`
		ParentID   string `json:"parentId,omitempty"`
		PreviousID string `json:"previousId,omitempty"`
		NextID     string `json:"nextId,omitempty"`
	}{}

	if err := json.Unmarshal(data, &marshalNode); err != nil {
		return err
	}

	return nil
}

// Size() returns the Size of the node + its children
// it returns 1 + recursive size of its children
func (n *Node[K, V]) Size() (size int) {
	if n.Previous != nil {
		size += n.Previous.Size()
	}
	if n.Next != nil {
		size += n.Next.Size()
	}
	return 1 + size
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

func (n *Node[K, V]) Print(wantedDepth, actualDepth uint) (nodes []*Node[K, V]) {
	if wantedDepth != 0 {
		if wantedDepth == actualDepth {
			nodes = append(nodes, n)
			return nodes
		} else {
			if n.Previous != nil {
				nodes = n.Previous.Print(wantedDepth, actualDepth+1)
			}
			if n.Next != nil {
				nodes = append(nodes, n.Next.Print(wantedDepth, actualDepth+1)...)
			}
			return nodes
		}
	}

	if n.Previous != nil {
		nodes = n.Previous.Print(wantedDepth, actualDepth)
	}
	nodes = append(nodes, n)

	if n.Next != nil {
		nodes = append(nodes, n.Next.Print(wantedDepth, actualDepth)...)
	}
	return
}

// Put() add a new Node in the tree, preserving the order and the balance of the Tree
func (n *Node[K, V]) Put(key K, value V) (newRootNode *Node[K, V]) {
	switch {
	case key > n.Key: //key is bigger than the n.Key
		if n.Next != nil { //delegates to its Next (if exist)
			return n.Next.Put(key, value)
		}
		//otherwise : create a new Node and affect to its next
		n.Next = &Node[K, V]{Key: key, Parent: n, Value: value}
		return n.balance()

	case key < n.Key: //key is smaller than the n.Key
		if n.Previous != nil { //delegates to its Previous (if exist)
			return n.Previous.Put(key, value)
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

// GetFromTo() search in the node the value of the key between from and to and returns them
func (n *Node[K, V]) GetFromTo(from, to K, boundsIncluded bool) []*Node[K, V] {
	nodes := make([]*Node[K, V], 0)
	if n.Key > from && n.Previous != nil {
		nodes = append(nodes, n.Previous.GetFromTo(from, to, boundsIncluded)...)
	}

	if (n.Key > from && n.Key < to) || ((n.Key == to || n.Key == from) && boundsIncluded) {
		nodes = append(nodes, n)
	}

	if n.Key < to && n.Next != nil {
		nodes = append(nodes, n.Next.GetFromTo(from, to, boundsIncluded)...)
	}

	return nodes
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

// Delete() will delete the node if the key is found and returns the new RootNode
func (n *Node[K, V]) Delete() *Node[K, V] {

	switch {
	case n.Next == nil && n.Previous == nil: //The node to delete is a leaf... Simply delete it !
		if n.Parent == nil { //the node to delete is the only node (and the root node...) simply return nil informing the tree that there's no more node
			return nil
		}
		if n.Parent.Previous == n {
			n.Parent.Previous = nil
		} else {
			n.Parent.Next = nil
		}
		return n.Parent.balance()
	case (n.Next == nil && n.Previous != nil) || (n.Next != nil && n.Previous == nil): //The node has only one child
		if n.Parent == nil { //the node to delete is the rootnode, so simply return its only child has new root node
			if n.Previous != nil {
				n.Previous.Parent = nil
				newRoot := n.Previous
				n.Previous = nil
				return newRoot
			} else {
				n.Next.Parent = nil
				newRoot := n.Next
				n.Next = nil
				return newRoot
			}
		} else {
			//Get the child
			successor := n.Previous
			if n.Previous == nil {
				successor = n.Next
			}

			successor.Parent = n.Parent

			if n.Parent.Next == n {
				n.Parent.Next = successor
			} else {
				n.Parent.Previous = successor
			}
			n.Parent = nil
			n.Next = nil
			n.Previous = nil

			return successor.Parent.balance()
		}
	default: //the node to delete has two children
		return nil
	}

}

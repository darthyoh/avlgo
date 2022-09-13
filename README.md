# avl-bst
A simple implementation of a AVL tree in **GO**, using generics.

- [avl-bst](#avl-bst)
  - [Introduction](#introduction)
  - [Installation](#installation)
  - [Basic usage](#basic-usage)
  - [Implementation decisions](#implementation-decisions)

## Introduction

An AVL Tree is an ordered data structure where all nodes (a key / value container) are linked to their unique parent node, and eventualy to a left unique child (with a smaller key) and a unique right child (with a bigger key). The tree has only one root node.

You then use recursive functions to walk thru your tree until you find the key you were looking for (or nil if you came at the end (the leaf) of your tree without finding it)

Adding a new Node to the tree is pretty simple : you compare the key you want to insert with the key of the node : if the first is bigger, you pass the requested key to the *next child* (or add the new node as the next child of this node). The same is done if the key is smaller but with the *previous child*. The added node is **always** a leaf.

See what happen :

```
//Adding 4

            4

//Adding 2

            4
           /
          2

//Adding 6

            4
           / \
          2   6

//Adding 3

            4
           / \
          2   6
           \
            3

//Adding 5

            4
           / \
          /   \
         2     6
          \   /
           3 5

//Adding 8

            4
           / \
          /   \
         2     6
          \   / \
           3 5   8

//Adding 1

            4
           / \
          /   \
         2     6
        / \   / \
       1   3 5   8

//Adding 9

            4
           / \
          /   \
         2     6
        / \   / \
       1   3 5   8
                  \
                   9

//Adding 7

            4
           / \
          /   \
         2     6
        / \   / \
       1   3 5   8
                / \
               7   9
```

Searching in a such structure is easy and fast : you can get any value in a few iterations number (depending of the **Depth** of the tree). In our example, 4 iterations is needed to find every key. It combines the rapididy of a binary search in an array with the non-reallocation ability of a linked list.

But, what happens if we decided to add these values in another order ? Consider adding 0,1,2,3,4,5,6,7,8,9 :

```

0
 \
  1
   \
    2
     \
      3
       \
        4
         \
          5
           \
            6
             \
              7
               \
                8
                 \
                  9

```

Our tree has the same **size** (10) than above but with a **depth** of 10 ! The benefit of the tree is lost !!!

That's where the AVL Trees come !

In an AVL Tree, each node must be **balanced** : the difference between the depth of its Next child and the depth of its Previous child must be between -1 and 1. In other case, the tree should be re-arranged performing 1 or 2 rotation around the node :

```
0
 \
  1
   \
    2

// Is unbalanced : the 0 node has a balance of +2, so rotate the tree around 0 :

  1
 / \
0   2

// OK.... now, add 3
  1
 / \
0   2
     \
      3

// Fine.... Add 4
  1
 / \
0   2
     \
      3
       \
        4

// Arggg.... 2 is unbalanced : rotate around 2
  1
 / \
0   3
   / \ 
  2   4

// Fine ! And so on...
```  

When you add a new node (is a leaf) you're sure that this node in balanced, so you have to recursivly test the parent balance and perform 1 or 2 rotations. The rotation ensures the tree stay ordered and balanced.

So you have to test the balance of the tree after adding or removing node in the tree (in THIS implementation, there's something different happening when you remove a node, see [Implementation decisions](#implementation-decisions) to know why)


See [Wikipedia](https://en.wikipedia.org/wiki/AVL_tree) for more infos about what is an AVL Tree.

In this implementation, the `Tree` struct represents our AVL. It is composed of `Node` structs, each of them is linked to its `Parent`, and eventually to their `Previous` and `Next` child.



## Installation

Go to your project and download the dependency with :

```
cd myProject
go get github.com/darthyoh/avl-bst
```

Then, in your "main.go" file, you can import it with :

```
package main

import (
    "github.com/darthyoh/avl-bst"
)

func main() {
    //use the avl here
}
```

## Basic usage

You can use the `avl.NewTree()` utility to get a new `*Node[K,V]` and then use `Put()` or `PutOne()` methods to add some values :

```
// the Tree is generic about the Key and Value types
tree := avl.NewTree[int, int]()

tree.PutOne(0,0)

items := make([]struct {
		key   int
		value int
	}, 0)

ITEMS := 10

for i := 1; i < ITEMS; i++ {
	items = append(items, struct {
		key   int
		value int
	}{key: i, value: i})
}

tree.Put(items...)
```

Use the `Size()` method to read the size of the tree and the `Depth()` method for the depth :

```
fmt.Println(tree.Size()) // 10
fmt.Println(tree.Depth()) // 4
```

Use the `PrintKeys()` or `PrintValues` to get the ordered keys or values :

```
fmt.Println(reflect.DeepEqual([]int{0,1,2,3,4,5,6,7,8,9}), tree.PrintKeys()) //true
fmt.Println(reflect.DeepEqual([]int{0,1,2,3,4,5,6,7,8,9}), tree.PrintValues()) //true
```

Pass the above methods the depth you want to read :

```
fmt.Println(reflect.DeepEqual([]int{1,3,5,8}), tree.PrintKeys(3)) //true
fmt.Println(reflect.DeepEqual([]int{1,3,5,8}), tree.PrintValues(3)) //true
```

Use the `Delete()` method to delete some keys :

```
tree.Delete(5,6,7)
fmt.Println(tree.Size()) // 7
fmt.Println(tree.Depth()) // 4 ????? We will see later why !!!
```

*Restore* deleted values with `Put()` or `PutOne()` :

```
tree.PutOne(5)
tree.PutOne(6)
tree.PutOne(7)

fmt.Println(tree.Size()) // 10
fmt.Println(tree.Depth()) // 4
```

Delete again some values and `Flush()` the tree to balance it. Use `NeedFlush()` to see if the tree should be balanced 

```
fmt.Println(tree.NeedFlush()) //false
tree.Delete(5,6,7)
fmt.Println(tree.NeedFlush()) //true
tree.Flush()
fmt.Println(tree.Size()) // 7
fmt.Println(tree.Depth()) // 3 !!! The tree is balanced
```



## Implementation decisions

We decided to use `Put()`, `Get()` and `Delete()` method names, like classical **HTTP methods**


`Tree` and `Node` structs don't have some `Post()` method. Only `Put()`. A `Put()` call will add the key/value to the tree if not exists or **replace** the value if exists.


Putting some values will cause immediat re-balancing (with one or two rotations), meaning that the `Tree` couldn't be accessed.


Deleting some values will **not** really delete them : if fact, each `Node` has a special boolean tag called `Deleted`. When you `Delete()` the key, the corresponding node set its deleted parameter to `true`. It will not be printed by `Print()`, `PrintKeys()` and `PrintValues()`, and it will not be count by `Size()`. But, it continues to be present in the tree, that's why `Depth()` returns the same depth ! Why ?

Deleting a Node (aka re-link parent and childs to other nodes) will cause unbalance and force the methods to perform rotations, as the `Put()` method does. But after deleting a key, if you want to put it back, you have to re-balance the Tree again !!!!

We decided to self-balance the tree when putting a new value on it, but not when we deleting it. We simply mark it as *deleted* and inform the Tree that one value was deleted (increment counter). When you `Put()` a key that has been deleted before, the deleted flag comes back to false and the counter is decremented. The `NeedFlush()` method simply return if the counter is 0. In this way, the tree is not many time rebalanced when deleting / putting back a value. The developer has the responsability to call `Flush()` when he juges the tree has to be balanced.

Because `Add()` and `Delete()` modify the structure or this content, it should block the code : if a `Get()` method (or a `Size()` or `Depth()`) is running, adding or deleting should wait that the getting process is done. But getting datas in parallel are not a problem. That's why the Tree acts like a `sync.RWMutex` : reading functions `RLock()` and `defer RUnlock()`, and adding and deleting functions `Lock()` and `defer Unlock()`

There are no benefit using add and delete in goroutines !

Marshalling and Unmarshalling are enable :
- when marshalling, the `*Node` arborescence is flatten : the json provides an array of Node objects where pointers to Parent, Next and Previous are replaced with the memory allocation
- when unmarshalling, the `*Node` arborescence is turned back to a pointer architecture

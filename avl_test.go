package avl

import (
	"sync"
	"testing"
)

func TestEmptyTree(t *testing.T) {
	tree := NewTree[int, bool]()
	if tree.Size() != 0 {
		t.Errorf("Tree size is %d, want 0", tree.Size())
	}
	if tree.Depth() != 0 {
		t.Errorf("Tree depth is %d, want 0", tree.Depth())
	}
}

func TestAddOneingOneValue(t *testing.T) {
	tree := NewTree[int, bool]()
	tree.AddOne(1, true)
	if tree.Size() != 1 {
		t.Errorf("Tree size is %d, want 1", tree.Size())
	}
	if tree.Depth() != 1 {
		t.Errorf("Tree depth is %d, want 1", tree.Depth())
	}
	if tree.RootNode.Key != 1 {
		t.Errorf("RootNode is %d, want 1", tree.RootNode.Key)
	}
}

func TestAddOneingMoreValues(t *testing.T) {

	tree := NewTree[int, bool]()
	tree.AddOne(2, true)
	tree.AddOne(1, true)
	tree.AddOne(3, true)
	if tree.Size() != 3 {
		t.Errorf("Tree size is %d, want 3", tree.Size())
	}
	if tree.Depth() != 2 {
		t.Errorf("Tree depth is %d, want 2", tree.Depth())
	}
	if tree.RootNode.Key != 2 {
		t.Errorf("RootNode is %d, want 2", tree.RootNode.Key)
	}
	tree = NewTree[int, bool]()
	tree.AddOne(4, true)
	tree.AddOne(5, true)
	tree.AddOne(2, true)
	tree.AddOne(3, true)
	tree.AddOne(1, true)
	if tree.Size() != 5 {
		t.Errorf("Tree size is %d, want 3", tree.Size())
	}
	if tree.Depth() != 3 {
		t.Errorf("Tree depth is %d, want 2", tree.Depth())
	}
	if tree.RootNode.Key != 4 {
		t.Errorf("RootNode is %d, want 4", tree.RootNode.Key)
	}

}

func TestAddOneingMoreValuesThatUnbalanceTree(t *testing.T) {
	tree := NewTree[int, bool]()
	tree.AddOne(1, true)
	tree.AddOne(2, true)
	tree.AddOne(3, true)
	if tree.Size() != 3 {
		t.Errorf("Tree size is %d, want 3", tree.Size())
	}
	if tree.Depth() != 2 {
		t.Errorf("Tree depth is %d, want 2", tree.Depth())
	}
	if tree.RootNode.Key != 2 {
		t.Errorf("RootNode is %d, want 2", tree.RootNode.Key)
	}

	tree.AddOne(4, true)
	tree.AddOne(5, true)
	if tree.Size() != 5 {
		t.Errorf("Tree size is %d, want 5", tree.Size())
	}
	if tree.Depth() != 3 {
		t.Errorf("Tree depth is %d, want 3", tree.Depth())
	}
	if tree.RootNode.Key != 2 {
		t.Errorf("RootNode is %d, want 2", tree.RootNode.Key)
	}

	tree.AddOne(6, true)
	if tree.Size() != 6 {
		t.Errorf("Tree size is %d, want 6", tree.Size())
	}
	if tree.Depth() != 3 {
		t.Errorf("Tree depth is %d, want 3", tree.Depth())
	}
	if tree.RootNode.Key != 4 {
		t.Errorf("RootNode is %d, want 4", tree.RootNode.Key)
	}

	tree.AddOne(7, true)
	if tree.Size() != 7 {
		t.Errorf("Tree size is %d, want 7", tree.Size())
	}
	if tree.Depth() != 3 {
		t.Errorf("Tree depth is %d, want 3", tree.Depth())
	}
	if tree.RootNode.Key != 4 {
		t.Errorf("RootNode is %d, want 4", tree.RootNode.Key)
	}

	tree.AddOne(8, true)
	tree.AddOne(9, true)
	if tree.Size() != 9 {
		t.Errorf("Tree size is %d, want 9", tree.Size())
	}
	if tree.Depth() != 4 {
		t.Errorf("Tree depth is %d, want 4", tree.Depth())
	}
	if tree.RootNode.Key != 4 {
		t.Errorf("RootNode is %d, want 4", tree.RootNode.Key)
	}

	tree.AddOne(10, true)
	if tree.Size() != 10 {
		t.Errorf("Tree size is %d, want 10", tree.Size())
	}
	if tree.Depth() != 4 {
		t.Errorf("Tree depth is %d, want 4", tree.Depth())
	}
	if tree.RootNode.Key != 4 {
		t.Errorf("RootNode is %d, want 4", tree.RootNode.Key)
	}

	tree.AddOne(1, true)
	tree.AddOne(2, true)
	tree.AddOne(3, true)
	tree.AddOne(4, true)
	tree.AddOne(5, true)
	tree.AddOne(6, true)
	tree.AddOne(7, true)
	tree.AddOne(8, true)
	tree.AddOne(9, true)
	tree.AddOne(10, true)
	if tree.Size() != 10 {
		t.Errorf("Tree size is %d, want 10", tree.Size())
	}
	if tree.Depth() != 4 {
		t.Errorf("Tree depth is %d, want 4", tree.Depth())
	}
	if tree.RootNode.Key != 4 {
		t.Errorf("RootNode is %d, want 4", tree.RootNode.Key)
	}
}

func TestAddOneingMoreValuesThatUnbalanceTreeString(t *testing.T) {
	tree := NewTree[string, bool]()
	tree.AddOne("a", true)
	tree.AddOne("b", true)
	tree.AddOne("c", true)
	if tree.Size() != 3 {
		t.Errorf("Tree size is %d, want 3", tree.Size())
	}
	if tree.Depth() != 2 {
		t.Errorf("Tree depth is %d, want 2", tree.Depth())
	}
	tree.AddOne("d", true)
	tree.AddOne("e", true)
	if tree.Size() != 5 {
		t.Errorf("Tree size is %d, want 5", tree.Size())
	}
	if tree.Depth() != 3 {
		t.Errorf("Tree depth is %d, want 3", tree.Depth())
	}
	tree.AddOne("f", true)
	if tree.Size() != 6 {
		t.Errorf("Tree size is %d, want 6", tree.Size())
	}
	if tree.Depth() != 3 {
		t.Errorf("Tree depth is %d, want 3", tree.Depth())
	}
	tree.AddOne("g", true)
	if tree.Size() != 7 {
		t.Errorf("Tree size is %d, want 7", tree.Size())
	}
	if tree.Depth() != 3 {
		t.Errorf("Tree depth is %d, want 3", tree.Depth())
	}
	tree.AddOne("h", true)
	tree.AddOne("i", true)
	if tree.Size() != 9 {
		t.Errorf("Tree size is %d, want 9", tree.Size())
	}
	if tree.Depth() != 4 {
		t.Errorf("Tree depth is %d, want 4", tree.Depth())
	}
	tree.AddOne("j", true)
	if tree.Size() != 10 {
		t.Errorf("Tree size is %d, want 10", tree.Size())
	}
	if tree.Depth() != 4 {
		t.Errorf("Tree depth is %d, want 4", tree.Depth())
	}
	tree.AddOne("a", true)
	tree.AddOne("b", true)
	tree.AddOne("c", true)
	tree.AddOne("d", true)
	tree.AddOne("e", true)
	tree.AddOne("f", true)
	tree.AddOne("g", true)
	tree.AddOne("h", true)
	tree.AddOne("i", true)
	tree.AddOne("j", true)
	if tree.Size() != 10 {
		t.Errorf("Tree size is %d, want 10", tree.Size())
	}
	if tree.Depth() != 4 {
		t.Errorf("Tree depth is %d, want 4", tree.Depth())
	}
}

func TestDoubleRotation(t *testing.T) {
	tree := NewTree[int, bool]()
	tree.AddOne(7, true)
	tree.AddOne(8, true)
	tree.AddOne(4, true)
	tree.AddOne(1, true)
	tree.AddOne(5, true)
	if tree.Size() != 5 {
		t.Errorf("Tree size is %d, want 5", tree.Size())
	}
	if tree.Depth() != 3 {
		t.Errorf("Tree depth is %d, want 3", tree.Depth())
	}
	if tree.RootNode.Key != 7 {
		t.Errorf("RootNode is %d, want 7", tree.RootNode.Key)
	}
	// At this point, the Tree is balanced without any rotation

	//AddOneing 6 will cause double rotation
	tree.AddOne(6, true)
	if tree.Size() != 6 {
		t.Errorf("Tree size is %d, want 5", tree.Size())
	}
	if tree.Depth() != 3 {
		t.Errorf("Tree depth is %d, want 3", tree.Depth())
	}
	if tree.RootNode.Key != 5 {
		t.Errorf("RootNode is %d, want 5", tree.RootNode.Key)
	}
}

func TestGettingSomeValues(t *testing.T) {
	tree := NewTree[int, string]()
	tree.AddOne(7, "g")
	tree.AddOne(8, "h")
	tree.AddOne(1, "a")
	tree.AddOne(3, "c")

	if value, ok := tree.Get(3); !ok || value != "c" {
		if !ok {
			t.Errorf("Get should find 3")
		} else {
			t.Errorf("Get returns %s, want c", value)
		}
	}

	if value, ok := tree.Get(8); !ok || value != "h" {
		if !ok {
			t.Errorf("Get should find 8")
		} else {
			t.Errorf("Get returns %s, want h", value)
		}
	}

	if value, ok := tree.Get(7); !ok || value != "g" {
		if !ok {
			t.Errorf("Get should find 7")
		} else {
			t.Errorf("Get returns %s, want g", value)
		}
	}

	if value, ok := tree.Get(1); !ok || value != "a" {
		if !ok {
			t.Errorf("Get should find 1")
		} else {
			t.Errorf("Get returns %s, want a", value)
		}
	}

	if _, ok := tree.Get(2); ok {
		t.Errorf("Get shouldn't find anything for key 2")
	}

	if _, ok := tree.Get(4); ok {
		t.Errorf("Get shouldn't find anything for key 4")
	}

	//changing a key
	tree.AddOne(1, "z")

	if value, ok := tree.Get(1); !ok || value != "z" {
		if !ok {
			t.Errorf("Get should find %d, nothing found", tree.Size())
		} else {
			t.Errorf("Get returns %s, want z", value)
		}
	}
}

func TestRemovingNonPresentValues(t *testing.T) {
	tree := NewTree[int, string]()
	tree.AddOne(7, "g")
	tree.AddOne(8, "h")
	tree.AddOne(1, "a")
	tree.AddOne(3, "c")
	if tree.Size() != 4 {
		t.Errorf("Tree size is %d, want 4", tree.Size())
	}
	if tree.Depth() != 3 {
		t.Errorf("Tree depth is %d, want 3", tree.Depth())
	}
	if tree.RootNode.Key != 7 {
		t.Errorf("RootNode is %d, want 7", tree.RootNode.Key)
	}

	if ok := tree.RemoveOne(2); ok {
		t.Errorf("Removing 2 should return false")
	}
	if ok := tree.RemoveOne(9); ok {
		t.Errorf("Removing 2 should return false")
	}
}

func TestRemovingALeafWithoutUnbalance(t *testing.T) {
	tree := NewTree[int, bool]()
	tree.AddOne(3, true)
	tree.AddOne(1, true)
	tree.AddOne(6, true)
	tree.AddOne(5, true)
	tree.AddOne(7, true)

	if tree.Size() != 5 {
		t.Errorf("Tree size is %d, want 5", tree.Size())
	}
	if tree.Depth() != 3 {
		t.Errorf("Tree depth is %d, want 3", tree.Depth())
	}
	if tree.RootNode.Key != 3 {
		t.Errorf("RootNode is %d, want 3", tree.RootNode.Key)
	}

	if ok := tree.RemoveOne(7); !ok {
		t.Errorf("Removing 7 should return true")
	}
	if tree.Size() != 4 {
		t.Errorf("Tree size is %d, want 4", tree.Size())
	}
	if tree.Depth() != 3 {
		t.Errorf("Tree depth is %d, want 3", tree.Depth())
	}
	if tree.RootNode.Key != 3 {
		t.Errorf("RootNode is %d, want 3", tree.RootNode.Key)
	}

	if ok := tree.RemoveOne(5); !ok {
		t.Errorf("Removing 5 should return true")
	}
	if tree.Size() != 3 {
		t.Errorf("Tree size is %d, want 3", tree.Size())
	}
	if tree.Depth() != 2 {
		t.Errorf("Tree depth is %d, want 2", tree.Depth())
	}
	if tree.RootNode.Key != 3 {
		t.Errorf("RootNode is %d, want 3", tree.RootNode.Key)
	}

	if ok := tree.RemoveOne(6); !ok {
		t.Errorf("Removing 6 should return true")
	}
	if tree.Size() != 2 {
		t.Errorf("Tree size is %d, want 2", tree.Size())
	}
	if tree.Depth() != 2 {
		t.Errorf("Tree depth is %d, want 2", tree.Depth())
	}
	if tree.RootNode.Key != 3 {
		t.Errorf("RootNode is %d, want 3", tree.RootNode.Key)
	}

	if ok := tree.RemoveOne(1); !ok {
		t.Errorf("Removing 1 should return true")
	}
	if tree.Size() != 1 {
		t.Errorf("Tree size is %d, want 3", tree.Size())
	}
	if tree.Depth() != 1 {
		t.Errorf("Tree depth is %d, want 2", tree.Depth())
	}
	if tree.RootNode.Key != 3 {
		t.Errorf("RootNode is %d, want 3", tree.RootNode.Key)
	}

	if ok := tree.RemoveOne(3); !ok {
		t.Errorf("Removing 3 should return true")
	}
	if tree.Size() != 0 {
		t.Errorf("Tree size is %d, want 0", tree.Size())
	}
	if tree.Depth() != 0 {
		t.Errorf("Tree depth is %d, want 0", tree.Depth())
	}
	if tree.RootNode != nil {
		t.Errorf("RootNode is %v, want nil", tree.RootNode)
	}
}

func TestRemovingALeafWithUnbalance(t *testing.T) {
	tree := NewTree[int, bool]()
	tree.AddOne(3, true)
	tree.AddOne(1, true)
	tree.AddOne(6, true)
	tree.AddOne(5, true)
	tree.AddOne(7, true)

	if tree.Size() != 5 {
		t.Errorf("Tree size is %d, want 5", tree.Size())
	}
	if tree.Depth() != 3 {
		t.Errorf("Tree depth is %d, want 3", tree.Depth())
	}
	if tree.RootNode.Key != 3 {
		t.Errorf("RootNode is %d, want 3", tree.RootNode.Key)
	}

	if ok := tree.RemoveOne(1); !ok {
		t.Errorf("Removing 1 should return true")
	}
	if tree.Size() != 4 {
		t.Errorf("Tree size is %d, want 4", tree.Size())
	}
	if tree.Depth() != 3 {
		t.Errorf("Tree depth is %d, want 3", tree.Depth())
	}
	if tree.RootNode.Key != 6 {
		t.Errorf("RootNode is %d, want 6", tree.RootNode.Key)
	}
}

func TestRemovingSingleChildWithoutBalance(t *testing.T) {
	tree := NewTree[int, bool]()

	tree.AddOne(2, true)
	tree.AddOne(1, true)
	tree.AddOne(3, true)
	tree.AddOne(4, true)

	if ok := tree.RemoveOne(3); !ok {
		t.Errorf("Removing 3 should return true")
	}
	if tree.Size() != 3 {
		t.Errorf("Tree size is %d, want 3", tree.Size())
	}
	if tree.Depth() != 2 {
		t.Errorf("Tree depth is %d, want 2", tree.Depth())
	}
	if tree.RootNode.Key != 2 {
		t.Errorf("RootNode is %d, want 2", tree.RootNode.Key)
	}

	tree = NewTree[int, bool]()

	tree.AddOne(2, true)
	tree.AddOne(1, true)
	tree.AddOne(4, true)
	tree.AddOne(3, true)

	if ok := tree.RemoveOne(4); !ok {
		t.Errorf("Removing 4 should return true")
	}
	if tree.Size() != 3 {
		t.Errorf("Tree size is %d, want 3", tree.Size())
	}
	if tree.Depth() != 2 {
		t.Errorf("Tree depth is %d, want 2", tree.Depth())
	}
	if tree.RootNode.Key != 2 {
		t.Errorf("RootNode is %d, want 2", tree.RootNode.Key)
	}

	tree = NewTree[int, bool]()

	tree.AddOne(3, true)
	tree.AddOne(4, true)
	tree.AddOne(2, true)
	tree.AddOne(1, true)

	if ok := tree.RemoveOne(2); !ok {
		t.Errorf("Removing 2 should return true")
	}
	if tree.Size() != 3 {
		t.Errorf("Tree size is %d, want 3", tree.Size())
	}
	if tree.Depth() != 2 {
		t.Errorf("Tree depth is %d, want 2", tree.Depth())
	}
	if tree.RootNode.Key != 3 {
		t.Errorf("RootNode is %d, want 3", tree.RootNode.Key)
	}

	tree = NewTree[int, bool]()

	tree.AddOne(3, true)
	tree.AddOne(4, true)
	tree.AddOne(1, true)
	tree.AddOne(2, true)

	if ok := tree.RemoveOne(1); !ok {
		t.Errorf("Removing 1 should return true")
	}
	if tree.Size() != 3 {
		t.Errorf("Tree size is %d, want 3", tree.Size())
	}
	if tree.Depth() != 2 {
		t.Errorf("Tree depth is %d, want 2", tree.Depth())
	}
	if tree.RootNode.Key != 3 {
		t.Errorf("RootNode is %d, want 3", tree.RootNode.Key)
	}

	tree = NewTree[int, bool]()

	tree.AddOne(1, true)
	tree.AddOne(2, true)
	if ok := tree.RemoveOne(1); !ok {
		t.Errorf("Removing 1 should return true")
	}
	if tree.Size() != 1 {
		t.Errorf("Tree size is %d, want 1", tree.Size())
	}
	if tree.Depth() != 1 {
		t.Errorf("Tree depth is %d, want 1", tree.Depth())
	}
	if tree.RootNode == nil || tree.RootNode.Key != 2 {
		t.Errorf("RootNode is %v, want 2", tree.RootNode)
	}

	tree = NewTree[int, bool]()
	tree.AddOne(2, true)
	tree.AddOne(1, true)
	if ok := tree.RemoveOne(2); !ok {
		t.Errorf("Removing 2 should return true")
	}
	if tree.Size() != 1 {
		t.Errorf("Tree size is %d, want 1", tree.Size())
	}
	if tree.Depth() != 1 {
		t.Errorf("Tree depth is %d, want 1", tree.Depth())
	}
	if tree.RootNode.Key != 1 {
		t.Errorf("RootNode is %d, want 1", tree.RootNode.Key)
	}
}

func TestRemovingSingleChildWithBalance(t *testing.T) {
	tree := NewTree[int, bool]()
	tree.AddOne(9, true)
	tree.AddOne(5, true)
	tree.AddOne(10, true)
	tree.AddOne(6, true)
	tree.AddOne(12, true)
	tree.AddOne(4, true)
	tree.AddOne(3, true)

	if tree.Size() != 7 {
		t.Errorf("Tree size is %d, want 7", tree.Size())
	}
	if tree.Depth() != 4 {
		t.Errorf("Tree depth is %d, want 4", tree.Depth())
	}
	if tree.RootNode == nil || tree.RootNode.Key != 9 {
		t.Errorf("RootNode is %v, want 9", tree.RootNode)
	}

	if ok := tree.RemoveOne(10); !ok {
		t.Errorf("Removing 10 should return true")
	}
	if tree.Size() != 6 {
		t.Errorf("Tree size is %d, want 6", tree.Size())
	}
	if tree.Depth() != 3 {
		t.Errorf("Tree depth is %d, want 3", tree.Depth())
	}
	if tree.RootNode.Key != 5 {
		t.Errorf("RootNode is %d, want 5", tree.RootNode.Key)
	}

}

func TestRemovingWithBothChild(t *testing.T) {
	tree := NewTree[int, bool]()
	tree.AddOne(2, true)
	tree.AddOne(1, true)
	tree.AddOne(3, true)

	if ok := tree.RemoveOne(2); !ok {
		t.Errorf("Removing 2 should return true")
	}
	if tree.Size() != 2 {
		t.Errorf("Tree size is %d, want 2", tree.Size())
	}
	if tree.Depth() != 2 {
		t.Errorf("Tree depth is %d, want 2", tree.Depth())
	}
	if tree.RootNode.Key != 1 {
		t.Errorf("RootNode is %d, want 1", tree.RootNode.Key)
	}
}

func TestRemovingWithBothChildSuccessorPrevious(t *testing.T) {
	tree := NewTree[int, bool]()
	tree.AddOne(4, true)
	tree.AddOne(5, true)
	tree.AddOne(2, true)
	tree.AddOne(1, true)
	tree.AddOne(3, true)

	if ok := tree.RemoveOne(4); !ok {
		t.Errorf("Removing 4 should return true")
	}
	if tree.Size() != 4 {
		t.Errorf("Tree size is %d, want 4", tree.Size())
	}
	if tree.Depth() != 3 {
		t.Errorf("Tree depth is %d, want 3", tree.Depth())
	}
	if tree.RootNode.Key != 2 {
		t.Errorf("RootNode is %d, want 2", tree.RootNode.Key)
	}
}

func TestRemovingWithBothChildSuccessorNext(t *testing.T) {
	tree := NewTree[int, bool]()
	tree.AddOne(2, true)
	tree.AddOne(1, true)
	tree.AddOne(4, true)
	tree.AddOne(3, true)
	tree.AddOne(5, true)

	if ok := tree.RemoveOne(2); !ok {
		t.Errorf("Removing 2 should return true")
	}
	if tree.Size() != 4 {
		t.Errorf("Tree size is %d, want 4", tree.Size())
	}
	if tree.Depth() != 3 {
		t.Errorf("Tree depth is %d, want 3", tree.Depth())
	}
	if tree.RootNode.Key != 4 {
		t.Errorf("RootNode is %d, want 4", tree.RootNode.Key)
	}
}

func TestRemovingWithBothChildSuccessorPreviousMiddle(t *testing.T) {
	tree := NewTree[int, bool]()
	tree.AddOne(4, true)
	tree.AddOne(5, true)
	tree.AddOne(2, true)
	tree.AddOne(1, true)
	tree.AddOne(3, true)

	if ok := tree.RemoveOne(2); !ok {
		t.Errorf("Removing 2 should return true")
	}
	if tree.Size() != 4 {
		t.Errorf("Tree size is %d, want 4", tree.Size())
	}
	if tree.Depth() != 3 {
		t.Errorf("Tree depth is %d, want 3", tree.Depth())
	}
	if tree.RootNode.Key != 4 {
		t.Errorf("RootNode is %d, want 4", tree.RootNode.Key)
	}
}

func TestRemovingWithBothChildSuccessorNextMiddle(t *testing.T) {
	tree := NewTree[int, bool]()
	tree.AddOne(2, true)
	tree.AddOne(1, true)
	tree.AddOne(4, true)
	tree.AddOne(3, true)
	tree.AddOne(5, true)

	if ok := tree.RemoveOne(4); !ok {
		t.Errorf("Removing 4 should return true")
	}
	if tree.Size() != 4 {
		t.Errorf("Tree size is %d, want 4", tree.Size())
	}
	if tree.Depth() != 3 {
		t.Errorf("Tree depth is %d, want 3", tree.Depth())
	}
	if tree.RootNode.Key != 2 {
		t.Errorf("RootNode is %d, want 2", tree.RootNode.Key)
	}
}

func TestRemovingWithBothChildSuccessorWithRotations(t *testing.T) {
	tree := NewTree[int, bool]()
	tree.AddOne(5, true)
	tree.AddOne(2, true)
	tree.AddOne(9, true)
	tree.AddOne(7, true)
	tree.AddOne(3, true)
	tree.AddOne(1, true)
	tree.AddOne(4, true)

	if ok := tree.RemoveOne(5); !ok {
		t.Errorf("Removing 5 should return true")
	}
	if tree.Size() != 6 {
		t.Errorf("Tree size is %d, want 6", tree.Size())
	}
	if tree.Depth() != 3 {
		t.Errorf("Tree depth is %d, want 3", tree.Depth())
	}
	if tree.RootNode.Key != 3 {
		t.Errorf("RootNode is %d, want 3", tree.RootNode.Key)
	}
}

func TestAddOneingManyValues(t *testing.T) {

	tree := NewTree[int, int]()

	items := make([]struct {
		key   int
		value int
	}, 0)

	for i := 0; i < 2000; i++ {
		items = append(items, struct {
			key   int
			value int
		}{key: i, value: i})
	}

	tree.Add(items...)

	if tree.Size() != 2000 {
		t.Errorf("Tree size is %d, want 2000", tree.Size())
	}
	t.Logf("Tree depth is %d and RootNode key is %d", tree.Depth(), tree.RootNode.Key)
}

func TestAddOneingManyMoreValuesGetThemAndDeleteThem(t *testing.T) {

	tree := NewTree[int, int]()
	const RECORDS = 100000

	itemsToAdd := make([]struct {
		key   int
		value int
	}, 0)
	itemsToRemove := make([]int, 0)

	for i := 0; i < RECORDS; i++ {
		itemsToAdd = append(itemsToAdd, struct {
			key   int
			value int
		}{key: i, value: i})
		itemsToRemove = append(itemsToRemove, i)
	}

	adddedItems := tree.Add(itemsToAdd...)

	if tree.Size() != RECORDS {
		t.Errorf("Tree size is %d, want %d", tree.Size(), RECORDS)
	}
	if adddedItems != RECORDS {
		t.Errorf("%d added items, want %d", adddedItems, RECORDS)
	}
	t.Logf("Tree depth is %d and RootNode key is %d", tree.Depth(), tree.RootNode.Key)

	var wg sync.WaitGroup
	for i := 0; i < RECORDS; i++ {
		wg.Add(1)
		go func(i int) {
			if value, ok := tree.Get(i); !ok || value != i {
				t.Errorf("Getting %d return false, want true", i)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()

	removedItems := tree.Remove(itemsToRemove...)
	if tree.Size() != 0 {
		t.Errorf("Tree size is %d, want 0", tree.Size())
	}
	if removedItems != RECORDS {
		t.Errorf("%d removed items, want %d", removedItems, RECORDS)
	}

}

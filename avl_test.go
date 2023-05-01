package avlgo

import (
	"reflect"
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

func TestPutOneValue(t *testing.T) {
	tree := NewTree[int, bool]()
	tree.PutOne(1, true)
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

func TestPutOneingMoreValues(t *testing.T) {

	tree := NewTree[int, bool]()
	tree.PutOne(2, true)
	tree.PutOne(1, true)
	tree.PutOne(3, true)
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
	tree.PutOne(4, true)
	tree.PutOne(5, true)
	tree.PutOne(2, true)
	tree.PutOne(3, true)
	tree.PutOne(1, true)
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

func TestPutOneingMoreValuesThatUnbalanceTree(t *testing.T) {
	tree := NewTree[int, bool]()
	tree.PutOne(1, true)
	tree.PutOne(2, true)
	tree.PutOne(3, true)
	if tree.Size() != 3 {
		t.Errorf("Tree size is %d, want 3", tree.Size())
	}
	if tree.Depth() != 2 {
		t.Errorf("Tree depth is %d, want 2", tree.Depth())
	}
	if tree.RootNode.Key != 2 {
		t.Errorf("RootNode is %d, want 2", tree.RootNode.Key)
	}

	tree.PutOne(4, true)
	tree.PutOne(5, true)
	if tree.Size() != 5 {
		t.Errorf("Tree size is %d, want 5", tree.Size())
	}
	if tree.Depth() != 3 {
		t.Errorf("Tree depth is %d, want 3", tree.Depth())
	}
	if tree.RootNode.Key != 2 {
		t.Errorf("RootNode is %d, want 2", tree.RootNode.Key)
	}

	tree.PutOne(6, true)
	if tree.Size() != 6 {
		t.Errorf("Tree size is %d, want 6", tree.Size())
	}
	if tree.Depth() != 3 {
		t.Errorf("Tree depth is %d, want 3", tree.Depth())
	}
	if tree.RootNode.Key != 4 {
		t.Errorf("RootNode is %d, want 4", tree.RootNode.Key)
	}

	tree.PutOne(7, true)
	if tree.Size() != 7 {
		t.Errorf("Tree size is %d, want 7", tree.Size())
	}
	if tree.Depth() != 3 {
		t.Errorf("Tree depth is %d, want 3", tree.Depth())
	}
	if tree.RootNode.Key != 4 {
		t.Errorf("RootNode is %d, want 4", tree.RootNode.Key)
	}

	tree.PutOne(8, true)
	tree.PutOne(9, true)
	if tree.Size() != 9 {
		t.Errorf("Tree size is %d, want 9", tree.Size())
	}
	if tree.Depth() != 4 {
		t.Errorf("Tree depth is %d, want 4", tree.Depth())
	}
	if tree.RootNode.Key != 4 {
		t.Errorf("RootNode is %d, want 4", tree.RootNode.Key)
	}

	tree.PutOne(10, true)
	if tree.Size() != 10 {
		t.Errorf("Tree size is %d, want 10", tree.Size())
	}
	if tree.Depth() != 4 {
		t.Errorf("Tree depth is %d, want 4", tree.Depth())
	}
	if tree.RootNode.Key != 4 {
		t.Errorf("RootNode is %d, want 4", tree.RootNode.Key)
	}

	tree.PutOne(1, true)
	tree.PutOne(2, true)
	tree.PutOne(3, true)
	tree.PutOne(4, true)
	tree.PutOne(5, true)
	tree.PutOne(6, true)
	tree.PutOne(7, true)
	tree.PutOne(8, true)
	tree.PutOne(9, true)
	tree.PutOne(10, true)
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

func TestPutOneingMoreValuesThatUnbalanceTreeString(t *testing.T) {
	tree := NewTree[string, bool]()
	tree.PutOne("a", true)
	tree.PutOne("b", true)
	tree.PutOne("c", true)
	if tree.Size() != 3 {
		t.Errorf("Tree size is %d, want 3", tree.Size())
	}
	if tree.Depth() != 2 {
		t.Errorf("Tree depth is %d, want 2", tree.Depth())
	}
	tree.PutOne("d", true)
	tree.PutOne("e", true)
	if tree.Size() != 5 {
		t.Errorf("Tree size is %d, want 5", tree.Size())
	}
	if tree.Depth() != 3 {
		t.Errorf("Tree depth is %d, want 3", tree.Depth())
	}
	tree.PutOne("f", true)
	if tree.Size() != 6 {
		t.Errorf("Tree size is %d, want 6", tree.Size())
	}
	if tree.Depth() != 3 {
		t.Errorf("Tree depth is %d, want 3", tree.Depth())
	}
	tree.PutOne("g", true)
	if tree.Size() != 7 {
		t.Errorf("Tree size is %d, want 7", tree.Size())
	}
	if tree.Depth() != 3 {
		t.Errorf("Tree depth is %d, want 3", tree.Depth())
	}
	tree.PutOne("h", true)
	tree.PutOne("i", true)
	if tree.Size() != 9 {
		t.Errorf("Tree size is %d, want 9", tree.Size())
	}
	if tree.Depth() != 4 {
		t.Errorf("Tree depth is %d, want 4", tree.Depth())
	}
	tree.PutOne("j", true)
	if tree.Size() != 10 {
		t.Errorf("Tree size is %d, want 10", tree.Size())
	}
	if tree.Depth() != 4 {
		t.Errorf("Tree depth is %d, want 4", tree.Depth())
	}
	tree.PutOne("a", true)
	tree.PutOne("b", true)
	tree.PutOne("c", true)
	tree.PutOne("d", true)
	tree.PutOne("e", true)
	tree.PutOne("f", true)
	tree.PutOne("g", true)
	tree.PutOne("h", true)
	tree.PutOne("i", true)
	tree.PutOne("j", true)
	if tree.Size() != 10 {
		t.Errorf("Tree size is %d, want 10", tree.Size())
	}
	if tree.Depth() != 4 {
		t.Errorf("Tree depth is %d, want 4", tree.Depth())
	}
}

func TestAddWithDoubleRotation(t *testing.T) {
	tree := NewTree[int, bool]()
	tree.PutOne(7, true)
	tree.PutOne(8, true)
	tree.PutOne(4, true)
	tree.PutOne(1, true)
	tree.PutOne(5, true)
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

	//PutOneing 6 will cause double rotation
	tree.PutOne(6, true)
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
	tree.PutOne(7, "g")
	tree.PutOne(8, "h")
	tree.PutOne(1, "a")
	tree.PutOne(3, "c")

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
	tree.PutOne(1, "z")

	if value, ok := tree.Get(1); !ok || value != "z" {
		if !ok {
			t.Errorf("Get should find %d, nothing found", tree.Size())
		} else {
			t.Errorf("Get returns %s, want z", value)
		}
	}
}

func TestAddManyValues(t *testing.T) {
	tree := NewTree[int, int]()
	items := make([]struct {
		key   int
		value int
	}, 0)

	ITEMS := 8000

	for i := 0; i < ITEMS; i++ {
		items = append(items, struct {
			key   int
			value int
		}{key: i, value: i})
	}

	tree.Put(items...)

	if tree.Size() != ITEMS {
		t.Errorf("Tree size is %d, want %v", tree.Size(), ITEMS)
	}
}

func TestGetFromTo(t *testing.T) {
	tree := NewTree[int, int]()
	tree.PutOne(0, 0)
	tree.PutOne(1, 1)
	tree.PutOne(2, 2)
	tree.PutOne(3, 3)
	tree.PutOne(4, 4)
	tree.PutOne(5, 5)
	tree.PutOne(6, 6)

	values := tree.GetFromTo(2, 4, true)

	if !reflect.DeepEqual(values, []int{2, 3, 4}) {
		t.Errorf("values is %v, want %v", values, []int{2, 3, 4})
	}

	values = tree.GetFromTo(3, 6, true)

	if !reflect.DeepEqual(values, []int{3, 4, 5, 6}) {
		t.Errorf("values is %v, want %v", values, []int{3, 4, 5, 6})
	}

	values = tree.GetFromTo(0, 2, true)

	if !reflect.DeepEqual(values, []int{0, 1, 2}) {
		t.Errorf("values is %v, want %v", values, []int{0, 1, 2})
	}

	values = tree.GetFromTo(0, 6, true)

	if !reflect.DeepEqual(values, []int{0, 1, 2, 3, 4, 5, 6}) {
		t.Errorf("values is %v, want %v", values, []int{0, 1, 2, 3, 4, 5, 6})
	}

	values = tree.GetFromTo(-10, 60, true)

	if !reflect.DeepEqual(values, []int{0, 1, 2, 3, 4, 5, 6}) {
		t.Errorf("values is %v, want %v", values, []int{0, 1, 2, 3, 4, 5, 6})
	}

	values = tree.GetFromTo(2, 4, false)

	if !reflect.DeepEqual(values, []int{3}) {
		t.Errorf("values is %v, want %v", values, []int{3})
	}

	values = tree.GetFromTo(3, 6, false)

	if !reflect.DeepEqual(values, []int{4, 5}) {
		t.Errorf("values is %v, want %v", values, []int{4, 5})
	}

	values = tree.GetFromTo(0, 2, false)

	if !reflect.DeepEqual(values, []int{1}) {
		t.Errorf("values is %v, want %v", values, []int{1})
	}

	values = tree.GetFromTo(0, 6, false)

	if !reflect.DeepEqual(values, []int{1, 2, 3, 4, 5}) {
		t.Errorf("values is %v, want %v", values, []int{1, 2, 3, 4, 5})
	}

	values = tree.GetFromTo(-10, 60, false)

	if !reflect.DeepEqual(values, []int{0, 1, 2, 3, 4, 5, 6}) {
		t.Errorf("values is %v, want %v", values, []int{0, 1, 2, 3, 4, 5, 6})
	}

}

func TestDeletingSimpleValues(t *testing.T) {
	tree := NewTree[int, int]()
	tree.PutOne(0, 0)
	deleted := tree.Delete(0)
	if tree.Size() != 0 {
		t.Errorf("Tree size is %d, want 0", tree.Size())
	}
	if deleted != 1 {
		t.Errorf("Deleted nodes is %d, want 1", deleted)
	}
	tree.PutOne(0, 0)
	deleted = tree.Delete(1)
	if tree.Size() != 1 {
		t.Errorf("Tree size is %d, want 1", tree.Size())
	}
	if deleted != 0 {
		t.Errorf("Deleted nodes is %d, want 0", deleted)
	}
	tree.PutOne(1, 1)
	tree.PutOne(2, 2)
	tree.PutOne(3, 3)
	tree.PutOne(4, 4)
	tree.PutOne(5, 5)
	tree.PutOne(6, 6)
	tree.PutOne(7, 7)
	tree.PutOne(8, 8)
	tree.PutOne(9, 9)
	if tree.Size() != 10 {
		t.Errorf("Tree size is %d, want 10", tree.Size())
	}
	if tree.Depth() != 4 {
		t.Errorf("Tree depth is %d, want 4", tree.Size())
	}
	if tree.RootNode.Key != 3 {
		t.Errorf("RootNode is %d, want 3", tree.RootNode.Key)
	}
	deleted = tree.Delete(11)
	if tree.Size() != 10 {
		t.Errorf("Tree size is %d, want 10", tree.Size())
	}
	if deleted != 0 {
		t.Errorf("Deleted nodes is %d, want 0", deleted)
	}

	//deleting 2 (a leaf) should not change anything in the tree
	deleted = tree.Delete(2)
	if tree.Size() != 9 {
		t.Errorf("Tree size is %d, want 9", tree.Size())
	}
	if deleted != 1 {
		t.Errorf("Deleted nodes is %d, want 1", deleted)
	}
	if tree.Depth() != 4 {
		t.Errorf("Tree depth is %d, want 4", tree.Size())
	}
	if tree.RootNode.Key != 3 {
		t.Errorf("RootNode is %d, want 3", tree.RootNode.Key)
	}

	//deleting 1 will cause a re-balance
	deleted = tree.Delete(1)
	if tree.Size() != 8 {
		t.Errorf("Tree size is %d, want 8", tree.Size())
	}
	if deleted != 1 {
		t.Errorf("Deleted nodes is %d, want 1", deleted)
	}
	if tree.Depth() != 4 {
		t.Errorf("Tree depth is %d, want 4", tree.Size())
	}
	if tree.RootNode.Key != 7 {
		t.Errorf("RootNode is %d, want 7", tree.RootNode.Key)
	}

	//deleting 4 should not change anything in the tree
	deleted = tree.Delete(4)
	if tree.Size() != 7 {
		t.Errorf("Tree size is %d, want 7", tree.Size())
	}
	if deleted != 1 {
		t.Errorf("Deleted nodes is %d, want 1", deleted)
	}
	if tree.Depth() != 4 {
		t.Errorf("Tree depth is %d, want 4", tree.Size())
	}
	if tree.RootNode.Key != 7 {
		t.Errorf("RootNode is %d, want 7", tree.RootNode.Key)
	}

	//deleting 7 (root node) should create double rotation
	deleted = tree.Delete(7)
	if tree.Size() != 6 {
		t.Errorf("Tree size is %d, want 6", tree.Size())
	}
	if deleted != 1 {
		t.Errorf("Deleted nodes is %d, want 1", deleted)
	}
	if tree.Depth() != 3 {
		t.Errorf("Tree depth is %d, want 3", tree.Size())
	}
	if tree.RootNode.Key != 5 {
		t.Errorf("RootNode is %d, want 5", tree.RootNode.Key)
	}
	values := tree.PrintValues(0)
	if !reflect.DeepEqual(values, []int{0, 3, 5, 6, 8, 9}) {
		t.Errorf("values is %v, want %v", values, []int{0, 3, 5, 6, 8, 9})
	}

	values = tree.PrintValues(1)
	wanted := []int{5}
	if !reflect.DeepEqual(values, wanted) {
		t.Errorf("values is %v, want %v", values, wanted)
	}
	values = tree.PrintValues(2)
	wanted = []int{3, 8}
	if !reflect.DeepEqual(values, wanted) {
		t.Errorf("values is %v, want %v", values, wanted)
	}
	values = tree.PrintValues(3)
	wanted = []int{0, 6, 9}
	if !reflect.DeepEqual(values, wanted) {
		t.Errorf("values is %v, want %v", values, wanted)
	}
}

package tree

type Tree struct {
	Data  int
	Left  *Tree
	Right *Tree
}

func (t *Tree) Insert(key int) {
	if t.Data < key {
		// move to the right
		if t.Right == nil {
			t.Right = &Tree{Data: key}
		} else {
			t.Right.Insert(key)
		}
	} else if t.Data >= key {
		if t.Left == nil {
			t.Left = &Tree{Data: key}
		} else {
			t.Left.Insert(key)
		}
	}
}

func (t *Tree) Search(key int) bool {
	if t == nil {
		return false
	}

	if t.Data < key {
		// search to right
		return t.Right.Search(key)
	} else if t.Data > key {
		// search to left
		return t.Left.Search(key)
	}

	return true
}

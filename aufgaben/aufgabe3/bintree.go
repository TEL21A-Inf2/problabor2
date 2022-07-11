package aufgabe3

import "fmt"

type BinTree struct {
	Value       int
	Left, Right *BinTree
}

func NewBinTree() *BinTree {
	return &BinTree{0, nil, nil}
}

func (tree *BinTree) Empty() bool {
	return tree.Left == nil && tree.Right == nil
}

func (tree *BinTree) Add(value int) {
	if tree.Empty() {
		tree.Value = value
		tree.Left = NewBinTree()
		tree.Right = NewBinTree()
		return
	}
	if value < tree.Value {
		tree.Left.Add(value)
	} else {
		tree.Right.Add(value)
	}
}

func (tree *BinTree) String() string {
	if tree.Empty() {
		return "()"
	}
	return fmt.Sprintf("(%d %s %s)", tree.Value, tree.Left.String(), tree.Right.String())
}

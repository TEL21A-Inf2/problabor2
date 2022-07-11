package aufgabe3

import "fmt"

func ExampleBinTree_DepthOf_nonEmptyTreeExistingElement() {
	t1 := NewBinTree()

	t1.Add(42)
	t1.Add(23)
	t1.Add(55)
	t1.Add(77)

	fmt.Println(t1.DepthOf(42))
	fmt.Println(t1.DepthOf(23))
	fmt.Println(t1.DepthOf(55))
	fmt.Println(t1.DepthOf(77))

	// Output:
	// 0
	// 1
	// 1
	// 2
}

func ExampleBinTree_DepthOf_nonEmptyTreeNonExistingElement() {
	t1 := NewBinTree()

	t1.Add(42)
	t1.Add(23)
	t1.Add(55)
	t1.Add(77)

	fmt.Println(t1.DepthOf(41))
	fmt.Println(t1.DepthOf(22))

	// Output:
	// -1
	// -1
}

func ExampleBinTree_DepthOf_emptyTree() {
	t1 := NewBinTree()

	fmt.Println(t1.DepthOf(41))

	// Output:
	// -1
}

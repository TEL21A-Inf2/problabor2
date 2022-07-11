package aufgabe4

import "fmt"

func ExampleBinTree_HasAtLeastNElements() {
	t1 := NewBinTree()
	fmt.Println(t1.HasAtLeastNElements(0))
	fmt.Println(t1.HasAtLeastNElements(1))

	t1.Add(42)
	fmt.Println(t1.HasAtLeastNElements(1))
	fmt.Println(t1.HasAtLeastNElements(2))
	t1.Add(23)
	fmt.Println(t1.HasAtLeastNElements(2))
	fmt.Println(t1.HasAtLeastNElements(3))
	t1.Add(55)
	fmt.Println(t1.HasAtLeastNElements(3))
	fmt.Println(t1.HasAtLeastNElements(4))
	t1.Add(77)
	fmt.Println(t1.HasAtLeastNElements(4))
	fmt.Println(t1.HasAtLeastNElements(5))

	// Output:
	// true
	// false
	// true
	// false
	// true
	// false
	// true
	// false
	// true
	// false
}

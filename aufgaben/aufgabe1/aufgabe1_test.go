package aufgabe1

import "fmt"

func ExampleLinkedList_SubList_fromMiddle() {
	l1 := MakeLinkedList("A", "B", "C", "D", "E", "F", "G", "H")

	fmt.Println(l1.SubList(2, 5))

	// Output:
	// [C D E]
}

func ExampleLinkedList_SubList_first() {
	l1 := MakeLinkedList("A", "B", "C", "D", "E", "F", "G", "H")

	fmt.Println(l1.SubList(0, 1))

	// Output:
	// [A]
}

func ExampleLinkedList_SubList_last() {
	l1 := MakeLinkedList("A", "B", "C", "D", "E", "F", "G", "H")

	fmt.Println(l1.SubList(7, 8))

	// Output:
	// [H]
}

func ExampleLinkedList_SubList_beforeStart() {
	l1 := MakeLinkedList("A", "B", "C", "D", "E", "F", "G", "H")

	fmt.Println(l1.SubList(-2, 3))

	// Output:
	// [A B C]
}

func ExampleLinkedList_SubList_afterEnd() {
	l1 := MakeLinkedList("A", "B", "C", "D", "E", "F", "G", "H")

	fmt.Println(l1.SubList(6, 10))

	// Output:
	// [G H]
}

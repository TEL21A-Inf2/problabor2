package aufgabe5

import "fmt"

func ExampleBucketList_AppendAll() {
	b1 := NewBucketList()

	b1.Add("Hallo")
	b1.Add("Haus")
	b1.Add("Auto")
	b1.Add("Fahrrad")

	b1.AppendAll("42")
	fmt.Println(b1)

	// Output:
	// a: [Auto42]
	// f: [Fahrrad42]
	// h: [Hallo42 Haus42]
}

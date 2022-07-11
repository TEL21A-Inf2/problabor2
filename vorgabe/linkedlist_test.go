package vorgabe

import (
	"testing"

	"github.com/tel21a-inf2/probelabor2/testhelpers"
)

func TestLinkedList_MakeLinkedList(t *testing.T) {
	test := testhelpers.NewTest("Test für das Erzeugen neuer Listen mittels MakeLinkedList", t)

	l1 := MakeLinkedList("A", "B", "C")
	test.AssertValuesEqual(l1.Id, "A")
	test.AssertValuesEqual(l1.Next.Id, "B")
	test.AssertValuesEqual(l1.Next.Next.Id, "C")
	dummy := l1.Next.Next.Next
	test.AssertValuesEqual(dummy.Id, "")
	test.Assert(dummy.Next == nil, "Der Dummy hat noch einen Nachfolger!")

	l2 := MakeLinkedList()
	test.AssertValuesEqual(l2.Id, "")
}

func TestLinkedList_String(t *testing.T) {
	test := testhelpers.NewTest("Test für die String-Repräsentation von Listen", t)

	l1 := MakeLinkedList("A", "B", "C")
	test.AssertValuesEqual(l1.String(), "[A B C]")

	l2 := MakeLinkedList()
	test.AssertValuesEqual(l2.String(), "[]")
}

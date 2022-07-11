package vorgabe

import (
	"testing"

	"github.com/tel21a-inf2/probelabor2/testhelpers"
)

func TestBucketList_emptyList(t *testing.T) {
	test := testhelpers.NewTest("Test für Erzeugen einer leeren BucketList Liste", t)

	b1 := NewBucketList()
	for _, bucket := range b1 {
		test.AssertValuesEqual(0, len(bucket))
	}
}

func TestBucketList_nonEmptyList(t *testing.T) {
	test := testhelpers.NewTest("Test für das Hinzufügen von Elementen zur einer BucketList", t)

	b1 := NewBucketList()
	b1.Add("Hallo")
	b1.Add("Haus")
	b1.Add("Auto")
	b1.Add("Fahrrad")

	test.AssertStringListsEqual([]string{"Auto"}, b1[0])
	test.AssertStringListsEqual([]string{"Fahrrad"}, b1[5])
	test.AssertStringListsEqual([]string{"Hallo", "Haus"}, b1[7])
}

func TestBucketListString(t *testing.T) {
	test := testhelpers.NewTest("Test für die String-Repräsentatio einer BucketList", t)

	b1 := NewBucketList()
	b1.Add("Hallo")
	b1.Add("Haus")
	b1.Add("Auto")
	b1.Add("Fahrrad")

	test.AssertValuesEqual(`a: [Auto]
f: [Fahrrad]
h: [Hallo Haus]`, b1.String())

	b2 := NewBucketList()
	test.AssertValuesEqual("", b2.String())
}

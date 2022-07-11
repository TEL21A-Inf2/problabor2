package vorgabe

import (
	"testing"

	"github.com/tel21a-inf2/probelabor2/testhelpers"
)

func TestBinTree_Add(t *testing.T) {
	test := testhelpers.NewTest("Test für das Hinzufügen von Elementen zum Baum", t)

	t1 := NewBinTree()
	test.Assert(t1.Empty(), "Ein neuer Baum sollte leer sein.")

	t1.Add(42)
	test.AssertValuesEqual(42, t1.Value)
	test.Assert(t1.Left.Empty(), "Das linke Kind der Wurzel sollte leer sein.")
	test.Assert(t1.Right.Empty(), "Das rechte Kind der sollte leer sein.")

	t1.Add(23)
	t1.Add(55)
	t1.Add(77)
	test.AssertValuesEqual(42, t1.Value)
	test.AssertValuesEqual(23, t1.Left.Value)
	test.AssertValuesEqual(55, t1.Right.Value)
	test.AssertValuesEqual(77, t1.Right.Right.Value)
}

func TestBintree_String(t *testing.T) {
	test := testhelpers.NewTest("Test für die String-Repräsentation des Baumes", t)

	t1 := NewBinTree()
	test.AssertValuesEqual("()", t1.String())

	t1.Add(42)
	test.AssertValuesEqual("(42 () ())", t1.String())

	t1.Add(23)
	t1.Add(55)
	t1.Add(77)
	test.AssertValuesEqual("(42 (23 () ()) (55 () (77 () ())))", t1.String())
}

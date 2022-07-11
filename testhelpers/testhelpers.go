package testhelpers

import (
	"fmt"
	"testing"
)

type Test struct {
	t    *testing.T
	Name string

	count int
}

// Liefert ein neues Test-Objekt mit dem angegebenen Namen.
func NewTest(name string, t *testing.T) *Test {
	return &Test{t, name, 1}
}

// Prüft, ob zwei Werte gleich sind.
func (t *Test) AssertValuesEqual(expected, actual any) {
	t.Assert(expected == actual, "Werte sind nicht gleich!\n  Erwartet:    %v\n  Tatsächlich: %v\n", expected, actual)
}

// Prüft, ob zwei Listen gleich sind.
func (t *Test) AssertStringListsEqual(expected, actual []string) {
	for i, v := range expected {
		equal := true
		if v != actual[i] {
			equal = false
		}
		t.Assert(equal, "Listen sind nicht gleich!\n  Erwartet:    %v\n  Tatsächlich: %v\n", expected, actual)
	}
}

// Prüft einen booleschen Wert.
// Gibt die angegebene Fehlermeldung aus, wenn die Bedingung nicht zutrifft.
// Die Fehlermeldung darf ein Formatstring sein, es werden die weiteren Werte eingesetzt.
func (t *Test) Assert(value bool, message string, formatValues ...any) {
	if !value {
		t.AssertionFailure(message, formatValues...)
	}
	t.count++
}

func (t *Test) AssertionFailure(message string, formatValues ...any) {
	message = fmt.Sprintf(message, formatValues...)
	t.t.Errorf("\n[Assertion %d] %s", t.count, message)
}

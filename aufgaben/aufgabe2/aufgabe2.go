// Thema Aufgabe: Listen
// Erreichbare Punkte: 6
// Bewertung:

package aufgabe2

/* Aufgabenstellung:
 * In der Datei linkedlist.go  ist ein Datentyp für eine verkettete Liste definiert.
 *
 * Implementieren Sie für diesen Datentyp die unten vorgegebene Methode Reverse().
 * Die Methode soll die Liste umdrehen und den neuen Head liefern.
 * D.h. die Methode soll die Elemente erhalten, die Pointer aber so umbiegen, dass
 * die Liste umgedreht ist.
 *
 * Anmerkung: Bei dieser Aufgabe würde in der Bewertung auch geprüft werden,
 * ob wirklich nur Pointer verdreht wurden, d.h. ob die Adressen der Elemente noch
 * die gleichen sind. Diese Forderung ist aber relativ leicht einzuhalten,
 * indem man einfach keine neuen Elemente erzeugt ;-).
 */

// Dreht die Liste um und liefert den neuen Kopf.
func (list *LinkedList) Reverse() *LinkedList {
	if list.Empty() || list.Next.Empty() {
		return list
	}
	newEnd := list.Next
	newHead := list.Next.Reverse()
	list.Next = newEnd.Next
	newEnd.Next = list
	return newHead
}

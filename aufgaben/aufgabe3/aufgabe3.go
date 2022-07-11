// Thema der Aufgabe: Bäume
// Erreichbare Punkte: 6
// Bewertung:

package aufgabe3

/* Aufgabenstellung:
 * In der Datei bintrree.go ist ein Datentyp für einen binären Suchbaum definiert.
 *
 * Implementieren Sie für diesen Datentyp die unten vorgegebene Methode DepthOf().
 * Die Methode soll einen Wert erwarten und die Tiefe dieses Knotens liefern.
 * Die Wurzel hat dabei die Tiefe 0, deren Kinder die Tiefe 1, die Enkel Tiefe 2 usw.
 * Gibt es keinen Knoten mit diesem Wert, soll -1 geliefert werden.
 *
 */

// Liefert die Tiefe des Knotens mit dem Wert value.
func (tree *BinTree) DepthOf(value int) int {
	if tree.Empty() {
		return -1
	}
	if tree.Value == value {
		return 0
	}
	lDepth := tree.Left.DepthOf(value)
	rDepth := tree.Right.DepthOf(value)
	max := lDepth
	if rDepth > max {
		max = rDepth
	}
	if max >= 0 {
		return max + 1
	}
	return -1
}

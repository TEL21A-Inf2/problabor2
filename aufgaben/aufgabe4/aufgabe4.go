// Thema der Aufgabe: Bäume
// Erreichbare Punkte: 6
// Bewertung:

package aufgabe4

/* Aufgabenstellung:
 * In der Datei bintrree.go ist ein Datentyp für einen binären Suchbaum definiert.
 *
 * Implementieren Sie für diesen Datentyp die unten vorgegebene Methode HasAtLeastNElements().
 * Die Methode soll true liefern, falls der Baum mehr als N Elemente hat.
 */

// Liefert true, falls der Baum mindestens n Elemente hat.
func (tree *BinTree) HasAtLeastNElements(n int) bool {
	if n <= 0 {
		return true
	}
	if tree.Empty() {
		return false
	}
	for i := 0; i < n; i++ {
		if tree.Left.HasAtLeastNElements(i) && tree.Right.HasAtLeastNElements(n-i-1) {
			return true
		}
	}
	return false
}

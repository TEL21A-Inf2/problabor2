package aufgabe5

import (
	"fmt"
	"strings"
)

type BucketList [][]string

func NewBucketList() BucketList {
	return make(BucketList, 26)
}

func (b BucketList) Add(s string) {
	pos := strings.ToLower(s)[0] - 'a'
	b[pos] = append(b[pos], s)
}

func (b BucketList) String() string {
	resultStrings := make([]string, 0)

	current := 'a'

	for _, bucket := range b {
		if len(bucket) > 0 {
			resultStrings = append(resultStrings, fmt.Sprintf("%c: %v", current, bucket))
		}
		current++
	}
	return strings.Join(resultStrings, "\n")
}

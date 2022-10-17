//go:build !fribidi || js
// +build !fribidi js

package text

// Bidi maps the string from its logical order to the visual order to correctly display mixed LTR/RTL text. It returns a mapping of rune positions.
func Bidi(text string) (string, []int) {
	str := []rune(text)
	mapV2L := make([]int, len(str))
	for i := range str {
		mapV2L[i] = i
	}
	return text, mapV2L
}

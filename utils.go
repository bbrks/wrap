package wrap

import "unicode/utf8"

// isASCII returns true if the string contains only ASCII characters.
func isASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] >= utf8.RuneSelf {
			return false
		}
	}
	return true
}

// runeIndexToByte returns the byte index for a given rune index in s.
// If runeIndex exceeds the string length, returns len(s).
func runeIndexToByte(s string, runeIndex int) int {
	if runeIndex >= len(s) {
		if isASCII(s) {
			return len(s)
		}
	} else if isASCII(s[:runeIndex]) {
		return runeIndex
	}
	byteIndex := 0
	for i := 0; i < runeIndex && byteIndex < len(s); i++ {
		_, size := utf8.DecodeRuneInString(s[byteIndex:])
		byteIndex += size
	}
	return byteIndex
}

// runeIndexToByteWithShortCheck returns the byte index for a given rune index.
// Returns -1 if the string has fewer than runeIndex runes.
func runeIndexToByteWithShortCheck(s string, runeIndex int) int {
	if runeIndex > len(s) {
		// String can't have enough runes if byte length is less
		if isASCII(s) {
			return -1
		}
		// Count actual runes for non-ASCII
		if utf8.RuneCountInString(s) < runeIndex {
			return -1
		}
	} else if isASCII(s[:runeIndex]) {
		return runeIndex
	}
	byteIndex := 0
	for i := 0; i < runeIndex && byteIndex < len(s); i++ {
		_, size := utf8.DecodeRuneInString(s[byteIndex:])
		byteIndex += size
	}
	if byteIndex >= len(s) && utf8.RuneCountInString(s) < runeIndex {
		return -1
	}
	return byteIndex
}

package wrap

import "strings"

// Line will wrap a single line of text at the given length.
// If limit is less than 1, the string remains unchanged.
//
// If a word is longer than the given limit, it will not be broken to fit.
// See the examples for this scenario.
func Line(s string, limit int) string {
	if limit < 1 || len(s) < limit {
		return s
	}

	// Find the index of the last space within the limit.
	i := strings.LastIndex(s[:limit], " ")

	// Can't wrap within the limit, wrap at the next space instead.
	if i < 0 {
		i = strings.Index(s, " ")
		// Nothing left to do!
		if i < 0 {
			return s
		}
	}

	return s[:i] + "\n" + Line(s[i+1:], limit)
}

// LineWithPrefix will wrap a single line of text and prepend the given prefix,
// whilst staying within given limits.
func LineWithPrefix(s, prefix string, limit int) string {
	var ret string
	for _, str := range strings.Split(Line(s, limit-len(prefix)), "\n") {
		ret += prefix + str + "\n"
	}
	return ret
}

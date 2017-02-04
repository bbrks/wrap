package wrap

import "strings"

const (
	// breakpoints defines which characters should be able to break a line.
	breakpoints = " "
)

// Line will wrap a single line of text at the given length.
// If limit is less than 1, the string remains unchanged.
//
// If a word is longer than the given limit, it will not be broken to fit.
// See the examples for this scenario.
func Line(s string, limit int) string {
	if limit < 1 || len(s) < limit {
		return s
	}

	// Find the index of the last breakpoint within the limit.
	i := strings.LastIndexAny(s[:limit], breakpoints)

	// Can't wrap within the limit, wrap at the next breakpoint instead.
	if i < 0 {
		i = strings.IndexAny(s, breakpoints)
		// Nothing left to do!
		if i < 0 {
			return s
		}
	}

	// Recurse until we have nothing left to do.
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

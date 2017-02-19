package wrap

import "strings"

const (
	defaultBreakpoints = " -"
	defaultNewline     = "\n"
)

// Wrapper contains settings for customisable word-wrapping.
type Wrapper struct {
	// Breakpoints defines which characters should be able to break a line.
	// By default, this follows the usual English rules of spaces, and hyphens.
	// Default: " -"
	Breakpoints string

	// Newline defines which characters should be used to split and create new lines.
	// Default: "\n"
	Newline string
}

// NewWrapper returns a new instance of a Wrapper initialised with defaults.
func NewWrapper() Wrapper {
	return Wrapper{
		Breakpoints: defaultBreakpoints,
		Newline:     defaultNewline,
	}
}

// line will wrap a single line of text at the given length.
// If limit is less than 1, the string remains unchanged.
func (w Wrapper) line(s string, limit int) string {
	if limit < 1 || len(s) < limit {
		return s
	}

	// Find the index of the last breakpoint within the limit.
	i := strings.LastIndexAny(s[:limit], w.Breakpoints)

	// Can't wrap within the limit, wrap at the next breakpoint instead.
	if i < 0 {
		i = strings.IndexAny(s, w.Breakpoints)
		// Nothing left to do!
		if i < 0 {
			return s
		}
	}

	// Recurse until we have nothing left to do.
	return s[:i] + w.Newline + w.line(s[i+1:], limit)
}

// Wrap will wrap one or more lines of text at the given length.
// If limit is less than 1, the string remains unchanged.
func (w Wrapper) Wrap(s string, limit int) string {
	var ret string
	for _, str := range strings.Split(s, w.Newline) {
		ret += w.line(str, limit) + w.Newline
	}
	return ret
}

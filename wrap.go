package wrap

import "strings"

const (
	defaultBreakpoints   = " -"
	defaultNewline       = "\n"
	defaultNewlinePrefix = ""
	defaultTrimPrefix    = ""
	defaultTrimSuffix    = ""
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

	// NewlinePrefix is prefixed to any newly created lines. This can be useful
	// for wrapping a code-comment blocks and prefix new lines with "// ".
	// Default: ""
	NewlinePrefix string

	// TrimPrefix can be set to remove a prefix on each line.
	// This can be paired up with NewlinePrefix to create a block of C-style
	// comments (/* * */ ) from a long single-line comment.
	// Default: ""
	TrimPrefix string

	// TrimSuffix can be set to remove a suffix on each line.
	// Default: ""
	TrimSuffix string
}

// NewWrapper returns a new instance of a Wrapper initialised with defaults.
func NewWrapper() Wrapper {
	return Wrapper{
		Breakpoints:   defaultBreakpoints,
		Newline:       defaultNewline,
		NewlinePrefix: defaultNewlinePrefix,
		TrimPrefix:    defaultTrimPrefix,
		TrimSuffix:    defaultTrimSuffix,
	}
}

// line will wrap a single line of text at the given length.
// If limit is less than 1, the string remains unwrapped.
func (w Wrapper) line(s string, limit int) string {
	if limit < 1 || len(s) < limit {
		return w.NewlinePrefix + s
	}

	// Find the index of the last breakpoint within the limit.
	i := strings.LastIndexAny(s[:limit], w.Breakpoints)

	// Can't wrap within the limit, wrap at the next breakpoint instead.
	if i < 0 {
		i = strings.IndexAny(s, w.Breakpoints)
		// Nothing left to do!
		if i < 0 {
			return w.NewlinePrefix + s
		}
	}

	// Recurse until we have nothing left to do.
	return w.NewlinePrefix + s[:i] + w.Newline + w.line(s[i+1:], limit)
}

// Wrap will wrap one or more lines of text at the given length.
// If limit is less than 1, the string remains unwrapped.
func (w Wrapper) Wrap(s string, limit int) string {
	// Subtract the length of the prefix from the limit
	// so we don't break length limits with prefixes.
	limit -= len(w.NewlinePrefix)

	var ret string
	for _, str := range strings.Split(s, w.Newline) {
		str = strings.TrimPrefix(str, w.TrimPrefix)
		str = strings.TrimSuffix(str, w.TrimSuffix)
		ret += w.line(str, limit) + w.Newline
	}
	return ret
}

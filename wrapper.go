package wrap

import (
	"strings"
	"unicode/utf8"
)

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

	// OutputLinePrefix is prepended to any output lines. This can be useful
	// for wrapping code-comments and prefixing new lines with "// ".
	// Default: ""
	OutputLinePrefix string

	// OutputLineSuffix is appended to any output lines.
	// Default: ""
	OutputLineSuffix string

	// LimitIncludesPrefixSuffix can be set to false if you don't want prefixes
	// and suffixes to be included in the length limits.
	// Default: true
	LimitIncludesPrefixSuffix bool

	// TrimPrefix can be set to remove a prefix on each input line.
	// This can be paired up with OutputPrefix to create a block of C-style
	// comments (/* * */ ) from a long single-line comment.
	// Default: ""
	TrimInputPrefix string

	// TrimSuffix can be set to remove a suffix on each input line.
	// Default: ""
	TrimInputSuffix string

	// StripTrailingNewline can be set to true if you want the trailing
	// newline to be removed from the return value.
	// Default: false
	StripTrailingNewline bool

	// CutLongWords will cause a hard-wrap in the middle of a word if the word's length exceeds the given limit.
	CutLongWords bool
}

// NewWrapper returns a new instance of a Wrapper initialised with defaults.
func NewWrapper() Wrapper {
	return Wrapper{
		Breakpoints:               defaultBreakpoints,
		Newline:                   defaultNewline,
		LimitIncludesPrefixSuffix: true,
	}
}

// Wrap is shorthand for declaring a new default Wrapper calling its Wrap method
func Wrap(s string, limit int) string {
	return NewWrapper().Wrap(s, limit)
}

// Wrap will wrap one or more lines of text at the given length.
// If limit is less than 1, the string remains unwrapped.
func (w Wrapper) Wrap(s string, limit int) string {
	// Empty newline would cause infinite loop, use default
	if w.Newline == "" {
		w.Newline = defaultNewline
	}

	// Subtract the length of the prefix and suffix from the limit
	// so we don't break length limits when using them.
	if w.LimitIncludesPrefixSuffix {
		limit -= utf8.RuneCountInString(w.OutputLinePrefix) + utf8.RuneCountInString(w.OutputLineSuffix)
	}

	var sb strings.Builder
	growLimit := limit
	if growLimit < 1 {
		growLimit = 1
	}
	sb.Grow(len(s) + len(s)/growLimit*len(w.Newline))

	for {
		idx := strings.Index(s, w.Newline)
		var str string
		if idx < 0 {
			str = s
		} else {
			str = s[:idx]
		}
		str = strings.TrimPrefix(str, w.TrimInputPrefix)
		str = strings.TrimSuffix(str, w.TrimInputSuffix)
		w.lineBuilder(&sb, str, limit)
		if idx < 0 {
			if !w.StripTrailingNewline {
				sb.WriteString(w.Newline)
			}
			break
		}
		sb.WriteString(w.Newline)
		s = s[idx+len(w.Newline):]
	}

	return sb.String()
}

// lineBuilder writes a single wrapped line to the builder.
func (w Wrapper) lineBuilder(sb *strings.Builder, s string, limit int) {
	// Trim leading breakpoints to avoid empty or whitespace-only lines
	s = strings.TrimLeft(s, w.Breakpoints)

	// Fast path: if byte length is less than limit, rune count must also be less
	if limit < 1 || len(s) < limit+1 {
		sb.WriteString(w.OutputLinePrefix)
		sb.WriteString(s)
		sb.WriteString(w.OutputLineSuffix)
		return
	}

	// Convert rune limit to byte index for slicing (also checks rune count)
	limitByteIndex := runeIndexToByteWithShortCheck(s, limit+1)
	if limitByteIndex < 0 {
		// String is shorter than limit in runes
		sb.WriteString(w.OutputLinePrefix)
		sb.WriteString(s)
		sb.WriteString(w.OutputLineSuffix)
		return
	}

	// Find the index of the last breakpoint within the limit.
	i := strings.LastIndexAny(s[:limitByteIndex], w.Breakpoints)

	breakpointWidth := 1

	// Can't wrap within the limit
	if i < 0 {
		if w.CutLongWords {
			// wrap at the limit (convert rune index to byte index)
			i = runeIndexToByte(s, limit)
			breakpointWidth = 0
		} else {
			// wrap at the next breakpoint instead
			i = strings.IndexAny(s, w.Breakpoints)
			// Nothing left to do!
			if i < 0 {
				sb.WriteString(w.OutputLinePrefix)
				sb.WriteString(s)
				sb.WriteString(w.OutputLineSuffix)
				return
			}
		}
	}

	// Write this line (trim trailing breakpoints) and recurse
	sb.WriteString(w.OutputLinePrefix)
	sb.WriteString(strings.TrimRight(s[:i], w.Breakpoints))
	sb.WriteString(w.OutputLineSuffix)
	sb.WriteString(w.Newline)

	// Trim leading breakpoints from the next line to avoid leading whitespace
	remainder := s[i+breakpointWidth:]
	remainder = strings.TrimLeft(remainder, w.Breakpoints)

	w.lineBuilder(sb, remainder, limit)
}

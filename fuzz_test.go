package wrap_test

import (
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/bbrks/wrap/v2"
)

func FuzzWrap(f *testing.F) {
	// Seed corpus with interesting cases
	f.Add("hello world", 10)
	f.Add("hello-world-foo-bar", 8)
	f.Add("", 5)
	f.Add("a", 1)
	f.Add("日本語テスト", 3)
	f.Add("   leading spaces", 10)
	f.Add("trailing spaces   ", 10)
	f.Add("multiple   spaces", 5)
	f.Add("line1\nline2\nline3", 5)
	f.Add("no-break-points-here", 5)
	f.Add("😀😀😀😀", 2)

	// Edge cases
	f.Add("", 0)
	f.Add("", -1)
	f.Add("a", 0)
	f.Add("a", -1)
	f.Add("a", 1000000)
	f.Add(strings.Repeat("a", 10000), 10)
	f.Add(strings.Repeat("a ", 1000), 5)
	f.Add(strings.Repeat("日本語 ", 500), 3)

	// Pathological cases
	f.Add("\n\n\n", 1)
	f.Add("   \n   \n   ", 2)
	f.Add("-", 1)
	f.Add("---", 1)
	f.Add("- - - -", 1)
	f.Add("\t\t\t", 2)
	f.Add("a\nb\nc\nd\ne", 1)
	f.Add(strings.Repeat("\n", 100), 5)
	f.Add(strings.Repeat(" ", 100), 5)
	f.Add(strings.Repeat("-", 100), 5)

	// Mixed multibyte
	f.Add("aあbいcうdえeお", 2)
	f.Add("🇺🇸🇬🇧🇯🇵", 2)  // Flag emojis (multi-codepoint)
	f.Add("👨‍👩‍👧‍👦", 1) // Family emoji (ZWJ sequence)
	f.Add("e\u0301", 1) // e + combining acute accent
	f.Add("한글테스트", 3)   // Korean
	f.Add("مرحبا", 2)   // Arabic (RTL)
	f.Add("שלום", 2)    // Hebrew (RTL)

	// Breakpoint edge cases
	f.Add("a-b-c-d-e", 1)
	f.Add("a - b - c", 2)
	f.Add("----", 2)
	f.Add("    ", 2)
	f.Add("a--b", 2)
	f.Add("a  b", 2)
	f.Add("-a-", 2)
	f.Add(" a ", 2)

	f.Fuzz(func(t *testing.T, input string, limit int) {
		// Skip invalid UTF-8 inputs
		if !utf8.ValidString(input) {
			t.Skip()
		}

		w := wrap.NewWrapper()
		result := w.Wrap(input, limit)

		// Result must be valid UTF-8
		if !utf8.ValidString(result) {
			t.Errorf("result is not valid UTF-8: %q", result)
		}

		// Test with StripTrailingNewline
		w.StripTrailingNewline = true
		result2 := w.Wrap(input, limit)
		if !utf8.ValidString(result2) {
			t.Errorf("result with StripTrailingNewline is not valid UTF-8: %q", result2)
		}

		// Test with CutLongWords
		w.CutLongWords = true
		result3 := w.Wrap(input, limit)
		if !utf8.ValidString(result3) {
			t.Errorf("result with CutLongWords is not valid UTF-8: %q", result3)
		}

		// With CutLongWords and positive limit, lines should not exceed limit by more than 1
		// (the +1 accounts for preserved hyphens at line breaks)
		if limit > 0 {
			for _, line := range strings.Split(strings.TrimSuffix(result3, "\n"), "\n") {
				lineLen := utf8.RuneCountInString(line)
				// Allow +1 for hyphen preservation at end of line
				if lineLen > limit+1 {
					t.Errorf("line exceeds limit %d by more than 1: %q (len=%d)", limit, line, lineLen)
				}
			}
		}

		// Test with MinimumRaggedness
		w.CutLongWords = false
		w.MinimumRaggedness = true
		result4 := w.Wrap(input, limit)
		if !utf8.ValidString(result4) {
			t.Errorf("result with MinimumRaggedness is not valid UTF-8: %q", result4)
		}

		// Test MinimumRaggedness with CutLongWords
		w.CutLongWords = true
		result5 := w.Wrap(input, limit)
		if !utf8.ValidString(result5) {
			t.Errorf("result with MinimumRaggedness+CutLongWords is not valid UTF-8: %q", result5)
		}
	})
}

func FuzzWrapWithOptions(f *testing.F) {
	// Seed with combinations of input, limit, prefix, suffix
	f.Add("hello world", 10, "// ", "", " -")
	f.Add("hello world", 15, "/* ", " */", " -")
	f.Add("test input", 5, "> ", "", " ")
	f.Add("日本語", 4, "- ", "", " -")
	f.Add("a-b-c", 3, "", "", "-")
	f.Add("a|b|c", 3, "", "", "|")
	f.Add("test", 10, ">>>", "<<<", " ")
	f.Add("", 5, "prefix", "suffix", " -")
	f.Add("no breaks here", 5, "", "", "")
	f.Add(strings.Repeat("word ", 100), 20, "# ", "", " ")

	// Extreme prefix/suffix
	f.Add("x", 100, strings.Repeat("p", 50), strings.Repeat("s", 50), " ")
	f.Add("x", 10, strings.Repeat("p", 100), "", " ")

	// Custom breakpoints
	f.Add("a,b,c,d,e", 3, "", "", ",")
	f.Add("a::b::c", 4, "", "", ":")
	f.Add("path/to/file", 5, "", "", "/")
	f.Add("CamelCaseText", 5, "", "", "ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	f.Fuzz(func(t *testing.T, input string, limit int, prefix, suffix, breakpoints string) {
		// Skip invalid UTF-8 inputs
		if !utf8.ValidString(input) || !utf8.ValidString(prefix) || !utf8.ValidString(suffix) || !utf8.ValidString(breakpoints) {
			t.Skip()
		}

		w := wrap.NewWrapper()
		w.OutputLinePrefix = prefix
		w.OutputLineSuffix = suffix
		w.Breakpoints = breakpoints

		result := w.Wrap(input, limit)

		// Result must be valid UTF-8
		if !utf8.ValidString(result) {
			t.Errorf("result is not valid UTF-8: %q", result)
		}

		// Test all option combinations including MinimumRaggedness
		for _, strip := range []bool{false, true} {
			for _, cut := range []bool{false, true} {
				for _, includeLimit := range []bool{false, true} {
					for _, optimal := range []bool{false, true} {
						w.StripTrailingNewline = strip
						w.CutLongWords = cut
						w.LimitIncludesPrefixSuffix = includeLimit
						w.MinimumRaggedness = optimal

						result := w.Wrap(input, limit)
						if !utf8.ValidString(result) {
							t.Errorf("result is not valid UTF-8 with strip=%v cut=%v includeLimit=%v optimal=%v: %q",
								strip, cut, includeLimit, optimal, result)
						}
					}
				}
			}
		}
	})
}

func FuzzWrapCustomNewline(f *testing.F) {
	f.Add("hello world", 5, "\n")
	f.Add("hello world", 5, "\r\n")
	f.Add("hello world", 5, "<br>")
	f.Add("line1\nline2", 10, "\n")
	f.Add("line1\r\nline2", 10, "\r\n")
	f.Add("line1<br>line2", 10, "<br>")
	f.Add(strings.Repeat("word ", 50), 10, "|||")
	f.Add("日本語テスト", 3, "→")

	f.Fuzz(func(t *testing.T, input string, limit int, newline string) {
		if !utf8.ValidString(input) || !utf8.ValidString(newline) {
			t.Skip()
		}

		w := wrap.NewWrapper()
		w.Newline = newline

		result := w.Wrap(input, limit)

		if !utf8.ValidString(result) {
			t.Errorf("result is not valid UTF-8: %q", result)
		}

		// Test with various option combinations
		for _, strip := range []bool{false, true} {
			for _, cut := range []bool{false, true} {
				w.StripTrailingNewline = strip
				w.CutLongWords = cut

				result := w.Wrap(input, limit)
				if !utf8.ValidString(result) {
					t.Errorf("result is not valid UTF-8 with strip=%v cut=%v: %q", strip, cut, result)
				}
			}
		}
	})
}

func FuzzWrapTrimOptions(f *testing.F) {
	f.Add("// hello world", 10, "// ", "")
	f.Add("/* comment */", 20, "/* ", " */")
	f.Add("# heading", 15, "# ", "")
	f.Add("> quote", 10, "> ", "")
	f.Add("  indented  ", 10, "  ", "  ")
	f.Add("prefix text suffix", 10, "prefix ", " suffix")
	f.Add("// line1\n// line2", 20, "// ", "")

	f.Fuzz(func(t *testing.T, input string, limit int, trimPrefix, trimSuffix string) {
		if !utf8.ValidString(input) || !utf8.ValidString(trimPrefix) || !utf8.ValidString(trimSuffix) {
			t.Skip()
		}

		w := wrap.NewWrapper()
		w.TrimInputPrefix = trimPrefix
		w.TrimInputSuffix = trimSuffix

		result := w.Wrap(input, limit)

		if !utf8.ValidString(result) {
			t.Errorf("result is not valid UTF-8: %q", result)
		}

		// Combined with output prefix/suffix
		w.OutputLinePrefix = trimPrefix
		w.OutputLineSuffix = trimSuffix
		result2 := w.Wrap(input, limit)

		if !utf8.ValidString(result2) {
			t.Errorf("result with output prefix/suffix is not valid UTF-8: %q", result2)
		}
	})
}

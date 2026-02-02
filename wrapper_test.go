package wrap_test

import (
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/bbrks/wrap/v2"
)

// tests contains various line lengths to test our wrap functions.
var testLimits = []int{-5, 0, 5, 10, 25, 80, 120, 500}

// loremIpsums contains lorem ipsum of various line-lengths and word-lengths.
var loremIpsums = []string{
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed vulputate quam nibh, et faucibus enim gravida vel. Integer bibendum lectus et erat semper fermentum quis a risus. Fusce dignissim tempus metus non pretium. Nunc sagittis magna nec purus porttitor mollis. Pellentesque feugiat quam eget laoreet aliquet. Donec gravida congue massa, et sollicitudin turpis lacinia a. Fusce non tortor magna. Cras vel finibus tellus.",
	"Quisque facilisis dictum tellus vitae sagittis. Sed gravida nulla vel ultrices ultricies. Praesent vehicula ligula sit amet massa elementum, eget fringilla nunc ultricies. Fusce aliquet nunc ac lectus tempus sagittis. Phasellus molestie commodo leo, sit amet ultrices est. Integer vitae hendrerit neque, in pretium tellus. Nam egestas mauris id nunc sollicitudin ullamcorper. Integer eget accumsan nulla. Phasellus quis eros non leo condimentum fringilla quis sit amet tellus. Donec semper vulputate lacinia. In hac habitasse platea dictumst. Aliquam varius metus fringilla sapien cursus cursus.\n",
	"Curabitur tellus libero, feugiat vel mauris et, consequat auctor ipsum. Praesent sed pharetra dolor, at convallis lectus. Vivamus at ullamcorper sem. Sed euismod vel massa a dignissim. Proin auctor nibh at pretium facilisis. Ut aliquam erat lacus. Integer sit amet magna urna. Maecenas bibendum pretium mauris convallis semper. Nunc arcu tortor, pulvinar quis eros ut, mattis placerat tortor. Sed et lacus magna. Proin ultrices fermentum sem et placerat. Donec eget sapien mi. Maecenas maximus justo sed vulputate pulvinar. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Vestibulum accumsan, sapien sit amet suscipit dignissim, velit velit maximus elit, a cursus mi odio eu magna. Nunc nec fermentum nisi, non imperdiet purus.",
	"Vestibulum convallis magna arcu, sagittis porta mi luctus sit amet. Nunc tellus magna, fermentum et mi vitae, consectetur vestibulum nulla. Fusce ornare, augue vitae tempor pellentesque, orci orci fringilla tortor, porta feugiat justo purus nec sem. Interdum et malesuada fames ac ante ipsum primis in faucibus. Nulla pellentesque sed odio in aliquam. Fusce sed molestie velit. Curabitur id quam ac felis accumsan vehicula quis in ex.",
	"\nDuis ac ornare erat. Nulla in odio eget ante tristique dignissim a non erat. Sed non nisi vitae arcu dapibus porta vitae dignissim ante. Cras et fringilla turpis. Maecenas arcu nibh, tempus euismod pretium eget, hendrerit vitae arcu. Sed vel dolor quam. Etiam consequat sed dolor ut elementum. Quisque dictum tempor pretium. Sed eu sollicitudin mi, in commodo ante.",
	"£££ ££££££ £££££ ££££ ££££ ££ ££££ ££ £ ££ £££££££ ££ £££ £££££££££ ££ ££££ £££££ ££ ££££££££ £ ££££ £££\n",
	"",
}

func TestWrapper_Wrap(t *testing.T) {
	w, n := wrap.NewWrapper(), wrap.NewWrapper()
	n.StripTrailingNewline = true

	for _, l := range testLimits {
		for _, s := range loremIpsums {
			wrapped := w.Wrap(s, l)
			stripped := n.Wrap(s, l)

			for _, v := range strings.Split(wrapped, w.Newline) {
				if !strings.Contains(v, " ") {
					continue
				}
				if l < 1 {
					if strings.Trim(s, "\n") != v {
						t.Error("Wrapped value does not equal original")
					}
					continue
				}
				if utf8.RuneCountInString(v) > l {
					t.Errorf("Line length greater than %d: %s", l, v)
				}
			}

			if wrapped != stripped+"\n" {
				t.Error("Wrapped value did not strip newline")
			}
		}
	}
}

func TestWrapper_Wrap_CutLongWords(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		limit    int
		expected string
	}{
		{"short word", "A short woooord", 8, "A short\nwoooord"},
		{"perfect word", "A perfect wooooord", 8, "A\nperfect\nwooooord"},
		{"long word", "A long wooooooooooooord", 8, "A long\nwooooooo\noooooord"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := wrap.NewWrapper()
			w.CutLongWords = true
			w.StripTrailingNewline = true
			if got := w.Wrap(tt.input, tt.limit); got != tt.expected {
				t.Errorf("Wrap(%q, %d) = %q, want %q", tt.input, tt.limit, got, tt.expected)
			}
		})
	}
}

func TestCutLongWordsUTF8Multibyte(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		prefix string
		limit  int
	}{
		{"multibyte_char_U+0080", "\u00800", "ӽ00", 4},
		{"multibyte_char_U+00FF", "\u00FF\u00FF\u00FF", "", 2},
		{"emoji_4byte", "😀😀😀😀", "", 2},
		{"mixed_multibyte", "日本語テスト", "-> ", 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := wrap.NewWrapper()
			w.CutLongWords = true
			w.OutputLinePrefix = tt.prefix
			result := w.Wrap(tt.input, tt.limit)
			if !utf8.ValidString(result) {
				t.Errorf("Result is not valid UTF-8: %q (bytes: %v)", result, []byte(result))
			}
		})
	}
}

func TestWrapper_Breakpoints(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		breakpoints string
		limit       int
		expected    string
	}{
		{"space only", "hello world foo bar", " ", 10, "hello\nworld foo\nbar\n"},
		{"hyphen only", "hello-world-foo-bar", "-", 12, "hello-world\nfoo-bar\n"},
		{"custom breakpoint", "hello|world|foo|bar", "|", 12, "hello|world\nfoo|bar\n"},
		{"multiple custom breakpoints", "hello|world,foo;bar", "|,;", 12, "hello|world\nfoo;bar\n"},
		{"no breakpoints found", "helloworldfoobar", " ", 10, "helloworldfoobar\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := wrap.NewWrapper()
			w.Breakpoints = tt.breakpoints
			if got := w.Wrap(tt.input, tt.limit); got != tt.expected {
				t.Errorf("got %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestWrapper_Newline(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		newline  string
		limit    int
		expected string
	}{
		{"default newline", "hello world", "\n", 5, "hello\nworld\n"},
		{"CRLF newline", "hello world", "\r\n", 5, "hello\r\nworld\r\n"},
		{"custom newline", "hello world", "<br>", 5, "hello<br>world<br>"},
		{"input with CRLF", "line1\r\nline2", "\r\n", 80, "line1\r\nline2\r\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := wrap.NewWrapper()
			w.Newline = tt.newline
			if got := w.Wrap(tt.input, tt.limit); got != tt.expected {
				t.Errorf("got %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestWrapper_OutputLinePrefix(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		prefix   string
		limit    int
		expected string
	}{
		{"comment prefix", "hello world", "// ", 10, "// hello\n// world\n"},
		{"bullet prefix", "hello world", "* ", 10, "* hello\n* world\n"},
		{"empty prefix", "hello world", "", 5, "hello\nworld\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := wrap.NewWrapper()
			w.OutputLinePrefix = tt.prefix
			if got := w.Wrap(tt.input, tt.limit); got != tt.expected {
				t.Errorf("got %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestWrapper_OutputLineSuffix(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		suffix   string
		limit    int
		expected string
	}{
		{"trailing backslash", "hello world", " \\", 10, "hello \\\nworld \\\n"},
		{"HTML break", "hello world", "<br>", 10, "hello<br>\nworld<br>\n"},
		{"empty suffix", "hello world", "", 5, "hello\nworld\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := wrap.NewWrapper()
			w.OutputLineSuffix = tt.suffix
			if got := w.Wrap(tt.input, tt.limit); got != tt.expected {
				t.Errorf("got %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestWrapper_LimitIncludesPrefixSuffix(t *testing.T) {
	tests := []struct {
		name                      string
		limitIncludesPrefixSuffix bool
		expected                  string
	}{
		{"includes prefix/suffix in limit", true, ">> hello\n>> world\n>> foo\n"},
		{"excludes prefix/suffix from limit", false, ">> hello\n>> world foo\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := wrap.NewWrapper()
			w.OutputLinePrefix = ">> "
			w.LimitIncludesPrefixSuffix = tt.limitIncludesPrefixSuffix
			if got := w.Wrap("hello world foo", 10); got != tt.expected {
				t.Errorf("got %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestWrapper_TrimInputPrefix(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		trim     string
		limit    int
		expected string
	}{
		{"trim comment prefix", "// hello world", "// ", 80, "hello world\n"},
		{"trim multiline", "// line1\n// line2", "// ", 80, "line1\nline2\n"},
		{"no match", "hello world", "// ", 80, "hello world\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := wrap.NewWrapper()
			w.TrimInputPrefix = tt.trim
			if got := w.Wrap(tt.input, tt.limit); got != tt.expected {
				t.Errorf("got %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestWrapper_TrimInputSuffix(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		trim     string
		limit    int
		expected string
	}{
		{"trim trailing marker", "hello world //", " //", 80, "hello world\n"},
		{"trim multiline", "line1 //\nline2 //", " //", 80, "line1\nline2\n"},
		{"no match", "hello world", " //", 80, "hello world\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := wrap.NewWrapper()
			w.TrimInputSuffix = tt.trim
			if got := w.Wrap(tt.input, tt.limit); got != tt.expected {
				t.Errorf("got %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestWrapper_StripTrailingNewline(t *testing.T) {
	tests := []struct {
		name                 string
		stripTrailingNewline bool
		expected             string
	}{
		{"keep trailing newline", false, "hello\nworld\n"},
		{"strip trailing newline", true, "hello\nworld"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := wrap.NewWrapper()
			w.StripTrailingNewline = tt.stripTrailingNewline
			if got := w.Wrap("hello world", 5); got != tt.expected {
				t.Errorf("got %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestWrapper_CutLongWordsFeatures(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		limit    int
		expected string
	}{
		{"cut single long word", "abcdefghij", 5, "abcde\nfghij"},
		{"cut with normal words", "hi abcdefghij bye", 5, "hi\nabcde\nfghij\nbye"},
		{"no cut needed", "hello world", 10, "hello\nworld"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := wrap.NewWrapper()
			w.CutLongWords = true
			w.StripTrailingNewline = true
			if got := w.Wrap(tt.input, tt.limit); got != tt.expected {
				t.Errorf("got %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestWrapper_CombinedFeatures(t *testing.T) {
	tests := []struct {
		name            string
		input           string
		limit           int
		prefix          string
		suffix          string
		trimInputPrefix string
		expected        string
	}{
		{
			name:     "prefix and suffix together",
			input:    "hello world",
			limit:    15,
			prefix:   "/* ",
			suffix:   " */",
			expected: "/* hello */\n/* world */",
		},
		{
			name:            "trim input and add output prefix",
			input:           "# hello world",
			limit:           10,
			prefix:          "// ",
			trimInputPrefix: "# ",
			expected:        "// hello\n// world",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := wrap.NewWrapper()
			w.OutputLinePrefix = tt.prefix
			w.OutputLineSuffix = tt.suffix
			w.TrimInputPrefix = tt.trimInputPrefix
			w.StripTrailingNewline = true
			if got := w.Wrap(tt.input, tt.limit); got != tt.expected {
				t.Errorf("got %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestWrapper_EdgeCases(t *testing.T) {
	tests := []struct {
		name                 string
		input                string
		limit                int
		stripTrailingNewline bool
		expected             string
	}{
		{"empty string", "", 80, false, "\n"},
		{"empty string with strip", "", 80, true, ""},
		{"limit zero", "hello world", 0, true, "hello world"},
		{"limit negative", "hello world", -5, true, "hello world"},
		{"string shorter than limit", "hi", 80, true, "hi"},
		{"multiple consecutive newlines", "a\n\nb", 80, true, "a\n\nb"},
		{"only newlines", "\n\n", 80, true, "\n\n"},
		{"unicode multibyte with spaces", "日本 語テ スト", 3, true, "日本\n語テ\nスト"},
		{"unicode no breakpoint", "日本語テスト", 3, true, "日本語テスト"},
		{"limit equals word length", "hello world", 5, true, "hello\nworld"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := wrap.NewWrapper()
			w.StripTrailingNewline = tt.stripTrailingNewline
			if got := w.Wrap(tt.input, tt.limit); got != tt.expected {
				t.Errorf("got %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestWrap_Shorthand(t *testing.T) {
	if got := wrap.Wrap("hello world", 5); got != "hello\nworld\n" {
		t.Errorf("got %q, want %q", got, "hello\nworld\n")
	}
}

func TestWrapper_UTF8EdgeCases(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		limit        int
		cutLongWords bool
		expected     string
	}{
		{"non-ASCII with CutLongWords", "日本語テスト", 3, true, "日本語\nテスト"},
		{"short non-ASCII string", "日", 10, false, "日"},
		{"non-ASCII CutLongWords exact boundary", "日本語", 2, true, "日本\n語"},
		{"mixed ASCII and non-ASCII", "hello 世界 world", 8, false, "hello 世界\nworld"},
		{"non-ASCII longer than limit in bytes but not runes", "日本", 10, false, "日本"},
		{"non-ASCII exceeding limit needs wrap", "こんにちは 世界", 5, false, "こんにちは\n世界"},
		{"emoji with spaces", "😀😀 😀😀 😀😀", 4, false, "😀😀\n😀😀\n😀😀"},
		{"CutLongWords with emoji", "😀😀😀😀", 2, true, "😀😀\n😀😀"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := wrap.NewWrapper()
			w.CutLongWords = tt.cutLongWords
			w.StripTrailingNewline = true
			if got := w.Wrap(tt.input, tt.limit); got != tt.expected {
				t.Errorf("got %q, want %q", got, tt.expected)
			}
		})
	}
}

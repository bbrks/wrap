package wrap

import "testing"

func TestIsASCII(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"empty string", "", true},
		{"ASCII only", "hello world 123!@#", true},
		{"non-ASCII at end", "hello 世界", false},
		{"non-ASCII at start", "日本 hello", false},
		{"all non-ASCII", "日本語", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isASCII(tt.input); got != tt.expected {
				t.Errorf("isASCII(%q) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}

func TestRuneIndexToByte(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		runeIndex int
		expected  int
	}{
		{"ASCII runeIndex < len", "abcdef", 3, 3},
		{"ASCII runeIndex >= len", "abc", 5, 3},
		{"non-ASCII runeIndex < len", "日本語", 2, 6},
		{"non-ASCII runeIndex >= len", "日本", 5, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := runeIndexToByte(tt.input, tt.runeIndex); got != tt.expected {
				t.Errorf("runeIndexToByte(%q, %d) = %d, want %d", tt.input, tt.runeIndex, got, tt.expected)
			}
		})
	}
}

func TestRuneIndexToByteWithShortCheck(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		runeIndex int
		expected  int
	}{
		{"ASCII runeIndex <= len", "abcdef", 3, 3},
		{"ASCII runeIndex > len", "abc", 5, -1},
		{"non-ASCII runeIndex <= len", "日本語", 2, 6},
		{"non-ASCII runeIndex > len, not enough runes", "日本", 10, -1},
		{"non-ASCII loop terminates early", "日", 2, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := runeIndexToByteWithShortCheck(tt.input, tt.runeIndex); got != tt.expected {
				t.Errorf("runeIndexToByteWithShortCheck(%q, %d) = %d, want %d", tt.input, tt.runeIndex, got, tt.expected)
			}
		})
	}
}

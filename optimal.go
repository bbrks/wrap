package wrap

import (
	"strings"
	"unicode/utf8"
)

const infinity = 1e20

// wordWithSep stores a word along with the separator that followed it.
type wordWithSep struct {
	word string
	sep  string // separator after this word (empty for last word)
}

// lineBuilderOptimal writes wrapped lines using minimum raggedness algorithm.
func (w Wrapper) lineBuilderOptimal(sb *strings.Builder, s string, limit int) {
	if s == "" {
		sb.WriteString(w.OutputLinePrefix)
		sb.WriteString(w.OutputLineSuffix)
		return
	}

	lines := w.wrapOptimalLines(s, limit)
	for i, line := range lines {
		sb.WriteString(w.OutputLinePrefix)
		sb.WriteString(line)
		sb.WriteString(w.OutputLineSuffix)
		if i < len(lines)-1 {
			sb.WriteString(w.Newline)
		}
	}
}

// wrapOptimalLines wraps text using minimum raggedness algorithm.
// Returns a slice of lines. Uses SMAWK-based approach for O(n) time complexity.
func (w Wrapper) wrapOptimalLines(s string, limit int) []string {
	// Split into words, preserving separators
	words := w.splitWordsWithSep(s)
	if len(words) == 0 {
		return []string{""}
	}

	// Handle CutLongWords: split any words longer than limit
	if w.CutLongWords {
		words = w.cutLongWordsInListWithSep(words, limit)
	}

	count := len(words)

	// Precompute word and separator lengths for O(1) line width calculation
	wordLens := make([]int, count)
	sepLens := make([]int, count)
	for i, ws := range words {
		wordLens[i] = utf8.RuneCountInString(ws.word)
		sepLens[i] = utf8.RuneCountInString(ws.sep)
	}

	// Prefix sums for O(1) range queries
	// wordOffsets[j] = sum of word lengths for words[0:j]
	// sepOffsets[j] = sum of separator lengths for words[0:j]
	wordOffsets := make([]int, count+1)
	sepOffsets := make([]int, count+1)
	for i := 0; i < count; i++ {
		wordOffsets[i+1] = wordOffsets[i] + wordLens[i]
		sepOffsets[i+1] = sepOffsets[i] + sepLens[i]
	}

	// minima[j] = minimum cost to break words[0:j]
	minima := make([]float64, count+1)
	for i := 1; i <= count; i++ {
		minima[i] = infinity
	}

	// breaks[j] = optimal break point for line ending at word j
	breaks := make([]int, count+1)

	// cost calculates the cost of a line from word i to word j-1
	// Line width = sum of word lengths + separators between words (not after last word)
	cost := func(i, j int) float64 {
		if i >= j {
			return infinity
		}
		// Words from i to j-1: wordOffsets[j] - wordOffsets[i]
		// Separators from i to j-2: sepOffsets[j-1] - sepOffsets[i]
		lineWidth := wordOffsets[j] - wordOffsets[i]
		if j > i+1 {
			lineWidth += sepOffsets[j-1] - sepOffsets[i]
		}
		if lineWidth > limit {
			return infinity * float64(lineWidth-limit)
		}
		return minima[i] + float64((limit-lineWidth)*(limit-lineWidth))
	}

	// SMAWK-based algorithm using divide and conquer with online matrix
	var smawk func(rows []int, columns []int)
	smawk = func(rows []int, columns []int) {
		if len(columns) == 0 {
			return
		}

		// Reduce rows
		stack := make([]int, 0, len(rows))
		for _, row := range rows {
			for len(stack) > 0 {
				c := columns[len(stack)-1]
				if cost(stack[len(stack)-1], c) < cost(row, c) {
					break
				}
				stack = stack[:len(stack)-1]
			}
			if len(stack) < len(columns) {
				stack = append(stack, row)
			}
		}
		rows = stack

		// Recurse on odd columns
		if len(columns) > 1 {
			oddCols := make([]int, 0, (len(columns)+1)/2)
			for k := 1; k < len(columns); k += 2 {
				oddCols = append(oddCols, columns[k])
			}
			smawk(rows, oddCols)
		}

		// Fill in even columns
		rowIdx := 0
		for colIdx := 0; colIdx < len(columns); colIdx += 2 {
			col := columns[colIdx]
			var endRow int
			if colIdx+1 < len(columns) {
				endRow = breaks[columns[colIdx+1]]
			} else {
				endRow = rows[len(rows)-1]
			}

			for rowIdx < len(rows) {
				c := cost(rows[rowIdx], col)
				if c < minima[col] {
					minima[col] = c
					breaks[col] = rows[rowIdx]
				}
				if rows[rowIdx] >= endRow {
					break
				}
				rowIdx++
			}
		}
	}

	// Process using the online matrix approach from Aggarwal-Tokuyama
	n := count + 1
	offset := 0
	for i := 0; ; {
		r := n
		if pow := 1 << (i + 1); pow < n {
			r = pow
		}
		edge := (1 << i) + offset

		// Build row and column ranges
		rows := make([]int, edge-offset)
		for j := range rows {
			rows[j] = j + offset
		}
		cols := make([]int, r+offset-edge)
		for j := range cols {
			cols[j] = edge + j
		}

		smawk(rows, cols)

		// Check if we can skip ahead
		x := minima[r-1+offset]
		found := false
		for j := 1 << i; j < r-1; j++ {
			y := cost(j+offset, r-1+offset)
			if y <= x {
				n -= j
				i = 0
				offset += j
				found = true
				break
			}
		}
		if !found {
			if r == n {
				break
			}
			i++
		}
	}

	// If SMAWK didn't find a valid solution (minima still at infinity),
	// fall back to simple greedy line breaking
	if minima[count] >= infinity {
		return w.greedyWrapWithSep(words, limit)
	}

	// Reconstruct lines from break points, preserving original separators
	lines := make([]string, 0)
	j := count
	for j > 0 {
		i := breaks[j]
		// Safety check: if breaks[j] == j, we'd loop forever
		if i >= j {
			return w.greedyWrapWithSep(words, limit)
		}
		// Build line with original separators
		var sb strings.Builder
		for k := i; k < j; k++ {
			sb.WriteString(words[k].word)
			if k < j-1 {
				sb.WriteString(words[k].sep)
			}
		}
		lines = append(lines, sb.String())
		j = i
	}

	// Reverse lines (we built them backwards)
	for i, k := 0, len(lines)-1; i < k; i, k = i+1, k-1 {
		lines[i], lines[k] = lines[k], lines[i]
	}

	return lines
}

// greedyWrapWithSep provides a fallback greedy algorithm for cases where
// the SMAWK algorithm doesn't find a valid solution.
func (w Wrapper) greedyWrapWithSep(words []wordWithSep, limit int) []string {
	if len(words) == 0 {
		return []string{""}
	}

	var lines []string
	var sb strings.Builder
	lineLen := 0

	for i, ws := range words {
		wordLen := utf8.RuneCountInString(ws.word)
		sepLen := 0
		if i < len(words)-1 {
			sepLen = utf8.RuneCountInString(ws.sep)
		}

		if sb.Len() == 0 {
			// First word on line
			sb.WriteString(ws.word)
			lineLen = wordLen
		} else if lineLen+sepLen+wordLen <= limit {
			// Word fits on current line
			sb.WriteString(words[i-1].sep)
			sb.WriteString(ws.word)
			lineLen += sepLen + wordLen
		} else {
			// Word doesn't fit, start new line
			lines = append(lines, sb.String())
			sb.Reset()
			sb.WriteString(ws.word)
			lineLen = wordLen
		}
	}

	if sb.Len() > 0 {
		lines = append(lines, sb.String())
	}

	return lines
}

// splitWordsWithSep splits a string into words, preserving the separators between them.
func (w Wrapper) splitWordsWithSep(s string) []wordWithSep {
	s = strings.TrimLeft(s, w.Breakpoints)
	if s == "" {
		return nil
	}

	var words []wordWithSep
	var current strings.Builder
	var sep strings.Builder

	inWord := true
	for _, r := range s {
		if strings.ContainsRune(w.Breakpoints, r) {
			if inWord && current.Len() > 0 {
				inWord = false
			}
			if !inWord {
				sep.WriteRune(r)
			}
		} else {
			if !inWord {
				// End of separator, save word with its separator
				if current.Len() > 0 {
					words = append(words, wordWithSep{word: current.String(), sep: sep.String()})
					current.Reset()
					sep.Reset()
				}
				inWord = true
			}
			current.WriteRune(r)
		}
	}

	if current.Len() > 0 {
		words = append(words, wordWithSep{word: current.String(), sep: ""})
	}

	return words
}

// cutLongWordsInListWithSep splits any words longer than limit into chunks.
func (w Wrapper) cutLongWordsInListWithSep(words []wordWithSep, limit int) []wordWithSep {
	if limit < 1 {
		return words
	}

	result := make([]wordWithSep, 0, len(words))
	for _, ws := range words {
		wordLen := utf8.RuneCountInString(ws.word)
		if wordLen <= limit {
			result = append(result, ws)
			continue
		}

		// Split word into chunks of limit runes
		runes := []rune(ws.word)
		for i := 0; i < len(runes); i += limit {
			end := i + limit
			if end > len(runes) {
				end = len(runes)
			}
			chunk := wordWithSep{word: string(runes[i:end]), sep: ""}
			// Only the last chunk keeps the original separator
			if end == len(runes) {
				chunk.sep = ws.sep
			}
			result = append(result, chunk)
		}
	}
	return result
}



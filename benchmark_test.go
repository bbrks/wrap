package wrap_test

import (
	"testing"

	"github.com/bbrks/wrap"
)

func benchmarkWrap(b *testing.B, limit int) {
	w := wrap.NewWrapper()

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		w.Wrap(loremIpsums[0], limit)
	}
}

func BenchmarkWrap0(b *testing.B)   { benchmarkWrap(b, 0) }
func BenchmarkWrap5(b *testing.B)   { benchmarkWrap(b, 5) }
func BenchmarkWrap10(b *testing.B)  { benchmarkWrap(b, 10) }
func BenchmarkWrap25(b *testing.B)  { benchmarkWrap(b, 25) }
func BenchmarkWrap80(b *testing.B)  { benchmarkWrap(b, 80) }
func BenchmarkWrap120(b *testing.B) { benchmarkWrap(b, 120) }
func BenchmarkWrap500(b *testing.B) { benchmarkWrap(b, 500) }

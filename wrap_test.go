package wrap

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tests contains various line lengths to test our wrap functions.
var tests = []int{-5, 0, 5, 10, 25, 80, 120, 500}

// loremIpsums contains lorem ipsum of various line-lengths and word-lengths.
var loremIpsums = []string{
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed vulputate quam nibh, et faucibus enim gravida vel. Integer bibendum lectus et erat semper fermentum quis a risus. Fusce dignissim tempus metus non pretium. Nunc sagittis magna nec purus porttitor mollis. Pellentesque feugiat quam eget laoreet aliquet. Donec gravida congue massa, et sollicitudin turpis lacinia a. Fusce non tortor magna. Cras vel finibus tellus.",
	"Quisque facilisis dictum tellus vitae sagittis. Sed gravida nulla vel ultrices ultricies. Praesent vehicula ligula sit amet massa elementum, eget fringilla nunc ultricies. Fusce aliquet nunc ac lectus tempus sagittis. Phasellus molestie commodo leo, sit amet ultrices est. Integer vitae hendrerit neque, in pretium tellus. Nam egestas mauris id nunc sollicitudin ullamcorper. Integer eget accumsan nulla. Phasellus quis eros non leo condimentum fringilla quis sit amet tellus. Donec semper vulputate lacinia. In hac habitasse platea dictumst. Aliquam varius metus fringilla sapien cursus cursus.",
	"Curabitur tellus libero, feugiat vel mauris et, consequat auctor ipsum. Praesent sed pharetra dolor, at convallis lectus. Vivamus at ullamcorper sem. Sed euismod vel massa a dignissim. Proin auctor nibh at pretium facilisis. Ut aliquam erat lacus. Integer sit amet magna urna. Maecenas bibendum pretium mauris convallis semper. Nunc arcu tortor, pulvinar quis eros ut, mattis placerat tortor. Sed et lacus magna. Proin ultrices fermentum sem et placerat. Donec eget sapien mi. Maecenas maximus justo sed vulputate pulvinar. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Vestibulum accumsan, sapien sit amet suscipit dignissim, velit velit maximus elit, a cursus mi odio eu magna. Nunc nec fermentum nisi, non imperdiet purus.",
	"Vestibulum convallis magna arcu, sagittis porta mi luctus sit amet. Nunc tellus magna, fermentum et mi vitae, consectetur vestibulum nulla. Fusce ornare, augue vitae tempor pellentesque, orci orci fringilla tortor, porta feugiat justo purus nec sem. Interdum et malesuada fames ac ante ipsum primis in faucibus. Nulla pellentesque sed odio in aliquam. Fusce sed molestie velit. Curabitur id quam ac felis accumsan vehicula quis in ex.",
	"Duis ac ornare erat. Nulla in odio eget ante tristique dignissim a non erat. Sed non nisi vitae arcu dapibus porta vitae dignissim ante. Cras et fringilla turpis. Maecenas arcu nibh, tempus euismod pretium eget, hendrerit vitae arcu. Sed vel dolor quam. Etiam consequat sed dolor ut elementum. Quisque dictum tempor pretium. Sed eu sollicitudin mi, in commodo ante.",
}

func TestLine(t *testing.T) {
	// Test multiple line lengths.
	for _, l := range tests {

		// Test each input line individually.
		for _, s := range loremIpsums {
			wrapped := Line(s, l)

			// Assert that each output line is no longer than the limit.
			for _, v := range strings.Split(wrapped, "\n") {

				// Only check lines which contain more than one word.
				if !strings.Contains(v, " ") {
					continue
				}

				// If length < 1, the string remains unchaged.
				if l < 1 {
					assert.Equal(t, s, v)
					continue
				}

				assert.True(t, len(v) <= l,
					fmt.Sprintf("Line length greater than %d: %s", l, v))
			}

		}

	}
}

func TestLineWithPrefix(t *testing.T) {
	var prefix = "// "
	// Test multiple line lengths.
	for _, l := range tests {

		// Test each input line individually.
		for _, s := range loremIpsums {
			wrapped := LineWithPrefix(s, prefix, l)

			// Assert that each output line is no longer than the limit.
			for _, v := range strings.Split(wrapped, "\n") {
				if !strings.HasPrefix(s, prefix) {
					continue
				}

				// Only check lines which contain more than one word.
				if !strings.Contains(v, " ") {
					continue
				}

				// If length < 1, the string remains unchaged.
				if l < 1 {
					assert.Equal(t, prefix+s, v)
					continue
				}

				assert.True(t, len(v) <= l,
					fmt.Sprintf("Line length greater than %d: %s", l, v))
			}

		}

	}
}

func benchmarkLine(b *testing.B, limit int) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Line(loremIpsums[0], limit)
	}
}

func BenchmarkLine0(b *testing.B)   { benchmarkLine(b, 0) }
func BenchmarkLine5(b *testing.B)   { benchmarkLine(b, 5) }
func BenchmarkLine10(b *testing.B)  { benchmarkLine(b, 10) }
func BenchmarkLine25(b *testing.B)  { benchmarkLine(b, 25) }
func BenchmarkLine80(b *testing.B)  { benchmarkLine(b, 80) }
func BenchmarkLine120(b *testing.B) { benchmarkLine(b, 120) }
func BenchmarkLine500(b *testing.B) { benchmarkLine(b, 500) }

func ExampleLineWithPrefix() {
	var loremIpsum = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed vulputate quam nibh, et faucibus enim gravida vel. Integer bibendum lectus et erat semper fermentum quis a risus. Fusce dignissim tempus metus non pretium. Nunc sagittis magna nec purus porttitor mollis. Pellentesque feugiat quam eget laoreet aliquet. Donec gravida congue massa, et sollicitudin turpis lacinia a. Fusce non tortor magna. Cras vel finibus tellus."

	// Wrap when lines exceed 80 chars and prepend a comment prefix.
	fmt.Println(LineWithPrefix(loremIpsum, "// ", 80))
	// Output:
	// // Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed vulputate quam
	// // nibh, et faucibus enim gravida vel. Integer bibendum lectus et erat semper
	// // fermentum quis a risus. Fusce dignissim tempus metus non pretium. Nunc
	// // sagittis magna nec purus porttitor mollis. Pellentesque feugiat quam eget
	// // laoreet aliquet. Donec gravida congue massa, et sollicitudin turpis lacinia
	// // a. Fusce non tortor magna. Cras vel finibus tellus.
}

func ExampleLine() {
	var loremIpsum = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed vulputate quam nibh, et faucibus enim gravida vel. Integer bibendum lectus et erat semper fermentum quis a risus. Fusce dignissim tempus metus non pretium. Nunc sagittis magna nec purus porttitor mollis. Pellentesque feugiat quam eget laoreet aliquet. Donec gravida congue massa, et sollicitudin turpis lacinia a. Fusce non tortor magna. Cras vel finibus tellus."

	// Wrap when lines exceed 80 chars.
	fmt.Println(Line(loremIpsum, 80))
	// Output:
	// Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed vulputate quam
	// nibh, et faucibus enim gravida vel. Integer bibendum lectus et erat semper
	// fermentum quis a risus. Fusce dignissim tempus metus non pretium. Nunc sagittis
	// magna nec purus porttitor mollis. Pellentesque feugiat quam eget laoreet
	// aliquet. Donec gravida congue massa, et sollicitudin turpis lacinia a. Fusce
	// non tortor magna. Cras vel finibus tellus.
}

func ExampleLine_short() {
	var loremIpsum = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed vulputate quam nibh, et faucibus enim gravida vel. Integer bibendum lectus et erat semper fermentum quis a risus. Fusce dignissim tempus metus non pretium. Nunc sagittis magna nec purus porttitor mollis. Pellentesque feugiat quam eget laoreet aliquet. Donec gravida congue massa, et sollicitudin turpis lacinia a. Fusce non tortor magna. Cras vel finibus tellus."

	// Wrap on every word.
	fmt.Println(Line(loremIpsum, 1))
	// Output:
	// Lorem
	// ipsum
	// dolor
	// sit
	// amet,
	// consectetur
	// adipiscing
	// elit.
	// Sed
	// vulputate
	// quam
	// nibh,
	// et
	// faucibus
	// enim
	// gravida
	// vel.
	// Integer
	// bibendum
	// lectus
	// et
	// erat
	// semper
	// fermentum
	// quis
	// a
	// risus.
	// Fusce
	// dignissim
	// tempus
	// metus
	// non
	// pretium.
	// Nunc
	// sagittis
	// magna
	// nec
	// purus
	// porttitor
	// mollis.
	// Pellentesque
	// feugiat
	// quam
	// eget
	// laoreet
	// aliquet.
	// Donec
	// gravida
	// congue
	// massa,
	// et
	// sollicitudin
	// turpis
	// lacinia
	// a.
	// Fusce
	// non
	// tortor
	// magna.
	// Cras
	// vel
	// finibus
	// tellus.
}

package wrap_test

import (
	"fmt"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/bbrks/wrap"
	"github.com/stretchr/testify/assert"
)

var w = wrap.NewWrapper()

// tests contains various line lengths to test our wrap functions.
var tests = []int{-5, 0, 5, 10, 25, 80, 120, 500}

// loremIpsums contains lorem ipsum of various line-lengths and word-lengths.
var loremIpsums = []string{
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed vulputate quam nibh, et faucibus enim gravida vel. Integer bibendum lectus et erat semper fermentum quis a risus. Fusce dignissim tempus metus non pretium. Nunc sagittis magna nec purus porttitor mollis. Pellentesque feugiat quam eget laoreet aliquet. Donec gravida congue massa, et sollicitudin turpis lacinia a. Fusce non tortor magna. Cras vel finibus tellus.",
	"Quisque facilisis dictum tellus vitae sagittis. Sed gravida nulla vel ultrices ultricies. Praesent vehicula ligula sit amet massa elementum, eget fringilla nunc ultricies. Fusce aliquet nunc ac lectus tempus sagittis. Phasellus molestie commodo leo, sit amet ultrices est. Integer vitae hendrerit neque, in pretium tellus. Nam egestas mauris id nunc sollicitudin ullamcorper. Integer eget accumsan nulla. Phasellus quis eros non leo condimentum fringilla quis sit amet tellus. Donec semper vulputate lacinia. In hac habitasse platea dictumst. Aliquam varius metus fringilla sapien cursus cursus.",
	"Curabitur tellus libero, feugiat vel mauris et, consequat auctor ipsum. Praesent sed pharetra dolor, at convallis lectus. Vivamus at ullamcorper sem. Sed euismod vel massa a dignissim. Proin auctor nibh at pretium facilisis. Ut aliquam erat lacus. Integer sit amet magna urna. Maecenas bibendum pretium mauris convallis semper. Nunc arcu tortor, pulvinar quis eros ut, mattis placerat tortor. Sed et lacus magna. Proin ultrices fermentum sem et placerat. Donec eget sapien mi. Maecenas maximus justo sed vulputate pulvinar. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Vestibulum accumsan, sapien sit amet suscipit dignissim, velit velit maximus elit, a cursus mi odio eu magna. Nunc nec fermentum nisi, non imperdiet purus.",
	"Vestibulum convallis magna arcu, sagittis porta mi luctus sit amet. Nunc tellus magna, fermentum et mi vitae, consectetur vestibulum nulla. Fusce ornare, augue vitae tempor pellentesque, orci orci fringilla tortor, porta feugiat justo purus nec sem. Interdum et malesuada fames ac ante ipsum primis in faucibus. Nulla pellentesque sed odio in aliquam. Fusce sed molestie velit. Curabitur id quam ac felis accumsan vehicula quis in ex.",
	"Duis ac ornare erat. Nulla in odio eget ante tristique dignissim a non erat. Sed non nisi vitae arcu dapibus porta vitae dignissim ante. Cras et fringilla turpis. Maecenas arcu nibh, tempus euismod pretium eget, hendrerit vitae arcu. Sed vel dolor quam. Etiam consequat sed dolor ut elementum. Quisque dictum tempor pretium. Sed eu sollicitudin mi, in commodo ante.",
	"£££ ££££££ £££££ ££££ ££££ ££ ££££ ££ £ ££ £££££££ ££ £££ £££££££££ ££ ££££ £££££ ££ ££££££££ £ ££££ £££",
	"",
}

func TestWrapper_Wrap(t *testing.T) {
	// Test multiple line lengths.
	for _, l := range tests {

		// Test each input line individually.
		for _, s := range loremIpsums {
			wrapped := w.Wrap(s, l)

			// Assert that each output line is no longer than the limit.
			for _, v := range strings.Split(wrapped, w.Newline) {

				// Only check lines which contain more than one word.
				if !strings.Contains(v, " ") {
					continue
				}

				// If length < 1, the string remains unchaged.
				if l < 1 {
					assert.Equal(t, s, v)
					continue
				}

				assert.True(t, utf8.RuneCountInString(v) <= l,
					fmt.Sprintf("Line length greater than %d: %s", l, v))
			}

		}

	}
}

func benchmarkWrap(b *testing.B, limit int) {
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

func ExampleWrapper_Wrap() {
	var loremIpsum = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed vulputate quam nibh, et faucibus enim gravida vel. Integer bibendum lectus et erat semper fermentum quis a risus. Fusce dignissim tempus metus non pretium. Nunc sagittis magna nec purus porttitor mollis. Pellentesque feugiat quam eget laoreet aliquet. Donec gravida congue massa, et sollicitudin turpis lacinia a. Fusce non tortor magna. Cras vel finibus tellus."

	// Wrap when lines exceed 80 chars.
	fmt.Println(w.Wrap(loremIpsum, 80))
	// Output:
	// Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed vulputate quam
	// nibh, et faucibus enim gravida vel. Integer bibendum lectus et erat semper
	// fermentum quis a risus. Fusce dignissim tempus metus non pretium. Nunc sagittis
	// magna nec purus porttitor mollis. Pellentesque feugiat quam eget laoreet
	// aliquet. Donec gravida congue massa, et sollicitudin turpis lacinia a. Fusce non
	// tortor magna. Cras vel finibus tellus.
}

func ExampleWrapper_Wrap_paragraphs() {
	var loremIpsum = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. In pulvinar augue vel dui gravida, sed convallis ante aliquam. Morbi euismod felis in justo lobortis, eu egestas quam cursus. Ut ut tellus mattis, porttitor leo ut, porttitor ex. Nulla suscipit molestie ligula, quis porta nulla pellentesque ac. Cras ut vestibulum orci. Phasellus ante nisl, dignissim non nunc eget, dapibus convallis orci. Integer vel euismod mauris. Integer tortor elit, vestibulum eget augue vitae, vehicula commodo sapien. Integer iaculis maximus dui, vitae rutrum magna congue at. Praesent varius quam vitae rhoncus fringilla. Quisque ac ex sit amet enim aliquam rutrum in in tortor. Sed sit amet est finibus, congue purus et, ultrices quam. Aenean felis velit, ullamcorper at sagittis ut, aliquam eu mauris.

Phasellus vel lorem venenatis, condimentum risus quis, ultricies risus. Vivamus porttitor lorem sit amet bibendum congue. Mauris sem enim, rutrum in ipsum eget, porttitor placerat diam. Pellentesque ut pharetra augue. Maecenas in ante eget ex efficitur tincidunt. Cras ut ultrices nisl. Donec tristique tincidunt eros condimentum tempus. Morbi libero urna, pretium id turpis vel, cursus efficitur orci. Mauris ut elit felis. Duis ultrices nisl eget accumsan consectetur. Nullam blandit elit vel vulputate scelerisque. Nulla facilisi. Cras quis maximus odio. Nam orci est, tempor ac arcu eget, tincidunt consectetur risus. Donec quis faucibus velit.

Maecenas rhoncus semper nisi non luctus. Nam accumsan malesuada urna vel vehicula. Nullam quis dui in augue tristique sollicitudin. Praesent vulputate condimentum vestibulum. Morbi tincidunt consectetur velit non accumsan. Praesent sit amet vestibulum purus. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Nulla rhoncus urna ut aliquet congue. Sed ornare dignissim orci non imperdiet. Maecenas nec magna bibendum, cursus nisi ac, commodo arcu.

Sed auctor id leo at molestie. Donec sed cursus massa. Morbi semper lobortis dui. Sed mattis sem a molestie sodales. Cras consequat sapien semper, pretium nulla a, dignissim massa. Aliquam non ornare lacus. Cras gravida lorem tellus, et consectetur ante sodales ut.

Nunc mi enim, aliquam quis bibendum sed, commodo quis nulla. Aliquam vulputate arcu a volutpat semper. Donec nec mauris eros. Suspendisse velit ante, fermentum a odio non, porta dignissim nunc. Vestibulum condimentum at massa at malesuada. Etiam augue purus, interdum a est pretium, cursus interdum eros. Vestibulum et ligula dignissim, suscipit arcu et, congue sem. Integer posuere mauris id scelerisque sagittis. Proin cursus congue sem, nec pulvinar neque auctor eget. Suspendisse vitae mi ipsum. Nullam sed mauris posuere, accumsan ligula vitae, viverra tellus. Morbi quam turpis, sagittis vitae arcu vel, tempus sagittis neque. Vivamus dolor purus, blandit ac condimentum a, interdum in ipsum.`

	fmt.Println(w.Wrap(loremIpsum, 80))
	// Output:
	// Lorem ipsum dolor sit amet, consectetur adipiscing elit. In pulvinar augue vel
	// dui gravida, sed convallis ante aliquam. Morbi euismod felis in justo lobortis,
	// eu egestas quam cursus. Ut ut tellus mattis, porttitor leo ut, porttitor ex.
	// Nulla suscipit molestie ligula, quis porta nulla pellentesque ac. Cras ut
	// vestibulum orci. Phasellus ante nisl, dignissim non nunc eget, dapibus convallis
	// orci. Integer vel euismod mauris. Integer tortor elit, vestibulum eget augue
	// vitae, vehicula commodo sapien. Integer iaculis maximus dui, vitae rutrum magna
	// congue at. Praesent varius quam vitae rhoncus fringilla. Quisque ac ex sit amet
	// enim aliquam rutrum in in tortor. Sed sit amet est finibus, congue purus et,
	// ultrices quam. Aenean felis velit, ullamcorper at sagittis ut, aliquam eu
	// mauris.
	//
	// Phasellus vel lorem venenatis, condimentum risus quis, ultricies risus. Vivamus
	// porttitor lorem sit amet bibendum congue. Mauris sem enim, rutrum in ipsum eget,
	// porttitor placerat diam. Pellentesque ut pharetra augue. Maecenas in ante eget
	// ex efficitur tincidunt. Cras ut ultrices nisl. Donec tristique tincidunt eros
	// condimentum tempus. Morbi libero urna, pretium id turpis vel, cursus efficitur
	// orci. Mauris ut elit felis. Duis ultrices nisl eget accumsan consectetur. Nullam
	// blandit elit vel vulputate scelerisque. Nulla facilisi. Cras quis maximus odio.
	// Nam orci est, tempor ac arcu eget, tincidunt consectetur risus. Donec quis
	// faucibus velit.
	//
	// Maecenas rhoncus semper nisi non luctus. Nam accumsan malesuada urna vel
	// vehicula. Nullam quis dui in augue tristique sollicitudin. Praesent vulputate
	// condimentum vestibulum. Morbi tincidunt consectetur velit non accumsan. Praesent
	// sit amet vestibulum purus. Orci varius natoque penatibus et magnis dis
	// parturient montes, nascetur ridiculus mus. Nulla rhoncus urna ut aliquet congue.
	// Sed ornare dignissim orci non imperdiet. Maecenas nec magna bibendum, cursus
	// nisi ac, commodo arcu.
	//
	// Sed auctor id leo at molestie. Donec sed cursus massa. Morbi semper lobortis
	// dui. Sed mattis sem a molestie sodales. Cras consequat sapien semper, pretium
	// nulla a, dignissim massa. Aliquam non ornare lacus. Cras gravida lorem tellus,
	// et consectetur ante sodales ut.
	//
	// Nunc mi enim, aliquam quis bibendum sed, commodo quis nulla. Aliquam vulputate
	// arcu a volutpat semper. Donec nec mauris eros. Suspendisse velit ante, fermentum
	// a odio non, porta dignissim nunc. Vestibulum condimentum at massa at malesuada.
	// Etiam augue purus, interdum a est pretium, cursus interdum eros. Vestibulum et
	// ligula dignissim, suscipit arcu et, congue sem. Integer posuere mauris id
	// scelerisque sagittis. Proin cursus congue sem, nec pulvinar neque auctor eget.
	// Suspendisse vitae mi ipsum. Nullam sed mauris posuere, accumsan ligula vitae,
	// viverra tellus. Morbi quam turpis, sagittis vitae arcu vel, tempus sagittis
	// neque. Vivamus dolor purus, blandit ac condimentum a, interdum in ipsum.
}

func ExampleWrapper_Wrap_short() {
	var loremIpsum = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed vulputate quam nibh, et faucibus enim gravida vel. Integer bibendum lectus et erat semper fermentum quis a risus. Fusce dignissim tempus metus non pretium. Nunc sagittis magna nec purus porttitor mollis. Pellentesque feugiat quam eget laoreet aliquet. Donec gravida congue massa, et sollicitudin turpis lacinia a. Fusce non tortor magna. Cras vel finibus tellus."

	// Wrap on every word.
	fmt.Println(w.Wrap(loremIpsum, 1))
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

func ExampleWrapper_Wrap_hyphens() {
	var loremIpsum = `
In this particular example, I will spam a lot of hyphenated words, which should wrap at some point, and test the multi-breakpoint feature of this package.

The girl was accident-prone, good-looking, quick-thinking, carbon-neutral, bad-tempered, sport-mad, fair-haired, camera-ready, and finally open-mouthed.
`

	fmt.Println(w.Wrap(loremIpsum, 80))
	// Output:
	// In this particular example, I will spam a lot of hyphenated words, which should
	// wrap at some point, and test the multi-breakpoint feature of this package.
	//
	// The girl was accident-prone, good-looking, quick-thinking, carbon-neutral, bad
	// tempered, sport-mad, fair-haired, camera-ready, and finally open-mouthed.
}

func ExampleWrapper_Wrap_prefix() {
	var loremIpsum = "/* Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed vulputate quam nibh, et faucibus enim gravida vel. Integer bibendum lectus et erat semper fermentum quis a risus. Fusce dignissim tempus metus non pretium. Nunc sagittis magna nec purus porttitor mollis. Pellentesque feugiat quam eget laoreet aliquet. Donec gravida congue massa, et sollicitudin turpis lacinia a. Fusce non tortor magna. Cras vel finibus tellus. */"

	// Trim the single-line block comment symbols from each input line.
	w.TrimInputPrefix = "/* "
	w.TrimInputSuffix = " */"

	// Prefix each new line with a single-line comment symbol.
	w.OutputLinePrefix = "// "

	// Wrap when lines exceed 80 chars.
	fmt.Println(w.Wrap(loremIpsum, 80))
	// Output:
	// // Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed vulputate quam
	// // nibh, et faucibus enim gravida vel. Integer bibendum lectus et erat semper
	// // fermentum quis a risus. Fusce dignissim tempus metus non pretium. Nunc
	// // sagittis magna nec purus porttitor mollis. Pellentesque feugiat quam eget
	// // laoreet aliquet. Donec gravida congue massa, et sollicitudin turpis lacinia
	// // a. Fusce non tortor magna. Cras vel finibus tellus.
}

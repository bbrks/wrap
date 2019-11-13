package wrap_test

import (
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/bbrks/wrap/v2"
)

// tests contains various line lengths to test our wrap functions.
var tests = []int{-5, 0, 5, 10, 25, 80, 120, 500}

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

	// Test multiple line lengths.
	for _, l := range tests {

		// Test each input line individually.
		for _, s := range loremIpsums {
			wrapped := w.Wrap(s, l)
			stripped := n.Wrap(s, l)

			// Assert that each output line is no longer than the limit.
			for _, v := range strings.Split(wrapped, w.Newline) {

				// Only check lines which contain more than one word.
				if !strings.Contains(v, " ") {
					continue
				}

				// If length < 1, the string remains unchaged.
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

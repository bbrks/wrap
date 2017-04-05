package wrap_test

import (
	"fmt"

	"github.com/bbrks/wrap"
)

func ExampleWrap() {
	var loremIpsum = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed vulputate quam nibh, et faucibus enim gravida vel. Integer bibendum lectus et erat semper fermentum quis a risus. Fusce dignissim tempus metus non pretium. Nunc sagittis magna nec purus porttitor mollis. Pellentesque feugiat quam eget laoreet aliquet. Donec gravida congue massa, et sollicitudin turpis lacinia a. Fusce non tortor magna. Cras vel finibus tellus."

	// Wrap when lines exceed 80 chars.
	fmt.Println(wrap.Wrap(loremIpsum, 80))
	// Output:
	// Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed vulputate quam
	// nibh, et faucibus enim gravida vel. Integer bibendum lectus et erat semper
	// fermentum quis a risus. Fusce dignissim tempus metus non pretium. Nunc sagittis
	// magna nec purus porttitor mollis. Pellentesque feugiat quam eget laoreet
	// aliquet. Donec gravida congue massa, et sollicitudin turpis lacinia a. Fusce non
	// tortor magna. Cras vel finibus tellus.
}

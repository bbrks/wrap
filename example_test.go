package wrap_test

import (
	"fmt"

	"github.com/bbrks/wrap/v2"
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

func ExampleWrapper_Wrap() {
	var loremIpsum = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed vulputate quam nibh, et faucibus enim gravida vel. Integer bibendum lectus et erat semper fermentum quis a risus. Fusce dignissim tempus metus non pretium. Nunc sagittis magna nec purus porttitor mollis. Pellentesque feugiat quam eget laoreet aliquet. Donec gravida congue massa, et sollicitudin turpis lacinia a. Fusce non tortor magna. Cras vel finibus tellus."

	w := wrap.NewWrapper()

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

	w := wrap.NewWrapper()

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

func ExampleWrapper_Wrap_cutLongWords() {
	var loremIpsum = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed vulputate quam nibh, et faucibus enim gravida vel."

	w := wrap.NewWrapper()
	w.CutLongWords = true

	// Wrap at 10 chars and cut words longer.
	fmt.Println(w.Wrap(loremIpsum, 10))
	// Output:
	// Lorem
	// ipsum
	// dolor sit
	// amet,
	// consectetu
	// r
	// adipiscing
	// elit. Sed
	// vulputate
	// quam nibh,
	// et
	// faucibus
	// enim
	// gravida
	// vel.
}

func ExampleWrapper_Wrap_short() {
	var loremIpsum = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed vulputate quam nibh, et faucibus enim gravida vel. Integer bibendum lectus et erat semper fermentum quis a risus. Fusce dignissim tempus metus non pretium. Nunc sagittis magna nec purus porttitor mollis. Pellentesque feugiat quam eget laoreet aliquet. Donec gravida congue massa, et sollicitudin turpis lacinia a. Fusce non tortor magna. Cras vel finibus tellus."

	w := wrap.NewWrapper()

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
	var text = `My sister-in-law bought a state-of-the-art, energy-efficient, top-rated washing machine from a well-known, highly-recommended manufacturer based in the north-east. It was a last-minute, spur-of-the-moment decision, but the twenty-year guarantee and the user-friendly, self-cleaning features made it a no-brainer.`

	w := wrap.NewWrapper()

	fmt.Println(w.Wrap(text, 40))
	// Output:
	// My sister-in-law bought a state-of-the-
	// art, energy-efficient, top-rated washing
	// machine from a well-known, highly-
	// recommended manufacturer based in the
	// north-east. It was a last-minute, spur-
	// of-the-moment decision, but the twenty-
	// year guarantee and the user-friendly,
	// self-cleaning features made it a no-
	// brainer.
}

func ExampleWrapper_Wrap_prefix() {
	var loremIpsum = "/* Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed vulputate quam nibh, et faucibus enim gravida vel. Integer bibendum lectus et erat semper fermentum quis a risus. Fusce dignissim tempus metus non pretium. Nunc sagittis magna nec purus porttitor mollis. Pellentesque feugiat quam eget laoreet aliquet. Donec gravida congue massa, et sollicitudin turpis lacinia a. Fusce non tortor magna. Cras vel finibus tellus. */"

	w := wrap.NewWrapper()

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

func ExampleWrapper_Wrap_minimumRaggedness() {
	// This example demonstrates the difference between greedy and optimal wrapping.
	// The input is designed to show how optimal wrapping produces more balanced lines.
	var text = "a b c d e f g h i j k l m n o p"

	fmt.Println("Greedy (default):")
	w := wrap.NewWrapper()
	w.StripTrailingNewline = true
	fmt.Println(w.Wrap(text, 9))

	fmt.Println()
	fmt.Println("Optimal (MinimumRaggedness):")
	w.MinimumRaggedness = true
	fmt.Println(w.Wrap(text, 9))
	// Output:
	// Greedy (default):
	// a b c d e
	// f g h i j
	// k l m n o
	// p
	//
	// Optimal (MinimumRaggedness):
	// a b c d
	// e f g h
	// i j k l
	// m n o p
}

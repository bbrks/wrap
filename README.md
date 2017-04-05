# wrap [![Build Status](https://travis-ci.org/bbrks/wrap.svg)](https://travis-ci.org/bbrks/wrap) [![GoDoc](https://godoc.org/github.com/bbrks/wrap?status.svg)](https://godoc.org/github.com/bbrks/wrap) [![Go Report Card](https://goreportcard.com/badge/github.com/bbrks/wrap)](https://goreportcard.com/report/github.com/bbrks/wrap) [![GitHub tag](https://img.shields.io/github/tag/bbrks/wrap.svg)](https://github.com/bbrks/wrap/releases) [![license](https://img.shields.io/github/license/bbrks/wrap.svg)](https://github.com/bbrks/wrap/blob/master/LICENSE)

An efficient and flexible word-wrapping package for Go (golang)

## Usage

[embedmd]:# (wrap_test.go /\tvar loremIpsum/ /tellus.\n/)
```go
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
```

See [godoc.org/github.com/bbrks/wrap](https://godoc.org/github.com/bbrks/wrap) for examples using the `Wrapper` type to provide customisable breakpoints, prefixes, suffixes and more!

## Contributing

Issues/PRs are much appreciated!

Feature requests/improvements welcome.

## License
This project is licensed under the [MIT License](LICENSE.md).

package wrap

// Wrap is shorthand for declaring a new default Wrapper calling its Wrap method
func Wrap(s string, limit int) string {
	return NewWrapper().Wrap(s, limit)
}

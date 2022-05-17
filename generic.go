package utils

// Contains To use it, we have to make the following function call:
// Contains[string](["bar", "foo"], "foo") .
func Contains[T comparable](items []T, item T) bool {
	for _, v := range items {
		if v == item {
			return true
		}
	}
	return false
}

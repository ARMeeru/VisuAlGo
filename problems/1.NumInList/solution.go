package numinlist

// NumInList checks if a given number is present in a list of integers.
// It takes a slice of integers `list` and an integer `num` as input.
// It returns true if `num` is found in `list`, otherwise it returns false.
func NumInList(list []int, num int) bool {
	for _, v := range list {
		if v == num {
			return true
		}
	}
	return false
}
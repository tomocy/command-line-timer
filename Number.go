package main

type Number struct {
	size int
}

// Width calcutes the width of the number
func (num *Number) Width() int {
	return num.size / 3
}

// Height calculates the height of the number
func (num *Number) Height() int {
	return num.size * 2 / 3
}

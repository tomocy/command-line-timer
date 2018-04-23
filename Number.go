package main

import (
	"github.com/tomocy/command-line-timer/zero2nine"
)

// NumberInterface is interface for acting as a number
type NumberInterface interface {
	Width() int
	Height() int
	Show(x, y int)
}

// Number is the type for 0~9 number
type Number struct {
	size int
}

// Zero is the type for 0
type Zero struct {
	*Number
}

// One is the type for 1
type One struct {
	*Number
}

// Two is the type for 2
type Two struct {
	*Number
}

// Three is the type for 3
type Three struct {
	*Number
}

// Four is the type for 4
type Four struct {
	*Number
}

// Five is the type for 5
type Five struct {
	*Number
}

// Six is the type for 6
type Six struct {
	*Number
}

// Seven is the type for 7
type Seven struct {
	*Number
}

// Eight is the type for 8
type Eight struct {
	*Number
}

// Nine is the type for 9
type Nine struct {
	*Number
}

// Width calcutes the width of the number
func (num *Number) Width() int {
	return num.size * 3 / 4
}

// Height calculates the height of the number
func (num *Number) Height() int {
	var res = num.size * 4 / 5

	if res%2 == 0 {
		return res + 1
	}

	return res
}

// New returns a number instance implimenting NumberIntaface
func New(n int, size int) NumberInterface {
	z2n := zero2nine.FromInt(n)
	num := &Number{size: size}
	switch z2n {
	case zero2nine.Zero:
		return &Zero{num}
	case zero2nine.One:
		return &One{num}
	case zero2nine.Two:
		return &Two{num}
	case zero2nine.Three:
		return &Three{num}
	case zero2nine.Four:
		return &Four{num}
	case zero2nine.Five:
		return &Five{num}
	case zero2nine.Six:
		return &Six{num}
	case zero2nine.Seven:
		return &Seven{num}
	case zero2nine.Eight:
		return &Eight{num}
	case zero2nine.Nine:
		return &Nine{num}
	default:
		return nil
	}
}

func (zero *Zero) Show(x, y int) {
	const str = "="
	var originX = x
	var originY = y
	var width = zero.Width()
	var height = zero.Height()

	// horizontal line
	// top
	for i := 0; i <= width; i++ {
		drawLine(originX+i, originY, str)
	}

	// middle
	// for i := 0; i <= width; i++ {
	// 	drawLine(originX+i, originY+height/2, str)
	// }
	for i := 0; i <= 0; i++ {
		drawLine(originX+i, originY+height/2, str)
	}
	for i := width; i <= width; i++ {
		drawLine(originX+i, originY+height/2, str)
	}

	// bottom
	for i := 0; i <= width; i++ {
		drawLine(originX+i, originY+height-1, str)
	}

	// vertical line
	// left
	for i := 1; i < height/2; i++ {
		drawLine(originX, originY+i, str)
	}

	for i := height/2 + 1; i < height-1; i++ {
		drawLine(originX, originY+i, str)
	}

	// right
	for i := 1; i < height/2; i++ {
		drawLine(originX+width, originY+i, str)
	}

	for i := height/2 + 1; i < height-1; i++ {
		drawLine(originX+width, originY+i, str)
	}
}

func (one *One) Show(x, y int) {
	const str = "="
	var originX = x
	var originY = y
	var width = one.Width()
	var height = one.Height()

	// horizontal line
	// top
	for i := width; i <= width; i++ {
		drawLine(originX+i, originY, str)
	}

	// middle
	for i := width; i <= width; i++ {
		drawLine(originX+i, originY+height/2, str)
	}

	// bottom
	for i := width; i <= width; i++ {
		drawLine(originX+i, originY+height-1, str)
	}

	// vertical line
	// left
	// for i := 1; i < height/2; i++ {
	// 	drawLine(originX, originY+i, str)
	// }

	// for i := height/2 + 1; i < height-1; i++ {
	// 	drawLine(originX, originY+i, str)
	// }

	// right
	for i := 1; i < height/2; i++ {
		drawLine(originX+width, originY+i, str)
	}

	for i := height/2 + 1; i < height-1; i++ {
		drawLine(originX+width, originY+i, str)
	}
}

func (two *Two) Show(x, y int) {
	const str = "="
	var originX = x
	var originY = y
	var width = two.Width()
	var height = two.Height()

	// horizontal line
	// top
	for i := 0; i <= width; i++ {
		drawLine(originX+i, originY, str)
	}

	// middle
	for i := 0; i <= width; i++ {
		drawLine(originX+i, originY+height/2, str)
	}

	// bottom
	for i := 0; i <= width; i++ {
		drawLine(originX+i, originY+height-1, str)
	}

	// vertical line
	// left
	// for i := 1; i < height/2; i++ {
	// 	drawLine(originX, originY+i, str)
	// }

	for i := height/2 + 1; i < height-1; i++ {
		drawLine(originX, originY+i, str)
	}

	// right
	for i := 1; i < height/2; i++ {
		drawLine(originX+width, originY+i, str)
	}

	// for i := height/2 + 1; i < height-1; i++ {
	// 	drawLine(originX+width, originY+i, str)
	// }
}

func (three *Three) Show(x, y int) {
	const str = "="
	var originX = x
	var originY = y
	var width = three.Width()
	var height = three.Height()

	// horizontal line
	// top
	for i := 0; i <= width; i++ {
		drawLine(originX+i, originY, str)
	}

	// middle
	for i := 0; i <= width; i++ {
		drawLine(originX+i, originY+height/2, str)
	}

	// bottom
	for i := 0; i <= width; i++ {
		drawLine(originX+i, originY+height-1, str)
	}

	// vertical line
	// left
	// for i := 1; i < height/2; i++ {
	// 	drawLine(originX, originY+i, str)
	// }

	// for i := height/2 + 1; i < height-1; i++ {
	// 	drawLine(originX, originY+i, str)
	// }

	// right
	for i := 1; i < height/2; i++ {
		drawLine(originX+width, originY+i, str)
	}

	for i := height/2 + 1; i < height-1; i++ {
		drawLine(originX+width, originY+i, str)
	}
}

func (four *Four) Show(x, y int) {
	const str = "="
	var originX = x
	var originY = y
	var width = four.Width()
	var height = four.Height()

	// horizontal line
	// top
	// for i := 0; i <= width; i++ {
	// 	drawLine(originX+i, originY, str)
	// }
	for i := 0; i <= 0; i++ {
		drawLine(originX+i, originY, str)
	}
	for i := width; i <= width; i++ {
		drawLine(originX+i, originY, str)
	}

	// middle
	for i := 0; i <= width; i++ {
		drawLine(originX+i, originY+height/2, str)
	}

	// bottom
	// for i := 0; i <= width; i++ {
	// 	drawLine(originX+i, originY+height-1, str)
	// }
	for i := width; i <= width; i++ {
		drawLine(originX+i, originY+height-1, str)
	}

	// vertical line
	// left
	for i := 1; i < height/2; i++ {
		drawLine(originX, originY+i, str)
	}

	// for i := height/2 + 1; i < height-1; i++ {
	// 	drawLine(originX, originY+i, str)
	// }

	// right
	for i := 1; i < height/2; i++ {
		drawLine(originX+width, originY+i, str)
	}

	for i := height/2 + 1; i < height-1; i++ {
		drawLine(originX+width, originY+i, str)
	}
}

func (five *Five) Show(x, y int) {
	const str = "="
	var originX = x
	var originY = y
	var width = five.Width()
	var height = five.Height()

	// horizontal line
	// top
	for i := 0; i <= width; i++ {
		drawLine(originX+i, originY, str)
	}

	// middle
	for i := 0; i <= width; i++ {
		drawLine(originX+i, originY+height/2, str)
	}

	// bottom
	for i := 0; i <= width; i++ {
		drawLine(originX+i, originY+height-1, str)
	}

	// vertical line
	// left
	for i := 1; i < height/2; i++ {
		drawLine(originX, originY+i, str)
	}

	// for i := height/2 + 1; i < height-1; i++ {
	// 	drawLine(originX, originY+i, str)
	// }

	// right
	// for i := 1; i < height/2; i++ {
	// 	drawLine(originX+width, originY+i, str)
	// }

	for i := height/2 + 1; i < height-1; i++ {
		drawLine(originX+width, originY+i, str)
	}
}

func (six *Six) Show(x, y int) {
	const str = "="
	var originX = x
	var originY = y
	var width = six.Width()
	var height = six.Height()

	// horizontal line
	// top
	for i := 0; i <= width; i++ {
		drawLine(originX+i, originY, str)
	}

	// middle
	for i := 0; i <= width; i++ {
		drawLine(originX+i, originY+height/2, str)
	}

	// bottom
	for i := 0; i <= width; i++ {
		drawLine(originX+i, originY+height-1, str)
	}

	// vertical line
	// left
	for i := 1; i < height/2; i++ {
		drawLine(originX, originY+i, str)
	}

	for i := height/2 + 1; i < height-1; i++ {
		drawLine(originX, originY+i, str)
	}

	// right
	// for i := 1; i < height/2; i++ {
	// 	drawLine(originX+width, originY+i, str)
	// }

	for i := height/2 + 1; i < height-1; i++ {
		drawLine(originX+width, originY+i, str)
	}
}

func (seven *Seven) Show(x, y int) {
	const str = "="
	var originX = x
	var originY = y
	var width = seven.Width()
	var height = seven.Height()

	// horizontal line
	// top
	for i := 0; i <= width; i++ {
		drawLine(originX+i, originY, str)
	}

	// middle
	// for i := 0; i <= width; i++ {
	// 	drawLine(originX+i, originY+height/2, str)
	// }
	for i := width; i <= width; i++ {
		drawLine(originX+i, originY+height/2, str)
	}

	// bottom
	// for i := 0; i <= width; i++ {
	// 	drawLine(originX+i, originY+height-1, str)
	// }
	for i := width; i <= width; i++ {
		drawLine(originX+i, originY+height-1, str)
	}

	// vertical line
	// left
	// for i := 1; i < height/2; i++ {
	// 	drawLine(originX, originY+i, str)
	// }

	// for i := height/2 + 1; i < height-1; i++ {
	// 	drawLine(originX, originY+i, str)
	// }

	// right
	for i := 1; i < height/2; i++ {
		drawLine(originX+width, originY+i, str)
	}

	for i := height/2 + 1; i < height-1; i++ {
		drawLine(originX+width, originY+i, str)
	}
}

func (eight *Eight) Show(x, y int) {
	const str = "="
	var originX = x
	var originY = y
	var width = eight.Width()
	var height = eight.Height()

	// horizontal line
	// top
	for i := 0; i <= width; i++ {
		drawLine(originX+i, originY, str)
	}

	// middle
	for i := 0; i <= width; i++ {
		drawLine(originX+i, originY+height/2, str)
	}

	// bottom
	for i := 0; i <= width; i++ {
		drawLine(originX+i, originY+height-1, str)
	}

	// vertical line
	// left
	for i := 1; i < height/2; i++ {
		drawLine(originX, originY+i, str)
	}

	for i := height/2 + 1; i < height-1; i++ {
		drawLine(originX, originY+i, str)
	}

	// right
	for i := 1; i < height/2; i++ {
		drawLine(originX+width, originY+i, str)
	}

	for i := height/2 + 1; i < height-1; i++ {
		drawLine(originX+width, originY+i, str)
	}
}

func (nine *Nine) Show(x, y int) {
	const str = "="
	var originX = x
	var originY = y
	var width = nine.Width()
	var height = nine.Height()

	// horizontal line
	// top
	for i := 0; i <= width; i++ {
		drawLine(originX+i, originY, str)
	}

	// middle
	for i := 0; i <= width; i++ {
		drawLine(originX+i, originY+height/2, str)
	}

	// bottom
	for i := 0; i <= width; i++ {
		drawLine(originX+i, originY+height-1, str)
	}

	// vertical line
	// left
	for i := 1; i < height/2; i++ {
		drawLine(originX, originY+i, str)
	}

	// for i := height/2 + 1; i < height-1; i++ {
	// 	drawLine(originX, originY+i, str)
	// }

	// right
	for i := 1; i < height/2; i++ {
		drawLine(originX+width, originY+i, str)
	}

	for i := height/2 + 1; i < height-1; i++ {
		drawLine(originX+width, originY+i, str)
	}
}

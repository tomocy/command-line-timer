package main

import (
	"github.com/tomocy/command-line-timer/zero2nine"
)

// Number is interface for acting as a number
type Number interface {
	Width() int
	Height() int
	Show(x, y int)
}

// Number is the type for 0~9 number
type number struct {
	size int
}

// zero is the type for 0
type zero struct {
	*number
}

// one is the type for 1
type one struct {
	*number
}

// two is the type for 2
type two struct {
	*number
}

// three is the type for 3
type three struct {
	*number
}

// four is the type for 4
type four struct {
	*number
}

// five is the type for 5
type five struct {
	*number
}

// six is the type for 6
type six struct {
	*number
}

// seven is the type for 7
type seven struct {
	*number
}

// eight is the type for 8
type eight struct {
	*number
}

// nine is the type for 9
type nine struct {
	*number
}

// Width calcutes the width of the number
func (num *number) Width() int {
	return num.size * 3 / 4
}

// Height calculates the height of the number
func (num *number) Height() int {
	var res = num.size * 4 / 5

	if res%2 == 0 {
		return res + 1
	}

	return res
}

// New returns a number instance implimenting Number
func New(n int, size int) Number {
	z2n := zero2nine.FromInt(n)
	num := &number{size: size}
	switch z2n {
	case zero2nine.Zero:
		return &zero{num}
	case zero2nine.One:
		return &one{num}
	case zero2nine.Two:
		return &two{num}
	case zero2nine.Three:
		return &three{num}
	case zero2nine.Four:
		return &four{num}
	case zero2nine.Five:
		return &five{num}
	case zero2nine.Six:
		return &six{num}
	case zero2nine.Seven:
		return &seven{num}
	case zero2nine.Eight:
		return &eight{num}
	case zero2nine.Nine:
		return &nine{num}
	default:
		return nil
	}
}

func (zero *zero) Show(x, y int) {
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

func (one *one) Show(x, y int) {
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

func (two *two) Show(x, y int) {
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

func (three *three) Show(x, y int) {
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

func (four *four) Show(x, y int) {
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

func (five *five) Show(x, y int) {
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

func (six *six) Show(x, y int) {
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

func (seven *seven) Show(x, y int) {
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

func (eight *eight) Show(x, y int) {
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

func (nine *nine) Show(x, y int) {
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

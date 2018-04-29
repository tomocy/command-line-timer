package terminal

import (
	"os/exec"
	"strconv"
)

func Width() int {
	out, err := exec.Command("tput", "cols").Output()
	if err != nil {
		panic(err)
	}

	// delete \n at the last index
	just := out[:len(out)-1]

	width, err := strconv.Atoi(string(just))
	if err != nil {
		panic(err)
	}

	return width
}

func Height() int {
	out, err := exec.Command("tput", "lines").Output()
	if err != nil {
		panic(err)
	}

	// delete \n at the last index
	just := out[:len(out)-1]

	height, err := strconv.Atoi(string(just))
	if err != nil {
		panic(err)
	}

	return height
}

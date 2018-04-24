package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"time"

	termbox "github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	defer termbox.Close()

	var now int = 0
	keyCh := make(chan termbox.Key, 1)
	timeCh := make(chan int, 1)

	go listenKeyEvent(keyCh)
	go timer(now, timeCh)
	go draw(timeCh)
	<-keyCh
}

func listenKeyEvent(kch chan termbox.Key) {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			kch <- ev.Key
		}
	}
}

func timer(now int, tch chan int) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			now++
			tch <- now
		}
	}
}

func draw(tch chan int) {
	show(0)

	for {
		t := <-tch
		show(t)
	}
}

func show(t int) {
	originX := terminalWidth() / 3
	originY := terminalHeight() / 3
	size := 8
	digitsNum := digitsNum(t, size)
	intvl := 2
	sepIntvl := 3

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	drawLine(0, 0, "Timer")
	drawLine(0, 1, "Enter esc to stop")

	// hour
	hX := originX
	hMgn := digitsNum["h2"].Width() + intvl
	digitsNum["h2"].Show(hX, originY)
	digitsNum["h"].Show(hX+hMgn, originY)

	// separator
	sepX := hX + hMgn + digitsNum["h"].Width() + sepIntvl
	sepY := originY + digitsNum["h"].Height()/2 - 1
	drawLine(sepX, sepY, "=")
	drawLine(sepX, sepY+2, "=")

	// minute
	mX := sepX + sepIntvl
	mMgn := digitsNum["m2"].Width() + intvl
	digitsNum["m2"].Show(mX, originY)
	digitsNum["m"].Show(mX+mMgn, originY)

	// second
	sX := mX + mMgn + digitsNum["m"].Width() + intvl
	sMgn := digitsNum["s2"].Width() + intvl
	digitsNum["s2"].Show(sX, originY)
	digitsNum["s"].Show(sX+sMgn, originY)

	termbox.Flush()
}

func terminalWidth() int {
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

func terminalHeight() int {
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

func digitsNum(t, size int) map[string]Number {
	digits := digitsInt(t)
	secSize := size * 7 / 11

	return map[string]Number{
		"s":  New(digits["s"], secSize),
		"s2": New(digits["s2"], secSize),
		"m":  New(digits["m"], size),
		"m2": New(digits["m2"], size),
		"h":  New(digits["h"], size),
		"h2": New(digits["h2"], size),
	}
}

func digitsInt(t int) map[string]int {
	ss := fmt.Sprintf("%02d", t%60)
	ms := fmt.Sprintf("%02d", t/60)
	hs := fmt.Sprintf("%02d", t/3600)
	s, err := strconv.Atoi(string(ss[1]))
	if err != nil {
		panic(err)
	}
	s2, err := strconv.Atoi(string(ss[0]))
	if err != nil {
		panic(err)
	}
	m, err := strconv.Atoi(string(ms[1]))
	if err != nil {
		panic(err)
	}
	m2, err := strconv.Atoi(string(ms[0]))
	if err != nil {
		panic(err)
	}
	h, err := strconv.Atoi(string(hs[1]))
	if err != nil {
		panic(err)
	}
	h2, err := strconv.Atoi(string(hs[0]))
	if err != nil {
		panic(err)
	}
	return map[string]int{
		"s":  s,
		"s2": s2,
		"m":  m,
		"m2": m2,
		"h":  h,
		"h2": h2,
	}
}

func drawLine(x, y int, str string) {
	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		termbox.SetCell(x+i, y, runes[i], termbox.ColorDefault, termbox.ColorDefault)
	}
}

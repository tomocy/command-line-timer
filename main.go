package main

import (
	"fmt"
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
	originX := 15
	originY := 15
	digitsNum := digitsNum(t)

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	drawLine(0, 0, "Timer")
	drawLine(0, 1, "Enter esc to stop")

	// hour
	hX := originX
	digitsNum["h2"].Show(hX, originY)
	digitsNum["h"].Show(hX+10, originY)

	// separator
	sepX := hX + 10 + 11
	sepY := originY + 3
	drawLine(sepX, sepY, "=")
	drawLine(sepX, sepY+2, "=")

	// minute
	mX := sepX + 3
	digitsNum["m2"].Show(mX, originY)
	digitsNum["m"].Show(mX+10, originY)

	// second
	sX := mX + 10 + 10
	digitsNum["s2"].Show(sX, originY)
	digitsNum["s"].Show(sX+7, originY)

	termbox.Flush()
}

func digitsNum(t int) map[string]NumberInterface {
	digits := digitsInt(t)
	secSize := 7
	size := 11

	return map[string]NumberInterface{
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

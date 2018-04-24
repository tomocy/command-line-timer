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
	digits := digitsInt(t)
	secSize := 7
	size := 11
	sNum := New(digits["s"], secSize)
	s2Num := New(digits["s2"], secSize)
	mNum := New(digits["m"], size)
	m2Num := New(digits["m2"], size)
	hNum := New(digits["h"], size)
	h2Num := New(digits["h2"], size)
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	drawLine(0, 0, "Timer")
	drawLine(0, 1, "Enter esc to stop")
	sNum.Show(56, 5)
	s2Num.Show(49, 5)

	mNum.Show(39, 5)
	m2Num.Show(29, 5)

	drawLine(26, 8, "=")
	drawLine(26, 10, "=")

	hNum.Show(15, 5)
	h2Num.Show(5, 5)
	termbox.Flush()
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

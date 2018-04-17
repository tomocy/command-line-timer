package main

import (
	"fmt"
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
	show("0")

	for {
		t := <-tch
		show(fmt.Sprint(t))
	}
}

func show(time string) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	drawLine(0, 0, "Timer")
	drawLine(0, 1, "Enter esc to stop")
	drawLine(0, 2, time)
	termbox.Flush()
}

func drawLine(x, y int, str string) {
	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		termbox.SetCell(x+i, y, runes[i], termbox.ColorDefault, termbox.ColorDefault)
	}
}

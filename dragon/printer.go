package main

import "fmt"

type printer struct {
	dragon   *Dragon
	receiver chan string
}

func newPrinter() *printer {
	return &printer{
		receiver: make(chan string),
	}
}

func (p *printer) setDragon(d *Dragon) {
	p.dragon = d
}

func (p *printer) printing() {
	go func() {
		for {
			msg := <-p.receiver

			fmt.Print("\033[H\033[2J")
			fmt.Println(msg)
		}
	}()
}

func (p *printer) print(msg string) {
	p.receiver <- msg
}

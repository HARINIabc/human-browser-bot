package browser

import (
	"fmt"

	"human-browser-bot/behavior"
)

type Browser struct{}

func New() *Browser {
	return &Browser{}
}

func (b *Browser) MoveMouse(x, y int) {
	fmt.Println("-> Browser: moving mouse")
	behavior.MoveMouseHuman(x, y)
}

func (b *Browser) TypeText(text string) {
	fmt.Println("-> Browser: typing text")
	behavior.TypeHuman(text)
}

func (b *Browser) Scroll() {
	fmt.Println("->Browser: scrolling")
	behavior.ScrollHuman()
}

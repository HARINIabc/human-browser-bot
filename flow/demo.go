package flow

import (
	"fmt"

	"human-browser-bot/browser"
)

func DemoFlow(b *browser.Browser) {
	fmt.Println("\n Simulating login flow")
	b.MoveMouse(200, 120)
	b.TypeText("HariniMandalam@example.com")
	b.TypeText("••••••••")

	fmt.Println("\n Simulating search")
	b.MoveMouse(400, 80)
	b.TypeText("software engineer")
	b.Scroll()

	fmt.Println("\n Simulating profile browsing")
	b.Scroll()
	b.MoveMouse(350, 300)
}

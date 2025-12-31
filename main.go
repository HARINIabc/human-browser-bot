package main

import (
	"fmt"

	"human-browser-bot/browser"
	"human-browser-bot/config"
	"human-browser-bot/flow"
	"human-browser-bot/limits"
	"human-browser-bot/scheduler"
)

func main() {
	fmt.Println("🤖 Bot starting...\n")

	cfg := config.Load()

	s := scheduler.New(cfg.StartHour, cfg.EndHour)
	l := limits.New(cfg.MaxActions)
	b := browser.New()

	if s.IsAllowed() && l.Allow() {
		flow.DemoFlow(b)
	} else {
		fmt.Println("Not allowed to run demo now")
	}

	fmt.Println("\n Demo finished")
}

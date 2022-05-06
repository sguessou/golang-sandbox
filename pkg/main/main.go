package main

import (
	"sguessou/pkg/display"
	"sguessou/pkg/msg"
)

func main() {
	msg.Hi()
	display.Display("Hello from display")
	msg.Exciting("an exciting message")
}

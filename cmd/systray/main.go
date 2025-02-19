package main

import (
	"github.com/energye/systray"
	"github.com/isaqueveras/zoity/internal"
)

func main() {
	systray.Run(internal.Setup, internal.Exit)
}

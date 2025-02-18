package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

func main() {
	g, err := gocui.NewGui(gocui.Output256)
	if err != nil {
		log.Fatal(err)
	}
	defer g.Close()

	// g.CurrentView().Highlight = true
	g.Cursor = true
	g.SelFgColor = gocui.ColorRed

	g.SetManagerFunc(func(g *gocui.Gui) error {
		maxX, maxY := g.Size()

		if v, err := g.SetView("projects", 1, 1, 2, 2); err != nil {
			v.Title = "Projects"

			fmt.Fprintln(v, "Projects")
		}

		if v, err := g.SetView("hello", maxX/2-7, maxY/2, maxX/2+7, maxY/2+2); err != nil {
			if _, err := g.SetCurrentView("hello"); err != nil {
				return err
			}

			fmt.Fprintln(v, "Hello world!")
		}

		return nil
	})

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		return gocui.ErrQuit
	}); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && !errors.Is(err, gocui.ErrQuit) {
		log.Panic(err.Error())
	}
}

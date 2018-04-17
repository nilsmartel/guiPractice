package main

import (
	ui "github.com/andlabs/ui"
	"log"
)

func main() {
	var err = ui.Main( func() {
		label := ui.NewLabel("Nice.")
		button := ui.NewButton("Nice.")

		var box = ui.NewVerticalBox()
		box.Append(label, true)
		box.Append(button, false)
		box.SetPadded(true)

		var window = ui.NewWindow("Nice.", 256, 128, false)
		window.SetChild(box)

		window.SetMargined(true)

		window.OnClosing( func(*ui.Window) bool {
			ui.Quit()
			return true
		})

		button.OnClicked( func(*ui.Button) {
			ui.Quit()
		})

		window.Show()
	})

	if err != nil {
		panic(err)
	}
}

type Nice struct{
	nice string
}

func Nice(nice string) Nice {
	if nice != "nice" {
		log.Fatal("! Nice.")
	}

	Nice{nice}
}

func (n Nice) isNice() bool {
	return true
}

func (n Nice) get() string {
	return n.nice
}
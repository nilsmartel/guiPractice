package main

import (
	ui "github.com/andlabs/ui"
)

func main() {
	err := ui.Main(func() {
		buttons := []*ui.Button{
			ui.NewButton("Go"),
			ui.NewButton("Rust"),
			ui.NewButton("Swift"),
			ui.NewButton("Kotlin"),
			ui.NewButton("C"),
			ui.NewButton("Java"),
			ui.NewButton("JavaScript"),
			ui.NewButton("Python"),
		}
		buttonId := 0
		boxContainer := [3]*ui.Box{}
		topContainer := ui.NewHorizontalBox()
		for i := 0; i <3; i++ {
			boxContainer[i] = ui.NewVerticalBox()
			boxContainer[i].Append(buttons[buttonId], false)
			boxContainer[i].Append(buttons[buttonId+1], false)
			topContainer.Append(boxContainer[i], true)
			buttonId+=2
		}

		textLabel := ui.NewLabel("")

		box := ui.NewVerticalBox()
		box.Append(topContainer, true)
		box.Append(textLabel, false)

		window := ui.NewWindow("Buttons", 200, 100, false)
		window.SetChild(box)

		for _,button := range buttons {
			button.OnClicked(func(btn *ui.Button) {
				textLabel.SetText(btn.Text())
			})
		}

		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()
	})
	if err != nil {
		panic(err)
	}
}
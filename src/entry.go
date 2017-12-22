package main

import (
	ui "github.com/andlabs/ui"
)

func main() {
	var err = ui.Main( func() {
		var entry = ui.NewEntry()
		var label = ui.NewLabel("")

		entry.OnChanged( func(entry *ui.Entry) {
			label.SetText(entry.Text())
		})

		var box = ui.NewHorizontalBox()
		box.Append(entry, false)
		box.Append(label, true)
		box.SetPadded(true)

		var window = ui.NewWindow("Entry", 256, 128, false)
		window.SetChild(box)

		window.OnClosing( func(*ui.Window) bool {
			ui.Quit()
			return true
		})

		window.Show()
	})

	if err != nil {
		panic(err)
	}
}
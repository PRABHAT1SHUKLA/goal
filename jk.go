package main

import (
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	table := tview.NewTable().
		SetBorders(true).
		SetSelectable(true, false)

	table.SetCell(0, 0, tview.NewTableCell("ID").SetAlign(tview.AlignCenter))
	table.SetCell(0, 1, tview.NewTableCell("Name").SetAlign(tview.AlignCenter))
	table.SetCell(1, 0, tview.NewTableCell("1"))
	table.SetCell(1, 1, tview.NewTableCell("Prabhat"))

	if err := app.SetRoot(table, true).Run(); err != nil {
		panic(err)
	}
}

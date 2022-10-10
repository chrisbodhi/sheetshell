// Demo code for the Table primitive.
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"golang.design/x/clipboard"
)

func main() {
	app := tview.NewApplication()
	table := tview.NewTable().
		SetBorders(true)
	envVarPairs := os.Environ()
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}
	var keys []string
	var vals []string
	// TODO: alpabetize
	for _, pair := range envVarPairs {
		kv := strings.SplitN(pair, "=", 2)
		k := kv[0]
		v := kv[1]
		keys = append(keys, k)
		vals = append(vals, v)
	}
	assert(len(keys), len(vals))
	rows := len(keys)
	for r := 0; r < rows; r++ {
		table.SetCell(r, 0,
			tview.NewTableCell(keys[r]).
				SetTextColor(tcell.ColorYellow).
				SetAlign(tview.AlignCenter))
		table.SetCell(r, 1,
			tview.NewTableCell(vals[r]).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignCenter)) // TODO: what to do with long boys?
	}
	// TODO: use a better fixed header, or just remove it
	table.Select(0, 0).SetFixed(1, 1).SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEscape {
			app.Stop()
		}
		if key == tcell.KeyEnter {
			table.SetSelectable(true, true)
		}
	}).SetSelectedFunc(func(row int, column int) {
		cell := table.GetCell(row, column).SetTextColor(tcell.ColorRed)
		clipboard.Write(clipboard.FmtText, []byte(cell.Text))
		table.SetSelectable(false, false)
	})
	if err := app.SetRoot(table, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func assert(i1, i2 int) {
	if i1 != i2 {
		panic(fmt.Sprintf("assertion failed, encountered mismatch of size %d", abs(i1-i2)))
	}
}

func abs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}

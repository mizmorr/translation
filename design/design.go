package design

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func Design() {

	myApp := app.New()
	myWindow := myApp.NewWindow("Syntax analyzer")
	myWindow.Resize(fyne.NewSize(800, 500))
	label := widget.NewLabel("identifiers")
	label2 := widget.NewLabel("numeric const")
	label3 := widget.NewLabel("symbol cons")

	data_numeric := [][]string{{"N0", "0"}}
	table_numeric_const := widget.NewTable(
		func() (int, int) { return 1, 2 },
		func() fyne.CanvasObject {
			return widget.NewLabel("numeric")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data_numeric[i.Row][i.Col])

		})
	data_symbol := [][]string{{"C0", "_"}}
	table_symbol_const := widget.NewTable(
		func() (int, int) { return 1, 2 },
		func() fyne.CanvasObject {
			return widget.NewLabel("symbol")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data_symbol[i.Row][i.Col])

		})
	data_identifier := [][]string{{"I0", "num0"}}
	table_identifiers := widget.NewTable(
		func() (int, int) { return 1, 2 },
		func() fyne.CanvasObject {
			return widget.NewLabel("identifiers")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data_identifier[i.Row][i.Col])

		})

	data := [][]string{{"1", "2"}, {"3", "4"}, {"1", "2"}, {"3", "4"}, {"1", "2"}, {"3", "4"}}
	input, result, _ := widget.NewMultiLineEntry(), widget.NewLabelWithStyle("", fyne.TextAlignCenter, fyne.TextStyle{Bold: true, TabWidth: 2}), widget.NewEntry()

	button_res := widget.NewButtonWithIcon("parse", theme.ConfirmIcon(), func() {
		fmt.Println("button pushed")
		text := input.Text
		result.SetText(text)
		table_numeric_const.Length = func() (int, int) { return len(data), len(data[0]) }
		table_numeric_const.UpdateCell = func(id widget.TableCellID, template fyne.CanvasObject) {
			label1 := template.(*widget.Label)
			label1.SetText(data[id.Row][id.Col])
		}
	})
	button_refresh := widget.NewButtonWithIcon("refresh", theme.ContentUndoIcon(), func() {
		fmt.Println("button pushed")
		result.SetText("")
		table_numeric_const.Length = func() (int, int) { return 1, 2 }

		table_numeric_const.UpdateCell = func(id widget.TableCellID, template fyne.CanvasObject) {
			label1 := template.(*widget.Label)
			label1.SetText(data_numeric[id.Row][id.Col])
		}
	})
	input.PlaceHolder = "Java code input"
	first_row := container.NewGridWithColumns(3, layout.NewSpacer(), container.NewMax(table_numeric_const), layout.NewSpacer())
	second_row := container.NewGridWithColumns(3, layout.NewSpacer(), label2, layout.NewSpacer())
	third_row := container.NewGridWithColumns(3, layout.NewSpacer(), container.NewMax(table_symbol_const), layout.NewSpacer())
	fourth_row := container.NewGridWithColumns(3, layout.NewSpacer(), label3, layout.NewSpacer())
	fifth_row := container.NewGridWithColumns(3, layout.NewSpacer(), container.NewMax(table_identifiers), layout.NewSpacer())
	sixth_row := container.NewGridWithColumns(3, layout.NewSpacer(), label, layout.NewSpacer())

	first_column := container.NewGridWithRows(3, container.NewPadded(input), container.NewGridWithColumns(3, layout.NewSpacer(), container.NewGridWithRows(3, layout.NewSpacer(), container.NewVBox(button_res, button_refresh), layout.NewSpacer())), container.NewPadded(result))
	second_column := container.NewGridWithRows(6, first_row, second_row, third_row, fourth_row, fifth_row, sixth_row)
	cont := container.NewGridWithColumns(2, first_column, second_column)

	myWindow.SetContent(cont)

	myWindow.ShowAndRun()
}

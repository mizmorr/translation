package design

import (
	"fmt"
	"strings"

	parse "gojson/parser"

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
	myWindow.Resize(fyne.NewSize(600, 500))
	myWindow.SetFixedSize(true)
	label := widget.NewLabel("Идентификаторы I")
	label2 := widget.NewLabel("Числовые константы N")
	label3 := widget.NewLabel("Символьные константы C")
	label4 := widget.NewLabel("Служебные слова W")
	label5 := widget.NewLabel("Операторы O")
	label6 := widget.NewLabel("Разделители R")

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

	table_symbol_const.SetColumnWidth(1, 100)
	table_symbol_const.SetColumnWidth(0, 40)
	table_numeric_const.SetColumnWidth(1, 100)
	table_numeric_const.SetColumnWidth(0, 40)
	table_identifiers.SetColumnWidth(0, 40)
	table_identifiers.SetColumnWidth(1, 100)

	input, result, _ := widget.NewMultiLineEntry(), widget.NewLabelWithStyle("", fyne.TextAlignLeading, fyne.TextStyle{Bold: true, TabWidth: 2}), widget.NewEntry()
	keywords := parse.Get_data(3)
	operators := parse.Get_data(4)
	separators := parse.Get_data(5)
	table_keywords := widget.NewTable(
		func() (int, int) { return len(keywords), len(keywords[0]) },
		func() fyne.CanvasObject {
			return widget.NewLabel("keywords")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(keywords[i.Row][i.Col])

		})
	table_operators := widget.NewTable(
		func() (int, int) { return len(operators), len(operators[0]) },
		func() fyne.CanvasObject {
			return widget.NewLabel("operators")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(operators[i.Row][i.Col])

		})
	table_separators := widget.NewTable(
		func() (int, int) { return len(separators), len(separators[0]) },
		func() fyne.CanvasObject {
			return widget.NewLabel("separators")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(separators[i.Row][i.Col])

		})
	button_res := widget.NewButtonWithIcon("parse", theme.ConfirmIcon(), func() {
		fmt.Println("button pushed")
		java_code := input.Text
		for _, java_string := range strings.Split(java_code, "\n") {
			parse.Parse(java_string)
		}
		result_string := parse.Get_result()
		identifiers := parse.Get_data(0)
		numeric := parse.Get_data(1)
		symbol := parse.Get_data(2)

		if len(numeric) > 0 {
			table_numeric_const.Length = func() (int, int) { return len(numeric), len(numeric[0]) }
			table_numeric_const.UpdateCell = func(id widget.TableCellID, template fyne.CanvasObject) {
				label1 := template.(*widget.Label)
				label1.SetText(numeric[id.Row][id.Col])
			}
		}
		if len(symbol) > 0 {
			table_symbol_const.Length = func() (int, int) { return len(symbol), len(symbol[0]) }
			table_symbol_const.UpdateCell = func(id widget.TableCellID, template fyne.CanvasObject) {
				label1 := template.(*widget.Label)
				label1.SetText(symbol[id.Row][id.Col])
			}
		}
		if len(identifiers) > 0 {
			table_identifiers.Length = func() (int, int) { return len(identifiers), len(identifiers[0]) }
			table_identifiers.UpdateCell = func(id widget.TableCellID, template fyne.CanvasObject) {
				label1 := template.(*widget.Label)
				label1.SetText(identifiers[id.Row][id.Col])
			}
		}
		result.SetText(result_string)
	})
	button_refresh := widget.NewButtonWithIcon("refresh", theme.ContentUndoIcon(), func() {
		result.SetText("")
		table_numeric_const.Length = func() (int, int) { return 1, 2 }

		table_numeric_const.UpdateCell = func(id widget.TableCellID, template fyne.CanvasObject) {
			label1 := template.(*widget.Label)
			label1.SetText(data_numeric[id.Row][id.Col])
		}

		table_symbol_const.Length = func() (int, int) { return 1, 2 }

		table_symbol_const.UpdateCell = func(id widget.TableCellID, template fyne.CanvasObject) {
			label1 := template.(*widget.Label)
			label1.SetText(data_symbol[id.Row][id.Col])
		}
		table_identifiers.Length = func() (int, int) { return 1, 2 }

		table_identifiers.UpdateCell = func(id widget.TableCellID, template fyne.CanvasObject) {
			label1 := template.(*widget.Label)
			label1.SetText(data_identifier[id.Row][id.Col])
		}
		parse.Cleaner()
	})
	input.PlaceHolder = "Java code input"
	first_row := container.NewGridWithColumns(3, layout.NewSpacer(), container.NewScroll(table_numeric_const))
	second_row := container.NewGridWithColumns(3, layout.NewSpacer(), label2, layout.NewSpacer())
	third_row := container.NewGridWithColumns(3, layout.NewSpacer(), container.NewScroll(table_symbol_const), layout.NewSpacer())
	fourth_row := container.NewGridWithColumns(3, layout.NewSpacer(), label3, layout.NewSpacer())
	fifth_row := container.NewGridWithColumns(3, layout.NewSpacer(), container.NewScroll(table_identifiers), layout.NewSpacer())
	sixth_row := container.NewGridWithColumns(3, layout.NewSpacer(), label, layout.NewSpacer())

	first_column := container.NewGridWithRows(3, container.NewPadded(input), container.NewGridWithColumns(3, layout.NewSpacer(), container.NewGridWithRows(3, layout.NewSpacer(), container.NewVBox(button_res, button_refresh), layout.NewSpacer())), container.NewScroll(result))
	second_column := container.NewGridWithRows(6, first_row, second_row, third_row, fourth_row, fifth_row, sixth_row)
	third_column := container.NewGridWithRows(6, container.NewScroll(table_keywords), label4, container.NewScroll(table_operators), label5, container.NewScroll(table_separators), label6)
	cont := container.NewGridWithColumns(3, first_column, second_column, third_column)
	myWindow.SetContent(cont)
	myApp.Settings().SetTheme(theme.LightTheme())
	myWindow.ShowAndRun()
}
